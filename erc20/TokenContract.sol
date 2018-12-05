pragma solidity ^0.5.0;

contract Token {
    uint256 public totalSupply;

    /* This generates a public event on the blockchain that will notify clients */
    event Transfer(address indexed from, address indexed to, uint256 value);

    function getHolderContract(address a) internal view returns (Holder) {
        bytes32 holder_init_code_hash = 0x17bf60e63d4e2cc380fc7b8d419f6ebf015f77b0916680115f61cc123093468f;
        address payable holder_address = address(uint256(keccak256(abi.encodePacked(byte(0xff),factory,bytes32(uint256(a)),holder_init_code_hash))));
        uint256 codeLength;
        assembly {codeLength := extcodesize(holder_address)}
        require(codeLength > 0); // Not creating a new holder contract here, because if it does have ETH for rent, it can get evicted
        return Holder(holder_address);
    }

    /* Send tokens */
    function transfer(address _to, uint256 _value) public returns (bool) {
        Holder fromContract = getHolderContract(msg.sender);
        Holder toContract = getHolderContract(_to);
        uint256 fromBalance = fromContract.getBalance();
        uint256 toBalance = toContract.getBalance();
        require(fromBalance >= _value);                // Check if the sender has enough
        require(toBalance + _value >= toBalance);      // Check for overflows
        fromContract.setBalance(fromBalance - _value); // Subtract from the sender
        toContract.setBalance(toBalance + _value);     // Add the same to the recipient
        emit Transfer(msg.sender, _to, _value);        // Notify anyone listening that this transfer took place
        return true;
    }

    /* Allow another contract to spend some tokens in your behalf */
    function approve(address _spender, uint256 _value) public returns (bool) {
        Holder fromContract = getHolderContract(msg.sender);
        fromContract.setAllowance(_spender, _value);
        return true;
    }

    /* A contract attempts to get the tokens */
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        Holder fromContract = getHolderContract(_from);
        Holder toContract = getHolderContract(_to);
        uint256 fromBalance = fromContract.getBalance();
        uint256 toBalance = toContract.getBalance();
        uint256 fromAllowance = fromContract.getAllowance(msg.sender);
        require(fromBalance >= _value);                      // Check if the sender has enough
        require(toBalance + _value >= toBalance);            // Check for overflows
        require(_value <= fromAllowance);                        // Check allowance
        fromContract.setBalance(fromBalance - _value);       // Subtract from the sender
        toContract.setBalance(toBalance + _value);           // Add the same to the recipient
        fromContract.setAllowance(msg.sender, fromAllowance - _value);
        emit Transfer(_from, _to, _value);
        return true;
    }

    function balanceOf(address a) public view returns (uint256) {
        Holder c = getHolderContract(a);
        return c.getBalance();
    }

    function allowance(address a, address _spender) public view returns (uint256) {
        Holder c = getHolderContract(a);
        return c.getAllowance(_spender);
    }

    address public minter;
    address public factory;

    constructor(address _minter, address _factory) public {
        minter = _minter;
        factory = _factory;
    }

    // Needs to accept ETH to pay rent
    function() external payable {
    }

    /* Allows the owner to mint more tokens */
    function mint(address _to, uint256 _value) public returns (bool) {
        require(msg.sender == minter);                       // Only the minter is allowed to mint
        Holder toContract = getHolderContract(_to);
        uint256 toBalance = toContract.getBalance();
        require(toBalance + _value >= toBalance);  // Check for overflows
        toContract.setBalance(toBalance + _value);
        totalSupply += _value;
        return true;
    }
}

