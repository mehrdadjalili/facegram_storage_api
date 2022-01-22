package database

import (
	"facegram_file_server/config"
	"github.com/fatih/color"
	"github.com/go-pg/pg"
)

var postgresSqlConnection *pg.DB

func init() {
	_, err := GetPostgresSqlConnection()
	if err == nil {
		c := color.New(color.FgGreen).PrintfFunc()
		c("Successfully connected to postgresql\n")
	} else {
		c := color.New(color.FgRed).PrintfFunc()
		c("Failed Initializing postgresql Connection!\n")
	}
}

func GetPostgresSqlConnection() (*pg.DB, error) {

	if postgresSqlConnection != nil {
		return postgresSqlConnection, nil
	}

	conf := config.GetPostgresSqlConfig()
	p := conf.Password
	u := conf.Username
	a := conf.Address + conf.Port
	d := conf.Database + "?sslmode=disable"
	url := "postgres://" + u + ":" + p + "@" + a + "/" + d

	opt, err := pg.ParseURL(url)

	if err != nil {
		return nil, err
	}

	database := pg.Connect(opt)
	_, err = database.Exec("set search_path=?", conf.Schema)
	if err != nil {
		return nil, err
	}

	return nil, err
}
