# TinyStrconv

TinyStrconv is a lightweight Go module designed for converting and formatting various data types, optimized for minimal code size. It provides an alternative to the standard `strconv` library, making it suitable for embedded projects where code size is a critical constraint.

## Goals

- **Minimal Code Size**: Avoid importing bulky standard libraries and implement necessary functions from scratch to keep the code size minimal.
- **Comprehensive Functionality**: Provide essential functions for formatting and parsing integers, unsigned integers, floats, booleans, and strings.
- **Ease of Use**: Maintain a simple API for easy integration into projects.

## Requirements

- Go 1.20 or higher

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

// If you want to discard the error:
str, _ := tinystrconv.IntToString(123, 10)
println("Integer to String:", str)
```

#### String to Integer

```go
num, _ := tinystrconv.StringToInt("123", 10)
println("String to Integer:", num)
```

### Converting Unsigned Integers

#### Unsigned Integer to String

```go
str, _ := tinystrconv.UintToString(123, 10)
println("Unsigned Integer to String:", str)
```

#### String to Unsigned Integer

```go
num, _ := tinystrconv.StringToUint("123", 10)
println("String to Unsigned Integer:", num)
```

### Converting Floats

#### Float to String

```go
str, _ := tinystrconv.FloatToString(3.14159, 2)
println("Float to String:", str)
```

#### String to Float

```go
num, _ := tinystrconv.StringToFloat("3.14")
println("String to Float:", num)
```

### Converting Booleans

#### Boolean to String

```go
str := tinystrconv.BoolToString(true)
println("Boolean to String:", str)
```

#### String to Boolean

```go
val, _ := tinystrconv.StringToBool("true")
println("String to Boolean:", val)
```

### Quoting Strings

#### Quoting a String

```go
quotedStr := tinystrconv.Quote("Hello, world!")
println("Quoted String:", quotedStr)
```

## Code Size Difference

Using TinyStrconv results in significantly smaller code size compared to the standard library. When built with TinyGo for a Pico target, using string conversion functions resulted in a code size increase of approximately **1.5kB** with `tinystrconv`, compared to **30kB** with the standard library `strconv`.

## Contribution

Contributions are welcome! Please feel free to submit a pull request or open an issue on GitHub.

## License

TinyStrconv is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
