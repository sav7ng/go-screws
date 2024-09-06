package qrcode

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/skip2/go-qrcode"
)

func TestQrCode(t *testing.T) {
	var png []byte
	png, err := qrcode.Encode("https://saving.com", qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}
	toString := base64.StdEncoding.EncodeToString(png)
	fmt.Printf(toString)
}

func TestQRCodeTool_GenerateQrcodeToBase64(t *testing.T) {
	QRCodeTool := &QRCodeTool{}
	toBase64, _ := QRCodeTool.GenerateQrcodeToBase64("https://saving.com", 256)
	fmt.Printf(*toBase64)
}