contract Holder {
    address public constant SENTINEL_TOKEN_CONTRACTS = address(0x1);
    mapping(address => address) internal tokenContracts;
    address payable internal owner;
    address internal factory; // This will be non-zero only until setOwner is successfully invoked
    mapping(address => uint256) internal balances;
    mapping(address => mapping(address => uint256)) internal allowances;

    constructor() public {
        factory = msg.sender;
        tokenContracts[SENTINEL_TOKEN_CONTRACTS] = SENTINEL_TOKEN_CONTRACTS;
    }

    function setOwner(address payable _owner, bytes32 init_code_hash) public {
        require(address(uint256(keccak256(abi.encodePacked(byte(0xff),factory,bytes32(uint256(_owner)),init_code_hash)))) == address(this));
        owner = _owner;
        factory = address(0);
    }

    // Needs to accept ETH to pay rent
    function() external payable {
    }

    function addTokenContract(address tokenContract) public {
        require(owner != address(0));
        require(msg.sender == owner);
        require(tokenContract != address(0) && tokenContract != SENTINEL_TOKEN_CONTRACTS);
        require(tokenContracts[tokenContract] == address(0)); // Not yet in the list
        tokenContracts[tokenContract] = tokenContracts[SENTINEL_TOKEN_CONTRACTS];
        tokenContracts[SENTINEL_TOKEN_CONTRACTS] = tokenContract;
    }

    function removeTokenContract(address tokenContract, address prev) public {
        require(owner != address(0));
        require(msg.sender == owner);
        require(tokenContract != address(0) && tokenContract != SENTINEL_TOKEN_CONTRACTS);
        require(tokenContracts[prev] == tokenContract);
        tokenContracts[prev] = tokenContracts[tokenContract];
        tokenContracts[tokenContract] = address(0);
    }

    // Removes all the tokens and returns ETH balance to the owner
    function destroy() public {
        require(owner != address(0));
        require(msg.sender == owner);
        //TODO: Notify all the contracts so they can reduce supply accordingly - these tokens are gone forever!!!
        selfdestruct(owner);
    }

    function getBalance() public view returns (uint256) {
        require(msg.sender != SENTINEL_TOKEN_CONTRACTS);
        require(tokenContracts[msg.sender] != address(0));
        return balances[msg.sender];
    }

    function setBalance(uint256 value) public {
        require(msg.sender != SENTINEL_TOKEN_CONTRACTS);
        require(tokenContracts[msg.sender] != address(0));
        balances[msg.sender] = value;
    }

    function getAllowance(address _spender) public view returns (uint256) {
        require(msg.sender != SENTINEL_TOKEN_CONTRACTS);
        require(tokenContracts[msg.sender] != address(0));
        return allowances[msg.sender][_spender];
    }

    function setAllowance(address _spender, uint256 value) public {
        require(msg.sender != SENTINEL_TOKEN_CONTRACTS);
        require(tokenContracts[msg.sender] != address(0));
        allowances[msg.sender][_spender] = value;
    }
}

contract Factory {
    event HolderCreated(Holder holder, address owner);
    event TokenCreated(Token token, address minter);

    function createToken(address _minter) public returns (Token) {
        Token token = new Token(_minter, address(this));
        emit TokenCreated(token, _minter);
        return token;
    }

    function createHolder(address payable _owner) public returns (Holder) {
        bytes memory holder_init_code = hex"608060405234801561001057600080fd5b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061118c806100e06000396000f3fe608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806312065fe01461009b5780632a847f46146100c6578063310ec4a7146101175780634df4d0391461017257806383197ef0146101c95780638852671e146101e057806399d3c20414610251578063eb5a662e146102ac578063fb1669ca14610311575b005b3480156100a757600080fd5b506100b061034c565b6040518082815260200191505060405180910390f35b3480156100d257600080fd5b50610115600480360360208110156100e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610469565b005b34801561012357600080fd5b506101706004803603604081101561013a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061078e565b005b34801561017e57600080fd5b506101876108e9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101d557600080fd5b506101de6108ee565b005b3480156101ec57600080fd5b5061024f6004803603604081101561020357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109e3565b005b34801561025d57600080fd5b506102aa6004803603604081101561027457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d07565b005b3480156102b857600080fd5b506102fb600480360360208110156102cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ee7565b6040518082815260200191505060405180910390f35b34801561031d57600080fd5b5061034a6004803603602081101561033457600080fd5b8101908080359060200190929190505050611043565b005b6000600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415151561038a57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561042457600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156104c757600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561052357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561058d5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b151561059857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561063157600080fd5b600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515156107ca57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561086457600080fd5b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b600181565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561094c57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109a857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151515610a4157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a9d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015610b075750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1515610b1257600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610baa57600080fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b3073ffffffffffffffffffffffffffffffffffffffff1660ff7f010000000000000000000000000000000000000000000000000000000000000002600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff166001028460405160200180857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401838152602001828152602001945050505050604051602081830303815290604052805190602001206001900473ffffffffffffffffffffffffffffffffffffffff16141515610e6057600080fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610f2557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151515610fbf57600080fd5b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415151561107f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561111957600080fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505056fea165627a7a72305820fcc4fbded726882960a22fcbb868fbd0701133626d09eab23d55ffc330bfdda60029";
        bytes32 holder_init_code_hash = 0x17bf60e63d4e2cc380fc7b8d419f6ebf015f77b0916680115f61cc123093468f;
        // The following line can be commented out to save gas
        require(holder_init_code_hash == keccak256(holder_init_code));
        address payable h;
        assembly{
            h := create2(0, add(holder_init_code, 32), mload(holder_init_code), _owner)
        }
        require(h != address(0));
        Holder holder = Holder(h);
        holder.setOwner(_owner, holder_init_code_hash);
        emit HolderCreated(holder, _owner);
        return holder;
    }
}
