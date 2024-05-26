/*
@Time : 2024/5/23 下午4:13
@Author : ljn
@File : conf_test
@Software: GoLand
*/

package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseConf(t *testing.T) {

	file, err := ioutil.ReadFile("conf.json")
	if err != nil {
		return
	}
	fmt.Println(string(file))
	var c Conf
	if err = json.Unmarshal(file, &c); err != nil {
		return
	}
	fmt.Println(c)
	t.Log(c.CommonRequest.CommonEq("getCateGoryList", nil))
}
