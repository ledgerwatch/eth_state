package token

import (
    "fmt"
    "context"
    "math/big"
    "testing"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
)

var (
    one_Ether = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)

    // Generate a new random account and a funded simulator
    deployer_key, _ = crypto.GenerateKey()
    deployer = bind.NewKeyedTransactor(deployer_key)

    alice_key, _ = crypto.GenerateKey()
    alice = bind.NewKeyedTransactor(alice_key)

    bob_key, _ = crypto.GenerateKey()
    bob = bind.NewKeyedTransactor(bob_key)

    charlie_key, _ = crypto.GenerateKey()
    charlie = bind.NewKeyedTransactor(charlie_key)
)

func newTestBackend() *backends.SimulatedBackend {
    return backends.NewSimulatedBackend(core.GenesisAlloc{
        deployer.From: {Balance: big.NewInt(10000000000)},
        alice.From: {Balance: big.NewInt(0).Mul(one_Ether, big.NewInt(100))}, // Alice has 100 ETH
        bob.From: {Balance: big.NewInt(0).Mul(one_Ether, big.NewInt(100))}, // Bob has 100 ETH
        charlie.From: {Balance: big.NewInt(0).Mul(one_Ether, big.NewInt(100))}, // Charlie has 100 ETH
        }, 0)
}

func mustDeployFactory(t *testing.T) (*backends.SimulatedBackend, *Factory) {
    b := newTestBackend()
    address, _, factory, err := DeployFactory(deployer, b)
    if err != nil {
        t.Errorf("Could not deploy ChannelBook contract")
    }
    fmt.Printf("Factory addrss: %x\n", address)
    b.Commit()
    return b, factory
}

func commitAndGetReceipt(t *testing.T, b *backends.SimulatedBackend, tx *types.Transaction, err error) (*types.Receipt) {
    if err != nil {
        t.Errorf("Error calling bindings: %v", err)
        return nil
    }
    b.Commit()
    ctx := context.Background()
    r, err := b.TransactionReceipt(ctx, tx.Hash())
    if err != nil || r == nil {
        t.Errorf("Could not get tx receipt for %x: %v", tx.Hash(), err)
    }
    return r
}

func TestCreateTokenContract(t *testing.T) {
    b, factory := mustDeployFactory(t)
    tx, err := factory.CreateToken(&bind.TransactOpts{From: alice.From, Signer: alice.Signer, Value: big.NewInt(0)}, alice.From)
    r := commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Alice could not create token contract, tx failed")
    }
}

func TestCreateHolderContract(t *testing.T) {
    b, factory := mustDeployFactory(t)
    tx, err := factory.CreateHolder(&bind.TransactOpts{From: bob.From, Signer: bob.Signer, Value: big.NewInt(0)}, bob.From)
    r := commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Bob could not create holder contract, tx failed")
    }
}

func TestMint(t *testing.T) {
    b, factory := mustDeployFactory(t)
    tx, err := factory.CreateToken(&bind.TransactOpts{From: alice.From, Signer: alice.Signer, Value: big.NewInt(0)}, alice.From)
     r := commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Alice could not create token contract, tx failed")
    }
    var tokenAddress common.Address
    {
        it, err := factory.FilterTokenCreated(&bind.FilterOpts{})
        if err != nil {
            t.Errorf("Could not filter TokenCreated events: %v", err)
        }
        if it.Next() {
            tokenAddress = it.Event.Token
        } else {
            t.Errorf("Expected TokenCreated event, got none")
        }
        it.Close()
    }
    fmt.Printf("Token address: %x\n", tokenAddress)
    token, err := NewToken(tokenAddress, b)
    if err != nil {
        t.Errorf("Could not bind Token to %x", tokenAddress)
    }
    tx, err = factory.CreateHolder(&bind.TransactOpts{From: bob.From, Signer: bob.Signer, Value: big.NewInt(0)}, bob.From)
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Bob could not create holder contract, tx failed")
    }
    var holderAddress common.Address
    {
        it, err := factory.FilterHolderCreated(&bind.FilterOpts{})
        if err != nil {
            t.Errorf("Could not filter HolderCreated events: %v", err)
        }
        if it.Next() {
            holderAddress = it.Event.Holder
        } else {
            t.Errorf("Expected HolderCreated event, got none")
        }
        it.Close()
    }
    fmt.Printf("Holder address: %x\n", holderAddress)
    // Bob allows token contract
    holder, err := NewHolder(holderAddress, b)
    if err != nil {
        t.Errorf("Could not bind Holder to %x", holderAddress)
    }
    tx, err = holder.AddTokenContract(&bind.TransactOpts{From: bob.From, Signer: bob.Signer, Value: big.NewInt(0)}, tokenAddress)
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Bob could allow token contract in his holder, tx failed")
    }
    // Alice mints 10 tokens for bob
    tx, err = token.Mint(&bind.TransactOpts{From: alice.From, Signer: alice.Signer, Value: big.NewInt(0)}, bob.From, big.NewInt(10))
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Alice could not mint tokens for Bob, tx failed")
    }
    // Check Bob's balance
    balance, err := token.BalanceOf(&bind.CallOpts{}, bob.From)
    if err != nil {
        t.Errorf("Could not check Bob's balance: %v", err)
    }
    if balance.Cmp(big.NewInt(10)) != 0 {
        t.Errorf("Bob's balance is %d, expected 10", balance)
    }
}

