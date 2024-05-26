/*
@Time : 2024/5/23 下午4:04
@Author : ljn
@File : writeRequest
@Software: GoLand
*/

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"log"
	"net/http"
)

type CommonRequest struct {
	User            string `json:"user"`
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	Abi             string `json:"contractAbi"`
	Url             string `json:"url"`
	ParsePut        string `json:"parsePut"`
}

func (c *CommonRequest) CommonEq(funcName string, funcParam []interface{}) string {
	requestData := map[string]interface{}{
		"user":            c.User,
		"contractName":    c.ContractName,
		"contractAddress": c.ContractAddress,
		"funcName":        funcName,
		"contractAbi":     json.RawMessage(c.Abi),
		"funcParam":       funcParam,
		"groupId":         1,
		"useCns":          false,
		"useAes":          false,
		"cnsName":         c.ContractName,
		"version":         "",
	}
	requestDataBytes, _ := json.Marshal(requestData)
	req, err := http.NewRequest(http.MethodPost, c.Url, bytes.NewBuffer(requestDataBytes))
	if err != nil {
		fmt.Println("创建HTTP请求错误:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送HTTP请求错误:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应主体错误:", err)
		return ""
	}
	return string(body)
}

func (c *CommonRequest) ParsePutResult(body, funcName string) (interface{}, error) {
	requestData := map[string]interface{}{
		"abiList":    json.RawMessage(c.Abi),
		"methodName": funcName,
		// decodeType为1表示解析input输入的参数，为2表示解析output输出的参数
		"decodeType": 2,
		"output":     gojsonq.New().JSONString(body).Find("output"),
	}
	requestDataBytes, _ := json.Marshal(requestData)
	req, err := http.NewRequest(http.MethodPost, c.ParsePut, bytes.NewBuffer(requestDataBytes))
	if err != nil {
		fmt.Println("创建HTTP请求错误:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送HTTP请求错误:", err)
		return nil, err
	}
	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应主体错误:", err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(d, &data)
	if err != nil {
		fmt.Println("json解析错误:", err)
	}
	var result interface{}
	for k, _ := range data {
		if len(k) != 0 {
			result = k[1 : len(k)-1]
		}
	}
	if result == nil {
		return nil, err
	}
	return result, nil
}

func (c *CommonRequest) IsSuccess(body string) bool {
	data := gojsonq.New().JSONString(body)
	val := data.Find("message")
	if val == "Success" {
		return true
	}
	log.Println(data.Find("errorMessage"))
	return false
}
