package languages

import (
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/emitter"
	"strings"
)

var goLangPrimitives = map[string]string{
	"String":    "string",
	"Char":      "byte",
	"UInt8":     "uint8",
	"UInt16":    "uint16",
	"UInt32":    "uint32",
	"UInt64":    "uint64",
	"Int8":      "int8",
	"Int16":     "int16",
	"Int32":     "int32",
	"Int64":     "int64",
	"Float":     "float",
	"Double":    "double",
	"TimeStamp": "time.Time",
	"Date":      "time.Time",
	"DateTime":  "time.Time",
	"Array":     "[]",
	"Map":       "map",
	"Bool":      "bool",
}

type GoLangEmitter struct{}

type SimpleSchemaWithGenerics[T any] struct {
	value1 T
	value2 []T
}

type ComplexSchema[TIn string, TOut any] struct {
	value1 *map[TIn]SimpleSchemaWithGenerics[TOut]
	value2 *[][]string
}

func (c *GoLangEmitter) Emit(tree *parser.EskemaTree) string {
	var builder strings.Builder

	for _, expr := range tree.Expr {
		builder.WriteString(c.emitExpression(expr))
		builder.WriteString("\n")
	}

	return builder.String()
}

func (c *GoLangEmitter) emitExpression(expr *parser.EskemaExpression) string {
	switch expr.Type {
	case parser.SchemaExpr:
		return c.emitSchema(expr.Data.(*parser.SchemaDefinition))
	case parser.EnumExpr:
		return c.emitEnum(expr.Data.(*parser.EnumDefinition))
	default:
		return ""
	}
}

func (c *GoLangEmitter) emitSchema(schema *parser.SchemaDefinition) string {
	var builder strings.Builder

	builder.WriteString("type ")
	builder.WriteString(schema.Id.Name)

	if len(schema.Generics) > 0 {
		builder.WriteString("[")

		for i, generic := range schema.Generics {
			isLast := i+1 == len(schema.Generics)

			builder.WriteString(c.emitType(generic))
			builder.WriteString(" any")

			if !isLast {
				builder.WriteString(", ")
			}
		}

		builder.WriteString("]")
	}

	builder.WriteString(" struct")

	builder.WriteString(" {\n")

	for _, field := range schema.Fields {
		builder.WriteString(Indent)

		builder.WriteString(c.emitField(field))

		builder.WriteString("\n")
	}

	builder.WriteString("}\n")

	return builder.String()
}

func (c *GoLangEmitter) emitField(field *parser.FieldExpression) string {
	var builder strings.Builder

	builder.WriteString(field.Id.Name)
	builder.WriteString(" ")

	if field.IsOptional {
		builder.WriteString("*")
	}

	builder.WriteString(c.emitType(field.Type))

	return builder.String()
}

func (c *GoLangEmitter) emitType(typeExpr *parser.TypeExpression) string {
	var builder strings.Builder

	primitive, isPrimitive := goLangPrimitives[typeExpr.Id.Name]

	if isPrimitive {
		builder.WriteString(primitive)
	} else {
		builder.WriteString(typeExpr.Id.Name)
	}

	for i, typ := range typeExpr.Generics {
		isFirst := i == 0
		hasMultipleElements := len(typeExpr.Generics) > 1
		isArray := typeExpr.Id.Name == "Array"

		if isFirst && (hasMultipleElements || !isArray) {
			builder.WriteString("[")
		}

		builder.WriteString(c.emitType(typ))

		if isFirst && (hasMultipleElements || !isArray) {
			builder.WriteString("]")
		}
	}

	return builder.String()
}

func (c *GoLangEmitter) emitEnum(enum *parser.EnumDefinition) string {
	var builder strings.Builder

	builder.WriteString("type ")
	builder.WriteString(enum.Id.Name)
	builder.WriteString(" int\n")

	builder.WriteString("const (\n")

	for i, value := range enum.Values {

		isFirst := i == 0

		builder.WriteString(Indent)
		builder.WriteString(c.emitLiteralValue(value))

		if isFirst {
			builder.WriteString(" ")
			builder.WriteString(enum.Id.Name)
			builder.WriteString(" = iota")
		}

		builder.WriteString("\n")
	}

	builder.WriteString(")\n")

	return builder.String()
}

func (c *GoLangEmitter) emitLiteralValue(enum string) string {
	return enum
}

func NewGoLangEmitter() emitter.LanguageCodeEmitter {
	return &GoLangEmitter{}
}
