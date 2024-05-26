/*
@Time : 2024/5/23 下午8:11
@Author : ljn
@File : heritage
@Software: GoLand
*/

package data

import (
	"back/internal/model"
	"back/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
	"strings"
	"sync"
	"time"
)

const (
	HeritageInheritorHashKey = "heritage_inheritor_hash_key"
	HeritageProjectHashKey   = "heritage_project_hash_key"
)

var mx sync.Mutex
var workerPool = make(chan struct{}, 10) // 假设最多同时处理10个消息

type heritageRepo struct {
	data *Data
}

func (h heritageRepo) ReceiveHeritageInheritor() {
	h.receiveConsume(h.data.c.Rabbitmq.Queue[0], h.createHeritageInheritorData)
}

func (h heritageRepo) ReceiveHeritageProject() {
	h.receiveConsume(h.data.c.Rabbitmq.Queue[1], h.createHeritageProjectData)
}
func (h heritageRepo) QueryHeritageProject(page, size int) map[string]interface{} {
	return h.queryPageSizeHeritage(page, size, model.HeritageTypeProject)
}

// QueryHeritageInheritor 取出待审核消息
func (h heritageRepo) QueryHeritageInheritor(page, size int) map[string]interface{} {
	return h.queryPageSizeHeritage(page, size, model.HeritageTypeInheritor)
}

func (h heritageRepo) PublishHeritageInheritor(inheritor *model.HeritageInheritor) error {
	return h.publishMsg(inheritor, h.data.c.Rabbitmq.Queue[0])
}

func (h heritageRepo) PublishHeritageProject(project *model.HeritageProject) error {
	return h.publishMsg(project, h.data.c.Rabbitmq.Queue[1])
}

func (h heritageRepo) GetCateGory() map[string]interface{} {
	return model.HeritageInheritorModel.ToMappingCateGory()
}

func (h heritageRepo) GetLevel() map[string]interface{} {
	return model.HeritageInheritorModel.ToMappingLevel()
}

func (h heritageRepo) QueryAllHeritageInheritor() ([]interface{}, error) {
	return h.queryAllHeritage(HeritageInheritorHashKey)
}

func (h heritageRepo) QueryAllHeritageProject() ([]interface{}, error) {
	return h.queryAllHeritage(HeritageProjectHashKey)
}

func (h heritageRepo) InitHeritageProject() {
	h.initHeritage("getHeritageProject", HeritageProjectHashKey, heritageProject)
}
func (h heritageRepo) InitHeritageInheritor() {
	h.initHeritage("getHeritageInheritorList", HeritageInheritorHashKey, heritageInheritor)
}

func (h heritageRepo) CreateHeritageInheritor(inheritor *model.HeritageInheritor) error {
	strBody := h.data.c.CommonRequest.CommonEq("createHeritageInheritor", inheritor.ToList())
	if !h.data.c.CommonRequest.IsSuccess(
		strBody,
	) {
		return fmt.Errorf("CreateHeritageInheritor error")
	}
	if err := h.createHeritage(strBody,
		"createHeritageInheritor", "queryHeritageInheritor"); err != nil {
		return err
	}
	return nil
}

func (h heritageRepo) CreateHeritageProject(project *model.HeritageProject) error {
	strBody := h.data.c.CommonRequest.CommonEq("createHeritageProject", project.ToList())
	if !h.data.c.CommonRequest.IsSuccess(
		strBody,
	) {
		return fmt.Errorf("CreateHeritageProject error")
	}
	if err := h.createHeritage(strBody,
		"createHeritageProject", "queryHeritageProject"); err != nil {
		return err
	}
	return nil
}

func NewHeritageRepo(data *Data) service.HeritageRepo {
	return &heritageRepo{
		data: data,
	}
}

func (h heritageRepo) initHeritage(funcName, cacheKey string, f func([]interface{}) map[string]interface{}) {
	res := h.data.c.CommonRequest.CommonEq(funcName, nil)
	res = res[2 : len(res)-2]
	if res == "[ ]" {
		err := h.data.rdb.Del(context.Background(), cacheKey).Err()
		if err != nil {
			panic(err)
		}
		return
	}
	var list [][]interface{}
	if err := json.Unmarshal([]byte(strings.ReplaceAll(res, "\\", "")), &list); err != nil {
		panic(err)
	}
	mx.Lock()
	for _, v := range list {
		data := f(v)
		by, _ := json.Marshal(data)
		_, err := h.data.rdb.HSet(context.Background(), cacheKey, v[0], string(by)).Result()
		if err != nil {
			panic(err)
		}
	}
	mx.Unlock()
}

