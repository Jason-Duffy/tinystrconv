# TinyStrconv

TinyStrconv is a lightweight Go module designed for converting and formatting various data types, optimized for minimal code size. It provides an alternative to the standard `strconv` and `fmt` libraries, making it suitable for embedded projects where code size is a critical constraint.

## Goals

- **Minimal Code Size**: Avoid importing bulky standard libraries and implement necessary functions from scratch to keep the code size minimal.
- **Comprehensive Functionality**: Provide essential functions for formatting and parsing integers, floats, booleans, and strings.
- **Ease of Use**: Maintain a simple API for easy integration into projects.

## Requirements

- Go 1.16 or higher

## Achievements

- Implemented conversion functions for various data types without using standard libraries like `strconv`, `fmt`, `strings`, or `unicode`.
- Comprehensive test coverage to ensure correctness and reliability.
- Optimized for embedded systems and scenarios where code size is crucial.
- Clear, concise function names (doing away with shorthand like Itoa())

## Installation

To install TinyStrconv, use the following command:

```sh
go get github.com/Jason-Duffy/tinystrconv
```

## Usage

### Importing the Package

```go
import "github.com/Jason-Duffy/tinystrconv"
```

### Converting Integers

#### Integer to String

```go
str, err := tinystrconv.IntToString(123, 10)
if err != nil {
    println("Error:", err.Error())
} else {
    println("Integer to String:", str)
}
```

#### String to Integer

```go
num, err := tinystrconv.StringToInt("123", 10)
if err != nil {
    println("Error:", err.Error())
} else {
    println("String to Integer:", num)
}
```

### Converting Floats

#### Float to String

```go
str, err := tinystrconv.FloatToString(3.14159, 2)
if err != nil {
    println("Error:", err.Error())
} else {
    println("Float to String:", str)
}
```

#### String to Float

```go
num, err := tinystrconv.StringToFloat("3.14")
if err != nil {
    println("Error:", err.Error())
} else {
    println("String to Float:", num)
}
```

### Converting Booleans

#### Boolean to String

```go
str := tinystrconv.BoolToString(true)
println("Boolean to String:", str)
```

#### String to Boolean

```go
val, err := tinystrconv.StringToBool("true")
if err != nil {
    println("Error:", err.Error())
} else {
    println("String to Boolean:", val)
}
```

### Quoting Strings

#### Quoting a String

```go
quotedStr := tinystrconv.Quote("Hello, world!")
println("Quoted String:", quotedStr)
```

### Formatting Strings

#### Using Format Specifiers

```go
formattedStr, err := tinystrconv.Format("Hello, %s! Value: %d", "world", 42)
if err != nil {
    println("Error:", err.Error())
} else {
    println("Formatted String:", formattedStr)
}
```

## Contribution

Contributions are welcome! Please feel free to submit a pull request or open an issue on GitHub.

## License

TinyStrconv is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

Feel free to modify this draft according to your needs or add any additional information you find relevant.
