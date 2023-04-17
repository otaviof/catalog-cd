package render

import (
	"fmt"
	"strings"
	"text/template"
)

// templateFuncMap map with the functions available on the template.
var templateFuncMap = template.FuncMap{
	"chomp":          chomp,
	"formatType":     formatType,
	"formatValue":    formatValue,
	"formatOptional": formatOptional,
}

// chomp removes new lines.
func chomp(s string) string {
	return strings.TrimSuffix(strings.ReplaceAll(s, "\n", " "), " ")
}

// formatType when type is not informed the type "string" returned.
func formatType(s interface{}) string {
	if s == nil || s.(string) == "" {
		return "string"
	}
	return s.(string)
}

// formatValues highlights the informed value is required or empty.
func formatValue(s interface{}) string {
	if s == nil {
		return "(required)"
	}
	if s.(string) == "" {
		return "\"\" (empty)"
	}
	return fmt.Sprintf("`%s`", s)
}

// formatOptional makes sure "false" is printed when the informed variable is nil.
func formatOptional(s interface{}) string {
	if s == nil || !s.(bool) {
		return "false"
	}
	return "true"
}
