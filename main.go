package main

import (
	"flag"
	"gin-vue-blog/model"
	"gin-vue-blog/routers"
)

func main() {
	withTLS := flag.Bool("tls", false, "enable HTTPS (TLS)")
	flag.Parse()

	model.InitDB()
	routers.InitRouter(*withTLS)
}
