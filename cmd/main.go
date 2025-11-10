package main

import (
	"log"

	"github.com/QingCold/community-server/internal/api"
	"github.com/QingCold/community-server/internal/config"
	"github.com/QingCold/community-server/internal/db"
)

func main() {
	config.LoadConfig("conf/config.yaml") // 读取配置
	db.InitDB()                           // 初始化 MySQL

	r := api.NewRouter()
	log.Println("Server started on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
