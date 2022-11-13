package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Second)
	// 返回的是指针类型
	user := flag.String("user", "root", "用户名")
	//port := flag.Int("user", 3306, "端口")
	ip := flag.String("ip", "localhost", "mysql.ip")

	flag.Parse()

	fmt.Println(*user, *ip)
}
