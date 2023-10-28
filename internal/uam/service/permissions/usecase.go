package permissions

import "github.com/prongbang/uam-service/internal/pkg/casbinx"

type UseCase interface {
	Enforce(requests ...any) (bool, error)
	BatchEnforce(requests [][]any) ([]bool, error)
	IsRoot(roles []string, permission string) bool
	Enforces(roles []string, permission string, action string) bool
}

type useCase struct {
	CasbinXs casbinx.CasbinXs
}

func (p *useCase) Enforces(roles []string, permission string, action string) bool {
	perms := [][]any{}
	for _, r := range roles {
		perms = append(perms, []any{r, permission, action})
	}
	allowed, _ := p.BatchEnforce(perms)
	return IsAllowed(allowed...)
}

func (p *useCase) IsRoot(roles []string, permission string) bool {
	return p.Enforces(roles, permission, All)
}

func (p *useCase) Enforce(requests ...any) (bool, error) {
	return p.CasbinXs.EnforcerRbac.Enforce(requests...)
}

func (p *useCase) BatchEnforce(requests [][]any) ([]bool, error) {
	return p.CasbinXs.EnforcerRbac.BatchEnforce(requests)
}

func NewUseCase(casbinXs casbinx.CasbinXs) UseCase {
	return &useCase{
		CasbinXs: casbinXs,
	}
}
