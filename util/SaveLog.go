package util

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"simple-payment/config"
)

type SaveLog struct {
}

func (u *SaveLog) Write(p []byte) (n int, err error) {

	conf := config.NewConfig()

	db, err := sqlx.Connect("postgres", conf.DataSourceName)
	if err != nil {
		fmt.Println("initDB failed", err.Error())
		panic(err)
	}
	if errConf := db.Ping(); errConf != nil {
		panic(errConf)
	}

	_, _ = db.Exec(CREATE_LOG, string(p))

	n = len(p)
	return n, err
}
