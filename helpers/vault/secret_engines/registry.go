package secret_engines

import (
	"fmt"

	"github.com/debu99/cicd-runner/helpers/vault"
	"github.com/debu99/cicd-runner/helpers/vault/internal/registry"
)

type Factory func(client vault.Client, path string) vault.SecretEngine

var factoriesRegistry = registry.New("secret engine")

func MustRegisterFactory(engineName string, factory Factory) {
	err := factoriesRegistry.Register(engineName, factory)
	if err != nil {
		panic(fmt.Sprintf("registering factory: %v", err))
	}
}

func GetFactory(engineName string) (Factory, error) {
	factory, err := factoriesRegistry.Get(engineName)
	if err != nil {
		return nil, err
	}

	return factory.(Factory), nil
}
