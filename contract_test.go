package sns_go_api

import "testing"

func TestResolve(t *testing.T) {
	tests := []struct {
		sns  string
		want string
	}{
		{"0x8C913aEc7443FE2018639133398955e0E17FB0C1", "baiyu.seedao"},
	}
	for _, tt := range tests {
		if got := Name(tt.sns); got != tt.want {
			t.Errorf("Name(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
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
	for _, tt := range tests {
		got := Resolves(tt.sns)

		if len(got) != len(tt.want) {
			t.Errorf("Resolves(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
		for i := 0; i < len(tt.want); i++ {
			if got[i] != tt.want[i] {
				t.Errorf("Resolves(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
			}
		}
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		sns  string
		want string
	}{
		{"baiyu.seedao", "0x8C913aEc7443FE2018639133398955e0E17FB0C1"},
	}
	for _, tt := range tests {
		if got := Resolve(tt.sns); got != tt.want {
			t.Errorf("Name(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
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
	for _, tt := range tests {
		got := Names(tt.addr)

		if len(got) != len(tt.want) {
			t.Errorf("Names(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
		}
		for i := 0; i < len(tt.want); i++ {
			if got[i] != tt.want[i] {
				t.Errorf("Names(%s)'s result: %v, want: %v", tt.addr, got, tt.want)
			}
		}
	}
}
