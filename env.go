/*
 * © $year $author
 *
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

// Int returns the value of the environment variable with the given name as an integer.
// If the environment variable is not set or cannot be parsed as an integer, the default value is returned.
func Int(name string, defaultValue int) int {
	// load the environment variable
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	// parse the environment variable
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// Float64 returns the value of the environment variable with the given name as a float64.
// If the environment variable is not set or cannot be parsed as a float64, the default value is returned.
func Float64(name string, defaultValue float64) float64 {
	// load the environment variable
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	// parse the environment variable
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return floatValue
}

// Duration returns the value of the environment variable with the given name as a time.Duration.
// If the environment variable is not set or cannot be parsed as a time.Duration, the default value is returned.
func Duration(name string, defaultValue time.Duration) time.Duration {
	// load the environment variable
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	// parse the environment variable
	durationValue, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return durationValue
}

// String returns the value of the environment variable with the given name as a string.
// If the environment variable is not set, the default value is returned.
func String(name string, defaultValue string) string {
	// load the environment variable
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}
