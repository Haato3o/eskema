package cli

import "errors"

var (
	ErrMissingFileName = errors.New("missing filename parameter")
	ErrMissingLanguage = errors.New("missing language parameter")
)

type EskemaArguments struct {
	FileName                      string
	Language                      string
	ShouldPrintAST                bool
	ShouldPrintSupportedLanguages bool
}

func (a *EskemaArguments) VerifyRequired() error {

	if a.FileName == "" {
		return ErrMissingFileName
	}

	if a.Language == "" {
		return ErrMissingLanguage
	}

	return nil
}
