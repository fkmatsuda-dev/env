# env
The env project is a library written in Go that provides functions to read values from environment variables. The library supports the following variable types:

* Int
* Duration
* Float64
* String

## All functions have two parameters:

* The first parameter is the name of the environment variable to be read.
* The second parameter is the default value to be returned if the variable cannot be read.
* The return type of the function will match the type of the default value provided.

## Installation
You can install the env library using the following command:

```bash
go get github.com/yourusername/env
```

## Usage
To use the env library, import the relevant function for the type of environment variable you wish to read.

```go
import (
    "github.com/yourusername/env"
)

// Reading an integer environment variable
myInt := env.ReadInt("MY_INT_VAR", 10)

// Reading a duration environment variable
myDuration := env.ReadDuration("MY_DURATION_VAR", time.Second * 60)

// Reading a float environment variable
myFloat := env.ReadFloat64("MY_FLOAT_VAR", 3.14)

// Reading a string environment variable
myString := env.ReadString("MY_STRING_VAR", "default_value")
```

If the environment variable cannot be read, the function will return the default value provided.

## License
This project is licensed under the MIT License. See the LICENSE file for more information.

## Contributing
If you would like to contribute to the env library, please submit a pull request on the GitHub repository.