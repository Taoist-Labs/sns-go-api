package sns_go_api

import (
	"testing"
)

func TestIsSafe(t *testing.T) {
	tests := []struct {
		sns  string
		want bool
	}{
		{"baiyu.seedao", true},
		{"xxxx.seedao", false},
	}
	for _, tt := range tests {
		if got := IsSafe(tt.sns); got != tt.want {
			t.Errorf("IsSafe(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
}

func TestSafe(t *testing.T) {
	tests := []struct {
		sns  []string
		want []string
	}{
		{[]string{"baiyu.seedao", "tanghan.seedao"}, []string{"baiyu.seedao", "tanghan.seedao"}},
		{[]string{"xxxx.seedao", "vitalik.seedao"}, []string{"", ""}},
		{[]string{"baiyu.seedao", "vitalik.seedao"}, []string{"baiyu.seedao", ""}},
	}
	for _, tt := range tests {
		got := Safe(tt.sns)

		if len(got) != len(tt.want) {
			t.Errorf("IsSafe(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
		for i := 0; i < len(tt.want); i++ {
			if got[i] != tt.want[i] {
				t.Errorf("IsSafe(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
			}
		}
	}
}

func TestIsAvailable(t *testing.T) {
	tests := []struct {
		sns  string
		want bool
	}{
		{"baiyu.seedao", false},
		{"vitalik.seedao", false},
		{"abcd.seedao", true},
	}
	for _, tt := range tests {
		if got := IsAvailable(tt.sns); got != tt.want {
			t.Errorf("IsAvailable(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
	}
}

func TestAvailable(t *testing.T) {
	tests := []struct {
		sns  []string
		want []string
	}{
		{[]string{"baiyu.seedao", "vitalik.seedao"}, []string{"", ""}},
		{[]string{"vitalik.seedao", "abcd.seedao"}, []string{"", "abcd.seedao"}},
		{[]string{"baiyu.seedao", "abcd.seedao"}, []string{"", "abcd.seedao"}},
	}
	for _, tt := range tests {
		got := Available(tt.sns)

		if len(got) != len(tt.want) {
			t.Errorf("IsSafe(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
		}
		for i := 0; i < len(tt.want); i++ {
			if got[i] != tt.want[i] {
				t.Errorf("IsSafe(%s)'s result: %v, want: %v", tt.sns, got, tt.want)
			}
		}
	}
}
