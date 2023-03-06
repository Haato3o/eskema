package languages

import (
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/emitter"
	"strings"
)

var cSharpPrimitives = map[string]string{
	"String":    "string",
	"Char":      "char",
	"UInt8":     "ubyte",
	"UInt16":    "ushort",
	"UInt32":    "uint",
	"UInt64":    "ulong",
	"Int8":      "byte",
	"Int16":     "short",
	"Int32":     "int",
	"Int64":     "long",
	"Float":     "float",
	"Double":    "double",
	"TimeStamp": "TimeSpan",
	"Date":      "DateTime",
	"DateTime":  "DateTime",
	"Array":     "List",
	"Map":       "Dictionary",
	"Bool":      "bool",
}

type CSharpEmitter struct{}

func (c *CSharpEmitter) Emit(tree *parser.EskemaTree) string {
	var builder strings.Builder

	for _, expr := range tree.Expr {
		builder.WriteString(c.emitExpression(expr))
		builder.WriteString("\n")
	}

	return builder.String()
}

func (c *CSharpEmitter) emitExpression(expr *parser.EskemaExpression) string {
	switch expr.Type {
	case parser.SchemaExpr:
		return c.emitSchema(expr.Data.(*parser.SchemaDefinition))
	case parser.EnumExpr:
		return c.emitEnum(expr.Data.(*parser.EnumDefinition))
	default:
		return ""
	}
}

func (c *CSharpEmitter) emitSchema(schema *parser.SchemaDefinition) string {
	var builder strings.Builder

	builder.WriteString("public record ")
	builder.WriteString(schema.Id.Name)

	if len(schema.Generics) > 0 {
		builder.WriteString("<")

		for i, generic := range schema.Generics {
			isLast := i+1 == len(schema.Generics)

			builder.WriteString(c.emitType(generic))

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

		builder.WriteString(c.emitField(field))

		if !isLast {
			builder.WriteString(",")
		}

		builder.WriteString("\n")
	}

	builder.WriteString(");\n")

	return builder.String()
}

func (c *CSharpEmitter) emitField(field *parser.FieldExpression) string {
	var builder strings.Builder

	builder.WriteString(c.emitType(field.Type))

	if field.IsOptional {
		builder.WriteString("?")
	}

	builder.WriteString(" ")
	builder.WriteString(field.Id.Name)

	return builder.String()
}

func (c *CSharpEmitter) emitType(typeExpr *parser.TypeExpression) string {
	var builder strings.Builder

	primitive, isPrimitive := cSharpPrimitives[typeExpr.Id.Name]

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
		builder.WriteString(c.emitType(typ))

		if isLast {
			builder.WriteString(">")
		} else {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}

func (c *CSharpEmitter) emitEnum(enum *parser.EnumDefinition) string {
	var builder strings.Builder

	builder.WriteString("enum ")
	builder.WriteString(enum.Id.Name)
	builder.WriteString(" {\n")

	for i, value := range enum.Values {

		builder.WriteString(Indent)
		builder.WriteString(c.emitLiteralValue(value))

		isLast := i+1 == len(enum.Values)

		if !isLast {
			builder.WriteString(",")
		}

		builder.WriteString("\n")
	}

	builder.WriteString("};\n")

	return builder.String()
}

func (c *CSharpEmitter) emitLiteralValue(enum string) string {
	return enum
}

func NewCSharpEmitter() emitter.LanguageCodeEmitter {
	return &CSharpEmitter{}
}
