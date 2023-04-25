/*
 * © 2023 fkmatsuda

 * This file is licensed under the terms of the MIT license. Permission is hereby
 * granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without
 * restriction, including without limitation the rights to use, copy, modify,
 * merge, publish, and/or distribute copies of the Software, and to permit persons
 * to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS," WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */

package env

import (
	"os"
	"strconv"
	"time"
)

// ChkString returns the value of the environment variable with the given name as a string and true if it is set, otherwise it returns empty string and false.
func ChkString(name string) (string, bool) {
	value := os.Getenv(name)
	return value, value != ""
}

// ChkInt returns the value of the environment variable with the given name as an integer and true if it is set, otherwise it returns 0 and false.
func ChkInt(name string) (int, bool, error) {
	// load the environment variable as a string
	value, ok := ChkString(name)
	if !ok {
		return 0, false, nil
	}
	// parse the environment variable as an integer
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, true, err
	}
	return intValue, true, nil
}

// ChkFloat64 returns the value of the environment variable with the given name as a float64 and true if it is set, otherwise it returns 0.0 and false.
func ChkFloat64(name string) (float64, bool, error) {
	// load the environment variable as a string
	value, ok := ChkString(name)
	if !ok {
		return 0.0, false, nil
	}
	// parse the environment variable as a float64
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0, true, err
	}
	return floatValue, true, nil
}

// ChkDuration returns the value of the environment variable with the given name as a time.Duration and true if it is set, otherwise it returns 0 and false.
func ChkDuration(name string) (time.Duration, bool, error) {
	// load the environment variable as a string
	value, ok := ChkString(name)
	if !ok {
		return 0, false, nil
	}
	// parse the environment variable as a time.Duration
	durationValue, err := time.ParseDuration(value)
	if err != nil {
		return 0, true, err
	}
	return durationValue, true, nil
}

// Int returns the value of the environment variable with the given name as an integer.
// If the environment variable is not set or cannot be parsed as an integer, the default value is returned.
func Int(name string, defaultValue int) (int, error) {
	// load the environment variable using the ChkInt function
	intValue, ok, err := ChkInt(name)
	if err != nil {
		return defaultValue, err
	}
	if !ok {
		return defaultValue, nil
	}
	return intValue, nil
}

// Float64 returns the value of the environment variable with the given name as a float64.
// If the environment variable is not set or cannot be parsed as a float64, the default value is returned.
func Float64(name string, defaultValue float64) (float64, error) {
	// load the environment variable using the ChkFloat64 function
	floatValue, ok, err := ChkFloat64(name)
	if err != nil {
		return defaultValue, err
	}
	if !ok {
		return defaultValue, nil
	}
	return floatValue, nil
}

// Duration returns the value of the environment variable with the given name as a time.Duration.
// If the environment variable is not set or cannot be parsed as a time.Duration, the default value is returned.
func Duration(name string, defaultValue time.Duration) (time.Duration, error) {
	// load the environment variable using the ChkDuration function
	durationValue, ok, err := ChkDuration(name)
	if err != nil {
		return defaultValue, err
	}
	if !ok {
		return defaultValue, nil
	}
	return durationValue, nil
}

// String returns the value of the environment variable with the given name as a string.
// If the environment variable is not set, the default value is returned.
func String(name string, defaultValue string) string {
	// load the environment variable using the ChkString function
	value, ok := ChkString(name)
	if !ok {
		return defaultValue
	}
	return value
}
