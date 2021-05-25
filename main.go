package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}
func main() {

	//1.注册一个给定模式的处理器函数到DefaultServeMux
	http.HandleFunc("/", sayHello)

	//2.设置监听的TCP地址并启动服务
	//参数1：TCP地址(IP+Port)
	//参数2：当设置为nil时表示使用DefaultServeMux
	err := http.ListenAndServe(":9999", nil)
	log.Fatal(err)
}
