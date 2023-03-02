package parser

type EskemaExprType int

const (
	SchemaExpr EskemaExprType = iota
	EnumExpr
)

type EskemaExpression struct {
	Type EskemaExprType
	Data interface{}
}

type IdentifierExpression struct {
	Name string
}

type SchemaDefinition struct {
	Id     IdentifierExpression
	Fields []*FieldExpression
}

type FieldExpression struct {
	Id         IdentifierExpression
	IsOptional bool
	Type       *TypeExpression
}

type TypeExpression struct {
	Id         IdentifierExpression
	Parameters []*TypeExpression
}

type EnumDefinition struct {
	Id     IdentifierExpression
	Values []string
}

type EskemaTree struct {
	Expr []*EskemaExpression
}
