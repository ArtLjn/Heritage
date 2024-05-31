/*
@Time : 2024/5/31 下午2:19
@Author : ljn
@File : response_test
@Software: GoLand
*/

package response

import (
	"fmt"
	"testing"
)

var r ResponseBuild

func TestResponse_1(t *testing.T) {
	//r.SetCode(200).SetMsg("hello").SetData("sss")
	fmt.Println(r)
}

func TestResponse_2(t *testing.T) {
	r.SetCode(200).SetMsg("ggg").SetData("sss")
	fmt.Println(r)
	r.SetCode(400)
	fmt.Println(r)
}
