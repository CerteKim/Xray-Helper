package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/certekim/xray-helper/app/config"
	hin "github.com/certekim/xray-helper/app/helper/inbound"
	hout "github.com/certekim/xray-helper/app/helper/outbound"
	"github.com/julienschmidt/httprouter"
)

var Router = httprouter.New()

func main() {
	dir := config.Conf["dir"].(string)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("Specific folder not found : " + dir)
		log.Println("Serve static website on : ./xray-webui")
		Router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			http.ServeFile(w, r, "xray-webui/index.html")
		})
		Router.ServeFiles("/helper/*filepath", http.Dir("xray-webui"))
	} else {
		log.Println("Server static website on : " + dir)
		Router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			http.ServeFile(w, r, dir+"/index.html")
		})
		Router.ServeFiles("/helper/*filepath", http.Dir(dir))
	}
	//helper
	Router.POST("/api/helper/inbound/add", hin.WriteInboundHandler)
	Router.GET("/api/helper/inbound/read/:tag", hin.ReadInboundHandler)
	Router.GET("/api/helper/inbound/delete/:tag", hin.DeleteInboundHandler)
	Router.GET("/api/helper/inbound/apply/:tag", hin.ApplyInboundHandler)
	Router.POST("/api/helper/outbound/add", hout.WriteOutboundHandler)
	Router.GET("/api/helper/outbound/read/:tag", hout.ReadOutboundHandler)
	Router.GET("/api/helper/outbound/delete/:tag", hout.DeleteOutboundHandler)
	Router.GET("/api/helper/outbound/apply/:tag", hout.ApplyOutboundHandler)
	//module
	/*
		Router.GET("/api/module/appid/read", appid.ReadHandler)
		Router.POST("/api/module/appid/write", appid.WriteHandler)
		Router.GET("/api/module/appid/apply", appid.ApplyHandler)
		Router.GET("/api/module/appid/query", appid.QueryHandler)
	*/
	port := strconv.FormatFloat(config.Conf["port"].(float64), 'f', 0, 64)
	log.Println("check your webui on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, Router))
}

func init() {
	log.Println("Starting Xray-helper")
}
