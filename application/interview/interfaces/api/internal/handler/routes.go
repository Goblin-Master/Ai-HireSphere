// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	resume "Ai-HireSphere/application/interview/interfaces/api/internal/handler/resume"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: InterviewHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 上传简历
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: resume.UploadResumeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/resume"),
	)
}