func TestMintAndTransfer(t *testing.T) {
    b, factory := mustDeployFactory(t)
    tx, err := factory.CreateToken(&bind.TransactOpts{From: alice.From, Signer: alice.Signer, Value: big.NewInt(0)}, alice.From)
     r := commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Alice could not create token contract, tx failed")
    }
    var tokenAddress common.Address
    {
        it, err := factory.FilterTokenCreated(&bind.FilterOpts{})
        if err != nil {
            t.Errorf("Could not filter TokenCreated events: %v", err)
        }
        if it.Next() {
            tokenAddress = it.Event.Token
        } else {
            t.Errorf("Expected TokenCreated event, got none")
        }
        it.Close()
    }
    fmt.Printf("Token address: %x\n", tokenAddress)
    token, err := NewToken(tokenAddress, b)
    if err != nil {
        t.Errorf("Could not bind Token to %x", tokenAddress)
    }
    tx, err = factory.CreateHolder(&bind.TransactOpts{From: bob.From, Signer: bob.Signer, Value: big.NewInt(0)}, bob.From)
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Bob could not create holder contract, tx failed")
    }
    var holderAddress common.Address
    {
        it, err := factory.FilterHolderCreated(&bind.FilterOpts{})
        if err != nil {
            t.Errorf("Could not filter HolderCreated events: %v", err)
        }
        if it.Next() {
            holderAddress = it.Event.Holder
        } else {
            t.Errorf("Expected HolderCreated event, got none")
        }
        it.Close()
    }
    fmt.Printf("Bob's holder address: %x\n", holderAddress)
    // Bob allows token contract
    holder, err := NewHolder(holderAddress, b)
    if err != nil {
        t.Errorf("Could not bind Holder to %x", holderAddress)
    }
    tx, err = holder.AddTokenContract(&bind.TransactOpts{From: bob.From, Signer: bob.Signer, Value: big.NewInt(0)}, tokenAddress)
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Bob could allow token contract in his holder, tx failed")
    }
    // Alice mints 10 tokens for bob
    tx, err = token.Mint(&bind.TransactOpts{From: alice.From, Signer: alice.Signer, Value: big.NewInt(0)}, bob.From, big.NewInt(10))
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Alice could not mint tokens for Bob, tx failed")
    }

    //----------------------
    // Now Charlie creates his holder
    tx, err = factory.CreateHolder(&bind.TransactOpts{From: charlie.From, Signer: charlie.Signer, Value: big.NewInt(0)}, charlie.From)
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Charlie could not create holder contract, tx failed")
    }
    var charlieHolderAddress common.Address
    {
        it, err := factory.FilterHolderCreated(&bind.FilterOpts{})
        if err != nil {
            t.Errorf("Could not filter HolderCreated events: %v", err)
        }
        if it.Next() && it.Next() {
            charlieHolderAddress = it.Event.Holder
        } else {
            t.Errorf("Expected HolderCreated event, got none")
        }
        it.Close()
    }
    fmt.Printf("Charlie's holder address: %x\n", charlieHolderAddress)
    // Bob allows token contract
    charlieHolder, err := NewHolder(charlieHolderAddress, b)
    if err != nil {
        t.Errorf("Could not bind Holder to %x", charlieHolderAddress)
    }
    tx, err = charlieHolder.AddTokenContract(&bind.TransactOpts{From: charlie.From, Signer: charlie.Signer, Value: big.NewInt(0)}, tokenAddress)
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Charlie could allow token contract in his holder, tx failed")
    }

    //-------------
    // Bob transfers to Charlie
    tx, err = token.Transfer(&bind.TransactOpts{From: bob.From, Signer: bob.Signer, Value: big.NewInt(0)}, charlie.From, big.NewInt(3))
    r = commitAndGetReceipt(t, b, tx, err)
    if r != nil && r.Status != types.ReceiptStatusSuccessful {
        t.Errorf("Bob could not transfer tokens to Charlie, tx failed")
    }
    // Check Charlie's balance
    balance, err := token.BalanceOf(&bind.CallOpts{}, charlie.From)
    if err != nil {
        t.Errorf("Could not check Charlie's balance: %v", err)
    }
    if balance.Cmp(big.NewInt(3)) != 0 {
        t.Errorf("Charlie's balance is %d, expected 3", balance)
    }
    // Check Bob's balance
    balance, err = token.BalanceOf(&bind.CallOpts{}, bob.From)
    if err != nil {
        t.Errorf("Could not check Bob's balance: %v", err)
    }
    if balance.Cmp(big.NewInt(7)) != 0 {
        t.Errorf("Bob's balance is %d, expected 7", balance)
    }
}
