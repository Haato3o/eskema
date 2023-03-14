package languages

import (
	"github.com/Haato3o/eskema/core/codestyle"
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/emitter"
	"strings"
)

var swiftPrimitives = map[string]string{
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
	"Map":       "Map",
	"Bool":      "Boolean",
}

type SwiftEmitter struct{}

func (s *SwiftEmitter) Emit(tree *parser.EskemaTree) string {
	var builder strings.Builder

	for _, expr := range tree.Expr {
		builder.WriteString(s.emitExpression(expr))
		builder.WriteString("\n\n")
	}

	return builder.String()
}

func (s *SwiftEmitter) emitExpression(expr *parser.EskemaExpression) string {
	switch expr.Type {
	case parser.SchemaExpr:
		return s.emitSchema(expr.Data.(*parser.SchemaDefinition))
	case parser.EnumExpr:
		return s.emitEnum(expr.Data.(*parser.EnumDefinition))
	default:
		return ""
	}
}

func (s *SwiftEmitter) emitSchema(schema *parser.SchemaDefinition) string {
	var builder strings.Builder

	builder.WriteString("public struct ")
	builder.WriteString(schema.Id.Name)
	builder.WriteString(": Decodable, Equatable {\n")

	// TODO: Emit generics

	for _, field := range schema.Fields {
		builder.WriteString(Indent)

		builder.WriteString(s.emitFieldDeclaration(field))

		builder.WriteString("\n")
	}

	if schema.ContainsNullableFields() {
		s.emitNullableConstructor(&builder, schema)
	}

	builder.WriteString("}")

	return builder.String()
}

func (s *SwiftEmitter) emitNullableConstructor(builder *strings.Builder, schema *parser.SchemaDefinition) {

	builder.WriteString("\n")

	builder.WriteString(Indent)
	builder.WriteString("public init(")

	for i, field := range schema.Fields {
		isLast := i+1 == len(schema.Fields)

		builder.WriteString(s.emitConstructorField(field))

		if !isLast {
			builder.WriteString(", ")
		}
	}

	builder.WriteString(") {\n")

	for _, field := range schema.Fields {
		builder.WriteString(Indent)
		builder.WriteString(Indent)
		builder.WriteString(s.emitFieldInitializer(field))
		builder.WriteString("\n")
	}

	builder.WriteString(Indent)
	builder.WriteString("}\n")
}

func (s *SwiftEmitter) emitField(field *parser.FieldExpression) string {
	var builder strings.Builder

	builder.WriteString(field.Id.Name)
	builder.WriteString(": ")
	builder.WriteString(s.emitType(field.Type))

	if field.IsOptional {
		builder.WriteString("?")
	}

	return builder.String()
}

func (s *SwiftEmitter) emitFieldDeclaration(field *parser.FieldExpression) string {
	var builder strings.Builder

	builder.WriteString("let ")
	builder.WriteString(s.emitField(field))

	return builder.String()
}

func (s *SwiftEmitter) emitConstructorField(field *parser.FieldExpression) string {
	emittedField := s.emitField(field)

	if field.IsOptional {
		emittedField = emittedField + " = nil"
	}

	return emittedField
}

func (s *SwiftEmitter) emitFieldInitializer(field *parser.FieldExpression) string {
	var builder strings.Builder

	builder.WriteString("self.")
	builder.WriteString(field.Id.Name)
	builder.WriteString(" = ")
	builder.WriteString(field.Id.Name)

	return builder.String()
}

func (s *SwiftEmitter) emitType(typeExpr *parser.TypeExpression) string {
	var builder strings.Builder

	primitive, isPrimitive := ktPrimitives[typeExpr.Id.Name]

	isMap := typeExpr.Id.Name == "Map"
	isArray := typeExpr.Id.Name == "Array"

	if isPrimitive {

		if isArray || isMap {
			builder.WriteString("[")
		} else {
			builder.WriteString(primitive)
		}
	} else {
		builder.WriteString(typeExpr.Id.Name)
	}

	for i, typ := range typeExpr.Generics {
		isFirst := i == 0
		isLast := i+1 == len(typeExpr.Generics)

		if isFirst && !isMap && !isArray {
			builder.WriteString("<")
		}

		builder.WriteString(s.emitType(typ))

		if isLast {
			if !isMap && !isArray {
				builder.WriteString(">")
			}
		} else {
			if isMap {
				builder.WriteString(" : ")
			} else {
				builder.WriteString(", ")
			}
		}
	}

	if isArray || isMap {
		builder.WriteString("]")
	}

	return builder.String()
}

func (s *SwiftEmitter) emitEnum(enum *parser.EnumDefinition) string {
	var builder strings.Builder

	builder.WriteString("public enum ")
	builder.WriteString(enum.Id.Name)
	builder.WriteString(": String, Decodable, Equatable {\n")

	for _, value := range enum.Values {

		builder.WriteString(Indent)
		builder.WriteString("case ")
		builder.WriteString(codestyle.ToCamelCase(value))
		builder.WriteString(" = \"")
		builder.WriteString(s.emitLiteralValue(value))
		builder.WriteString("\"\n")
	}

	builder.WriteString("}")

	return builder.String()
}

func (*SwiftEmitter) emitLiteralValue(enum string) string {
	return enum
}

func NewSwiftEmitter() emitter.LanguageCodeEmitter {
	return &SwiftEmitter{}
}
