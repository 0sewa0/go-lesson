package nils

import "fmt"


func MapsGoBoom() {
	var boomer map[string]string
	// boomer := make(map[string]string)
	// boomer := map[string]string{}

	boomer["b"] = "hihi"
}

func SlicesAreSafe() {
	var safe []string
	// boomer := make(map[string]string)

	for range safe {
		fmt.Print("nothing")
	}
	safe = append(safe, "something")
	for range safe {
		fmt.Print("something")
	}
}

type MyError int

func (i MyError) Error() string {
	return fmt.Sprintf("%d", i)
}
