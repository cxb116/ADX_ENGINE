package registration

// 注册中心
//
//func ServerEngineStart(ctx context.Context, serviceName, host, prot string,
//	registerHandlerFunc func()) (context context.Context, err error) {
//	registerHandlerFunc()
//
//	ctx = startService(ctx, serviceName, host, prot)
//	return ctx, nil
//}
//
//func startService(ctx context.Context, name string, host string, prot string) context.Context {
//
//	ctx, cancel := context.WithCancel(ctx)
//
//	var server http.Server
//	server.Addr = host + ":" + prot
//
//	go func() {
//		log.Println(server.ListenAndServe())
//		cancel()
//	}()
//
//	go func() {
//
//		fmt.Printf("%v started . 按任意键关闭服务.", name)
//		var s string
//		fmt.Scanln(&s)
//		server.Shutdown(ctx)
//		cancel()
//	}()
//	return ctx
//}
