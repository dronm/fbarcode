package fbarcode

import (
	"strconv"
	"testing"
)

func TestEan13(t *testing.T) {
	tests := []struct {
		Barcode     string
		ExpectedSeq string
	}{
		{"0072104006868", "0000::00802:2088200:8208:202882020"},
		{"0000000038744", "0000::000::000::000::0:20800082:20"},
		{"0001010604219", "0000::0802288022808:2000:288200:20"},
		{"0072104005489", "0000::00802:2088200:8280:208028220"},
		{"0001010604394", "0000::0802288022808:2000:282:08020"},
		{"0072104004888", "0000::00802:2088200:8200:28:00:020"},
	}
	for _, tt := range tests {
		got := Ean13(tt.Barcode)
		if got != tt.ExpectedSeq {
			t.Fatalf("Barcode %s, expected: %s, got: %s", tt.Barcode, tt.ExpectedSeq, got)
		}
		n := len(tt.Barcode)
		gotChecksum, err := Checksum(tt.Barcode[:n-1]+" ", len(tt.Barcode))
		if err != nil {
			t.Fatalf("Checksum() failed: %v", err)
		}

		expChecksum, err := strconv.Atoi(tt.Barcode[n-1:])
		if err != nil {
			t.Fatalf("convert.Atoi() failed: %v", err)
		}
		if gotChecksum != expChecksum {
			t.Fatalf("Expected checksum: %d, got: %d", expChecksum, gotChecksum)
		}
	}
}
