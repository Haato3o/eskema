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

type KotlinEmitter struct {
	buffer strings.Builder
}

func (k *KotlinEmitter) Emit(tree *parser.EskemaTree) string {
	k.buffer.WriteString("package com.example\n\n")

	for _, expr := range tree.Expr {
		k.emitExpression(expr)
		k.buffer.WriteString("\n")
	}

	return k.buffer.String()
}

func (k *KotlinEmitter) emitExpression(expr *parser.EskemaExpression) {
	switch expr.Type {
	case parser.SchemaExpr:
		k.emitSchema(expr.Data.(*parser.SchemaDefinition))
		break
	case parser.EnumExpr:
		k.emitEnum(expr.Data.(*parser.EnumDefinition))
		break
	default:
		break
	}
}

func (k *KotlinEmitter) emitSchema(schema *parser.SchemaDefinition) {
	k.buffer.WriteString("data class ")
	k.buffer.WriteString(schema.Id.Name)

	if len(schema.Generics) > 0 {
		k.buffer.WriteString("<")

		for i, generic := range schema.Generics {
			isLast := i+1 == len(schema.Generics)

			k.emitType(generic)

			if !isLast {
				k.buffer.WriteString(", ")
			}
		}

		k.buffer.WriteString(">")
	}

	k.buffer.WriteString("(\n")

	for i, field := range schema.Fields {

		isLast := i+1 == len(schema.Fields)

		k.buffer.WriteString(Indent)

		k.emitField(field)

		if !isLast {
			k.buffer.WriteString(",")
		}

		k.buffer.WriteString("\n")
	}

	k.buffer.WriteString(")\n")
}

func (k *KotlinEmitter) emitField(field *parser.FieldExpression) {
	k.buffer.WriteString("val ")
	k.buffer.WriteString(field.Id.Name)
	k.buffer.WriteString(": ")
	k.emitType(field.Type)

	if field.IsOptional {
		k.buffer.WriteString("?")
	}
}

func (k *KotlinEmitter) emitType(typeExpr *parser.TypeExpression) {
	primitive, isPrimitive := ktPrimitives[typeExpr.Id.Name]

	if isPrimitive {
		k.buffer.WriteString(primitive)
	} else {
		k.buffer.WriteString(typeExpr.Id.Name)
	}

	for i, typ := range typeExpr.Generics {
		isFirst := i == 0
		isLast := i+1 == len(typeExpr.Generics)

		if isFirst {
			k.buffer.WriteString("<")

		}

		k.emitType(typ)

		if isLast {
			k.buffer.WriteString(">")
		} else {
			k.buffer.WriteString(", ")
		}
	}
}

func (k *KotlinEmitter) emitEnum(enum *parser.EnumDefinition) {
	k.buffer.WriteString("enum class ")
	k.buffer.WriteString(enum.Id.Name)
	k.buffer.WriteString(" {\n")

	for i, value := range enum.Values {

		k.buffer.WriteString(Indent)

		k.emitLiteralValue(value)

		isLast := i+1 == len(enum.Values)

		if !isLast {
			k.buffer.WriteString(",")
		}

		k.buffer.WriteString("\n")
	}

	k.buffer.WriteString("}\n")
}

func (k *KotlinEmitter) emitLiteralValue(enum string) {
	k.buffer.WriteString(enum)
}

func NewKotlinEmitter() emitter.LanguageCodeEmitter {
	return &KotlinEmitter{}
}
