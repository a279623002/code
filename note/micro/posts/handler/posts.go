package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/store"
	"time"

	posts "posts/proto"
)

type Posts struct{}

type Post struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	CreateTimestamp int64    `json:"create_timestamp"`
	UpdateTimestamp int64    `json:"update_timestamp"`
	TapNames        []string `json:"tagNames"`
}
// store
// select: micro store read --table=posts --prefix post
// select key: micro store list --table=posts
// delete: micro store delete --table=posts key
// posts
// add: micro posts save --post_id="1" --post_title="how to micro" --post_content="simply put, micro is awesome"
// select: micro posts query [key]
// delete: micro posts delete key
func (p *Posts) Save(ctx context.Context, req *posts.SaveRequest, rsb *posts.SaveResponse) error {
	if len(req.Post.Id) == 0 || len(req.Post.Title) == 0 || len(req.Post.Content) == 0 {
		return errors.BadRequest("posts.Save", "ID, title or content is missing")
	}
	records, err := store.Read(fmt.Sprintf("id:%v", req.Post.Id))
	if err != nil && err != store.ErrNotFound {
		return err
	}

	post := &Post{
		ID:              req.Post.Id,
		Title:           req.Post.Title,
		Content:         req.Post.Content,
		CreateTimestamp: time.Now().Unix(),
		UpdateTimestamp: time.Now().Unix(),
		TapNames:        req.Post.TagNames,
	}

	if len(records) > 0 {
		record := records[0]
		oldPost := &Post{}
		if err := json.Unmarshal(record.Value, oldPost); err != nil {
			return err
		}
		post.CreateTimestamp = oldPost.CreateTimestamp
	}

	return p.savePost(ctx, post)
}

func (p *Posts) savePost(ctx context.Context, post *Post) error {
	bytes, err := json.Marshal(post)
	if err != nil {
		return err
	}
	return store.Write(&store.Record{
		Key:      fmt.Sprintf("id:%v", post.ID),
		Value:    bytes,
	})
}

func (p *Posts) Query(ctx context.Context, req *posts.QueryRequest, rsp *posts.QueryResponse) error {
	var key string
	var opts []store.Option
	if len(req.Id) > 0 {
		key = fmt.Sprintf("id:%v", req.Id)
	}
	if req.Limit > 0 {
		opts = append(opts, store.Limit(uint(req.Limit)))
	} else {
		opts = append(opts, store.Limit(20))
	}
	if req.Offset > 0 {
		opts = append(opts, store.Offset(uint(req.Offset)))
	}
	if key == "" {
		opts = append(opts, store.Prefix("id"))
	}
	records, err := store.Read(key, opts...)
	if err != nil {
		return err
	}

	rsp.Posts = make([]*posts.Post, len(records))
	for i, record := range records {
		postRecord := &Post{}
		if err := json.Unmarshal(record.Value, postRecord); err != nil {
			return err
		}
		rsp.Posts[i] = &posts.Post{
			Id:                   postRecord.ID,
			Title:                postRecord.Title,
			Content:              postRecord.Content,
			TagNames:             postRecord.TapNames,
		}
	}
	return nil
}

func (p *Posts) Delete(ctx context.Context, req *posts.DeleteRequest, rsp *posts.DeleteResponse) error {

	err := store.Delete(fmt.Sprintf("id:%v", req.Id))
	if err == nil {
		rsp.Code = 1
	}
	return err
}
