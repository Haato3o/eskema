package emitter

import "github.com/Haato3o/eskema/core/parser"

type LanguageCodeEmitter interface {
	Emit(tree *parser.EskemaTree) string
}
