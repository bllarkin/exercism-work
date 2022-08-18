package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	recRune := '\u2757'
	searchRune := '\U0001F50D'
	weatherRune := '\u2600'

	for _, char := range log {
		if char == recRune {
			return "recommendation"
		}
		if char == searchRune {
			return "search"
		}
		if char == weatherRune {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
