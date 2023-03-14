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

type GoLangEmitter struct {
	buffer strings.Builder
}

func (g *GoLangEmitter) Emit(tree *parser.EskemaTree) string {
	g.buffer.WriteString("package example\n\n")

	for _, expr := range tree.Expr {
		g.emitExpression(expr)
		g.buffer.WriteString("\n")
	}

	return g.buffer.String()
}

func (g *GoLangEmitter) emitExpression(expr *parser.EskemaExpression) {
	switch expr.Type {
	case parser.SchemaExpr:
		g.emitSchema(expr.Data.(*parser.SchemaDefinition))
		break
	case parser.EnumExpr:
		g.emitEnum(expr.Data.(*parser.EnumDefinition))
		break
	default:
		break
	}
}

func (g *GoLangEmitter) emitSchema(schema *parser.SchemaDefinition) {
	g.buffer.WriteString("type ")
	g.buffer.WriteString(schema.Id.Name)

	if len(schema.Generics) > 0 {
		g.buffer.WriteString("[")

		for i, generic := range schema.Generics {
			isLast := i+1 == len(schema.Generics)

			g.emitType(generic)
			g.buffer.WriteString(" any")

			if !isLast {
				g.buffer.WriteString(", ")
			}
		}

		g.buffer.WriteString("]")
	}

	g.buffer.WriteString(" struct")

	g.buffer.WriteString(" {\n")

	for _, field := range schema.Fields {
		g.buffer.WriteString(Indent)

		g.emitField(field)

		g.buffer.WriteString("\n")
	}

	g.buffer.WriteString("}\n")
}

func (g *GoLangEmitter) emitField(field *parser.FieldExpression) {
	g.buffer.WriteString(field.Id.Name)
	g.buffer.WriteString(" ")

	if field.IsOptional {
		g.buffer.WriteString("*")
	}

	g.emitType(field.Type)
}

func (g *GoLangEmitter) emitType(typeExpr *parser.TypeExpression) {
	primitive, isPrimitive := goLangPrimitives[typeExpr.Id.Name]

	if isPrimitive {
		g.buffer.WriteString(primitive)
	} else {
		g.buffer.WriteString(typeExpr.Id.Name)
	}

	for i, typ := range typeExpr.Generics {
		isFirst := i == 0
		hasMultipleElements := len(typeExpr.Generics) > 1
		isArray := typeExpr.Id.Name == "Array"

		if isFirst && (hasMultipleElements || !isArray) {
			g.buffer.WriteString("[")
		}

		g.emitType(typ)

		if isFirst && (hasMultipleElements || !isArray) {
			g.buffer.WriteString("]")
		}
	}
}

func (g *GoLangEmitter) emitEnum(enum *parser.EnumDefinition) {
	g.buffer.WriteString("type ")
	g.buffer.WriteString(enum.Id.Name)
	g.buffer.WriteString(" int\n")

	g.buffer.WriteString("const (\n")

	for i, value := range enum.Values {

		isFirst := i == 0

		g.buffer.WriteString(Indent)
		g.emitLiteralValue(value)

		if isFirst {
			g.buffer.WriteString(" ")
			g.buffer.WriteString(enum.Id.Name)
			g.buffer.WriteString(" = iota")
		}

		g.buffer.WriteString("\n")
	}

	g.buffer.WriteString(")\n")
}

func (g *GoLangEmitter) emitLiteralValue(enum string) {
	g.buffer.WriteString(enum)
}

func NewGoLangEmitter() emitter.LanguageCodeEmitter {
	return &GoLangEmitter{}
}
