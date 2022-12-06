package api

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/configs"
)

var Enforcer *casbin.Enforcer

func NewEnforcer(config configs.Config) {
	casbinEnforcer, err := casbin.NewEnforcer("casbin.conf", config.CasbinPolicy)
	if err != nil {
		log.Fatal(err)
	}
	Enforcer = casbinEnforcer
}
