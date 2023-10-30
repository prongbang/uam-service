package permissions

import "github.com/prongbang/uam-service/internal/pkg/casbinx"

type UseCase interface {
	RbacEnforce(requests ...any) (bool, error)
	RbacBatchEnforce(requests [][]any) ([]bool, error)
	RbacIsRoot(roles []string, permission string) bool
	RbacEnforces(roles []string, permission string, action string) bool
	RestEnforce(requests ...any) (bool, error)
	RestBatchEnforce(requests [][]any) ([]bool, error)
	RestIsRoot(roles []string, permission string) bool
	RestEnforces(roles []string, permission string, action string) bool
}

type useCase struct {
	CasbinXs casbinx.CasbinXs
}

func (p *useCase) RbacEnforces(roles []string, permission string, action string) bool {
	perms := [][]any{}
	for _, r := range roles {
		perms = append(perms, []any{r, permission, action})
	}
	allowed, _ := p.RbacBatchEnforce(perms)
	return IsAllowed(allowed...)
}

func (p *useCase) RbacIsRoot(roles []string, permission string) bool {
	return p.RbacEnforces(roles, permission, All)
}

func (p *useCase) RbacEnforce(requests ...any) (bool, error) {
	return p.CasbinXs.EnforcerRbac.Enforce(requests...)
}

func (p *useCase) RbacBatchEnforce(requests [][]any) ([]bool, error) {
	return p.CasbinXs.EnforcerRbac.BatchEnforce(requests)
}

func (p *useCase) RestEnforce(requests ...any) (bool, error) {
	return p.CasbinXs.EnforcerRest.Enforce(requests...)
}

func (p *useCase) RestBatchEnforce(requests [][]any) ([]bool, error) {
	return p.CasbinXs.EnforcerRest.BatchEnforce(requests)
}

func (p *useCase) RestIsRoot(roles []string, permission string) bool {
	return p.RestEnforces(roles, permission, All)
}

func (p *useCase) RestEnforces(roles []string, permission string, action string) bool {
	perms := [][]any{}
	for _, r := range roles {
		perms = append(perms, []any{r, permission, action})
	}
	allowed, _ := p.RestBatchEnforce(perms)
	return IsAllowed(allowed...)
}

func NewUseCase(casbinXs casbinx.CasbinXs) UseCase {
	return &useCase{
		CasbinXs: casbinXs,
	}
}
