package generator

import (
	"fmt"
	"strings"
)

func GenerateJSClass(className string, def Definition) string {
	var builder strings.Builder

	for _, prop := range def.Properties {
		if prop.Ref != "" {
			refTypeName := extractRefTypeName(prop.Ref)
			builder.WriteString(fmt.Sprintf("import {%s} from \"./%s\"\n", refTypeName, refTypeName))
		}
	}

	builder.WriteString(fmt.Sprintf("export class %s {\n", className))
	builder.WriteString("  constructor() {\n")

	// Constructor
	for _, propName := range def.Required {
		builder.WriteString(fmt.Sprintf("    this._%s = \"\";\n", toCamelCase(propName)))
	}
	builder.WriteString("  }\n\n")

	// fromJson method
	builder.WriteString("  static fromJson(json) {\n")
	builder.WriteString(fmt.Sprintf("    let %s = new %s();\n", strings.ToLower(className), className))
	for propName, prop := range def.Properties {
		camelCaseName := toCamelCase(propName)
		if prop.Ref != "" {
			// Extract referenced type name from $ref
			refTypeName := extractRefTypeName(prop.Ref)
			builder.WriteString(fmt.Sprintf("    %s.%s = %s.fromJson(json.%s);\n", strings.ToLower(className), camelCaseName, refTypeName, propName))
		} else {
			builder.WriteString(fmt.Sprintf("    %s.%s = json.%s;\n", strings.ToLower(className), camelCaseName, propName))
		}
	}
	builder.WriteString(fmt.Sprintf("    return %s;\n", strings.ToLower(className)))
	builder.WriteString("  }\n")

	// Getters and Setters
	for _, propName := range def.Required {
		camelCaseName := toCamelCase(propName)
		builder.WriteString(fmt.Sprintf("\n  get %s() {\n", camelCaseName))
		builder.WriteString(fmt.Sprintf("    return this._%s;\n", camelCaseName))
		builder.WriteString("  }\n\n")
		builder.WriteString(fmt.Sprintf("  set %s(value) {\n", camelCaseName))
		builder.WriteString(fmt.Sprintf("    this._%s = value;\n", camelCaseName))
		builder.WriteString("  }\n")
	}

	builder.WriteString("}\n")
	return builder.String()
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

// Extracts the type name from a $ref string
func extractRefTypeName(ref string) string {
	parts := strings.Split(ref, "/")
	if len(parts) > 0 {
		return toCamelCase(parts[len(parts)-1])
	}
	return ""
}
