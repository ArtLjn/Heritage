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
	"back/util"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"github.com/thedevsaddam/gojsonq"
	"log"
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
	t    util.TokenManager
}

func (h heritageRepo) QueryHeritageInheritorBlockChainDataByNumber(number string) []interface{} {
	return h.queryHeritageBlockChainRecord("queryHeritageInheritor", number)
}
func (h heritageRepo) QueryHeritageProjectBlockChainDataByNumber(number string) []interface{} {
	return h.queryHeritageBlockChainRecord("queryHeritageProject", number)
}
func (h heritageRepo) UpdateHeritageInheritorLevel(number string, rank uint) error {
	return h.updateHeritageLevel("updateHeritageInheritorLevel", number, rank, h.InitHeritageInheritor)
}
func (h heritageRepo) UpdateHeritageProjectLevel(number string, rank uint) error {
	return h.updateHeritageLevel("updateHeritageProjectLevel", number, rank, h.InitHeritageProject)
}
func (h heritageRepo) QueryHeritageInheritorByLocate(page, size, raw int, locate string) (map[string]interface{}, error) {
	return h.queryHeritageByLocate(page, size, raw, locate, model.HeritageInheritorDb{}.TableName(), []model.HeritageInheritorDb{})
}
func (h heritageRepo) QueryHeritageProjectByLocate(page, size, raw int, locate string) (map[string]interface{}, error) {
	return h.queryHeritageByLocate(page, size, raw, locate, model.HeritageProjectDb{}.TableName(), []model.HeritageProjectDb{})
}
func (h heritageRepo) QueryHeritageTaskById(id int) (m model.Heritage, err error) {
	return m, h.data.db.Table(model.Heritage{}.TableName()).Where("id = ?", id).First(&m).Error
}
func (h heritageRepo) DeleteHeritageTaskById(id int) error {
	return h.data.db.Table(model.Heritage{}.TableName()).Where("id = ?", id).Delete(&model.Heritage{}).Error
}
func (h heritageRepo) ReceiveHeritageInheritor() {
	h.receiveConsume(h.data.c.Rabbitmq.Queue[0], h.createHeritageInheritorData)
}
func (h heritageRepo) ReceiveHeritageProject() {
	h.receiveConsume(h.data.c.Rabbitmq.Queue[1], h.createHeritageProjectData)
}
func (h heritageRepo) QueryHeritageProject(page, size int, header string) map[string]interface{} {
	return h.queryPageSizeHeritage(page, size, model.HeritageTypeProject, header)
}
func (h heritageRepo) QueryHeritageInheritor(page, size int, header string) map[string]interface{} {
	return h.queryPageSizeHeritage(page, size, model.HeritageTypeInheritor, header)
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
	h.initHeritage("getHeritageProject", HeritageProjectHashKey, heritageProject, h.heritageProjectDb)
}
func (h heritageRepo) InitHeritageInheritor() {
	h.initHeritage("getHeritageInheritorList", HeritageInheritorHashKey, heritageInheritor, h.heritageInheritorDb)
}
func (h heritageRepo) CreateHeritageInheritor(inheritor *model.HeritageInheritor) ([]interface{}, error) {
	strBody := h.data.c.CommonRequest.CommonEq("createHeritageInheritor", inheritor.ToList())
	if !h.data.c.CommonRequest.IsSuccess(
		strBody,
	) {
		return nil, fmt.Errorf("CreateHeritageInheritor error")
	}
	response, err := h.data.c.CommonRequest.ParsePutResult(strBody, "createHeritageInheritor")
	if err != nil {
		return nil, err
	}

	if err := h.createHeritage(response, "queryHeritageInheritor",
		HeritageInheritorHashKey, heritageInheritor, h.InitHeritageInheritor); err != nil {
		return nil, err
	}
	return []interface{}{response}, nil
}

