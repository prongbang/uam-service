package interceptor

import (
	"context"
	"fmt"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/pkg/casbinx"
	"github.com/prongbang/uam-service/internal/pkg/token"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/pkg/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type JWEInterceptor interface {
	Intercept(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error)
}

type jweInterceptor struct {
	CasbinXs casbinx.CasbinXs
}

func (j *jweInterceptor) Intercept(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Perform your checks or operations here
	roles := []string{}
	if tk := token.Parse(req); tk != nil {
		payload, err := token.Verification(tk.Token)
		if err != nil {
			fmt.Println("[ERROR] JWE Interceptor", err)
			return nil, status.New(codes.Unauthenticated, core.TranslateCtx(ctx, localizations.CommonUnauthenticated)).Err()
		}
		roles = payload.Roles
	} else {
		roles = []string{role.RoleAnonymous}
	}

	// Check permissions by Roles
	allowed := false
	for _, r := range roles {
		result, e := j.CasbinXs.EnforcerGrpc.Enforce(strings.ToLower(r), info.FullMethod, casbinx.GRPC)
		if result && e == nil {
			allowed = true
			break
		}
	}

	// Call the next handler in the chain
	if allowed {
		return handler(ctx, req)
	}

	fmt.Println("[ERROR] JWE Interceptor Permission Denied")
	return nil, status.New(codes.PermissionDenied, core.TranslateCtx(ctx, localizations.CommonPermissionDenied)).Err()
}

func NewJWEInterceptor(casbinXs casbinx.CasbinXs) JWEInterceptor {
	return &jweInterceptor{
		CasbinXs: casbinXs,
	}
}
