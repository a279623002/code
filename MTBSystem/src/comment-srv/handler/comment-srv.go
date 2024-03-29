package handler

import (
	"comment-srv/db"
	"comment-srv/entity"
	pb "comment-srv/proto"
	"context"
	"errors"
)

type CommentServiceExtHandler struct {
}

func NewCommentServiceExtHandler() *CommentServiceExtHandler {
	return &CommentServiceExtHandler{

	}
}

// 获取评论
func (c *CommentServiceExtHandler) HotComment(ctx context.Context, req *pb.HotCommentReq, rsp *pb.HotCommentRsp) error {

	movieId := req.MovieId

	comments, err := db.SelectHotComment(movieId)
	if err != nil {
		//c.logger.Error("err", zap.Error(err))
		//return errors.ErrorCommentFailed
		return errors.New("操作异常")
	}
	records := []*pb.CommentRecord{}
	for _, comment := range comments {
		record := pb.CommentRecord{
			Title:     comment.Title,
			Content:   comment.Content,
			HeadImg:   comment.HeadImg,
			Nickname:  comment.NickName,
			CreateAt:  comment.CreateAt,
			UpNum:     comment.UpNum,
			CommentID: comment.CommentId,
		}
		records = append(records, &record)
	}

	plus := pb.CommentPlus{
		Total: int64(len(comments)),
		List:  records,
	}

	data := pb.CommentData{
		Plus: &plus,
	}
	rsp.Data = &data
	return nil
}

func (f *CommentServiceExtHandler) MakeComment(ctx context.Context, req *pb.MakeCommentReq, rsp *pb.MakeCommentRsp) error {

	comment := entity.Comment{
		FilmId:    req.MovieId,
		Title:     req.Title,
		Content:   req.Content,
		HeadImg:   req.HeadImg,
		NickName:  req.Nickname,
		UserId:    req.UserId,
	}
	err := db.InsertHotComment(&comment)
	if err != nil {
		return err
	}
	return nil
}

func (f *CommentServiceExtHandler) UpNumComment(ctx context.Context, req *pb.UpNumCommentReq, rsp *pb.UpNumCommentRsp) error {

	err := db.UpdateHotComment(req.CommentID)
	if err != nil {
		return errors.New("操作异常")
	}
	upNum, err := db.SelectUpNum(req.CommentID)
	if err != nil {
		return errors.New("操作异常")
	}
	rsp.UpNum = upNum
	return nil
}

func (c *CommentServiceExtHandler) MyComments(ctx context.Context, req *pb.MyCommentsReq, rsp *pb.MyCommentsRsp) error {

	comments, err := db.SelectMyComment(req.UserId)
	if err != nil {
		return errors.New("操作异常")
	}
	commentsPB := []*pb.MyComment{}
	for _, comment := range comments {

		img, title, err := db.SelectFilmImageAndName(comment.FilmId)
		if err != nil {
			return errors.New("操作异常")
		}
		commentPB := pb.MyComment{
			Content:   comment.Content,
			UpNum:     comment.UpNum,
			CommentID: comment.CommentId,
			FilmImage: img,
			FilmName:  title,
			Score:     comment.Title,
		}
		commentsPB = append(commentsPB, &commentPB)
	}
	rsp.MyComments = commentsPB
	return nil
}

func (c *CommentServiceExtHandler) DeleteComment(ctx context.Context, req *pb.DeleteCommentReq, rsp *pb.DeleteCommentRsp) error {

	err := db.DeleteComment(req.CommentID)
	if err != nil {
		return errors.New("操作异常")
	}
	return nil
}
