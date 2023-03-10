// Code generated by hertz generator.

package favorite

import (
	"context"

	favorite "offer_tiktok/biz/model/interact/favorite"
	"offer_tiktok/biz/pack"
	favorite_service "offer_tiktok/biz/service/favorite"
	"offer_tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteAction .
// @router /douyin/favortie/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)

	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, favorite.DouyinFavoriteActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	ok, err := favorite_service.NewFavoriteService(ctx, c).FavoriteAction(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, favorite.DouyinFavoriteActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	if !ok {
		resp := pack.BuildBaseResp(errno.FavoriteActionErr)
		c.JSON(consts.StatusOK, favorite.DouyinFavoriteActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, favorite.DouyinFavoriteActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	favoritelist, err := favorite_service.NewFavoriteService(ctx, c).GetFavoriteList(&req)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := favorite.DouyinFavoriteListResponse{
		VideoList:  favoritelist,
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	}
	c.JSON(consts.StatusOK, resp)
}
