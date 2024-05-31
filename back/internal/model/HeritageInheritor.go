/*
@Time : 2024/5/23 下午8:17
@Author : ljn
@File : HeritageInheritor
@Software: GoLand
*/

package model

type HeritageInheritorDb struct {
	Number       string `json:"number"`
	Inheritor    string `json:"inheritor"`
	InheritorImg string `json:"inheritorImg" gorm:"column:inheritorImg"`
	Project      string `json:"project"`
	CreateTime   string `json:"createTime" gorm:"column:createTime"`
	CateGory     string `json:"category"  gorm:"column:category"`
	Details      string `json:"details"`
	Locate       string `json:"locate"`
	Rank         string `json:"rank"`
}

func (h *HeritageInheritorDb) TableName() string {
	return "heritage_inheritor"
}

type HeritageInheritor struct {
	Inheritor    string `json:"inheritor"`
	InheritorImg string `json:"inheritorImg"`
	Project      string `json:"project"`
	CateGory     uint8  `json:"cateGory"`
	Level        uint8  `json:"level"`
	Details      string `json:"details"`
	Locate       string `json:"locate"`
}

func (h *HeritageInheritor) ToList() []interface{} {
	return []interface{}{
		h.Inheritor,
		h.InheritorImg,
		h.Project,
		h.CateGory,
		h.Level,
		h.Details,
		h.Locate,
	}
}

func (h *HeritageInheritor) TableName() string {
	return "heritage_inheritor"
}

var HeritageInheritorModel = new(HeritageInheritor)

// Level
const (
	National   = iota
	Provincial = 1
	City       = 2
	Human      = 3
)

// CateGory
const (
	FolkLiterature      = iota // 民间文学
	TraditionalMusic    = 1    // 传统音乐
	TraditionalDance    = 2    //传统舞蹈
	TraditionalDrama    = 3    // 传统戏剧
	QuYi                = 4    // 曲艺
	TraditionalSports   = 5    //传统体育
	TraditionalArt      = 6    // 传统美术
	TraditionalSkill    = 7    //传统技艺
	TraditionalMedicine = 8    //传统医药
	FolkCustom          = 9    //民俗
)

func (h *HeritageInheritor) GetCateGory() string {
	switch h.CateGory {
	case FolkLiterature:
		return "民间文学"
	case TraditionalMusic:
		return "传统音乐"
	case TraditionalDance:
		return "传统舞蹈"
	case TraditionalDrama:
		return "传统戏剧"
	case QuYi:
		return "曲艺"
	case TraditionalSports:
		return "传统体育"
	case TraditionalArt:
		return "传统美术"
	case TraditionalSkill:
		return "传统技艺"
	case TraditionalMedicine:
		return "传统医药"
	case FolkCustom:
		return "民俗"
	default:
		return "未知"
	}
}

func (h *HeritageInheritor) GetLevel() string {
	switch h.Level {
	case National:
		return "国家级"
	case Provincial:
		return "省级"
	case City:
		return "市级"
	case Human:
		return "人类非物质文化遗产"
	default:
		return "未知"
	}
}

func (h *HeritageInheritor) ToMappingCateGory() map[string]interface{} {
	return map[string]interface{}{
		"民间文学": FolkLiterature,
		"传统音乐": TraditionalMusic,
		"传统舞蹈": TraditionalDance,
		"传统戏剧": TraditionalDrama,
		"曲艺":   QuYi,
		"传统体育": TraditionalSports,
		"传统美术": TraditionalArt,
		"传统技艺": TraditionalSkill,
		"传统医药": TraditionalMedicine,
		"民俗":   FolkCustom,
	}
}

func (h *HeritageInheritor) ToMappingLevel() map[string]interface{} {
	return map[string]interface{}{
		"国家级":       National,
		"省级":        Provincial,
		"市级":        City,
		"人类非物质文化遗产": Human,
	}
}
