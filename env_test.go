package env_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/fkmatsuda-dev/env"
)

func TestInt(t *testing.T) {
	t.Run("when the environment variable is set", func(t *testing.T) {
		value := 42
		t.Setenv("MY_VAR", strconv.Itoa(value))

		result := env.Int("MY_VAR", 0)
		if result != value {
			t.Errorf("Int() returned %v, expected %v", result, value)
		}
	})

	t.Run("when the environment variable is not set", func(t *testing.T) {
		value := 0
		save := os.Getenv("MY_VAR")
		err := os.Unsetenv("MY_VAR")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			err := os.Setenv("MY_VAR", save)
			if err != nil {
				t.Fatal(err)
			}
		}()

		result := env.Int("MY_VAR", value)
		if result != value {
			t.Errorf("Int() returned %v, expected %v", result, value)
		}
	})

	t.Run("when the environment variable is not an integer", func(t *testing.T) {
		defaultValue := 42
		t.Setenv("MY_VAR", "not an integer")

		result := env.Int("MY_VAR", defaultValue)
		if result != defaultValue {
			t.Errorf("Int() returned %v, expected %v", result, defaultValue)
		}
	})
}

func TestFloat64(t *testing.T) {
	t.Run("when the environment variable is set", func(t *testing.T) {
		value := 3.14
		t.Setenv("MY_VAR", strconv.FormatFloat(value, 'f', -1, 64))

		result := env.Float64("MY_VAR", 0)
		if result != value {
			t.Errorf("Float64() returned %v, expected %v", result, value)
		}
	})

	t.Run("when the environment variable is not set", func(t *testing.T) {
		value := 0.0
		save := os.Getenv("MY_VAR")
		err := os.Unsetenv("MY_VAR")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			err := os.Setenv("MY_VAR", save)
			if err != nil {
				t.Fatal(err)
			}
		}()

		result := env.Float64("MY_VAR", value)
		if result != value {
			t.Errorf("Float64() returned %v, expected %v", result, value)
		}
	})

	t.Run("when the environment variable is not a float", func(t *testing.T) {
		defaultValue := 3.14
		t.Setenv("MY_VAR", "not a float")

		result := env.Float64("MY_VAR", defaultValue)
		if result != defaultValue {
			t.Errorf("Float64() returned %v, expected %v", result, defaultValue)
		}
	})
}

func TestDuration(t *testing.T) {
	t.Run("when the environment variable is set", func(t *testing.T) {
		value, _ := time.ParseDuration("1h30m")
		t.Setenv("MY_VAR", value.String())

		result := env.Duration("MY_VAR", 0)
		if result != value {
			t.Errorf("Duration() returned %v, expected %v", result, value)
		}
	})

	t.Run("when the environment variable is not set", func(t *testing.T) {
		value, _ := time.ParseDuration("1m")
		save := os.Getenv("MY_VAR")
		err := os.Unsetenv("MY_VAR")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			err := os.Setenv("MY_VAR", save)
			if err != nil {
				t.Fatal(err)
			}
		}()

		result := env.Duration("MY_VAR", value)
		if result != value {
			t.Errorf("Duration() returned %v, expected %v", result, value)
		}
	})

	t.Run("when the environment variable is not a duration", func(t *testing.T) {
		defaultValue, _ := time.ParseDuration("1h30m")
		t.Setenv("MY_VAR", "not a duration")

		result := env.Duration("MY_VAR", defaultValue)
		if result != defaultValue {
			t.Errorf("Duration() returned %v, expected %v", result, defaultValue)
		}
	})
}

func TestString(t *testing.T) {
	t.Run("when the environment variable is set", func(t *testing.T) {
		value := "hello"
		t.Setenv("MY_VAR", value)

		result := env.String("MY_VAR", "")
		if result != value {
			t.Errorf("String() returned %v, expected %v", result, value)
		}
	})

	t.Run("when the environment variable is not set", func(t *testing.T) {
		value := ""
		save := os.Getenv("MY_VAR")
		err := os.Unsetenv("MY_VAR")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			err := os.Setenv("MY_VAR", save)
			if err != nil {
				t.Fatal(err)
			}
		}()

		result := env.String("MY_VAR", value)
		if result != value {
			t.Errorf("String() returned %v, expected %v", result, value)
		}
	})
}
