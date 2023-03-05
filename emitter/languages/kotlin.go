package languages

import (
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/emitter"
	"strings"
)

const Indent = "    "

var primitives = map[string]string{
	"String":    "String",
	"Char":      "Char",
	"UInt8":     "UByte",
	"UInt16":    "UShort",
	"UInt32":    "UInt",
	"UInt64":    "ULong",
	"Int8":      "Byte",
	"Int16":     "Short",
	"Int32":     "Int",
	"Int64":     "Long",
	"Float":     "Float",
	"Double":    "Double",
	"TimeStamp": "Instant",
	"Date":      "LocalDate",
	"DateTime":  "LocalDateTime",
	"Array":     "List",
	"Map":       "Map",
	"Bool":      "Boolean",
}

type KotlinEmitter struct{}

func (*KotlinEmitter) Emit(tree *parser.EskemaTree) string {
	var builder strings.Builder

	for _, expr := range tree.Expr {
		builder.WriteString(emitExpression(expr))
		builder.WriteString("\n")
	}

	return builder.String()
}

func emitExpression(expr *parser.EskemaExpression) string {
	switch expr.Type {
	case parser.SchemaExpr:
		return emitSchema(expr.Data.(*parser.SchemaDefinition))
	case parser.EnumExpr:
		return emitEnum(expr.Data.(*parser.EnumDefinition))
	default:
		return ""
	}
}

func emitSchema(schema *parser.SchemaDefinition) string {
	var builder strings.Builder

	builder.WriteString("data class ")
	builder.WriteString(schema.Id.Name)

	if len(schema.Generics) > 0 {
		builder.WriteString("<")

		for _, generic := range schema.Generics {
			builder.WriteString(emitType(generic))
		}

		builder.WriteString(">")
	}

	builder.WriteString("(\n")

	for i, field := range schema.Fields {

		isLast := i+1 == len(schema.Fields)

		builder.WriteString(Indent)

		builder.WriteString(emitField(field))

		if !isLast {
			builder.WriteString(",")
		}

		builder.WriteString("\n")
	}

	builder.WriteString(")\n")

	return builder.String()
}

func emitField(field *parser.FieldExpression) string {
	var builder strings.Builder

	builder.WriteString("val ")
	builder.WriteString(field.Id.Name)
	builder.WriteString(": ")
	builder.WriteString(emitType(field.Type))

	if field.IsOptional {
		builder.WriteString("?")
	}

	return builder.String()
}

func emitType(typeExpr *parser.TypeExpression) string {
	var builder strings.Builder

	primitive, isPrimitive := primitives[typeExpr.Id.Name]

	if isPrimitive {
		builder.WriteString(primitive)
	} else {
		builder.WriteString(typeExpr.Id.Name)
	}

	for i, typ := range typeExpr.Generics {
		isFirst := i == 0
		isLast := i+1 == len(typeExpr.Generics)

		if isFirst {
			builder.WriteString("<")

		}
		builder.WriteString(emitType(typ))

		if isLast {
			builder.WriteString(">")
		} else {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}

func emitEnum(enum *parser.EnumDefinition) string {
	var builder strings.Builder

	builder.WriteString("enum class ")
	builder.WriteString(enum.Id.Name)
	builder.WriteString(" {\n")

	for i, value := range enum.Values {

		builder.WriteString(Indent)
		builder.WriteString(emitLiteralValue(value))

		isLast := i+1 == len(enum.Values)

		if !isLast {
			builder.WriteString(",")
		}

		builder.WriteString("\n")
	}

	builder.WriteString("}\n")

	return builder.String()
}

func emitLiteralValue(enum string) string {
	return enum
}

func NewKotlinEmitter() emitter.LanguageCodeEmitter {
	return &KotlinEmitter{}
}
