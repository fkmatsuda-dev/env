# env
The env project is a library written in Go that provides functions to read values from environment variables. The library supports the following variable types:

* Int
* Duration
* Float64
* String

### This functions have two parameters:

* The first parameter is the name of the environment variable to be read.
* The second parameter is the default value to be returned if the variable cannot be read.

### return
* The return type of the function will match the type of the default value provided.

Functions Int, Duration and Float64 will return an error if a conversion error occurs.

Duration variables can be specified in the following formats:
* 1h
* 1h30m
* 1m
* 1s
* 1ms
* 1us
* 1ns

The above values are just examples, you need to adjust the values as per your need.

## Installation
You can install the env library using the following command:

```bash
go get github.com/fkmatsuda-dev/env
```

## Usage
To use the env library, import the relevant function for the type of environment variable you wish to read.

```go
import (
    "github.com/fkmatsuda-dev/env"
)

// Reading an integer environment variable
myInt, err := env.Int("MY_INT_VAR", 10)
if err != nil {
// handle error
}

// Reading a duration environment variable
myDuration, err := env.Duration("MY_DURATION_VAR", time.Second * 60)
if err != nil {
// handle error
}

// Reading a float environment variable
myFloat, err := env.Float64("MY_FLOAT_VAR", 3.14)
if err != nil {
// handle error
}

// Reading a string environment variable
myString := env.String("MY_STRING_VAR", "default_value")
```

If the environment variable cannot be read, the function will return the default value provided.

## Adicional functions
There are alternative versions of these functions without a default value but with an additional Boolean return to tell whether or not the variable can be read.

* ChkInt
* ChkDuration
* ChkFloat64
* ChkString

```go
import (
    "github.com/fkmatsuda-dev/env"
)

// Reading an integer environment variable
myInt, ok, err := env.ChkInt("MY_INT_VAR")
if err != nil {
// handle error
}
if !ok {
// handle missing environment variable
}

// Reading a duration environment variable
myDuration, ok, err := env.ChkDuration("MY_DURATION_VAR")
if err != nil {
// handle error
}
if !ok {
// handle missing environment variable
}

// Reading a float environment variable
myFloat, ok, err := env.ChkFloat64("MY_FLOAT_VAR")
if err != nil {
// handle error
}
if !ok {
// handle missing environment variable
}

// Reading a string environment variable
myString, ok := env.ChkString("MY_STRING_VAR")
if !ok {
// handle missing environment variable
}
```


## License
This project is licensed under the MIT License. See the LICENSE file for more information.

## Contributing
If you would like to contribute to the env library, feel free to send a pull request or create an issue on GitHub. Your feedback is very important for us to improve the project together.