package qrcode

import (
	"encoding/base64"
	"image/color"

	"github.com/skip2/go-qrcode"
)

type QRCodeTool struct{}

// GenerateQrcodeToBase64
//
// Generate a QR code and convert it to Base 64
func (s *QRCodeTool) GenerateQrcodeToBase64(content string, size int) (*string, error) {
	var png *[]byte
	png, err := s.GenerateQrcodeToBytes(content, size)
	base64 := base64.StdEncoding.EncodeToString(*png)
	return &base64, err
}

// GenerateQrcodeToBytes
//
// Encode a QR Code and return a raw PNG image.
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently returned. Negative values for size cause a
// variable sized image to be returned: See the documentation for Image().
// To serve over HTTP, remember to send a Content-Type: image/png header.
func (s *QRCodeTool) GenerateQrcodeToBytes(content string, size int) (*[]byte, error) {
	var png []byte
	png, err := qrcode.Encode(content, qrcode.Medium, size)
	return &png, err
}

// GenerateQrcodeToFile
//
// WriteFile encodes, then writes a QR Code to the given filename in PNG format.
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func (s *QRCodeTool) GenerateQrcodeToFile(content, filename string, size int) error {
	err := qrcode.WriteFile(content, qrcode.Medium, size, filename)
	return err
}

// GenerateColorQrcodeToFile
//
// WriteColorFile encodes, then writes a QR Code to the given filename in PNG format.
// With WriteColorFile you can also specify the colors you want to use.
// size is both the image width and height in pixels. If size is too small then
// a larger image is silently written. Negative values for size cause a variable
// sized image to be written: See the documentation for Image().
func (s *QRCodeTool) GenerateColorQrcodeToFile(content, filename string, background,
	foreground color.Color, size int) error {
	err := qrcode.WriteColorFile(content, qrcode.Medium, size, background, foreground, filename)
	return err
}
