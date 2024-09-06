package catch

import (
	"testing"
)

func TestDmp(t *testing.T) {
	tool := &CatchTool{}
	tool.Dmp()
	panic("s")
}
