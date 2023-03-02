package syntax

import (
	"io"
	"os"
)

type EskemaLexer struct {
	fileName string
	stream   io.ReadSeekCloser
	current  int64
	column   int64
	line     int64
}

func (l *EskemaLexer) discard() {
	_, _ = l.consume()
}

func (l *EskemaLexer) consume() (byte, error) {
	currentCharacter := make([]byte, 1)
	_, err := l.stream.Read(currentCharacter)
	l.current++
	l.column++

	return currentCharacter[0], err
}

func (l *EskemaLexer) prev() {
	_, _ = l.stream.Seek(-1, io.SeekCurrent)
	l.current--
	l.column--
}

func (l *EskemaLexer) peek() (byte, error) {
	value, err := l.consume()
	l.prev()

	return value, err
}

func (l *EskemaLexer) Lex() *Token {
	metadata := &Metadata{
		Filename: l.fileName,
		Offset:   l.current,
		Line:     l.line,
		Column:   l.column,
	}

	currentCharacter, err := l.peek()

	if err == io.EOF {
		return &Token{
			Metadata: metadata,
			Value:    "\x00",
			Type:     EndOfFileToken,
		}
	} else if err != nil {
		// TODO: Return error here
		return nil
	}

	if isSpecial, specialToken := IsSpecialToken(currentCharacter); isSpecial {

		l.discard()

		if specialToken == NewLineToken {
			l.line++
			l.column = 0
		}

		if specialToken == NewLineToken ||
			specialToken == WhitespaceToken ||
			specialToken == InvalidToken {
			return nil
		}

		return &Token{
			Metadata: metadata,
			Value:    string(currentCharacter),
			Type:     specialToken,
		}
	}

	start := l.current
	for {
		// TODO: Handle EOF

		if err != nil {
			break
		}

		if isSpecial, _ := IsSpecialToken(currentCharacter); isSpecial {
			l.prev()
			break
		}

		currentCharacter, err = l.consume()
	}

	length := l.current - start
	buffer := make([]byte, length)
	_, _ = l.stream.Seek(start, io.SeekStart)
	_, _ = l.stream.Read(buffer)
	literal := string(buffer)

	if isKeyword, _ := IsKeyword(literal); isKeyword {
		return &Token{
			Metadata: metadata,
			Value:    literal,
			Type:     KeywordToken,
		}
	}

	if isPrimitiveType, _ := IsPrimitiveType(literal); isPrimitiveType {
		return &Token{
			Metadata: metadata,
			Value:    literal,
			Type:     PrimitiveTypeToken,
		}
	}

	return &Token{
		Metadata: metadata,
		Value:    literal,
		Type:     LiteralToken,
	}
}

func NewLexer(path string) *EskemaLexer {
	file, _ := os.OpenFile(path, os.O_RDONLY, os.ModeType)

	return &EskemaLexer{
		fileName: path,
		stream:   file,
		current:  0,
		line:     0,
	}
}
