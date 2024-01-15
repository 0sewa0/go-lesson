package concurrency

// Share by communicating

func CountChars(char rune, source string) int {
	numberOfRoutines := len(source) / 10
	sections := toSections(source, numberOfRoutines)
	ch := make(chan int, numberOfRoutines)

	for _, section := range sections {
		go countForSection(char, section, ch)
	}

	count := 0
	for range sections {
		count += <-ch
	}

	return count
}

func countForSection(char rune, source string, ch chan<- int) {
	count := 0
	for _, c := range source {
		if c == char {
			count += 1
		}
	}
	ch <- count
}
