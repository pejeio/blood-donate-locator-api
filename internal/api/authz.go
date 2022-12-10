package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
)

func NewEnforcer(config *configs.Config) (*casbin.Enforcer, error) {
	casbinEnforcer, err := casbin.NewEnforcer(config.CasbinConfFile, config.CasbinPolicyFile)
	if err != nil {
		return nil, err
	}
	return casbinEnforcer, nil
}
