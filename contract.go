package sns_go_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

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

func Resolve(sns, indexerHost, rpc, publicResolverAddr string) (addr string) {
	node := namehash.Namehash(sns)

	var err error
	err, addr = getTimeout(fmt.Sprintf("%s/sns/resolver/get_address?node=%s", indexerHost, node))
	if err == nil {
		return
	}

	c, err := ethclient.Dial(rpc)
	if err != nil {
		return ""
	}
	s, err := snsgen.NewSns(common.HexToAddress(publicResolverAddr), c)
	if err != nil {
		return ""
	}

	// convert 'namehash' result to [32]bytes
	b, _ := common.ParseHexOrString(node)
	nb := [32]byte{}
	copy(nb[:], b)

	commonAddr, err := s.Addr(nil, nb)
	if err != nil {
		return ""
	}

	addr = commonAddr.Hex()
	return
}

type addrOutput struct {
	Addr common.Address
}

func Resolves(sns []string, indexerHost, rpc, publicResolverAddr string) (addr []string) {
	nodes := make([]string, len(sns))
	for i, s := range sns {
		nodes[i] = namehash.Namehash(s)
	}

	var err error
	err, addr = postTimeout(fmt.Sprintf("%s/sns/resolver/get_addresses", indexerHost), nodes)
	if err == nil {
		return
	}

	m, err := multicall.Dial(context.Background(), rpc)
	if err != nil {
		return
	}

	c, err := multicall.NewContract(publicResolverABI, publicResolverAddr)
	if err != nil {
		return
	}

	var calls []*multicall.Call
	for _, n := range nodes {
		// convert 'namehash' result to [32]bytes
		b, _ := common.ParseHexOrString(n)
		nb := [32]byte{}
		copy(nb[:], b)

		calls = append(calls, c.NewCall(new(addrOutput), "addr", nb))
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

func Name(addr, indexerHost, rpc, publicResolverAddr string) (sns string) {
	node := namehash.Namehash(fmt.Sprintf("%s.addr.reverse", strings.ToLower(addr[2:])))

	var err error
	err, sns = getTimeout(fmt.Sprintf("%s/sns/resolver/get_name?node=%s", indexerHost, node))
	if err == nil {
		return
	}

	c, err := ethclient.Dial(rpc)
	if err != nil {
		return ""
	}
	s, err := snsgen.NewSns(common.HexToAddress(publicResolverAddr), c)
	if err != nil {
		return ""
	}

	// append `.addr.reverse` for `addr`
	// convert 'namehash' result to [32]bytes
	b, _ := common.ParseHexOrString(node)
	nb := [32]byte{}
	copy(nb[:], b)

	sns, err = s.Name(nil, nb)
	if err != nil {
		return ""
	}

	return sns
}

type nameOutput struct {
	Name string
}

func Names(addr []string, indexerHost, rpc, publicResolverAddr string) (sns []string) {
	nodes := make([]string, len(addr))
	for i, a := range addr {
		nodes[i] = namehash.Namehash(fmt.Sprintf("%s.addr.reverse", strings.ToLower(a[2:])))
	}

	var err error
	err, sns = postTimeout(fmt.Sprintf("%s/sns/resolver/get_names", indexerHost), nodes)
	if err == nil {
		return
	}

	m, err := multicall.Dial(context.Background(), rpc)
	if err != nil {
		return
	}

	c, err := multicall.NewContract(publicResolverABI, publicResolverAddr)
	if err != nil {
		return
	}

	var calls []*multicall.Call
	for _, n := range nodes {
		// convert 'namehash' result to [32]bytes
		b, _ := common.ParseHexOrString(n)
		nb := [32]byte{}
		copy(nb[:], b)

		calls = append(calls, c.NewCall(new(nameOutput), "name", nb))
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

func getTimeout(url string) (error, string) {
	client := &http.Client{Timeout: 1 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request indexer failed: %d", resp.StatusCode), ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}

	return nil, string(body)
}

func postTimeout(url string, param []string) (error, []string) {
	client := &http.Client{Timeout: 1 * time.Second}

	body, err := json.Marshal(param)
	if err != nil {
		return err, nil
	}

	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request indexer failed: %d", resp.StatusCode), nil
	}

	var result []string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err, nil
	}

	return nil, result
}