func (h heritageRepo) CreateHeritageProject(project *model.HeritageProject) ([]interface{}, error) {
	strBody := h.data.c.CommonRequest.CommonEq("createHeritageProject", project.ToList())
	if !h.data.c.CommonRequest.IsSuccess(
		strBody,
	) {
		return nil, fmt.Errorf("CreateHeritageProject error")
	}
	response, err := h.data.c.CommonRequest.ParsePutResult(strBody, "createHeritageProject")
	if err != nil {
		return nil, err
	}

	if err = h.createHeritage(response, "queryHeritageProject",
		HeritageProjectHashKey, heritageProject, h.InitHeritageProject); err != nil {
		return nil, err
	}
	return []interface{}{response}, nil
}
func (h heritageRepo) UpdateHeritageTask(level uint, id int, number string) error {
	if err := h.data.db.Model(model.Heritage{}).Where("id =?", id).Updates(map[string]interface{}{
		"pass_level": level,
		"number":     number,
	}).Error; err != nil {
		log.Printf("update error %v", err)
		return err
	}
	return nil
}
func NewHeritageRepo(data *Data) service.HeritageRepo {
	return &heritageRepo{
		data: data,
		t:    util.NewToken(data.rdb),
	}
}

func (h heritageRepo) updateHeritageLevel(funcName, number string, rank uint, m func()) error {
	transRes := h.data.c.CommonRequest.CommonEq(funcName, []interface{}{
		number, rank - 1,
	})
	// 成功更新区块链数据
	if h.data.c.CommonRequest.IsSuccess(transRes) {
		// 刷新备份数据库数据
		go m()
	}
	return nil
}

