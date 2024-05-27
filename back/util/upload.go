/*
@Time : 2024/5/27 下午7:32
@Author : ljn
@File : upload
@Software: GoLand
*/

package util

import (
	"back/internal/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var u *UploadRepo

func NewUploadRepo(data ...string) {
	u = &UploadRepo{
		UploadPath: data[0],
		Domain:     data[1],
		MaxSize:    data[2],
	}
}

type UploadRepo struct {
	UploadPath string `json:"upload_path"` //上传路径
	Domain     string `json:"domain"`      // 域名
	MaxSize    string `json:"max_size"`    // 最大文件大小（字节）
}

func GinUploadImg(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	size, _ := ParseSize(u.MaxSize)
	res := response.ResponseBuild{}
	if err != nil {
		res.NewBuildJsonError(ctx)
		return
	} else if file.Size > size {
		info := fmt.Sprintf("上传文件大小不的大于%s", u.MaxSize)
		res.SetCode(400).SetMsg(info).SetData(nil).Build(ctx)
		return
	}
	originName := file.Filename
	lastDotIndex := strings.LastIndex(originName, ".")
	if lastDotIndex == -1 {
		res.SetCode(400).SetMsg("文件后缀获取失败").SetData(nil).Build(ctx)
		return
	}
	// 提取后缀名
	extension := originName[lastDotIndex+1:]
	newFileName := fmt.Sprintf("%s%s%s", uuid.New().String()[:8], ".", extension)
	savePath := filepath.Join(u.UploadPath, newFileName)
	getLastIndex(&u.Domain)
	if err = ctx.SaveUploadedFile(file, savePath); err != nil {
		res.SetCode(400).SetMsg("上传文件失败").SetData(nil).Build(ctx)
		return
	}
	uri := fmt.Sprintf("%s%s", u.Domain, newFileName)
	res.SetCode(200).SetMsg("上传成功").SetData(uri).Build(ctx)
}

func getLastIndex(uri *string) {
	n := len(*uri)
	lastIndex := (*uri)[n-1 : n]
	if lastIndex != "/" {
		*uri = fmt.Sprintf("%s%s", *uri, "/")
	}
}

// ParseSize parses a size string like "10MB" and returns the number of bytes.
func ParseSize(sizeStr string) (int64, error) {
	// Remove any potential spaces from the string.
	sizeStr = strings.TrimSpace(sizeStr)

	// Use regular expression to match the number and unit.
	re := regexp.MustCompile(`^(\d+)([A-Za-z]+)$`)
	matches := re.FindStringSubmatch(sizeStr)
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid size format: %s", sizeStr)
	}

	// Parse the number as an int64.
	size, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, err
	}

	// Convert the unit to bytes.
	var bytes int64
	switch strings.ToUpper(matches[2]) {
	case "B":
		bytes = size
	case "KB":
		bytes = size * 1024
	case "MB":
		bytes = size * 1024 * 1024
	case "GB":
		bytes = size * 1024 * 1024 * 1024
	default:
		return 0, fmt.Errorf("unsupported size unit: %s", matches[2])
	}

	return bytes, nil
}
