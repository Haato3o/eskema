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
	"TimeStamp": "DateTime",
	"Date":      "DateTime",
	"DateTime":  "DateTime",
	"Array":     "List",
	"Map":       "Dictionary",
	"Bool":      "bool",
}

type CSharpEmitter struct {
	buffer strings.Builder
}

func (c *CSharpEmitter) Emit(tree *parser.EskemaTree) string {
	c.buffer.WriteString("namespace Example;\n\n")

	for _, expr := range tree.Expr {
		c.emitExpression(expr)
		c.buffer.WriteString("\n")
	}

	return c.buffer.String()
}

func (c *CSharpEmitter) emitExpression(expr *parser.EskemaExpression) {
	switch expr.Type {
	case parser.SchemaExpr:
		c.emitSchema(expr.Data.(*parser.SchemaDefinition))
		break
	case parser.EnumExpr:
		c.emitEnum(expr.Data.(*parser.EnumDefinition))
		break
	}
}

func (c *CSharpEmitter) emitSchema(schema *parser.SchemaDefinition) {
	c.buffer.WriteString("public record ")
	c.buffer.WriteString(schema.Id.Name)

	if len(schema.Generics) > 0 {
		c.buffer.WriteString("<")

		for i, generic := range schema.Generics {
			isLast := i+1 == len(schema.Generics)

			c.emitType(generic)

			if !isLast {
				c.buffer.WriteString(", ")
			}
		}

		c.buffer.WriteString(">")
	}

	c.buffer.WriteString("(\n")

	for i, field := range schema.Fields {

		isLast := i+1 == len(schema.Fields)

		c.buffer.WriteString(Indent)

		c.emitField(field)

		if !isLast {
			c.buffer.WriteString(",")
		}

		c.buffer.WriteString("\n")
	}

	c.buffer.WriteString(");\n")
}

func (c *CSharpEmitter) emitField(field *parser.FieldExpression) {
	c.emitType(field.Type)

	if field.IsOptional {
		c.buffer.WriteString("?")
	}

	c.buffer.WriteString(" ")
	c.buffer.WriteString(field.Id.Name)
}

func (c *CSharpEmitter) emitType(typeExpr *parser.TypeExpression) {
	primitive, isPrimitive := cSharpPrimitives[typeExpr.Id.Name]

	if isPrimitive {
		c.buffer.WriteString(primitive)
	} else {
		c.buffer.WriteString(typeExpr.Id.Name)
	}

	for i, typ := range typeExpr.Generics {
		isFirst := i == 0
		isLast := i+1 == len(typeExpr.Generics)

		if isFirst {
			c.buffer.WriteString("<")
		}

		c.emitType(typ)

		if isLast {
			c.buffer.WriteString(">")
		} else {
			c.buffer.WriteString(", ")
		}
	}
}

func (c *CSharpEmitter) emitEnum(enum *parser.EnumDefinition) {
	c.buffer.WriteString("enum ")
	c.buffer.WriteString(enum.Id.Name)
	c.buffer.WriteString(" {\n")

	for i, value := range enum.Values {

		c.buffer.WriteString(Indent)
		c.emitLiteralValue(value)

		isLast := i+1 == len(enum.Values)

		if !isLast {
			c.buffer.WriteString(",")
		}

		c.buffer.WriteString("\n")
	}

	c.buffer.WriteString("};\n")
}

func (c *CSharpEmitter) emitLiteralValue(enum string) {
	c.buffer.WriteString(enum)
}

func NewCSharpEmitter() emitter.LanguageCodeEmitter {
	return &CSharpEmitter{}
}
