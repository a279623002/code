package db

import (
	"database/sql"
)

func UpdateOrderScore(orderNum string, score int64) error {
	_, err := db.Exec("UPDATE `film_order` SET `order_score` = ? WHERE `orderNum` = ?", score, orderNum)
	return err
}

func SelectOrderByUidMid(movieId int64, userId int64) (string, error) {

	var orderNUM string = "";
	err := db.Get(orderNUM, "SELECT order_num FROM film_order WHERE user_id = ? AND movie_id = ?", userId, movieId)
	if err != sql.ErrNoRows {
		return "", nil
	}
	return orderNUM, err
}
