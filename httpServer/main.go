package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhangtangzhao.github/httpServer/metrics"
)

func index(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"hello world")
}

func startServer(srv *http.Server) error{
	 fmt.Println("http server start....")
	http.HandleFunc("/",index)
	 http.HandleFunc("/hello",rootHandle)
	 http.Handle("/metrics",promhttp.Handler())
	error := srv.ListenAndServe()
	return error
}

func rootHandle(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("entering root handler")
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
	io.WriteString(writer, "hello")
}

func main(){
	rootctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(rootctx)
	srv := &http.Server{Addr: ":8080"}
	g.Go(func() error {
		return startServer(srv)
	})

	g.Go(func() error {
		<- ctx.Done()
		fmt.Println("http server down...")
		return srv.Shutdown(ctx)
	})

	chanel := make(chan os.Signal)
	signal.Notify(chanel,syscall.SIGINT,syscall.SIGTERM,syscall.SIGKILL)

	g.Go(func() error {
		for  {
			select {
			case <- ctx.Done():
				return ctx.Err()
			case s:= <- chanel:
				switch s{
				case syscall.SIGINT,syscall.SIGTERM,syscall.SIGKILL:
					cancel()
				default:
					fmt.Println("unsignal syscall...")
				}
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil && err != context.Canceled {
		fmt.Printf("errgroup err : %+v\n",err.Error())
	}

	fmt.Println("httpserver shutdown .....")
}
