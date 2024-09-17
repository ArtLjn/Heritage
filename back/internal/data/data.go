/*
@Time : 2024/5/21 下午9:46
@Author : ljn
@File : data
@Software: GoLand
*/

package data

import (
	"back/conf"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	c   *conf.Conf
	db  *gorm.DB
	rdb *redis.Client
	mq  *MQ
}

func NewData(c *conf.Conf) *Data {
	mq := ConnMq(c.Rabbitmq.Url)
	mq.StartMq(c.Rabbitmq.Exchange, c.Rabbitmq.Queue...)
	return &Data{
		c:   c,
		db:  NewDB(c),
		rdb: NewRDB(c),
		mq:  mq,
	}
}

func NewDB(conf *conf.Conf) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func NewRDB(conf *conf.Conf) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     conf.Redis.Addr,
			DB:       conf.Redis.Db,
			Password: conf.Redis.Password,
		},
	)
}

// NewClearCacheFunc 现在返回一个函数
func NewClearCacheFunc(d *Data, done chan bool) func() {
	return func() {
		defer func() {
			done <- true
		}()
		if err := d.rdb.Del(context.Background(), "heritageAccount").Err(); err != nil {
			panic("redis del cache failed: " + err.Error())
		}
	}
}
