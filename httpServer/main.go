package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func index(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"hello world")
}

func startServer(srv *http.Server) error{
	http.HandleFunc("/",index)
	fmt.Println("http server start....")
	error := srv.ListenAndServe()
	return error
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
	signal.Notify(chanel,syscall.SIGINT,syscall.SIGTERM)

	g.Go(func() error {
		for  {
			select {
			case <- ctx.Done():
				return ctx.Err()
			case s:= <- chanel:
				switch s{
				case syscall.SIGINT,syscall.SIGTERM:
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
