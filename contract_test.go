package sns_go_api

import "testing"

const (
	rpc                = "https://eth-sepolia.g.alchemy.com/v2/H43zK7UnIN2v7u2ZoTbizIPnXkylKIZl"
	indexerHost        = "https://test-spp-indexer.seedao.tech"
	publicResolverAddr = "0x4ffCfd37C362B415E4c4A607815f5dB6A297Ed8A"
)

func TestName(t *testing.T) {
	tests := []struct {
		sns  string
		want string
	}{
		{"0x8C913aEc7443FE2018639133398955e0E17FB0C1", "baiyu.seedao"},
	}

	// Query Indexer Success
	for _, tt := range tests {
		if got := Name(tt.sns, indexerHost, "", ""); got != tt.want {
			t.Errorf("Query Indexer Success: Name(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
	// Query Contract Success
	for _, tt := range tests {
		if got := Name(tt.sns, "", rpc, publicResolverAddr); got != tt.want {
			t.Errorf("Query Contract Success: Name(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
}

func TestResolves(t *testing.T) {
	tests := []struct {
		sns  []string
		want []string
	}{
		{[]string{"baiyu.seedao", "tanghan.seedao"}, []string{"0x8C913aEc7443FE2018639133398955e0E17FB0C1", "0x78f625A65Fc316D32d98d249b698fb509A6d98f2"}},
	}

	// Query Indexer Success
	for _, tt := range tests {
		got := Resolves(tt.sns, indexerHost, "", "")

		if len(got) != len(tt.want) {
			t.Errorf("Query Indexer Success: Resolves(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
		for i := 0; i < len(tt.want); i++ {
			if got[i] != tt.want[i] {
				t.Errorf("Query Indexer Success: Resolves(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
			}
		}
	}
	// Query Contract Success
	for _, tt := range tests {
		got := Resolves(tt.sns, "", rpc, publicResolverAddr)

		if len(got) != len(tt.want) {
			t.Errorf("Query Contract Success: Resolves(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
		for i := 0; i < len(tt.want); i++ {
			if got[i] != tt.want[i] {
				t.Errorf("Query Contract Success: Resolves(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
			}
		}
	}
}

func TestResolve(t *testing.T) {
	tests := []struct {
		sns  string
		want string
	}{
		{"baiyu.seedao", "0x8C913aEc7443FE2018639133398955e0E17FB0C1"},
	}

	// Query Indexer Success
	for _, tt := range tests {
		if got := Resolve(tt.sns, indexerHost, "", ""); got != tt.want {
			t.Errorf("Query Indexer Success: : Name(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
	// Query Contract Success
	for _, tt := range tests {
		if got := Resolve(tt.sns, "", rpc, publicResolverAddr); got != tt.want {
			t.Errorf("Query Contract Success: Name(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
}

func TestNames(t *testing.T) {
	tests := []struct {
		addr []string
		want []string
	}{
		{[]string{"0x8C913aEc7443FE2018639133398955e0E17FB0C1", "0x78f625A65Fc316D32d98d249b698fb509A6d98f2"}, []string{"baiyu.seedao", "tanghan.seedao"}},
	}

	// Query Indexer Success
	for _, tt := range tests {
		got := Names(tt.addr, indexerHost, "", "")

		if len(got) != len(tt.want) {
			t.Errorf("Query Indexer Success: Names(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
		}
		for i := 0; i < len(tt.want); i++ {
			if got[i] != tt.want[i] {
				t.Errorf("Query Indexer Success: Names(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
			}
		}
	}
	// Query Contract Success
	for _, tt := range tests {
		got := Names(tt.addr, "", rpc, publicResolverAddr)

		if len(got) != len(tt.want) {
			t.Errorf("Query Contract Success: Names(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
		}
		for i := 0; i < len(tt.want); i++ {
			if got[i] != tt.want[i] {
				t.Errorf("Query Contract Success: Names(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
			}
		}
	}
}
