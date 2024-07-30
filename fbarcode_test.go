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
		{"0072104006868", "0000::0002:88200:00:2800::0822808228020"},
		{"0000000038744", "0000::000::000::000::088220200:800:0:20"},
		{"0001010604219", "0000::000::0208822088202:800820:2808220"},
		{"0072104005489", "0000::0002:88200:00:2800::020:08280:020"},
		{"0001010604394", "0000::000::0208822088202:80882020282820"},
		{"0072104004888", "0000::0002:88200:00:2800::080282:00:020"},
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
