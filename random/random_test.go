package random

import (
	"testing"
)

func TestGetRandomChars(t *testing.T) {
	tool := &RandomTool{}
	chars := tool.GetRandomChars(10)
	println(chars)
}
