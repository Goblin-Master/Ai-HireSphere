// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/data",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/captcha/send",
				Handler: CaptchaSendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/captcha/verify",
				Handler: CaptchaVerifyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/base"),
	)
}
