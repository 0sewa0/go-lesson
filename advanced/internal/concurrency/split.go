package concurrency

func toSections(str string, splits int) []string {
	sections := []string{}
	for i := 0; i < len(str); i += splits {
		end := i + splits
		if end > len(str) {
			end = len(str)
		}
		sections = append(sections, str[i:end])
	}
	return sections
}
