package model

import (
	"github.com/elgris/sqrl"
	"time"
)

type GuitarScorePic struct {
	ID        int64     `db:"id" json:"id"`
	UID       string    `db:"uid" json:"uid"`
	ScoreUID  string    `db:"score_uid" json:"scoreUid"`
	URL       string    `db:"url" json:"url"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"` // 创建时间
}

func GetGuitarScorePics(scoreUID string) ([]GuitarScorePic, error) {
	sql, args, err := sqrl.Select("*").From("GuitarScorePic").Where(sqrl.Eq{"score_uid": scoreUID}).ToSql()
	if err != nil {
		return nil, err
	}
	var ret []GuitarScorePic
	err = db.Select(&ret, sql, args...)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func CreateGuitarScorePic(pic *GuitarScorePic) error {
	sql, args, err := sqrl.
		Insert("GuitarScorePic").
		Columns("uid", "score_uid", "url", "created_at").
		Values(pic.UID, pic.ScoreUID, pic.URL, pic.CreatedAt).
		ToSql()
	if err != nil {
		return err
	}
	result, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}
	pic.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func DeleteGuitarScorePic(pic *GuitarScorePic) error {
	sql, args, err := sqrl.
		Delete("GuitarScorePic").
		Where(sqrl.Eq{"uid": pic.UID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = db.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}
