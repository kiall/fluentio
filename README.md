# Fluent IO
A Golang library for fluent that implements the standard writer io interface.

[![Go Reference](https://pkg.go.dev/badge/go.gllm.dev/fluentio.svg)](https://pkg.go.dev/go.gllm.dev/fluentio) [![Go Report Card](https://goreportcard.com/badge/go.gllm.dev/fluentio)](https://goreportcard.com/report/go.gllm.dev/fluentio)

## Installation
```bash
go get go.gllm.dev/fluentio
```

## Usage
```go
package main

import (
    "fmt"
    "go.gllm.dev/fluentio"
)

func main() {
    // Create a new fluentio.Writer
    f, err := fluentio.New(fluentio.WithBasicConfig("127.0.0.1", 24224))
    if err != nil {
        os.Exit(1) 
    }
    
    // Write to fluent
    log := `{"message": "Hello World!"}`
    i, err := f.Write([]byte(log))
    if err != nil {
        os.Exit(1)
    }
    
    fmt.Printf("Wrote %d bytes to fluent", i)  
}
```

## Configuration
The fluentio.Writer can be configured with the following options:
```go
// WithBasicConfig creates a new fluentio.Writer with a basic configuration
fluentio.New(fluentio.WithBasicConfig("127.0.0.1", 24224))

// WithFluentConfig creates a new fluentio.Writer with a standard fluent Config
// See github.com/fluent/fluent-logger-golang for more information
fluentio.New(fluentio.WithFluentConfig(fluent.Config{
    FluentPort: 24224,
    FluentHost: "127.0.0.1",
    Async:      true,
    BufferLimit: 1024,
    Timeout:    3 * time.Second,
    // ...
}))

// WithTag sets the tag for the fluentio.Writer
fluentio.New(fluentio.WithTag("my.tag"))
```

## Dependencies
This library depends on the following libraries:
- [Standard fluentd Go library](github.com/fluent/fluent-logger-golang)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://github.com/gllm-dev/fluentio/blob/master/LICENSE)
