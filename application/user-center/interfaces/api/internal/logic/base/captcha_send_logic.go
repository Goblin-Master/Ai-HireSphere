package base

import (
	"Ai-HireSphere/common/model/enums"
	"context"

	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaSendLogic {
	return &CaptchaSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaSendLogic) CaptchaSend(req *types.CaptchaSendReq) error {
	err := l.svcCtx.BaseApp.CaptchaSend(l.ctx, enums.CaptchaWayType(req.Way), enums.CaptchaMethodType(req.Method), req.Target)

	return err
}
