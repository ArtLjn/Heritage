/*
@Time : 2024/5/25 下午12:35
@Author : ljn
@File : Account
@Software: GoLand
*/

package model

type Account struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Address  string `json:"address"`
	Password string `json:"password"`
	City     string `json:"city"`
	Level    uint   `json:"level"`
}

func (a *Account) ToFieldList() []string {
	return []string{
		"ID",
		"Address",
		"Password",
		"City",
		"Level",
	}
}
func (a *Account) TableName() string {
	return "account"
}

type AccountLogin struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

const (
	NationalLevel   = 0 // 国家级行政机构
	ProvincialLevel = 1 // 省级行政机构
	CityLevel       = 2 // 市级行政机构
)

func (a *Account) IsNational() bool {
	return a.Level == NationalLevel
}

func (a *Account) IsProvincial() bool {
	return a.Level == ProvincialLevel
}

func (a *Account) IsCity() bool {
	return a.Level == CityLevel
}
