/*
@Time : 2024/5/23 下午8:17
@Author : ljn
@File : HeritageProject
@Software: GoLand
*/

package model

type HeritageProject struct {
	Name     string `json:"name"`
	Category uint8  `json:"category"`
	Level    uint8  `json:"level"`
	Locate   string `json:"locate"`
	Details  string `json:"details"`
}

func (h *HeritageProject) TableName() string {
	return "heritage_project"
}

func (h *HeritageProject) ToList() []interface{} {
	return []interface{}{
		h.Name,
		h.Category,
		h.Level,
		h.Locate,
		h.Details,
	}
}

var HeritageProjectModel = new(HeritageProject)
