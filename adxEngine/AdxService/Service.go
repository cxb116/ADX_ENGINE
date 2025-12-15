package AdxService

import (
	"context"
	"fmt"
	"github.com/cxb116/ADX_ENGINE/registerServer/regService"
	"github.com/cxb116/ADX_ENGINE/registerServer/registry"
	"log"
	"net/http"
)

func ServerEngineStart(ctx context.Context, host, prot string, reg registry.Registration,
	registerHandlerFunc func()) (context context.Context, err error) {
	registerHandlerFunc()

	ctx = startService(ctx, reg.ServiceName, host, prot)
	err = regService.RegisterClient(reg)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func startService(ctx context.Context, name registry.ServiceName, host string, prot string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var server http.Server
	server.Addr = host + ":" + prot

	go func() {
		log.Println(server.ListenAndServe())
		cancel()
	}()

	go func() {

		fmt.Printf("%v started . 按任意键关闭服务.", name)
		var s string
		fmt.Scanln(&s)
		server.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
