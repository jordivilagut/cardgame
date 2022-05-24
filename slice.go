package main

type slice []string

func (slice slice) contains(str string) bool {
	for _, x := range slice {
		if x == str {
			return true
		}
	}

	return false
}
