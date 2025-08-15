package qrc_test

import (
	"fmt"
	"testing"

	"github.com/boombuler/barcode/qr"
)

func TestEncode(t *testing.T) {
	input := "Hello, world!"
	code, err := qr.Encode(input, qr.M, qr.Auto)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("meta data: \n", code.Metadata())
}
