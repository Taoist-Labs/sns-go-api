
install `solc` and `abigen` first:

```
$ npm install -g solc
$ go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

and then execute code generation:

```
$ cd sns
$ abigen --abi public_resolver.json --pkg sns --type sns --out public_resolver.go
```
