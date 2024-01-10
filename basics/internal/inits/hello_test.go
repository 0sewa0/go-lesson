package inits

import (
	"testing"
)

func TestInit(t *testing.T){
	if PublicValue == "" {
		t.Error("init didn't run")
	}
	if privateValue == "" {
		t.Error("init didn't run")
	}
}