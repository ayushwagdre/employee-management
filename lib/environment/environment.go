package environment

import "os"

type Environment string

const (
	DevEnv Environment = "development"
)

type EnvProvider interface {
	CurrEnv() Environment
}

type osEnv struct {
	key string
}

func NewOsEnv() *osEnv {
	return &osEnv{key: "ENV"}
}

func (e Environment) String() string {
	return string(e)
}

func (e *osEnv) CurrEnv() Environment {
	env := Environment(os.Getenv(e.key))
	switch env {
	case DevEnv:
		return env
	default:
		panic("current os env is not a valid value")
	}
}
