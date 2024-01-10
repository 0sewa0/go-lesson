package inits

import "fmt"

var (
	PublicValue string
	privateValue string
)

func init(){
	fmt.Println("Hello from init")
	PublicValue = "init-1"
	privateValue = "init-2"
}