var heritageInheritor = func(v []interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"number":       v[0],
		"Inheritor":    v[1],
		"InheritorImg": v[2],
		"project":      v[3],
		"creationTime": v[4],
		"category":     v[5],
		"rank":         v[6],
		"details":      v[7],
		"locate":       v[8],
	}
	return data
}

var heritageProject = func(v []interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"number":   v[0],
		"name":     v[1],
		"category": v[2],
		"rank":     v[3],
		"locate":   v[4],
		"details":  v[5],
	}
	return data
}

func (h heritageRepo) publishMsg(data interface{}, queue string) error {
	by, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	err = h.data.mq.PublishMsg(h.data.c.Rabbitmq.Exchange, queue, by)
	if err != nil {
		return err
	}
	return nil
}

func (h heritageRepo) createHeritage(strBody, funcName, queryFuncName string) error {
	response, err := h.data.c.CommonRequest.ParsePutResult(strBody, funcName)
	if err != nil {
		return err
	}
	res := h.data.c.CommonRequest.CommonEq(queryFuncName, []interface{}{
		response,
	})
	if len(res) != 0 {
		var list []interface{}
		if err = json.Unmarshal([]byte(res), &list); err != nil {
			return err
		}
		if len(list) == 0 {
			return fmt.Errorf(funcName, " error")
		}
		data := heritageInheritor(list)
		by, _ := json.Marshal(data)
		_, err = h.data.rdb.HSet(context.Background(), HeritageInheritorHashKey, response, string(by)).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func (h heritageRepo) queryAllHeritage(key string) ([]interface{}, error) {
	body, err := h.data.rdb.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	var list []interface{}
	if len(body) != 0 {
		var d map[string]interface{}
	Tag:
		for _, v := range body {
			if err = json.Unmarshal([]byte(v), &d); err != nil {
				log.Printf("%s error %v", key, err)
				continue Tag
			}
			list = append(list, d)
		}
	}
	return list, nil
}

func (h heritageRepo) queryTotal(queue string) (int, error) {
	q, err := h.data.mq.Channel.QueueDeclarePassive(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return 0, err
	}
	if q.Messages == 0 {
		return 0, fmt.Errorf("no message")
	}

	return q.Messages, nil
}

func (h heritageRepo) receiveConsume(queue string, f func(string) error) {
	msg, err := h.data.mq.Channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("consume error %v", err)
		return
	}

	forever := make(chan bool)
	go func() {
		var wg sync.WaitGroup
		for d := range msg {
			wg.Add(1)
			workerPool <- struct{}{} // 获取信号量，控制并发数
			go func(delivery amqp.Delivery) {
				defer wg.Done()
				defer func() {
					<-workerPool
					body := string(delivery.Body)
					err = f(body)
					if err == nil {
						// 消息确认
						err = d.Ack(false)
						if err != nil {
							log.Printf("ack error %v", err)
						}
					}
				}()
			}(d)
			wg.Wait()
		}
	}()
	<-forever
}

func (h heritageRepo) createHeritageData(data string, raw int) error {
	m := model.Heritage{
		UUID:       uuid.New().String()[:8],
		Field:      data,
		Type:       raw,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := h.data.db.Table(m.TableName()).Create(&m).Error
	if err != nil {
		log.Printf("create heritage error %v", err)
		return err
	}
	return nil
}

func (h heritageRepo) createHeritageInheritorData(data string) error {
	return h.createHeritageData(data, model.HeritageTypeInheritor)
}

func (h heritageRepo) createHeritageProjectData(data string) error {
	return h.createHeritageData(data, model.HeritageTypeProject)
}
func (h heritageRepo) queryPageSizeHeritage(page, size, raw int) map[string]interface{} {
	var list []model.Heritage
	// 计算偏移量
	offset := (page - 1) * size
	// 执行查询
	h.data.db.Table(model.Heritage{}.TableName()).Where("type = ?", raw).
		Order("create_time asc").Offset(offset).Limit(size).Find(&list)

	// 计算总记录数
	var total int64
	h.data.db.Table(model.Heritage{}.TableName()).Where("type = ?", raw).Count(&total)
	// 创建返回的map
	m := make(map[string]interface{})
	m["total"] = total
	m["totalPages"] = (total + int64(size) - 1) / int64(size) // 计算总页数
	m["currentPage"] = page
	m["pageSize"] = size
	m["list"] = list

	return m
}
