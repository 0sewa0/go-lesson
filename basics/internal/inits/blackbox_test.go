package inits_test

import (
	"testing"

	"go.lesson/basics/internal/inits"
)

func TestInit(t *testing.T){
	if inits.PublicValue == "" {
		t.Error("init didn't run")
	}
	// if inits.privateValue == "" {
	// 	t.Error("init didn't run")
	// }
}