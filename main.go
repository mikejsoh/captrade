package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mikejsoh/captrade/db"
	"github.com/mikejsoh/captrade/routers"
)

func main() {
	db.Init()
	routers.Init()
}
