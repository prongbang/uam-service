package casbinx

import "github.com/casbin/casbin/v2"

type CasbinXs struct {
	EnforcerRbac *casbin.Enforcer
	EnforcerRest *casbin.Enforcer
	EnforcerGrpc *casbin.Enforcer
}

func New(
	enforcerRbac *casbin.Enforcer,
	enforcerRest *casbin.Enforcer,
	enforceGrpc *casbin.Enforcer,
) CasbinXs {
	return CasbinXs{
		EnforcerRbac: enforcerRbac,
		EnforcerRest: enforcerRest,
		EnforcerGrpc: enforceGrpc,
	}
}
