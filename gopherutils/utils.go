package gopherutils

import (
	"fmt"
	"strings"
)

func RemovePrefix(command string) string {
	return command[1:]
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
