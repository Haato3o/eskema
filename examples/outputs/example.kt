package com.example

enum class State {
    TEST_1,
    TEST_2,
    TEST_3
}

data class SimpleSchema(
    val value1: String,
    val value2: Map<String, Int>,
    val value3: Boolean
)

data class SimpleSchemaWithGenerics<T>(
    val value1: T,
    val value2: List<T>
)

data class ComplexSchema<TIn, TOut>(
    val value1: Map<TIn, SimpleSchemaWithGenerics<TOut>>?,
    val value2: List<List<String>>?
)

