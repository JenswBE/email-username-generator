package config

import "fmt"

type EnvVarIsEmptyError struct {
	envVar string
}

func (e EnvVarIsEmptyError) Error() string {
	return fmt.Sprintf("env var %s is mandatory but explicitly set to an empty value", e.envVar)
}
