package db

import (
	"database/sql"
	"film-srv/entity"
)

// 获取影片剧照
func SelectFilimImages(movie_id int64) ([]*entity.Image, error) {
	images := []*entity.Image{}
	err := db.Select(&images, "SELECT `image_id`,`movie_id`,`image_url` FROM `film_image` WHERE `movie_id` = ?", movie_id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return images, err
}

