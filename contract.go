package sns_go_api

import (
	"context"
	"fmt"
	"strings"

	snsgen "github.com/Taoist-Labs/sns-go-api/sns"
	namehash "github.com/Taoist-Labs/sns-go-namehash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/forta-network/go-multicall"
)

const publicResolverABI = `[
  {
    "inputs": [
      {
        "internalType": "bytes32",
        "name": "node",
        "type": "bytes32"
      }
    ],
    "name": "addr",
    "outputs": [
      {
        "internalType": "address payable",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes32",
        "name": "node",
        "type": "bytes32"
      }
    ],
    "name": "name",
    "outputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]`

func Resolve(sns, publicResolverAddr, rpc string) (addr string) {
	c, err := ethclient.Dial(rpc)
	if err != nil {
		return ""
	}
	s, err := snsgen.NewSns(common.HexToAddress(publicResolverAddr), c)
	if err != nil {
		return ""
	}

	// convert 'namehash' result to [32]bytes
	bytes, _ := common.ParseHexOrString(namehash.Namehash(sns))
	node := [32]byte{}
	copy(node[:], bytes)

	commonAddr, err := s.Addr(nil, node)
	if err != nil {
		return ""
	}

	addr = commonAddr.Hex()
	return
}

type addrOutput struct {
	Addr common.Address
}

func Resolves(sns []string, publicResolverAddr, rpc string) (addr []string) {
	m, err := multicall.Dial(context.Background(), rpc)
	if err != nil {
		return
	}

	c, err := multicall.NewContract(publicResolverABI, publicResolverAddr)
	if err != nil {
		return
	}

	var calls []*multicall.Call
	for _, s := range sns {
		// convert 'namehash' result to [32]bytes
		bytes, _ := common.ParseHexOrString(namehash.Namehash(s))
		node := [32]byte{}
		copy(node[:], bytes)

		calls = append(calls, c.NewCall(new(addrOutput), "addr", node))
	}

	result, err := m.Call(nil, calls...)
	if err != nil {
		return
	}

	for _, r := range result {
		addr = append(addr, r.Outputs.(*addrOutput).Addr.Hex())
	}
	return
}

func Name(addr, publicResolverAddr, rpc string) (sns string) {
	c, err := ethclient.Dial(rpc)
	if err != nil {
		return ""
	}
	s, err := snsgen.NewSns(common.HexToAddress(publicResolverAddr), c)
	if err != nil {
		return ""
	}

	// append `.addr.reverse` for `addr`
	addr = fmt.Sprintf("%s.addr.reverse", strings.ToLower(addr[2:]))
	// convert 'namehash' result to [32]bytes
	bytes, _ := common.ParseHexOrString(namehash.Namehash(addr))
	node := [32]byte{}
	copy(node[:], bytes)

	sns, err = s.Name(nil, node)
	if err != nil {
		return ""
	}

	return sns
}

type nameOutput struct {
	Name string
}

func Names(addr []string, publicResolverAddr, rpc string) (sns []string) {
	m, err := multicall.Dial(context.Background(), rpc)
	if err != nil {
		return
	}

	c, err := multicall.NewContract(publicResolverABI, publicResolverAddr)
	if err != nil {
		return
	}

	var calls []*multicall.Call
	for _, a := range addr {
		// append `.addr.reverse` for `addr`
		reverseAddr := fmt.Sprintf("%s.addr.reverse", strings.ToLower(a[2:]))
		// convert 'namehash' result to [32]bytes
		bytes, _ := common.ParseHexOrString(namehash.Namehash(reverseAddr))
		node := [32]byte{}
		copy(node[:], bytes)

		calls = append(calls, c.NewCall(new(nameOutput), "name", node))
	}

	result, err := m.Call(nil, calls...)
	if err != nil {
		return
	}

	for _, r := range result {
		sns = append(sns, r.Outputs.(*nameOutput).Name)
	}
	return
}
