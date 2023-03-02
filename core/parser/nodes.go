package parser

type IdentifierExpression struct {
	Name string
}

type SchemaDefinition struct {
	Fields []*FieldExpression
}

type FieldExpression struct {
	Identifier *IdentifierExpression
	IsOptional bool
	Type       *TypeExpression
}

type TypeExpression struct {
	Name       string
	Parameters []*TypeExpression
}

type EnumDefinition struct {
	Values []string
}
