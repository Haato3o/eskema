## Metadata

Eskema supports annotations, this is useful when you want to control exactly how the emitters should behave for individual fields.

```
schema Example
{
    @PropertyCase("snake_case")
    MyField String,
    
    @PropertyCase("PascalCase")
    MyOtherField String
};
```

By running the language's emitter, they should respect the annotations. The output for the previous example when using the Kotlin emitter would be the following:

```kt
data class Example(
    @SerializedName("my_field") val myField: String,
    @SerializedName("MyOtherField") val myOtherField: String 
)
```