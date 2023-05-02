package handler

import (
	comment "comment-srv/proto"
	"config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"net/http"
	"strconv"
)

var (
	serviceComment        = config.Namespace + config.ServiceNameComment
	endpointHotComment    = "Comment.HotComment"
	endpointMakeComment   = "Comment.MakeComment"
	endpointUpNumComment  = "Comment.UpNumComment"
	endpointMyComments    = "Comment.MyComments"
	endpointDeleteComment = "Comment.DeleteComment"
)

// @Summary 删除评论
// @Tags 评论中心
// @Description 删除评论
// @Accept json
// @Produce json
// @Param commentID query int true "评论id"
// @Success 1 {object} model.Response "success"
// @Router /comment/deleteComment [post]
func DeleteComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Query("commentID"))
	grpcReq := &comment.DeleteCommentReq{
		CommentID: int64(commentID),
	}
	grpcRsp := &comment.DeleteCommentRsp{}

	req := client.NewRequest(serviceComment, endpointDeleteComment, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 我的评论
// @Tags 评论中心
// @Description 我的评论
// @Accept json
// @Produce json
// @Param userId query int true "用户id"
// @Success 1 {object} model.Response "success"
// @Router /comment/myComments [post]
func MyComments(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("userId"))
	grpcReq := &comment.MyCommentsReq{
		UserId: int64(userId),
	}
	grpcRsp := &comment.MyCommentsRsp{}

	req := client.NewRequest(serviceComment, endpointMyComments, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 评论点赞
// @Tags 评论中心
// @Description 评论点赞
// @Accept json
// @Produce json
// @Param commentID query int true "评论id"
// @Success 1 {object} model.Response "success"
// @Router /comment/upNumComment [post]
func UpNumComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Query("commentID"))
	grpcReq := &comment.UpNumCommentReq{
		CommentID: int64(commentID),
	}
	grpcRsp := &comment.UpNumCommentRsp{}

	req := client.NewRequest(serviceComment, endpointUpNumComment, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 评论
// @Tags 评论中心
// @Description 评论
// @Accept json
// @Produce json
// @Param movieId query int true "影片id"
// @Param userId query int true "用户id"
// @Param title query string true "标题"
// @Param headImg query string true "内容"
// @Param nickname query string true "昵称"
// @Success 1 {object} model.Response "success"
// @Router /comment/makeComment [post]
func MakeComment(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")
	headImg := c.Query("headImg")
	nickname := c.Query("nickname")
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	userId, _ := strconv.Atoi(c.Query("userId"))
	grpcReq := &comment.MakeCommentReq{
		MovieId:  int64(movieId),
		UserId:   int64(userId),
		Title:    title,
		HeadImg:  headImg,
		Nickname: nickname,
		Content:  content,
	}
	grpcRsp := &comment.MakeCommentRsp{}

	req := client.NewRequest(serviceComment, endpointMakeComment, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

// @Summary 获取评论
// @Tags 评论中心
// @Description 获取评论
// @Accept json
// @Produce json
// @Param movieId query int true "影片id"
// @Success 1 {object} model.Response "success"
// @Router /comment/hotComment [post]
func HotComment(c *gin.Context) {
	movieId, _ := strconv.Atoi(c.Query("movieId"))
	grpcReq := &comment.HotCommentReq{
		MovieId: int64(movieId),
	}
	grpcRsp := &comment.HotCommentRsp{}

	req := client.NewRequest(serviceComment, endpointHotComment, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}
