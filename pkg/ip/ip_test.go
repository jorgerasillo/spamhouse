package ip

import (
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid ip",
			args: args{
				ip: "1.2.3.4",
			},
			want: "4.3.2.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.ip); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidIPV4(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty ip",
			args: args{
				ip: "",
			},
			want: false,
		},
		{
			name: "valid ipv4",
			args: args{
				ip: "1.2.3.4",
			},
			want: true,
		},
		{
			name: "invalid ipv4",
			args: args{
				ip: "1.2.3.44444444",
			},
			want: false,
		},
		{
			name: "ipv6 should be false",
			args: args{
				ip: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidIPV4(tt.args.ip); got != tt.want {
				t.Errorf("IsValidIPV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
