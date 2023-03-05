package languages

import (
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/emitter"
	"strings"
)

const Indent = "    "

var ktPrimitives = map[string]string{
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

func (k *KotlinEmitter) Emit(tree *parser.EskemaTree) string {
	var builder strings.Builder

	for _, expr := range tree.Expr {
		builder.WriteString(k.emitExpression(expr))
		builder.WriteString("\n")
	}

	return builder.String()
}

func (k *KotlinEmitter) emitExpression(expr *parser.EskemaExpression) string {
	switch expr.Type {
	case parser.SchemaExpr:
		return k.emitSchema(expr.Data.(*parser.SchemaDefinition))
	case parser.EnumExpr:
		return k.emitEnum(expr.Data.(*parser.EnumDefinition))
	default:
		return ""
	}
}

func (k *KotlinEmitter) emitSchema(schema *parser.SchemaDefinition) string {
	var builder strings.Builder

	builder.WriteString("data class ")
	builder.WriteString(schema.Id.Name)

	if len(schema.Generics) > 0 {
		builder.WriteString("<")

		for i, generic := range schema.Generics {
			isLast := i+1 == len(schema.Generics)

			builder.WriteString(k.emitType(generic))

			if !isLast {
				builder.WriteString(", ")
			}
		}

		builder.WriteString(">")
	}

	builder.WriteString("(\n")

	for i, field := range schema.Fields {

		isLast := i+1 == len(schema.Fields)

		builder.WriteString(Indent)

		builder.WriteString(k.emitField(field))

		if !isLast {
			builder.WriteString(",")
		}

		builder.WriteString("\n")
	}

	builder.WriteString(")\n")

	return builder.String()
}

func (k *KotlinEmitter) emitField(field *parser.FieldExpression) string {
	var builder strings.Builder

	builder.WriteString("val ")
	builder.WriteString(field.Id.Name)
	builder.WriteString(": ")
	builder.WriteString(k.emitType(field.Type))

	if field.IsOptional {
		builder.WriteString("?")
	}

	return builder.String()
}

func (k *KotlinEmitter) emitType(typeExpr *parser.TypeExpression) string {
	var builder strings.Builder

	primitive, isPrimitive := ktPrimitives[typeExpr.Id.Name]

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
		builder.WriteString(k.emitType(typ))

		if isLast {
			builder.WriteString(">")
		} else {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}

func (k *KotlinEmitter) emitEnum(enum *parser.EnumDefinition) string {
	var builder strings.Builder

	builder.WriteString("enum class ")
	builder.WriteString(enum.Id.Name)
	builder.WriteString(" {\n")

	for i, value := range enum.Values {

		builder.WriteString(Indent)
		builder.WriteString(k.emitLiteralValue(value))

		isLast := i+1 == len(enum.Values)

		if !isLast {
			builder.WriteString(",")
		}

		builder.WriteString("\n")
	}

	builder.WriteString("}\n")

	return builder.String()
}

func (*KotlinEmitter) emitLiteralValue(enum string) string {
	return enum
}

func NewKotlinEmitter() emitter.LanguageCodeEmitter {
	return &KotlinEmitter{}
}
