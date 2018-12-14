package middleware

import (
	"time"
)

type Setting struct {
	ID        int64     `form:"id" database:"id"`
	Key       string    `form:"key"`
	Value     string    `form:"value"`
	Enable    bool      `form:"is_enable"`
	CreatedAt time.Time `form:"created_at"`
	UpdatedAt time.Time `form:"updated_at"`
}

func (setting *Setting) InsertSetting() error {
	_, err := db.Exec("INSERT INTO settings (key, value, is_enable, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING ID",
		setting.Key, setting.Value, setting.Enable, time.Now(), time.Now())
	return err
}

func (setting *Setting) UpdateSetting() error {
	_, err := db.Exec("")
	return err
}

func (setting *Setting) DeleteSetting() error {
	_, err := db.Exec("")
	return err
}

func GetSettingByKey(key string) (*Setting, error) {
	return db.Query("")
}
