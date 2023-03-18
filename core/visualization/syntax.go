package visualization

import (
	"fmt"
	"github.com/Haato3o/eskema/core/parser"
)

type TreeOrder int

const (
	First TreeOrder = iota
	Middle
	Last
)

const (
	TreeEmptyCharacter  = " "
	TreeCharacter       = "├──"
	TreeMiddleCharacter = "│"
	TreeEndCharacter    = "└──"
)

func VisualizeTree(tree *parser.EskemaTree) {
	result := ""

	for i, expr := range tree.Expr {
		result += buildExpression(expr, getOrder(i, len(tree.Expr)))
	}

	println(result)
}

func buildExpression(expr *parser.EskemaExpression, order TreeOrder) string {

	switch expr.Type {
	case parser.EnumExpr:
		return buildEnum(expr.Data.(*parser.EnumDefinition), order)
	case parser.SchemaExpr:
		return buildSchema(expr.Data.(*parser.SchemaDefinition), order)
	default:
		return ""
	}
}

func buildSchema(schema *parser.SchemaDefinition, order TreeOrder) string {
	level := fmt.Sprintf("%s   ", getTreeRootConnector(order))

	baseString := fmt.Sprintf("%s schema: %s\n", getParentConnector(order), schema.Id.Name)

	for i, generic := range schema.Generics {
		baseString += buildType(generic, level, getOrder(i, len(schema.Generics)+len(schema.Fields)))
	}

	for i, field := range schema.Fields {
		baseString += buildField(field, level, getOrder(i, len(schema.Fields)))
	}

	return baseString
}

func buildType(typeExpr *parser.TypeExpression, level string, order TreeOrder) string {
	childLevel := fmt.Sprintf("%s%s   ", level, getTreeRootConnector(order))
	currentLevel := fmt.Sprintf("%s%s", level, getParentConnector(order))

	baseString := fmt.Sprintf("%s type: '%s'\n", currentLevel, typeExpr.Id.Name)

	for i, generic := range typeExpr.Generics {
		baseString += buildType(generic, childLevel, getOrder(i, len(typeExpr.Generics)))
	}

	return baseString
}

func buildAnnotation(annotation *parser.AnnotationExpression, level string, order TreeOrder) string {
	currentLevel := fmt.Sprintf("%s%s", level, getParentConnector(order))

	return fmt.Sprintf("%s annotation: %s = %s\n", currentLevel, annotation.Id.Name, annotation.Value)
}

func buildField(field *parser.FieldExpression, level string, order TreeOrder) string {
	childLevel := fmt.Sprintf("%s%s   ", level, getTreeRootConnector(order))
	currentLevel := fmt.Sprintf("%s%s", level, getParentConnector(order))

	optional := "[Required]"

	if field.IsOptional {
		optional = "[Nullable]"
	}

	baseString := fmt.Sprintf("%s field: %s %s\n", currentLevel, field.Id.Name, optional)

	for i, annotation := range field.Annotations {

		baseString += buildAnnotation(annotation, childLevel, getOrder(i, len(field.Annotations)+1))
	}

	baseString += buildType(field.Type, childLevel, Last)

	return baseString
}

func buildEnum(enum *parser.EnumDefinition, order TreeOrder) string {
	level := fmt.Sprintf("%s   ", getTreeRootConnector(order))

	baseString := fmt.Sprintf("%s enum: %s\n", getParentConnector(order), enum.Id.Name)

	for i, value := range enum.Values {

		isLast := i == (len(enum.Values) - 1)

		baseString += buildValue(value, level, isLast)
	}

	return baseString
}

func buildValue(value string, level string, isLast bool) string {
	connector := TreeCharacter

	if isLast {
		connector = TreeEndCharacter
	}

	return fmt.Sprintf("%s%s %s\n", level, connector, value)
}

func getOrder(index int, maxIndexes int) TreeOrder {
	switch index + 1 {
	case maxIndexes:
		return Last
	case 1:
		return First
	default:
		return Middle
	}
}

func getTreeRootConnector(order TreeOrder) string {
	switch order {
	case First:
		return TreeMiddleCharacter
	case Middle:
		return TreeMiddleCharacter
	default:
		return TreeEmptyCharacter
	}
}

func getParentConnector(order TreeOrder) string {
	switch order {
	case First:
		return TreeCharacter
	case Last:
		return TreeEndCharacter
	default:
		return TreeCharacter
	}
}
