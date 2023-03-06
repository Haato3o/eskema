namespace Example;

enum State {
    TEST_1,
    TEST_2,
    TEST_3
};

public record SimpleSchema(
    string value1,
    Dictionary<string, int> value2,
    bool value3
);

public record SimpleSchemaWithGenerics<T>(
    T value1,
    List<T> value2
);

public record ComplexSchema<TIn, TOut>(
    Dictionary<TIn, SimpleSchemaWithGenerics<TOut>>? value1,
    List<List<string>>? value2
);

