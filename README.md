# Generate Passphrase for Go

[![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/aldy505/go-generate-passphrase?include_prereleases)](https://github.com/aldy505/go-generate-passphrase/releases) [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/aldy505/go-generate-passphrase/Build%20test/master)](https://github.com/aldy505/go-generate-passphrase/actions) [![codecov](https://codecov.io/gh/aldy505/go-generate-passphrase/branch/master/graph/badge.svg?token=DV7VhMgdAp)](https://codecov.io/gh/aldy505/go-generate-passphrase) [![GitHub](https://img.shields.io/github/license/aldy505/go-generate-passphrase)](https://github.com/aldy505/go-generate-passphrase/blob/master/LICENSE)

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
  pass, err := passphrase.Generate(&passphrase.Options{})
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(pass)

  // To use with available options. See available options below.
  pass, err := passphrase.Generate(&passphrase.Options{
    Length:  6,
    Numbers: true,
  })
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(pass)

  // Generate multiple passphrase
  multiPass, err := passphrase.GenerateMultiple(5, &passphrase.Options{})
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
