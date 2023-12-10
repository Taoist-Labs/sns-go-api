
install `solc` and `abigen` first:

```
$ npm install -g solc
$ go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

and then execute code generation:

```
$ cd sns
$ abigen --abi public_resolver.json --pkg sns --type publicResolver --out public_resolver.go
$ abigen --abi base_registrar.json --pkg sns --type baseRegistrar --out base_registrar.go
```
