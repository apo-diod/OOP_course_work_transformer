package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = "6121"
const HOST = "127.0.0.1"
const SENDER_PORT = "6122"

func main() {
	pairs := make(map[string]string)
	r := gin.Default()
	r.POST("/use_module", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		id := req["id"].(string)
		data := req["data"]
		toUsebts, _ := json.Marshal(data)
		toUse := string(toUsebts)
		log.Println(id, toUse)
		go useModule(id, toUse)
		ctx.JSON(http.StatusOK, nil)
	})

	r.POST("/callback", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		id := req["id"].(string)
		data := req["data"]
		responsedatabts, _ := json.Marshal(data)
		responsedata := string(responsedatabts)
		log.Println(id, responsedata)
		req["id"] = pairs[id]
		request, _ := json.Marshal(req)
		log.Println(string(request))
		http.Post("http://"+HOST+":"+SENDER_PORT+"/use_module", "application/json", bytes.NewBuffer(request))
		ctx.JSON(http.StatusOK, nil)
	})

	r.POST("/add_module", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		module := req["module"].(string)
		settingsmap := req["settings"]
		settingsbts, _ := json.Marshal(settingsmap)
		settings := string(settingsbts)
		id := newModule(module, settings)
		if id == "0" {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	})

	r.POST("/link", func(ctx *gin.Context) {
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)
		first := req["first"].(string)
		second := req["second"].(string)
		pairs[first] = second
		ctx.JSON(http.StatusOK, nil)
	})

	r.Run(":" + PORT)
}
