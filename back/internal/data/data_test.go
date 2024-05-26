/*
@Time : 2024/5/24 下午9:20
@Author : ljn
@File : data_test
@Software: GoLand
*/

package data

import (
	"back/conf"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestRedisPageSize(t *testing.T) {
	var c *conf.Conf
	file, err := ioutil.ReadFile("../../conf/conf.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(file, &c); err != nil {
		return
	}
	rdb := NewRDB(
		c,
	)
	// 分页参数
	pageSize := 2
	pageNum := 1

	// 计算偏移量
	offset := (pageNum - 1) * pageSize
	// 获取HASH的所有fields
	var fields []string
	cmd := rdb.HKeys(context.Background(), HeritageInheritorHashKey)
	if err := cmd.Err(); err != nil {
		log.Fatalf("Error getting keys: %v", err)
	}
	fields = cmd.Val()
	fmt.Println(fields, offset)
	//// 根据分页参数截取fields
	//start := offset
	//end := offset + pageSize
	//if end > len(fields) {
	//	end = len(fields)
	//}

}
