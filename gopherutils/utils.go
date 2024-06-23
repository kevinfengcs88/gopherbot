package gopherutils

import (
	"fmt"
	// "regexp"
	"strings"
)

func RemovePrefix(command string) string {
	return command[1:]
}

func CleanedBytesToString(b []byte) string {
	return strings.TrimSpace(string(b))
}

func Redify(input string) string {
	var builder strings.Builder
	builder.WriteString("```diff\n")
	builder.WriteString(fmt.Sprintf("- %s\n", input))
	builder.WriteString("```")
	return builder.String()
}

func Greenify(input string) string {
	var builder strings.Builder
	builder.WriteString("```diff\n")
	builder.WriteString(fmt.Sprintf("+ %s\n", input))
	builder.WriteString("```")
	return builder.String()
}

// func ParseLog(line string) string {
// 	pattern := regexp.MustCompile(`\d{4}`)
// 	return "foo"
// }

// func ParseServerProgress(line string) string {
// }
