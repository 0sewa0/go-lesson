package returns


func Annoying() (string, string, int) {
	return "elephant", "Bob", 5
}

type Animal struct{
	Type string
	Name string
	Weight int
}

func NotAnnoying() Animal {
	return Animal{
		"T-Rex",
		"Jeff",
		9,
	}
}