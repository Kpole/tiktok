// Code generated by hertz generator.

package user

import (
	"context"

	user "offer_tiktok/biz/model/basic/user"
	"offer_tiktok/biz/pack"
	service "offer_tiktok/biz/service/user"
	"offer_tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserRegister .
// @router /douyin/user/register/  [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	hlog.CtxInfof(ctx, "OK")
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	token, user_id, err := service.NewUserRegisterService(ctx).UserRegister(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.DouyinUserRegisterResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Token:      token,
		UserId:     user_id,
	})
}

// UserLogin .
// @router /douyin/user/login/  [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserLoginResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	token, user_id, err := service.NewUserLoginService(ctx).UserLogin(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.DouyinUserLoginResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, user.DouyinUserLoginResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Token:      token,
		UserId:     user_id,
	})
}

// User .
// @router /douyin/user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.DouyinUserResponse)

	c.JSON(consts.StatusOK, resp)
}
