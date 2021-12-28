package main

import (
	"fmt"
	"net/http"

	"gochenweb/pkg/setting"
	"gochenweb/routers"
)

func main() {
	router := routers.InitRouters()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPORT),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
