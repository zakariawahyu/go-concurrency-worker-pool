package main

import (
	"github.com/zakariawahyu/go-concurrency-worker-pool/config"
	"github.com/zakariawahyu/go-concurrency-worker-pool/pkg/mysql"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db, err := mysql.NewDBConnection(cfg)
	if err != nil {
		panic(err)
	}

	log.Println(db)
}
