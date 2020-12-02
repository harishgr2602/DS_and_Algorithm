package main

func match(source string, pattern string) int {
	iSource := 0
	iPattern := 0
	sourceLen := len(source)
	patternLen := len(pattern)
	for iSource = 0; iSource < sourceLen; iSource++ {
		if source[iSource] == pattern[iPattern] {
			iPattern++
		}
		if iPattern == patternLen {
			return 1
		}
	}
	return 0
}
