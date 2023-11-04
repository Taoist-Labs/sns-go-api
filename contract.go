package sns_go_api

import (
	"fmt"
	"strings"

	snsgen "github.com/Taoist-Labs/sns-go-api/sns"
	namehash "github.com/Taoist-Labs/sns-go-namehash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	RPC                = "https://eth-sepolia.g.alchemy.com/v2/H43zK7UnIN2v7u2ZoTbizIPnXkylKIZl"
	publicResolverAddr = "0x4ffCfd37C362B415E4c4A607815f5dB6A297Ed8A"
)

func Resolve(sns string) (addr string) {
	return ResolveWithRPC(sns, RPC)
}

func ResolveWithRPC(sns, rpc string) (addr string) {
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

func Name(addr string) (sns string) {
	return NameWithRPC(addr, RPC)
}

func NameWithRPC(addr, rpc string) (sns string) {
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
