package main

import (
	"context"
	"fmt"
	"github.com/cxb116/ADX_ENGINE/registerEngine/regService"

	"log"
	"net/http"
)

func main() {

	http.Handle("/", &regService.RegisterService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server

	srv.Addr = ":80"
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Register service started. Press any key to exit...")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down...")
	//ctx, err := registration.ServerEngineStart(context.Background(),
	//	"SspEngine Service",
	//	"127.0.0.1",
	//	"8080",
	//	sspService.RegisterHandler,
	//)
	//if err != nil {
	//	fmt.Println("SspEngine start error:", err)
	//}
	//fmt.Println("SspEngine start success")
	//<-ctx.Done()
	//
	//fmt.Println("Shutting down SspEngine service .")
}
