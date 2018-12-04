pragma solidity ^0.5.0;

contract Token {
    uint256 public totalSupply;

    /* This creates an array with all balances */
    mapping (address => uint256) public balanceOf;
    mapping (address => mapping (address => uint256)) public allowance;

    /* This generates a public event on the blockchain that will notify clients */
    event Transfer(address indexed from, address indexed to, uint256 value);

    function getHolderContract(address a) internal view returns (Holder) {
        bytes32 holder_init_code_hash = "\x00";
        address factory = address(0x9494949);
        address payable holder_address = address(uint256(keccak256(abi.encodePacked(factory,a,holder_init_code_hash))));
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
    function approve(address _spender, uint256 _value) public returns (bool success) {
        Holder fromContract = getHolderContract(msg.sender);
        fromContract.setAllowance(_spender, _value);
        return true;
    }

    /* A contract attempts to get the tokens */
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
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

    address public minter;

    constructor() public {
        minter = msg.sender;
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
        tokenContracts[SENTINEL_TOKEN_CONTRACTS] = SENTINEL_TOKEN_CONTRACTS;
    }

    function setOwner(address payable _owner, bytes32 init_code_hash) public {
        if (address(uint256(keccak256(abi.encodePacked(factory,_owner,init_code_hash)))) == address(this)) {
            owner = _owner;
            factory = address(0);
        }
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

contract HolderFactory {
    function createHolder(address payable _owner) public returns (Holder) {
        bytes memory holder_init_code = "\x00";
        uint code_length = holder_init_code.length;
        bytes32 holder_init_code_hash = "\x00";
        address payable h;
        assembly{
            h := create2(0,_owner,holder_init_code,code_length)
        }
        require(h != address(0));
        Holder holder = Holder(h);
        holder.setOwner(_owner, holder_init_code_hash);
        return holder;
    }
}