// initHeritage: copy区块链数据写入数据库
// params:
// funcName: 合约方法名称
// cacheKey: redis缓存key名称
// f: 将列表转为json格式
// g: 存入数据库方法
func (h heritageRepo) initHeritage(funcName, cacheKey string,
	f func([]interface{}) map[string]interface{},
	g func(map[string]interface{})) {
	res := h.data.c.CommonRequest.CommonEq(funcName, nil)
	var list []interface{}
	var list2 [][]interface{}
	// 解析从区块链上获得的数据
	if err := json.Unmarshal([]byte(res), &list); err != nil {
		panic(err)
	} else if len(list) == 0 {
		// 如果列表为空 删除缓存
		err = h.data.rdb.Del(context.Background(), cacheKey).Err()
		if err != nil {
			panic(err)
		}
		// 格式化数列表结构
	} else if err = json.Unmarshal([]byte(list[0].(string)), &list2); err != nil {
		panic(err)
	}
	// 加上互斥锁
	mx.Lock()
	for _, v := range list2 {
		// 列表->JSON
		data := f(v)
		by, _ := json.Marshal(data)
		// save -> db
		g(data)
		// save -> rdb
		_, err := h.data.rdb.HSet(context.Background(), cacheKey, v[0], string(by)).Result()
		if err != nil {
			fmt.Println(err)
			continue
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

func (h heritageRepo) heritageInheritorDb(b map[string]interface{}) {
	var m model.HeritageInheritorDb
	d, err := json.Marshal(&b)
	if err != nil {
		log.Printf("marshal error %v", err)
		return
	} else if err = json.Unmarshal(d, &m); err != nil {
		log.Printf("unmarshal error %v", err)
		return
	} else if err = h.data.db.Save(&m).Error; err != nil {
		log.Printf("create error %v", err)
		return
	}

}

func (h heritageRepo) heritageProjectDb(b map[string]interface{}) {
	d, err := json.Marshal(&b)
	var m model.HeritageProjectDb
	if err != nil {
		log.Printf("marshal error %v", err)
		return
	} else if err = json.Unmarshal(d, &m); err != nil {
		log.Printf("unmarshal error %v", err)
		return
	} else if err = h.data.db.Save(&m).Error; err != nil {
		log.Printf("create error %v", err)
		return
	}
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

func (h heritageRepo) queryHeritageBlockChainRecord(funcName, number string) []interface{} {
	res := h.data.c.CommonRequest.CommonEq(funcName, []interface{}{
		number,
	})
	if len(res) != 0 {
		var list []interface{}
		if err := json.Unmarshal([]byte(res), &list); err != nil {
			log.Printf("query Error %v", err)
			return nil
		}
		return list
	}
	return nil
}

func (h heritageRepo) createHeritage(response interface{}, queryFuncName, key string,
	f func(v []interface{}) map[string]interface{}, m func()) error {
	list := h.queryHeritageBlockChainRecord(queryFuncName, response.(string))
	if len(list) == 0 {
		return fmt.Errorf("query Errory")
	}
	data := f(list)
	by, _ := json.Marshal(data)
	_, err := h.data.rdb.HSet(context.Background(), key, response, string(by)).Result()
	go m()
	if err != nil {
		return err
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
		for _, v := range body {
			if err = json.Unmarshal([]byte(v), &d); err != nil {
				log.Printf("%s error %v", key, err)
				continue
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
		Locate:     gojsonq.New().FromString(data).Find("locate").(string),
		ApplyLevel: uint8(gojsonq.New().FromString(data).Find("level").(float64)),
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

func (h heritageRepo) queryPageSizeHeritage(page, size, raw int, header string) map[string]interface{} {
	err := h.t.VerifyToken(header)
	if err != nil {
		return nil
	}

	locate := util.GetLoginName(header)
	if locate == "" {
		return nil
	}

	var (
		list  []model.Heritage
		total int64
		acc   model.Account
	)

	// 首先计算总记录数
	query := h.data.db.Table(model.Heritage{}.TableName()).Where("type = ?", raw)

	if err := h.data.db.Model(model.Account{}).Where("city = ?", locate).First(&acc).Error; err != nil {
		log.Printf("该城市不存在")
		return nil
	}

	// 根据申请级别调整查询逻辑
	if locate == "国家" {
		// 国家级用户审核：可以审核申请级别为"国家级"及以上的项目，但这些项目需要已经通过省级审核
		query = query.Where("apply_level IN (?)", []int{model.National, model.Human}).Where("pass_level = ?", model.Provincial)
	} else if locate == "人类非物质遗产" {
		// 人类非物质遗产项目审核：只能审核申请级别为"人类非物质遗产"的项目，并且 pass_level 需要为"国家级"
		query = query.Where("apply_level = ?", model.Human).Where("pass_level = ?", model.National)
	} else if acc.Level == model.Provincial {
		// 省级用户审核：可以审核申请级别为"省级"及以上的项目，但这些项目需要已经通过市级审核
		query = query.Where("apply_level IN (?)", []int{model.Provincial, model.National, model.Human}).Where("pass_level = ?", model.City)
	} else {
		// 市级用户只能审核申请级别为"市级"的项目
		query = query.Where("locate = ?", locate)
	}

	// Use Debug() to print the generated SQL and see if there are any issues
	query = query.Debug()

	// 先计算总记录数
	if err = query.Count(&total).Error; err != nil {
		log.Printf("query heritage count error: %v", err)
		return nil
	}

	// 再查询分页数据
	offset := (page - 1) * size
	if err = query.Order("create_time asc").Offset(offset).Limit(size).Find(&list).Error; err != nil {
		log.Printf("query heritage error %v", err)
		return nil
	}

	// 创建返回的map
	m := make(map[string]interface{})
	m["total"] = total
	m["totalPages"] = (total + int64(size) - 1) / int64(size) // 计算总页数
	m["currentPage"] = page
	m["pageSize"] = size
	m["list"] = list

	return m
}

func (h heritageRepo) queryHeritageByLocate(page, size, raw int, locate, tableName string, list interface{}) (map[string]interface{}, error) {
	offset := (page - 1) * size
	var total int64

	query := h.data.db.Table(tableName)

	// 根据 raw 参数判断查询方式
	if raw == 1 {
		// 使用 LIKE 查询 locate
		keyword := "%" + locate + "%"
		query = query.Where("locate LIKE ?", keyword)
	} else {
		// 使用精确查询 locate
		query = query.Where("locate = ?", locate)
	}

	// 先执行 Count 查询总数
	if err := query.Count(&total).Error; err != nil {
		log.Printf("query heritage count error: %v", err)
		return nil, err
	}

	// 再执行分页查询
	if err := query.Offset(offset).Limit(size).Find(&list).Error; err != nil {
		log.Printf("query heritage error: %v", err)
		return nil, err
	}

	// 构建返回的结果
	m := make(map[string]interface{})
	m["total"] = total
	m["totalPages"] = (total + int64(size) - 1) / int64(size) // 计算总页数
	m["currentPage"] = page
	m["list"] = list

	return m, nil
}
