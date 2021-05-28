# Generate Passphrase for Go

[![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/aldy505/go-generate-passphrase?include_prereleases)](https://github.com/aldy505/go-generate-passphrase/releases) [![Go Reference](https://pkg.go.dev/badge/github.com/aldy505/go-generate-passphrase.svg)](https://pkg.go.dev/github.com/aldy505/go-generate-passphrase) [![GitHub](https://img.shields.io/github/license/aldy505/go-generate-passphrase)](https://github.com/aldy505/go-generate-passphrase/blob/master/LICENSE) [![codecov](https://codecov.io/gh/aldy505/go-generate-passphrase/branch/master/graph/badge.svg?token=DV7VhMgdAp)](https://codecov.io/gh/aldy505/go-generate-passphrase) [![CodeFactor](https://www.codefactor.io/repository/github/aldy505/go-generate-passphrase/badge)](https://www.codefactor.io/repository/github/aldy505/go-generate-passphrase) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/6bb40d1557574b96b8bd478b179c9803)](https://www.codacy.com/gh/aldy505/go-generate-passphrase/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=aldy505/go-generate-passphrase&amp;utm_campaign=Badge_Grade) [![Build test](https://github.com/aldy505/go-generate-passphrase/actions/workflows/build.yml/badge.svg)](https://github.com/aldy505/go-generate-passphrase/actions/workflows/build.yml) [![Test and coverage](https://github.com/aldy505/go-generate-passphrase/actions/workflows/codecov.yml/badge.svg)](https://github.com/aldy505/go-generate-passphrase/actions/workflows/codecov.yml)

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
  pass, err := passphrase.Generate(passphrase.Options{})
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Println(pass)

  // To use with available options. See available options below.
  pass, err := passphrase.Generate(passphrase.Options{
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

Yes!

## License

MIT License

Copyright (c) 2021-present Reinaldy Rafli

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
