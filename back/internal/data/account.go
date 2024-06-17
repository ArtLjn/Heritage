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
	"crypto/ecdsa"
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
