package cli

import (
	"errors"
	"fmt"
	"github.com/Haato3o/eskema/emitter"
	"github.com/Haato3o/eskema/emitter/languages"
	"log"
	"strings"
)

const (
	ErrUnsupportedLanguage = "language '%s' is not supported"
)

var supportedLanguages = map[string]emitter.LanguageCodeEmitter{
	"kotlin": languages.NewKotlinEmitter(),
	"csharp": languages.NewCSharpEmitter(),
}

func GetLanguageEmitter(language string) (emitter.LanguageCodeEmitter, error) {
	if lang, isSupported := supportedLanguages[language]; isSupported {
		return lang, nil
	}

	return nil, errors.New(fmt.Sprintf(ErrUnsupportedLanguage, language))
}

func PrintSupportedLanguages() {
	var builder strings.Builder

	for lang, _ := range supportedLanguages {
		builder.WriteString("   - ")
		builder.WriteString(lang)
		builder.WriteString("\n")
	}

	log.Println(builder.String())
}
