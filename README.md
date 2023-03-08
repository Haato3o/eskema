![eskema-banner](https://user-images.githubusercontent.com/35552782/223580390-63e2fb75-4baa-463f-92e2-9d43556e9375.png)

## About

Eskema is a command-line tool that generates class/data class for other programming languages based on a schema. This is useful for when you have multiple applications in different programming languages that can communicate between them with well defined contracts.

## Installation

TODO

### Usage

To use Eskema, you need to provide a schema file that defines the structure of the class/data class you want to generate. Here's an example:

```
enum Status
{
  ONLINE,
  BUSY,
  AWAY,
  OFFLINE
};

schema User 
{
  userId: Int64,
  name: String,
  email: String,
  status: Status
  friends: Array<Long>?,
};
```

To generate a class from the previous schema, you can run the following command:

```sh
eskema --filename example.skm --language kotlin --output example.kt
```

This will generate the following code:

```kt
enum Status {
  ONLINE,
  BUSY,
  AWAY,
  OFFLINE
}

data class User(
  val userId: Long,
  val name: String,
  val email: String,
  val status: Status,
  val friends: List<Long>?
)
```

Eskema currently supports the following languages:

- C#
- Kotlin
- GoLang

## Contributing
If you want to contribute to Eskema, please read our contributing guidelines before submitting a pull request.

## License
Eskema is licensed under the MIT License. See [LICENSE](https://github.com/Haato3o/eskema/blob/main/LICENSE) for more information.
