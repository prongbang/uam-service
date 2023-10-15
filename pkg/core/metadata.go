package core

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/localizations"
	"google.golang.org/grpc/metadata"
)

func Metadata(context context.Context) metadata.MD {
	if md, ok := metadata.FromIncomingContext(context); ok {
		return md
	}
	return metadata.MD{}
}

func AcceptLanguage(ctx context.Context) string {
	locale := Metadata(ctx)[fiber.HeaderAcceptLanguage]
	if len(locale) > 0 {
		return locale[0]
	}
	return localizations.En
}

func TranslateCtx(ctx context.Context, key string) string {
	locale := AcceptLanguage(ctx)
	return localizations.Translate(locale, key)
}
