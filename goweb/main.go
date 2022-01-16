package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

/**
 *  用于健康检查
 */
func Healthz(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	str := fmt.Sprint("server startup is normal")
	w.Write([]byte(str))
}

/**
 *  用于返回所有请求头信息
 */
func Header(w http.ResponseWriter, r *http.Request){
	h := r.Header.Clone()
	str := fmt.Sprintf("server interface route header data")
	for k, _ := range h {
		v := h.Get(k)
		w.Header().Set(k,v)
		str = fmt.Sprintf("%s \n header[%v]%v",str,k,v)
	}
	ip,port,err := net.SplitHostPort(r.RemoteAddr)
	if err != nil{
		log.Printf("host error %v",err)
	}else{
		if ip == "::1"{
			ip = "127.0.0.1"
		}
		log.Printf("host[%s] post[%s] httpStatus[%s]",ip,port,200)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(str))
}

/**
 *  用于返回系统变量VERSION
 */
func Version(w http.ResponseWriter, r *http.Request){
	v := os.Getenv("VERSION")
	w.Header().Set("VERSION",v)
	str := fmt.Sprintf("Header[VERSION]%s",v)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(str))
}

func Index(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	str := fmt.Sprint("hello goweb")
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/healthz",Healthz)
	http.HandleFunc("/header",Header)
	http.HandleFunc("/version",Version)
	http.HandleFunc("/",Index)
	http.ListenAndServe(":8080",nil)
}
