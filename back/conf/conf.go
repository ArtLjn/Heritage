/*
@Time : 2024/5/23 下午4:11
@Author : ljn
@File : conf
@Software: GoLand
*/

package conf

import (
	"back/util"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

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

	util.UploadRepo `json:"upload"`

	AdminAccount struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"adminAccount"`
}

func LoadEnvConfig(config *Conf, filename string) error {
	// 读取配置文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	// 使用环境变量覆盖配置文件内容
	if host := os.Getenv("SERVER_HOST"); host != "" {
		config.Server.Host = host
	}
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.Database.Port = dbPort
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.Database.User = dbUser
	}
	if dbPassword := os.Getenv("MYSQL_PASSWORD"); dbPassword != "" {
		config.Database.Password = dbPassword
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.Database.DbName = dbName
	}
	if redisAddr := os.Getenv("REDIS_ADDR"); redisAddr != "" {
		config.Redis.Addr = redisAddr
	}
	if redisPassword := os.Getenv("REDIS_PASSWORD"); redisPassword != "" {
		config.Redis.Password = redisPassword
	}
	if rabbitmqUrl := os.Getenv("RABBITMQ_URL"); rabbitmqUrl != "" {
		config.Rabbitmq.Url = rabbitmqUrl
	}
	if adminUsername := os.Getenv("ADMIN_NAME"); adminUsername != "" {
		config.AdminAccount.Username = adminUsername
	}
	if adminPassword := os.Getenv("ADMIN_PASSWORD"); adminPassword != "" {
		config.AdminAccount.Password = adminPassword
	}
	if contractAddress := os.Getenv("CONTRACT_ADDRESS"); contractAddress != "" {
		config.CommonRequest.ContractAddress = contractAddress
	}
	if contractUser := os.Getenv("CONTRACT_USER"); contractUser != "" {
		config.CommonRequest.User = contractUser
	}
	if ip := os.Getenv("IP"); ip != "" {
		config.CommonRequest.Url = strings.Replace(config.CommonRequest.Url, "localhost", ip, -1)
		config.CommonRequest.ParsePut = strings.Replace(config.CommonRequest.ParsePut, "localhost", ip, -1)
		config.UploadRepo.Domain = strings.Replace(config.UploadRepo.Domain, "localhost", ip, -1)
	}
	return nil
}
