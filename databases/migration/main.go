package main

import (
	"fmt"
	"os"
	"xbrother/databases/entities"
	"xbrother/databases/mysql"
)

func main() {
	if err := mysql.ConnectMysql(); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	if err := mysql.Engine.AutoMigrate(&entities.TBBoard{}); err != nil {
		panic(err)
	}

	if err := mysql.Engine.AutoMigrate(&entities.TBMember{}); err != nil {
		panic(err)
	}

	if err := mysql.Engine.AutoMigrate(&entities.TBLines{}); err != nil {
		panic(err)
	}
}
