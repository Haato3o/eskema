package example

type State int
const (
    TEST_1 State = iota
    TEST_2
    TEST_3
)

type SimpleSchema struct {
    value1 string
    value2 map[string]int32
    value3 bool
}

type SimpleSchemaWithGenerics[T any] struct {
    value1 T
    value2 []T
}

type ComplexSchema[TIn any, TOut any] struct {
    value1 *map[TIn]SimpleSchemaWithGenerics[TOut]
    value2 *[][]string
}

