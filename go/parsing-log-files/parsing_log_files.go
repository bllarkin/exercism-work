package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`(?i)^\[(trc|dbg|inf|wrn|err|ftl)\](.*)`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`(?i)<[~*=-]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)\".*(password).*\"`)

	var totalLines int

	for _, text := range lines {
		if re.FindStringSubmatch(text) != nil {
			totalLines++
		}
	}

	return totalLines
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	prefixString := "[USR]"

	re := regexp.MustCompile(`User\s+\S+`)

	for i, text := range lines {
		userText := re.FindString(text)
		if userText != "" {
			userName := regexp.MustCompile(`\s+`).Split(userText, 2)[1]
			lines[i] = prefixString + " " + userName + " " + text
		}
	}

	return lines
}
