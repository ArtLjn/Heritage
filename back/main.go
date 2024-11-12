/*
@Time : 2024/5/21 下午8:59
@Author : ljn
@File : main
@Software: GoLand
*/

package main

import (
	"back/common/custom_log"
	"back/conf"
	"back/internal/data"
	"back/internal/service"
	"back/router"
	"back/util"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	shoutDownClearCache func()
	finish              = make(chan bool)
)

func main() {
	r := gin.Default()
	go custom_log.InitGinLog("Heritage")
	var c conf.Conf
	wireApp(&c, r)
	srv := &http.Server{
		Addr:    ":" + c.Server.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		log.Println("Server Running on port: " + c.Server.Port)
		if err := srv.ListenAndServe(); !(err == nil || errors.Is(err, http.ErrServerClosed)) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go shoutDownClearCache() // 在 goroutine 中运行清除缓存的函数
	<-finish
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func wireApp(c *conf.Conf, r *gin.Engine) {
	cityFile, err := ioutil.ReadFile("conf/city.json")
	if err != nil {
		log.Fatal(err)
	}
	var cityJson = make(map[string][]string)
	err = conf.LoadEnvConfig(c, "conf/conf.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(cityFile, &cityJson); err != nil {
		log.Fatal(err)
	}

	dataData := data.NewData(c)
	// 注册停止服务删除缓存方法
	shoutDownClearCache = data.NewClearCacheFunc(dataData, finish)
	accountRepo := data.NewAccountRepo(dataData)
	heritageRepo := data.NewHeritageRepo(dataData)
	heritageService := service.NewHeritageService(heritageRepo, accountRepo)
	var mx sync.Mutex
	go func() {
		mx.Lock()
		defer mx.Unlock()
		heritageRepo.InitHeritageInheritor()
		heritageRepo.InitHeritageProject()
	}()
	go heritageRepo.ReceiveHeritageInheritor()
	go heritageRepo.ReceiveHeritageProject()
	go accountRepo.InitAccount(cityJson) // 构建账户数据

	accountService := service.NewAccountService(accountRepo)
	r.StaticFS("/img", http.Dir(c.UploadRepo.UploadPath))
	util.NewUploadRepo(c.UploadRepo.UploadPath, c.UploadRepo.Domain, c.UploadRepo.MaxSize)
	router.InitRouter(r, c, heritageService, accountService)
}
