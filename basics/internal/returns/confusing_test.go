package returns

import (
	"testing"
)

func TestXxx(t *testing.T) {
	a := Animal{
		Name: "old",
	}
	BadRename(a, "new")
	if a.Name != "old" {
		t.Error("")
	}

	GoodRename(&a, "new")
	if a.Name != "new" {
		t.Error("")
	}
}