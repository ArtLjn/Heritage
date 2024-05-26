package util

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

// Token 管理模块 ** 存储|过期｜续期

const TokenMangerKey = "token_manager"

type TokenManager interface {
	SaveToken(string) string
	LogOutToken(string)
	VerifyToken(string) error
}

type Token struct {
	rdb *redis.Client
}

func NewToken(rdb *redis.Client) TokenManager {
	return &Token{
		rdb: rdb,
	}
}

func (t *Token) SaveToken(username string) string {
	token, err := Sign(username)
	if err != nil {
		log.Println("❌ 生成token失败")
		return ""
	}
	_, err = t.rdb.HSet(context.Background(), TokenMangerKey, username, token).Result()
	if err != nil {
		log.Println("❌ 保存token失败")
		return ""
	}
	return token
}

func (t *Token) LogOutToken(username string) {
	_, err := t.rdb.HDel(context.Background(), TokenMangerKey, username).Result()
	if err != nil {
		log.Println("❌ 删除token失败")
	}
}

func (t *Token) VerifyToken(token string) error {
	username := GetLoginName(token)
	if username == "" {
		log.Println("❌ token无效")
		return fmt.Errorf("❌ token无效")
	}
	cacheToken, err := t.rdb.HGet(context.Background(), TokenMangerKey, username).Result()
	if err != nil {
		log.Println("❌ token不存在")
		return fmt.Errorf("❌ token不存在")
	}
	if token != cacheToken {
		log.Println("❌ token无效")
		return fmt.Errorf("❌ token无效")
	}
	return nil
}
