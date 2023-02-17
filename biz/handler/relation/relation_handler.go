// Code generated by hertz generator.

package relation

import (
	"context"

	relation "offer_tiktok/biz/model/social/relation"
	"offer_tiktok/biz/pack"
	"offer_tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	followerList_service "offer_tiktok/biz/service/relation/follower"
	friendList_service "offer_tiktok/biz/service/relation/friend"
)

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.DouyinRelationActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowList .
// @router /douyin/relation/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.DouyinRelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		// c.String(consts.StatusBadRequest, err.Error())
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	followerList, err := followerList_service.NewFollowerListService(ctx, c).GetFollowerList(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
	} else {
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
			UserList:   followerList,
		})
	}
}

// RelationFriendList .
// @router /douyin/relation/friend/list/ [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		// c.String(consts.StatusBadRequest, err.Error())
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	friendList, err := friendList_service.NewFriendListService(ctx, c).GetFriendList(&req)
	if err != nil {
		resp := pack.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
	} else {
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
			UserList:   friendList,
		})
	}
}
