package authz

import (
	"github.com/casbin/casbin/v2"
	log "github.com/sirupsen/logrus"
)

func NewEnforcer(modelPath string, policyPath string) *casbin.Enforcer {
	casbinEnforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatal(err)
	}
	return casbinEnforcer
}
