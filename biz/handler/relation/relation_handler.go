// Code generated by hertz generator.

package relation

import (
	"context"
	"offer_tiktok/pkg/errno"
	"offer_tiktok/pkg/utils"

	service "offer_tiktok/biz/service/relation"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	relation "offer_tiktok/biz/model/social/relation"
)

// RelationAction users follow or unfollow other users.
//
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	_, err = service.NewRelationService(ctx, c).FollowAction(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, relation.DouyinRelationActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}

// RelationFollowList get list of all users followed by the logged_in user.
//
// @router /douyin/relation/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)

	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	FollowInfo, err := service.NewRelationService(ctx, c).GetFollowList(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	c.JSON(consts.StatusOK, relation.DouyinRelationFollowListResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserList:   FollowInfo,
	})
}

// RelationFollowerList get the list of all followers following the logged-in user.
//
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		// c.String(consts.StatusBadRequest, err.Error())
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	followerList, err := service.NewRelationService(ctx, c).GetFollowerList(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
	} else {
		c.JSON(consts.StatusOK, relation.DouyinRelationFollowerListResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
			UserList:   followerList,
		})
	}
}

// RelationFriendList get A list of all friends who follow the logged_in user.
//
// @router /douyin/relation/friend/list/ [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		// c.String(consts.StatusBadRequest, err.Error())
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
		return
	}

	friendList, err := service.NewRelationService(ctx, c).GetFriendList(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			UserList:   nil,
		})
	} else {
		c.JSON(consts.StatusOK, relation.DouyinRelationFriendListResponse{
			StatusCode: errno.SuccessCode,
			StatusMsg:  errno.SuccessMsg,
			UserList:   friendList,
		})
	}
}
