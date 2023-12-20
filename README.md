# Sentinel Environment Variable Plugin

## Overview

This repository contains a simple Go package named `env` that serves as a plugin for interacting with environment variables. It is designed to be used with the Sentinel framework.

## Installation

To use this package, you need to have Go installed. If you haven't set up your Go environment, please visit [Go Installation Guide](https://golang.org/doc/install).

Once Go is installed, you can get the package by running:

```bash
go get -u github.com/your-username/your-repository/env
```

## Usage

### Initializing the Plugin

```go
import (
	"github.com/your-username/your-repository/env"
)
```

To create an instance of the plugin:

```go
plugin := env.New()
```

### Available Functions

#### 1. `get(key string) (interface{}, error)`

This function retrieves the value of the specified environment variable.

Example:

```go
value, err := plugin.Func("get")("YOUR_VARIABLE_NAME")
if err != nil {
    // Handle error
} else {
    // Use the retrieved value
}
```

#### 2. `list() (map[string]string, error)`

This function returns a map of all environment variables.

Example:

```go
allEnv, err := plugin.Func("list")()
if err != nil {
    // Handle error
} else {
    // Use the map of environment variables
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- [HashiCorp Sentinel SDK](https://github.com/hashicorp/sentinel-sdk)

---