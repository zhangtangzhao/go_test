package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"strings"
)
/**
redis-benchmark -h 127.0.0.1 -p 6379 -q -d 100  -t get
 redis-benchmark -h 127.0.0.1 -p 6379 -q -d 1000  -t get
 */
func main(){
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	start,err := used_memory(c1)
	if err != nil {
		fmt.Println(err)
		return
	}
	args := []interface{}{}
	for i := 0 ; i < 200000; i++{
		args = append(args, i)
		args = append(args, "test2")
	}
	c1.Do("mset",args...)
	end,err := used_memory(c1)
	if err != nil {
		fmt.Println(err)
		return
	}
	num := end - start
	fmt.Printf("消耗了%s字节",num)

}

func used_memory(c1 redis.Conn ) (int64 ,error){
	info_memory, err := redis.String(c1.Do("info","memory" ))
	if err != nil {
		fmt.Println("err while getting:", err)
		return 0, err
	}
	info_arr := strings.Split(info_memory,"\n")
	used_memory := strings.Split(info_arr[1],":")
	return strconv.ParseInt(strings.TrimSpace(used_memory[1]) , 10, 64)

}