package token

import (
    "context"
    "math/big"
    "testing"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
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
)

func newTestBackend() *backends.SimulatedBackend {
    return backends.NewSimulatedBackend(core.GenesisAlloc{
        deployer.From: {Balance: big.NewInt(10000000000)},
        alice.From: {Balance: big.NewInt(0).Mul(one_Ether, big.NewInt(100))}, // Alice has 100 ETH
        bob.From: {Balance: big.NewInt(0).Mul(one_Ether, big.NewInt(100))}, // Bob has 100 ETH
        }, 0)
}

func mustDeployFactory(t *testing.T) (*backends.SimulatedBackend, *Factory) {
    b := newTestBackend()
    _, _, factory, err := DeployFactory(deployer, b)
    if err != nil {
        t.Errorf("Could not deploy ChannelBook contract")
    }
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
