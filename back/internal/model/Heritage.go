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
}

func (Heritage) TableName() string {
	return "heritage"
}

const (
	HeritageTypeInheritor = 1
	HeritageTypeProject   = 2
)
