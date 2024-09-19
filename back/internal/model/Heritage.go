/*
@Time : 2024/5/26 下午12:00
@Author : ljn
@File : Heritage
@Software: GoLand
*/

package model

type Heritage struct {
	Id         int    `gorm:"primaryKey"  json:"id"`
	UUID       string `json:"uuid"`
	Field      string `json:"field"`
	Type       int    `json:"type"`
	CreateTime string `json:"create_time"`
	Locate     string `json:"locate"`
	ApplyLevel uint8  `json:"apply_level"`
	PassLevel  uint8  `json:"pass_level"`
	Number     string `json:"number"`
}

func (Heritage) TableName() string {
	return "heritage_task"
}

const (
	HeritageTypeInheritor = 1
	HeritageTypeProject   = 2
)
