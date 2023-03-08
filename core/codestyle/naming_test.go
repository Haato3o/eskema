package codestyle

import (
	"fmt"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{"testCase", "test_case"},
		{"TestCase", "test_case"},
		{"test_case", "test_case"},
		{"TEST_CASE", "test_case"},
		{"test", "test"},
		{"TEST", "test"},
	}

	for _, testCase := range testCases {
		name := fmt.Sprintf("should convert %s into %s", testCase.Input, testCase.Expected)

		t.Run(name, func(t *testing.T) {
			actual := ToSnakeCase(testCase.Input)

			if actual != testCase.Expected {
				t.Errorf("got %s, expected %s", actual, testCase.Expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{"testCase", "testCase"},
		{"TestCase", "testCase"},
		{"test_case", "testCase"},
		{"TEST_CASE", "testCase"},
		{"test", "test"},
		{"TEST", "test"},
	}

	for _, testCase := range testCases {
		name := fmt.Sprintf("should convert %s into %s", testCase.Input, testCase.Expected)

		t.Run(name, func(t *testing.T) {
			actual := ToCamelCase(testCase.Input)

			if actual != testCase.Expected {
				t.Errorf("got %s, expected %s", actual, testCase.Expected)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{"testCase", "TestCase"},
		{"TestCase", "TestCase"},
		{"test_case", "TestCase"},
		{"TEST_CASE", "TestCase"},
		{"test", "Test"},
		{"TEST", "Test"},
	}

	for _, testCase := range testCases {
		name := fmt.Sprintf("should convert %s into %s", testCase.Input, testCase.Expected)

		t.Run(name, func(t *testing.T) {
			actual := ToPascalCase(testCase.Input)

			if actual != testCase.Expected {
				t.Errorf("got %s, expected %s", actual, testCase.Expected)
			}
		})
	}
}
