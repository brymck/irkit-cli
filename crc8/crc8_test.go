package crc8

import (
	"testing"
)

func TestCrc8FromByte(t *testing.T) {
	type args struct {
		data byte
		crc  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "Test security", args: args{data: 0x08, crc: 0x00}, want: 0xb9},
		{name: "Test Wi-Fi is set", args: args{data: 0x01, crc: 0x7c}, want: 0xb4},
		{name: "Test Wi-Fi is valid", args: args{data: 0x00, crc: 0xb4}, want: 0x7b},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crc8FromByte(tt.args.data, tt.args.crc); got != tt.want {
				t.Errorf("Crc8FromByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrc8FromString(t *testing.T) {
	type args struct {
		data string
		size int
		crc  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "Test SSID", args: args{data: "foo", size: 33, crc: 0xb9}, want: 0xc8},
		{name: "Test password", args: args{data: "bar", size: 64, crc: 0xc8}, want: 0x7c},
		{name: "Test device key", args: args{data: "", size: 33, crc: 0x7b}, want: 0x48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crc8FromString(tt.args.data, tt.args.size, tt.args.crc); got != tt.want {
				t.Errorf("Crc8FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
