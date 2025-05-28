# fiber-scalar

A simple scalar handler for [Go Fiber v2](https://github.com/gofiber/fiber).

## Features

-   Serves the swagger docs of your API in scalar fashion

## Installation

```bash
go get github.com/oSethoum/fiber-scalar
```

## Usage

```go
import scalar "github.com/oSethoum/fiber-scalar"
```

### Example

```go
package main

import (
	"github.com/gofiber/fiber/v2"
	scalar "github.com/oSethoum/fiber-scalar"
)

func main() {
    app := fiber.New()

    app.Get("/docs", scalar.Handler(&scalar.Options{
            SpecFile: "./docs/swagger.json",
            Layout:   scalar.LayoutClassic,
            Theme:    scalar.ThemeSolarized,
            DarkMode: true,
            // other options can go here
        }))

    app.Listen(":5000")
}
```

## License

MIT
