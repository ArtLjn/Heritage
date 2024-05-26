/*
@Time : 2024/5/23 下午4:11
@Author : ljn
@File : conf
@Software: GoLand
*/

package conf

import "back/util"

type Conf struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	CommonRequest util.CommonRequest `json:"commonRequest"`
	Database      struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"dbName"`
	} `json:"database"`

	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`

	Rabbitmq struct {
		Url      string   `json:"url"`
		Exchange string   `json:"exchange"`
		Queue    []string `json:"queue"`
	}
}
