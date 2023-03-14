public enum State: String, Decodable, Equatable {
    case test1 = "TEST_1"
    case test2 = "TEST_2"
    case test3 = "TEST_3"
}

public struct SimpleSchema: Decodable, Equatable {
    let value1: String
    let value2: [String : Int]
    let value3: Boolean
}

public struct SimpleSchemaWithGenerics: Decodable, Equatable {
    let value1: T
    let value2: [T]
}

public struct ComplexSchema: Decodable, Equatable {
    let value1: [TIn : SimpleSchemaWithGenerics<TOut>]?
    let value2: [[String]]?

    public init(value1: [TIn : SimpleSchemaWithGenerics<TOut>]? = nil, value2: [[String]]? = nil) {
        self.value1 = value1
        self.value2 = value2
    }
}

