/*
@Time : 2024/5/25 下午12:32
@Author : ljn
@File : account
@Software: GoLand
*/

package data

import (
	"back/internal/model"
	"back/internal/service"
	"back/util"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"log"
)

type accountRepo struct {
	data *Data
	t    util.TokenManager
}

func (a accountRepo) QueryAccountByCity(city string) model.Account {
	var m model.Account
	if err := a.data.db.Where("city = ?", city).First(&m).Error; err != nil {
		fmt.Println(err)
		return model.Account{}
	}
	return m
}

func (a accountRepo) GetAllAccount() []model.Account {
	var acc []model.Account
	if err := a.data.db.Find(&acc).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return acc
}

func (a accountRepo) SetCacheList(key string, o []model.Account) {
	for _, value := range o {
		serializedValue, err := json.Marshal(value)
		if err != nil {
			log.Println(err)
			continue
		}
		i, e := a.data.rdb.RPush(context.Background(), key, serializedValue).Result()
		if e != nil {
			log.Println("Failed to push serialized value to Redis list:", e, i)
			continue
		}
	}
	log.Println("SetCacheList", key)
}

func (a accountRepo) VerifyToken(key string) error {
	return a.t.VerifyToken(key)
}

func (a accountRepo) LogOut(key string) {
	a.t.LogOutToken(key)
}

func (a accountRepo) Login(login model.AccountLogin) (data []interface{}, err error) {
	acc := &model.Account{}
	admin := a.data.c.AdminAccount
	if admin.Username == login.Address && admin.Password == login.Password {
		return []interface{}{"admin", nil, a.t.SaveToken(login.Address)}, nil
	}
	if err = a.data.db.Where("address = ? && password = ?",
		login.Address, login.Password).First(&acc).Error; err != nil {
		return nil, err
	}
	if len(acc.City) != 0 {
		token := a.t.SaveToken(acc.City)
		return []interface{}{acc.City, acc.Level, token}, nil
	}
	return nil, fmt.Errorf("登录失败")
}

func (a accountRepo) RestoreList(key string) []interface{} {
	var newList []interface{}
	serializedValues, err := a.data.rdb.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		log.Println("Failed to retrieve values from Redis list:", err)
		panic(err)
	}
	for _, serializedValue := range serializedValues {
		value := make(map[string]interface{})
		if e := json.Unmarshal([]byte(serializedValue), &value); e != nil {
			log.Println("Failed to deserialize value from JSON:", e)
			continue
		}
		newList = append(newList, &value)
	}
	return newList
}
func (a accountRepo) GetCacheAccountList() []model.Account {
	var (
		acd []model.Account
		acs []interface{}
	)
	acs = a.RestoreList("heritageAccount")
	if len(acs) != 0 {
		b, e := json.Marshal(&acs)
		if e != nil {
			panic(e)
		}
		if err := json.Unmarshal(b, &acd); err != nil {
			panic(err)
		}
		return acd
	}
	return nil
}

func (a accountRepo) InitAccount(city map[string][]string) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	mx.Lock()
	defer mx.Unlock()

	var existingAccounts []model.Account
	if err := a.data.db.Find(&existingAccounts).Error; err != nil {
		log.Fatal(err)
	}

	existingCityMap := make(map[string]bool)
	for _, account := range existingAccounts {
		existingCityMap[account.City] = true
	}
	var newAccounts []model.Account

	for k, v := range city {
		if _, exists := existingCityMap[k]; !exists {
			newAccounts = append(newAccounts, model.Account{
				Address:  GenerateAccount()["address"],
				Password: uuid.New().String()[:8],
				City:     k,
				Level:    model.ProvincialLevel,
			})
		}
		for _, c := range v {
			prefix := fmt.Sprintf("%s-%s", k, c)
			if _, exists := existingCityMap[prefix]; !exists {
				newAccounts = append(newAccounts, model.Account{
					Address:  GenerateAccount()["address"],
					Password: uuid.New().String()[:8],
					City:     prefix,
					Level:    model.CityLevel,
				})
			}
		}
	}
	if len(newAccounts) > 0 {
		if err := a.data.db.Create(&newAccounts).Error; err != nil {
			log.Fatal(err)
		}
	}
}

func NewAccountRepo(data *Data) service.AccountRepo {
	return &accountRepo{data: data, t: util.NewToken(data.rdb)}
}

func GenerateAccount() map[string]string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	_ = crypto.FromECDSAPub(publicKeyECDSA)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	data := map[string]string{
		"address":    address,
		"privateKey": hexutil.Encode(privateKeyBytes)[2:],
	}
	return data
}
