/*
@Time : 2024/5/21 下午8:59
@Author : ljn
@File : main
@Software: GoLand
*/

package main

import (
	"back/conf"
	"back/internal/data"
	"back/internal/service"
	"back/router"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func main() {
	r := gin.Default()
	go conf.InitGinLog("Heritage")
	var c conf.Conf
	wireApp(&c, r)
	err := r.Run(fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port))
	if err != nil {
		panic(err)
	}
}

func wireApp(c *conf.Conf, r *gin.Engine) {
	file, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(file, &c); err != nil {
		return
	}
	dataData := data.NewData(c)
	heritageRepo := data.NewHeritageRepo(dataData)
	heritageService := service.NewHeritageService(heritageRepo)
	go heritageRepo.InitHeritageInheritor()
	go heritageRepo.InitHeritageProject()
	go heritageRepo.ReceiveHeritageInheritor()
	go heritageRepo.ReceiveHeritageProject()
	accountRepo := data.NewAccountRepo(dataData)
	// 构建账户数据
	go accountRepo.InitAccount()
	accountService := service.NewAccountService(accountRepo)
	router.InitRouter(r, heritageService, accountService)
}
