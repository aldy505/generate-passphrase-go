# Generate Passphrase for Go

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/aldy505/go-generate-passphrase) ![GitHub branch checks state](https://img.shields.io/github/checks-status/aldy505/go-generate-passphrase/master) ![Codecov](https://img.shields.io/codecov/c/github/aldy505/go-generate-passphrase) ![GitHub](https://img.shields.io/github/license/aldy505/go-generate-passphrase)  

Ported from [Generate Passphrase for Nodejs](https://github.com/aldy505/generate-passphrase).

## Installation

Make sure you have Go v1.14 or higher.

```bash
go get github.com/aldy505/go-generate-passphrase
```

## Usage

```go
import (
  "fmt"
  passphrase "github.com/aldy505/go-generate-passphrase"
)

func main() {
  // Generate a single passphrase
  pass, err := passphrase.Generate(&generateOptions{})
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(pass)

  // Generate multiple passphrase
  multiPass, err := passphrase.GenerateMultiple(5, &generateOptions{})
  if err != nil {
    fmt.Println(err.Error())
  }
  for i := 0; i < len(multiPass); i++ {
    fmt.Println(multiPass[i])
  }
}
```

## Options

| Key | Type | Default |
| --- | --- | --- |
| Length | `Int` | `4` |
| Separator | `String` | `-` |
| Numbers | `Bool` | `false` |
| Uppercase | `Bool` | `false` |
| Titlecase | `Bool` | `false` |
| Pattern | `String` | `Nil` |

A few things to note:

* Uppercase is more prioritized than titlecase. So if you have both options set to true, it will be words full of uppercase.
* Pattern option is more prioritized than length, because you've set the passphrase pattern, hence the module is using the length from your pattern.

## Contributing

yes. please.

## License

MIT