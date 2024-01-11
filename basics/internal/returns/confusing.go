package returns


func Annoying() (string, string, int) {
	return "elephant", "Bob", 5
}

func LessAnnoying() (type_ string, name string, weight int) {
	type_ = "parrot"
	name = "Steven"
	weight = 1
	return
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