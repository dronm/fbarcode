package fbarcode

import (
	"fmt"
	"strconv"
	"strings"
)

func CodeChar(a string) string {
	var s strings.Builder
	if a == "211412" {
		s.WriteString("a")
	} else if a == "211214" {
		s.WriteString("B")
	} else if a == "211232" {
		s.WriteString("C")
	} else if a == "2331112" {
		s.WriteString("@")
	} else {
		for i := 0; i < len(a)/2; i++ {
			v := a[2*i : 2*i+2]

			if v == "11" {
				s.WriteString("0")
			} else if v == "21" {
				s.WriteString("1")
			} else if v == "31" {
				s.WriteString("2")
			} else if v == "41" {
				s.WriteString("3")
			} else if v == "12" {
				s.WriteString("4")
			} else if v == "22" {
				s.WriteString("5")
			} else if v == "32" {
				s.WriteString("6")
			} else if v == "42" {
				s.WriteString("7")
			} else if v == "13" {
				s.WriteString("8")
			} else if v == "23" {
				s.WriteString("9")
			} else if v == "33" {
				s.WriteString(":")
			} else if v == "43" {
				s.WriteString(";")
			} else if v == "14" {
				s.WriteString("<")
			} else if v == "24" {
				s.WriteString("=")
			} else if v == "34" {
				s.WriteString(">")
			} else if v == "44" {
				s.WriteString("?")
			}
		}
	}
	return s.String()
}

func Code2of5Ch(a string) string {
	res := ""
	if a == "0" {
		res = "11331"
	} else if a == "1" {
		res = "31113"
	} else if a == "2" {
		res = "13113"
	} else if a == "3" {
		res = "33111"
	} else if a == "4" {
		res = "11313"
	} else if a == "5" {
		res = "31311"
	} else if a == "6" {
		res = "13311"
	} else if a == "7" {
		res = "11133"
	} else if a == "8" {
		res = "31131"
	} else if a == "9" {
		res = "13131"
	}

	return res
}

func Interleaved2of5Pair(pair string) string {
	if len(pair) < 2 {
		return ""
	}
	s1 := Code2of5Ch(pair[0:1])
	s2 := Code2of5Ch(pair[1:2])
	var s strings.Builder
	for i := 0; i < len(s1); i++ {
		s.WriteString(s1[i:i+1] + s2[i:i+1])
	}

	return s.String()
}

// Checksum returns barcode checksum. Ean should conain barcode for computing checksum.
func Checksum(ean string, size int) (int, error) {
	if len(ean) != size {
		return -1, fmt.Errorf("incorrect ean %v to compute a checksum", ean)
	}

	code := ean[:size-1]
	multiplyWhenEven := size%2 == 0
	sum := 0

	for i, v := range code {
		value, err := strconv.Atoi(string(v))

		if err != nil {
			return -1, fmt.Errorf("contains non-digit: %q", v)
		}

		if (i%2 == 0) == multiplyWhenEven {
			sum += 3 * value
		} else {
			sum += value
		}
	}

	return (10 - sum%10) % 10, nil
}

func Ean13(barcode string) string {
	if len(barcode) != 13 {
		barcode = "0" + barcode
	}
	var res strings.Builder
	res.WriteString(CodeChar("1111")) //start simbol
	for i := 0; i < len(barcode)/2; i++ {
		sm := barcode[i*2 : i*2+2]
		interl := Interleaved2of5Pair(sm)
		code := CodeChar(interl)
		// fmt.Println("Take 2 simb from ", i*2, sm, "interleave", interl, "code", code)
		res.WriteString(code)
	}
	res.WriteString(CodeChar("3111")) //stop simbol
	return res.String()
}
