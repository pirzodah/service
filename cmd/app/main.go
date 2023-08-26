package main

import (
	"avesto-service/cmd/app/controllers/http"
	_ "avesto-service/docs"
	config "avesto-service/internal/configs"
	"avesto-service/internal/database/clients"
	"avesto-service/internal/entity"
	"avesto-service/internal/usecase"
	pkghttp "avesto-service/pkg/controller/http"
	pkgpostgres "avesto-service/pkg/storage/postgres/pgx"
	"context"
	"fmt"
	"log"
	"os"

	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ctx, cancelFunc := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	cfg := config.InitConfig()
	client, err := pkgpostgres.NewClient(ctx, cfg)
	if err != nil {
		panic(err)
	}
	//client, err := go_mssqldb.NewMSSQLClient(ctx, cfg)
	//if err != nil {
	//	panic(err)
	//}

	database := clients.New(client)
	entities := entity.New(database)
	useCase := usecase.New(entities)
	srv := pkghttp.NewServer()
	controller := http.NewController(useCase, srv)

	wg.Add(1)
	go func() {
		defer wg.Done()

		quitCh := make(chan os.Signal, 1)
		signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		sig := <-quitCh
		log.Println(sig)
		cancelFunc()

		ctx, cancelFunc = context.WithTimeout(ctx, time.Second*10)
		defer cancelFunc()

		if err := controller.Shutdown(ctx); err != nil {
			log.Println()
		}
		log.Println("Server - < finished goroutines")
	}()
	fmt.Println("Serve Started on port:", cfg.Port)
	if err = controller.Serve(ctx, &cfg); err != nil {
		panic(err)
	}
	wg.Wait()

	log.Println("Server - finished main goroutines")
}
