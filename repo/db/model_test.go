// +build !integration

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	ip, _ := NewIP("1.2.3.4")
	type args struct {
		ip IPAddress
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid ip",
			args: args{
				ip: ip,
			},
			want: "4.3.2.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.ip.Reverse(); got != tt.want {
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
		name    string
		args    args
		want    IPAddress
		err     error
		wantErr bool
	}{
		{
			name: "empty ip",
			args: args{
				ip: "",
			},
			want:    IPAddress{},
			wantErr: true,
		},
		{
			name: "valid ipv4",
			args: args{
				ip: "1.2.3.4",
			},
			want: IPAddress{
				IP: "1.2.3.4",
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "invalid ipv4",
			args: args{
				ip: "1.2.3.44444444",
			},
			want:    IPAddress{},
			err:     ErrInvalidIPAddress,
			wantErr: true,
		},
		{
			name: "ipv6 should be false",
			args: args{
				ip: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			want:    IPAddress{},
			err:     ErrInvalidIPAddress,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIP(tt.args.ip)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, got, tt.want)
				assert.NoError(t, err)
			}

		})
	}
}
