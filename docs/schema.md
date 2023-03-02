### Eskema - Schemas

Eskema allows you to define a single schema for data transfer between services or apps and have it automatically generate a class on your preferred language based on the schema.

### Defining a schema

To define a schema, all you need to do is use the `schema` keyword and give it a name, like you would do in any other programming language.

```
enum MyStateEnum
{
    FIRST_VALUE,
    SECOND_VALUE,
    THIRD_VALUE
};

schema MySubSchema
{
    state: MyStateEnum,
    field: Int32,
    optionalField: String?,
    customList: Array<Int32>
};

schema MySchema
{
    firstField: Float,
    createdAt: TimeStamp, 
    subSchema: MySubSchema?
};
```