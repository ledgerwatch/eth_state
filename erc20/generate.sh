solc --bin TokenContract.sol -o bin --overwrite
$GOPATH/src/github.com/ethereum/go-ethereum/build/bin/abigen --sol TokenContract.sol --pkg token --out token.go
