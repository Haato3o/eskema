package syntax

import (
	"reflect"
	"testing"
)

func TestTokenStream(t *testing.T) {
	testCase := []*Token{
		{&Metadata{}, "Test", LiteralToken},
		{&Metadata{}, "Test2", LiteralToken},
		{&Metadata{}, "Test3", LiteralToken},
	}

	target := newTokenStream(testCase)

	t.Run("should consume and return current token", func(t *testing.T) {
		actual := target.Next()
		expected := testCase[0]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, expected %v", actual, expected)
		}
	})

	t.Run("should not consume and return current token", func(t *testing.T) {
		actual := target.PeekCurrent()
		expected := testCase[1]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, expected %v", actual, expected)
		}
	})

	t.Run("should go back to previous token and return previous token", func(t *testing.T) {
		actual := target.Prev()
		expected := testCase[0]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, expected %v", actual, expected)
		}
	})

	t.Run("should not consume and return next token", func(t *testing.T) {
		actual := target.Peek()
		expected := testCase[1]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, expected %v", actual, expected)
		}
	})

	t.Run("should not consume return last token if index is higher than tokens quantity", func(t *testing.T) {
		actual := target.PeekAt(3)
		expected := testCase[2]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, expected %v", actual, expected)
		}
	})

	t.Run("should not consume and return token at given index", func(t *testing.T) {
		actual := target.PeekAt(1)
		expected := testCase[1]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v, expected %v", actual, expected)
		}
	})
}
