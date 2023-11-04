package sns_go_api

import "testing"

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
