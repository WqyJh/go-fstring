# go-fstring [![GoDoc][doc-img]][doc]

The fstring package in Go provides a simple yet powerful way to interpolate strings, inspired by Python's f-string syntax.

## Usage

Import the go-fstring package into your Go file:

```go
import "github.com/WqyJh/go-fstring"
```

You can then use the Format function to interpolate strings. The Format function takes a template string and a map of values to interpolate into the template:

```go
template := "Hello, {name}! You have {count} unread messages."
values := map[string]any{"name": "Alice", "count": 10}
result, err := fstring.Format(template, values)
if err != nil {
    // handle error
}
fmt.Println(result) // Output: Hello, Alice! You have 10 unread messages.
```

## License

Released under the [MIT License](LICENSE).

[doc-img]: https://godoc.org/github.com/WqyJh/go-fstring?status.svg
[doc]: https://godoc.org/github.com/WqyJh/go-fstring
