package gateway

import "testing"

func TestMatchDomain(t *testing.T) {
	tests := []struct {
		domain  string
		pattern string
		want    bool
	}{
		{"example.com", "**.com", true},
		{"sub.example.com", "**.com", true},
		{"sub.sub.example.com", "**.example.com", true},
		{"example.com", "**.example.com", true},
		{"example.net", "**.com", false},
		{"sub.example.com", "*.example.com", true},
		{"sub.sub.example.com", "*.example.com", false},
        {"a.example.com", "?.example.com", true},
        {"ab.example.com", "?.example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			if got := matchDomain(tt.pattern, tt.domain); got != tt.want {
				t.Errorf("MatchDomain(%q, %q) = %v, want %v", tt.domain, tt.pattern, got, tt.want)
			}
		})
	}
}
