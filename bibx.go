package bibx

import (
	"bufio"
	"io"
)

func Extract(input io.Reader) []string {
	result := []string{}
	inside := false
	current := ""
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Text() == "```bibtex" {
			inside = true
			continue
		} else if scanner.Text() == "```" && inside {
			inside = false
			result = append(result, current)
			current = ""
		}

		if inside {
			current += scanner.Text() + "\n"
		}
	}
	return result
}

func Merge(extracted []string) string {
	output := ""
	for _, item := range extracted {
		output += item + "\n"
	}
	return output
}
