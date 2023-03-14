package languages

import (
	"github.com/Haato3o/eskema/core/codestyle"
	"github.com/Haato3o/eskema/core/parser"
	"github.com/Haato3o/eskema/emitter"
	"strings"
)

var swiftPrimitives = map[string]string{
	"String":    "String",
	"Char":      "Character",
	"UInt8":     "UInt8",
	"UInt16":    "UInt16",
	"UInt32":    "UInt32",
	"UInt64":    "UInt64",
	"Int8":      "Int8",
	"Int16":     "Int16",
	"Int32":     "Int32",
	"Int64":     "Int64",
	"Float":     "Float",
	"Double":    "Double",
	"TimeStamp": "String",
	"Date":      "String",
	"DateTime":  "String",
	"Array":     "[]",
	"Map":       "[:]",
	"Bool":      "Bool",
}

type SwiftEmitter struct {
	buffer strings.Builder
}

func (s *SwiftEmitter) Emit(tree *parser.EskemaTree) string {
	for _, expr := range tree.Expr {
		s.emitExpression(expr)
		s.buffer.WriteString("\n\n")
	}

	return s.buffer.String()
}

func (s *SwiftEmitter) emitExpression(expr *parser.EskemaExpression) {
	switch expr.Type {
	case parser.SchemaExpr:
		s.emitSchema(expr.Data.(*parser.SchemaDefinition))
		break
	case parser.EnumExpr:
		s.emitEnum(expr.Data.(*parser.EnumDefinition))
		break
	default:
		break
	}
}

func (s *SwiftEmitter) emitSchema(schema *parser.SchemaDefinition) {
	s.buffer.WriteString("public struct ")
	s.buffer.WriteString(schema.Id.Name)
	s.buffer.WriteString(": Decodable, Equatable {\n")

	// TODO: Emit generics

	for _, field := range schema.Fields {
		s.buffer.WriteString(Indent)

		s.emitFieldDeclaration(field)

		s.buffer.WriteString("\n")
	}

	if schema.ContainsNullableFields() {
		s.emitNullableConstructor(schema)
	}

	s.buffer.WriteString("}")
}

func (s *SwiftEmitter) emitNullableConstructor(schema *parser.SchemaDefinition) {

	s.buffer.WriteString("\n")

	s.buffer.WriteString(Indent)
	s.buffer.WriteString("public init(")

	for i, field := range schema.Fields {
		isLast := i+1 == len(schema.Fields)

		s.emitConstructorField(field)

		if !isLast {
			s.buffer.WriteString(", ")
		}
	}

	s.buffer.WriteString(") {\n")

	for _, field := range schema.Fields {
		s.buffer.WriteString(Indent)
		s.buffer.WriteString(Indent)
		s.emitFieldInitializer(field)
		s.buffer.WriteString("\n")
	}

	s.buffer.WriteString(Indent)
	s.buffer.WriteString("}\n")
}

func (s *SwiftEmitter) emitField(field *parser.FieldExpression) {
	s.buffer.WriteString(field.Id.Name)
	s.buffer.WriteString(": ")
	s.emitType(field.Type)

	if field.IsOptional {
		s.buffer.WriteString("?")
	}
}

func (s *SwiftEmitter) emitFieldDeclaration(field *parser.FieldExpression) {
	s.buffer.WriteString("let ")
	s.emitField(field)
}

func (s *SwiftEmitter) emitConstructorField(field *parser.FieldExpression) {
	s.emitField(field)

	if field.IsOptional {
		s.buffer.WriteString(" = nil")
	}
}

func (s *SwiftEmitter) emitFieldInitializer(field *parser.FieldExpression) {
	s.buffer.WriteString("self.")
	s.buffer.WriteString(field.Id.Name)
	s.buffer.WriteString(" = ")
	s.buffer.WriteString(field.Id.Name)
}

func (s *SwiftEmitter) emitType(typeExpr *parser.TypeExpression) {
	primitive, isPrimitive := swiftPrimitives[typeExpr.Id.Name]

	isMap := typeExpr.Id.Name == "Map"
	isArray := typeExpr.Id.Name == "Array"

	if isPrimitive {

		if isArray || isMap {
			s.buffer.WriteString("[")
		} else {
			s.buffer.WriteString(primitive)
		}
	} else {
		s.buffer.WriteString(typeExpr.Id.Name)
	}

	for i, typ := range typeExpr.Generics {
		isFirst := i == 0
		isLast := i+1 == len(typeExpr.Generics)

		if isFirst && !isMap && !isArray {
			s.buffer.WriteString("<")
		}

		s.emitType(typ)

		if isLast {
			if !isMap && !isArray {
				s.buffer.WriteString(">")
			}
		} else {
			if isMap {
				s.buffer.WriteString(" : ")
			} else {
				s.buffer.WriteString(", ")
			}
		}
	}

	if isArray || isMap {
		s.buffer.WriteString("]")
	}
}

func (s *SwiftEmitter) emitEnum(enum *parser.EnumDefinition) {
	s.buffer.WriteString("public enum ")
	s.buffer.WriteString(enum.Id.Name)
	s.buffer.WriteString(": String, Decodable, Equatable {\n")

	for _, value := range enum.Values {

		s.buffer.WriteString(Indent)
		s.buffer.WriteString("case ")
		s.buffer.WriteString(codestyle.ToCamelCase(value))
		s.buffer.WriteString(" = \"")
		s.emitLiteralValue(value)
		s.buffer.WriteString("\"\n")
	}

	s.buffer.WriteString("}")
}

func (s *SwiftEmitter) emitLiteralValue(enum string) {
	s.buffer.WriteString(enum)
}

func NewSwiftEmitter() emitter.LanguageCodeEmitter {
	return &SwiftEmitter{}
}
