// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: interview.proto

package server

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/rpc/internal/logic"
	"Ai-HireSphere/application/interview/interfaces/rpc/internal/svc"
	"Ai-HireSphere/common/call/interview"
)

type InterviewServer struct {
	svcCtx *svc.ServiceContext
	interview.UnsafeInterviewRpcServer
}

func NewInterviewServer(svcCtx *svc.ServiceContext) *InterviewServer {
	return &InterviewServer{
		svcCtx: svcCtx,
	}
}

func (s *InterviewServer) Ping(ctx context.Context, in *interview.Request) (*interview.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
