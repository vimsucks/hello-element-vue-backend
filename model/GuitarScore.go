package model

import (
	"fmt"
	"github.com/elgris/sqrl"
	"time"
)

type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}

// 吉他乐谱
type GuitarScore struct {
	ID        int64            `db:"id" json:"id"`
	UID       string           `db:"uid" json:"uid"`
	Title     string           `db:"title" json:"title"` // 乐谱标题，如小星星
	Pics      []GuitarScorePic `json:"pics"`
	CreatedAt time.Time        `db:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time        `db:"updated_at" json:"updatedAt"`
}

func GetAllGuitarScore() ([]GuitarScore, error) {
	sql, _, err := sqrl.Select("*").From("GuitarScore").ToSql()
	if err != nil {
		return nil, err
	}
	var ret []GuitarScore
	err = db.Select(&ret, sql)
	if err != nil {
		return nil, err
	}
	for i := range ret {
		ret[i].Pics, err = GetGuitarScorePics(ret[i].UID)
		if err != nil {
			fmt.Println(err)
		}
	}
	return ret, nil
}

func GetGuitarScore(uid string) (GuitarScore, error) {
	var score GuitarScore
	sql, args, err := sqrl.Select("*").
		From("GuitarScore").
		Where(sqrl.Eq{"uid": uid}).
		ToSql()
	if err != nil {
		return score, err
	}
	err = db.Get(&score, sql, args...)
	if err != nil {
		return score, err
	}
	return score, nil
}

// 创建吉他乐谱，同时创建该乐谱的 Pic
// 调用方需要确保 UID 被正确生成
func CreateGuitarScore(score *GuitarScore) error {
	// TODO: 这里是不是要用事物？
	sql, args, err := sqrl.
		Insert("GuitarScore").
		Columns("uid", "title", "created_at", "updated_at").
		Values(score.UID, score.Title, score.CreatedAt, score.UpdatedAt).
		ToSql()
	if err != nil {
		return err
	}
	result, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}
	score.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	for _, pic := range score.Pics {
		err = CreateGuitarScorePic(&pic)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteGuitarScore(score *GuitarScore) error {
	sql, args, err := sqrl.
		Delete("GuitarScore").
		Where(sqrl.Eq{"uid": score.UID}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = db.Exec(sql, args...)
	if err != nil {
		return err
	}
	for _, pic := range score.Pics {
		err = DeleteGuitarScorePic(&pic)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}