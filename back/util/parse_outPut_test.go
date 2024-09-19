/*
@Time : 2024/5/24 下午1:57
@Author : ljn
@File : parse_outPut_test
@Software: GoLand
*/

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	url         = "http://127.0.0.1:5002/WeBASE-Front/tool/decode"
	outPut      = "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000008505a42456e513374000000000000000000000000000000000000000000000000"
	contractAbi = `[{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"categoryList","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"Inheritor","type":"string"},{"internalType":"string","name":"InheritorImg","type":"string"},{"internalType":"string","name":"project","type":"string"},{"internalType":"enum data.CateGory","name":"category","type":"uint8"},{"internalType":"enum data.ApplyLevel","name":"rank","type":"uint8"},{"internalType":"string","name":"details","type":"string"},{"internalType":"string","name":"locate","type":"string"}],"name":"createHeritageInheritor","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"enum data.CateGory","name":"category","type":"uint8"},{"internalType":"enum data.ApplyLevel","name":"rank","type":"uint8"},{"internalType":"string","name":"locate","type":"string"},{"internalType":"string","name":"details","type":"string"}],"name":"createHeritageProject","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getCateGoryList","outputs":[{"internalType":"string[]","name":"","type":"string[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getHeritageInheritorList","outputs":[{"components":[{"internalType":"string","name":"number","type":"string"},{"internalType":"string","name":"Inheritor","type":"string"},{"internalType":"string","name":"InheritorImg","type":"string"},{"internalType":"string","name":"project","type":"string"},{"internalType":"uint256","name":"creationTime","type":"uint256"},{"internalType":"string","name":"category","type":"string"},{"internalType":"string","name":"rank","type":"string"},{"internalType":"string","name":"details","type":"string"},{"internalType":"string","name":"locate","type":"string"}],"internalType":"struct data.HeritageInheritor[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getHeritageProject","outputs":[{"components":[{"internalType":"string","name":"number","type":"string"},{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"category","type":"string"},{"internalType":"string","name":"rank","type":"string"},{"internalType":"string","name":"locate","type":"string"},{"internalType":"string","name":"details","type":"string"}],"internalType":"struct data.HeritageProject[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getRankList","outputs":[{"internalType":"string[]","name":"","type":"string[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"heritageInheritorList","outputs":[{"internalType":"string","name":"number","type":"string"},{"internalType":"string","name":"Inheritor","type":"string"},{"internalType":"string","name":"InheritorImg","type":"string"},{"internalType":"string","name":"project","type":"string"},{"internalType":"uint256","name":"creationTime","type":"uint256"},{"internalType":"string","name":"category","type":"string"},{"internalType":"string","name":"rank","type":"string"},{"internalType":"string","name":"details","type":"string"},{"internalType":"string","name":"locate","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"heritageProjectList","outputs":[{"internalType":"string","name":"number","type":"string"},{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"category","type":"string"},{"internalType":"string","name":"rank","type":"string"},{"internalType":"string","name":"locate","type":"string"},{"internalType":"string","name":"details","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"rankList","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

func TestParseOutPut(t *testing.T) {
	requestData := map[string]interface{}{
		"abiList":    json.RawMessage(contractAbi),
		"methodName": "createHeritageProject",
		// decodeType为1表示解析input输入的参数，为2表示解析output输出的参数
		"decodeType": 2,
		"output":     outPut,
	}
	requestDataBytes, _ := json.Marshal(requestData)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestDataBytes))
	if err != nil {
		fmt.Println("创建HTTP请求错误:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送HTTP请求错误:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应主体错误:", err)
	}
	str := string(body)
	fmt.Println(str)
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("json解析错误:", err)
	}
	fmt.Println(data)

	for k, _ := range data {
		if len(k) != 0 {
			f := k[1 : len(k)-1]
			fmt.Println(f)
		}
	}
}
