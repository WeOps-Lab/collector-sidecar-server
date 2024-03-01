package ping

import "testing"

func TestIsLocalIp(t *testing.T) {
	tests := []struct {
		name   string
		host   string
		expect bool
	}{
		{
			name:   "Test with non-local IP",
			host:   "192.168.1.1:9090",
			expect: false,
		},
		{
			name:   "Test with local IP",
			host:   "127.0.0.1:8080",
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLocalIP(tt.host); got != tt.expect {
				t.Errorf("IsLocalIp() = %v, want %v", got, tt.expect)
			}
		})
	}
}
