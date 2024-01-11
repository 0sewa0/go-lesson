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

func BadRename(animal Animal, newName string) {
	animal.Name = newName
}

func GoodRename(animal *Animal, newName string) {
	animal.Name = newName
}