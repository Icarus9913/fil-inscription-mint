package sdk

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUTF8ToHex(t *testing.T) {
	utf8Str := `X;data:,{"p":"fil-20","op":"mint","tick":"fils","amt":"1000"}`
	hex := UTF8ToHex(utf8Str)
	fmt.Println(hex)
}

func TestHexToUTF8(t *testing.T) {
	hexStr := "0x583b646174613a2c7b2270223a2266696c2d3230222c226f70223a226d696e74222c227469636b223a2266696c73222c22616d74223a2231303030227d"
	utf8, err := HexToUTF8(hexStr)
	require.NoError(t, err)
	fmt.Println(utf8)
}

func TestStringToBase64(t *testing.T) {
	str := `X;data:,{"p":"fil-20","op":"mint","tick":"fils","amt":"1000"}`
	stringToBase64 := StringToBase64(str)
	fmt.Println(stringToBase64)
}

// XPdata:,{"p":"fil-20","op":"deploy","tick":"fils","lim":"1000","max":"1500000000"}
func TestBase64ToString(t *testing.T) {
	base64Str := "WFBkYXRhOix7InAiOiJmaWwtMjAiLCJvcCI6ImRlcGxveSIsInRpY2siOiJmaWxzIiwibGltIjoiMTAwMCIsIm1heCI6IjE1MDAwMDAwMDAifQ=="
	base64ToString, err := Base64ToString(base64Str)
	require.NoError(t, err)
	fmt.Println(base64ToString)
}
