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

func (a accountRepo) InitAccount() {
	city := make(map[string][]string)
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	if err := json.Unmarshal([]byte(cityJson), &city); err != nil {
		panic(err)
	}
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

const cityJson = `{"上海":["上海市"],"云南省":["昆明市","曲靖市","玉溪市","保山市","昭通市","丽江市","普洱市","临沧市","楚雄彝族自治州","红河哈尼族彝族自治州","文山壮族苗族自治州","西双版纳傣族自治州","大理白族自治州","德宏傣族景颇族自治州","怒江傈僳族自治州","迪庆藏族自治州"],"内蒙古自治区":["呼和浩特市","包头市","乌海市","赤峰市","通辽市","鄂尔多斯市","呼伦贝尔市","巴彦淖尔市","乌兰察布市","兴安盟","锡林郭勒盟","阿拉善盟"],"北京":["北京市"],"台湾":["台北市","高雄市","台南市","台中市","金门县","南投县","基隆市","新竹市","嘉义市","新北市","宜兰县","新竹县","桃园县","苗栗县","彰化县","嘉义县","云林县","屏东县","台东县","花莲县","澎湖县","连江县"],"吉林省":["长春市","吉林市","四平市","辽源市","通化市","白山市","松原市","白城市","延边朝鲜族自治州"],"四川省":["成都市","自贡市","攀枝花市","泸州市","德阳市","绵阳市","广元市","遂宁市","内江市","乐山市","南充市","眉山市","宜宾市","广安市","达州市","雅安市","巴中市","资阳市","阿坝藏族羌族自治州","甘孜藏族自治州","凉山彝族自治州"],"天津":["天津市"],"宁夏回族自治区":["银川市","石嘴山市","吴忠市","固原市","中卫市"],"安徽省":["合肥市","芜湖市","蚌埠市","淮南市","马鞍山市","淮北市","铜陵市","安庆市","黄山市","滁州市","阜阳市","宿州市","六安市","亳州市","池州市","宣城市"],"山东省":["济南市","青岛市","淄博市","枣庄市","东营市","烟台市","潍坊市","济宁市","泰安市","威海市","日照市","莱芜市","临沂市","德州市","聊城市","滨州市","菏泽市"],"山西省":["太原市","大同市","阳泉市","长治市","晋城市","朔州市","晋中市","运城市","忻州市","临汾市","吕梁市"],"广东省":["广州市","韶关市","深圳市","珠海市","汕头市","佛山市","江门市","湛江市","茂名市","肇庆市","惠州市","梅州市","汕尾市","河源市","阳江市","清远市","东莞市","中山市","东沙群岛","潮州市","揭阳市","云浮市"],"广西壮族自治区":["南宁市","柳州市","桂林市","梧州市","北海市","防城港市","钦州市","贵港市","玉林市","百色市","贺州市","河池市","来宾市","崇左市"],"新疆维吾尔自治区":["乌鲁木齐市","克拉玛依市","吐鲁番市","哈密地区","昌吉回族自治州","博尔塔拉蒙古自治州","巴音郭楞蒙古自治州","阿克苏地区","克孜勒苏柯尔克孜自治州","喀什地区","和田地区","伊犁哈萨克自治州","塔城地区","阿勒泰地区"],"江苏省":["南京市","无锡市","徐州市","常州市","苏州市","南通市","连云港市","淮安市","盐城市","扬州市","镇江市","泰州市","宿迁市"],"江西省":["南昌市","景德镇市","萍乡市","九江市","新余市","鹰潭市","赣州市","吉安市","宜春市","抚州市","上饶市"],"河北省":["石家庄市","唐山市","秦皇岛市","邯郸市","邢台市","保定市","张家口市","承德市","沧州市","廊坊市","衡水市"],"河南省":["郑州市","开封市","洛阳市","平顶山市","安阳市","鹤壁市","新乡市","焦作市","濮阳市","许昌市","漯河市","三门峡市","南阳市","商丘市","信阳市","周口市","驻马店市"],"浙江省":["杭州市","宁波市","温州市","嘉兴市","湖州市","绍兴市","金华市","衢州市","舟山市","台州市","丽水市"],"海南省":["海口市","三亚市","三沙市"],"湖北省":["武汉市","黄石市","十堰市","宜昌市","襄阳市","鄂州市","荆门市","孝感市","荆州市","黄冈市","咸宁市","随州市","恩施土家族苗族自治州"],"湖南省":["长沙市","株洲市","湘潭市","衡阳市","邵阳市","岳阳市","常德市","张家界市","益阳市","郴州市","永州市","怀化市","娄底市","湘西土家族苗族自治州"],"澳门特别行政区":["澳门半岛","离岛"],"甘肃省":["兰州市","嘉峪关市","金昌市","白银市","天水市","武威市","张掖市","平凉市","酒泉市","庆阳市","定西市","陇南市","临夏回族自治州","甘南藏族自治州"],"福建省":["福州市","厦门市","莆田市","三明市","泉州市","漳州市","南平市","龙岩市","宁德市"],"西藏自治区":["拉萨市","昌都市","山南地区","日喀则市","那曲地区","阿里地区","林芝市"],"贵州省":["贵阳市","六盘水市","遵义市","安顺市","铜仁市","黔西南布依族苗族自治州","毕节市","黔东南苗族侗族自治州","黔南布依族苗族自治州"],"辽宁省":["沈阳市","大连市","鞍山市","抚顺市","本溪市","丹东市","锦州市","营口市","阜新市","辽阳市","盘锦市","铁岭市","朝阳市","葫芦岛市"],"重庆":["重庆市"],"陕西省":["西安市","铜川市","宝鸡市","咸阳市","渭南市","延安市","汉中市","榆林市","安康市","商洛市"],"青海省":["西宁市","海东市","海北藏族自治州","黄南藏族自治州","海南藏族自治州","果洛藏族自治州","玉树藏族自治州","海西蒙古族藏族自治州"],"香港特别行政区":["香港岛","九龙","新界"],"黑龙江省":["哈尔滨市","齐齐哈尔市","鸡西市","鹤岗市","双鸭山市","大庆市","伊春市","佳木斯市","七台河市","牡丹江市","黑河市","绥化市","大兴安岭地区"]}`

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
