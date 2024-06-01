<script>
import heritageInheritor from "@/components/common/heritageInheritor.vue";
import heritageProject from "@/components/common/heritageProject.vue";
export default {
  data() {
    return{
      activeName: 'first',
      options:
          [{"children":[{"label":"长春市"},{"label":"吉林市"},{"label":"四平市"},{"label":"辽源市"},{"label":"通化市"},{"label":"白山市"},{"label":"松原市"},{"label":"白城市"},{"label":"延边朝鲜族自治州"}],"label":"吉林省"},{"children":[{"label":"贵阳市"},{"label":"六盘水市"},{"label":"遵义市"},{"label":"安顺市"},{"label":"铜仁市"},{"label":"黔西南布依族苗族自治州"},{"label":"毕节市"},{"label":"黔东南苗族侗族自治州"},{"label":"黔南布依族苗族自治州"}],"label":"贵州省"},{"children":[{"label":"台北市"},{"label":"高雄市"},{"label":"台南市"},{"label":"台中市"},{"label":"金门县"},{"label":"南投县"},{"label":"基隆市"},{"label":"新竹市"},{"label":"嘉义市"},{"label":"新北市"},{"label":"宜兰县"},{"label":"新竹县"},{"label":"桃园县"},{"label":"苗栗县"},{"label":"彰化县"},{"label":"嘉义县"},{"label":"云林县"},{"label":"屏东县"},{"label":"台东县"},{"label":"花莲县"},{"label":"澎湖县"},{"label":"连江县"}],"label":"台湾"},{"children":[{"label":"沈阳市"},{"label":"大连市"},{"label":"鞍山市"},{"label":"抚顺市"},{"label":"本溪市"},{"label":"丹东市"},{"label":"锦州市"},{"label":"营口市"},{"label":"阜新市"},{"label":"辽阳市"},{"label":"盘锦市"},{"label":"铁岭市"},{"label":"朝阳市"},{"label":"葫芦岛市"}],"label":"辽宁省"},{"children":[{"label":"南宁市"},{"label":"柳州市"},{"label":"桂林市"},{"label":"梧州市"},{"label":"北海市"},{"label":"防城港市"},{"label":"钦州市"},{"label":"贵港市"},{"label":"玉林市"},{"label":"百色市"},{"label":"贺州市"},{"label":"河池市"},{"label":"来宾市"},{"label":"崇左市"}],"label":"广西壮族自治区"},{"children":[{"label":"广州市"},{"label":"韶关市"},{"label":"深圳市"},{"label":"珠海市"},{"label":"汕头市"},{"label":"佛山市"},{"label":"江门市"},{"label":"湛江市"},{"label":"茂名市"},{"label":"肇庆市"},{"label":"惠州市"},{"label":"梅州市"},{"label":"汕尾市"},{"label":"河源市"},{"label":"阳江市"},{"label":"清远市"},{"label":"东莞市"},{"label":"中山市"},{"label":"东沙群岛"},{"label":"潮州市"},{"label":"揭阳市"},{"label":"云浮市"}],"label":"广东省"},{"children":[{"label":"澳门半岛"},{"label":"离岛"}],"label":"澳门特别行政区"},{"children":[{"label":"南京市"},{"label":"无锡市"},{"label":"徐州市"},{"label":"常州市"},{"label":"苏州市"},{"label":"南通市"},{"label":"连云港市"},{"label":"淮安市"},{"label":"盐城市"},{"label":"扬州市"},{"label":"镇江市"},{"label":"泰州市"},{"label":"宿迁市"}],"label":"江苏省"},{"children":[{"label":"南昌市"},{"label":"景德镇市"},{"label":"萍乡市"},{"label":"九江市"},{"label":"新余市"},{"label":"鹰潭市"},{"label":"赣州市"},{"label":"吉安市"},{"label":"宜春市"},{"label":"抚州市"},{"label":"上饶市"}],"label":"江西省"},{"children":[{"label":"海口市"},{"label":"三亚市"},{"label":"三沙市"}],"label":"海南省"},{"children":[{"label":"成都市"},{"label":"自贡市"},{"label":"攀枝花市"},{"label":"泸州市"},{"label":"德阳市"},{"label":"绵阳市"},{"label":"广元市"},{"label":"遂宁市"},{"label":"内江市"},{"label":"乐山市"},{"label":"南充市"},{"label":"眉山市"},{"label":"宜宾市"},{"label":"广安市"},{"label":"达州市"},{"label":"雅安市"},{"label":"巴中市"},{"label":"资阳市"},{"label":"阿坝藏族羌族自治州"},{"label":"甘孜藏族自治州"},{"label":"凉山彝族自治州"}],"label":"四川省"},{"children":[{"label":"香港岛"},{"label":"九龙"},{"label":"新界"}],"label":"香港特别行政区"},{"children":[{"label":"太原市"},{"label":"大同市"},{"label":"阳泉市"},{"label":"长治市"},{"label":"晋城市"},{"label":"朔州市"},{"label":"晋中市"},{"label":"运城市"},{"label":"忻州市"},{"label":"临汾市"},{"label":"吕梁市"}],"label":"山西省"},{"children":[{"label":"上海市"}],"label":"上海"},{"children":[{"label":"合肥市"},{"label":"芜湖市"},{"label":"蚌埠市"},{"label":"淮南市"},{"label":"马鞍山市"},{"label":"淮北市"},{"label":"铜陵市"},{"label":"安庆市"},{"label":"黄山市"},{"label":"滁州市"},{"label":"阜阳市"},{"label":"宿州市"},{"label":"六安市"},{"label":"亳州市"},{"label":"池州市"},{"label":"宣城市"}],"label":"安徽省"},{"children":[{"label":"西安市"},{"label":"铜川市"},{"label":"宝鸡市"},{"label":"咸阳市"},{"label":"渭南市"},{"label":"延安市"},{"label":"汉中市"},{"label":"榆林市"},{"label":"安康市"},{"label":"商洛市"}],"label":"陕西省"},{"children":[{"label":"北京市"}],"label":"北京"},{"children":[{"label":"哈尔滨市"},{"label":"齐齐哈尔市"},{"label":"鸡西市"},{"label":"鹤岗市"},{"label":"双鸭山市"},{"label":"大庆市"},{"label":"伊春市"},{"label":"佳木斯市"},{"label":"七台河市"},{"label":"牡丹江市"},{"label":"黑河市"},{"label":"绥化市"},{"label":"大兴安岭地区"}],"label":"黑龙江省"},{"children":[{"label":"郑州市"},{"label":"开封市"},{"label":"洛阳市"},{"label":"平顶山市"},{"label":"安阳市"},{"label":"鹤壁市"},{"label":"新乡市"},{"label":"焦作市"},{"label":"濮阳市"},{"label":"许昌市"},{"label":"漯河市"},{"label":"三门峡市"},{"label":"南阳市"},{"label":"商丘市"},{"label":"信阳市"},{"label":"周口市"},{"label":"驻马店市"}],"label":"河南省"},{"children":[{"label":"长沙市"},{"label":"株洲市"},{"label":"湘潭市"},{"label":"衡阳市"},{"label":"邵阳市"},{"label":"岳阳市"},{"label":"常德市"},{"label":"张家界市"},{"label":"益阳市"},{"label":"郴州市"},{"label":"永州市"},{"label":"怀化市"},{"label":"娄底市"},{"label":"湘西土家族苗族自治州"}],"label":"湖南省"},{"children":[{"label":"昆明市"},{"label":"曲靖市"},{"label":"玉溪市"},{"label":"保山市"},{"label":"昭通市"},{"label":"丽江市"},{"label":"普洱市"},{"label":"临沧市"},{"label":"楚雄彝族自治州"},{"label":"红河哈尼族彝族自治州"},{"label":"文山壮族苗族自治州"},{"label":"西双版纳傣族自治州"},{"label":"大理白族自治州"},{"label":"德宏傣族景颇族自治州"},{"label":"怒江傈僳族自治州"},{"label":"迪庆藏族自治州"}],"label":"云南省"},{"children":[{"label":"拉萨市"},{"label":"昌都市"},{"label":"山南地区"},{"label":"日喀则市"},{"label":"那曲地区"},{"label":"阿里地区"},{"label":"林芝市"}],"label":"西藏自治区"},{"children":[{"label":"兰州市"},{"label":"嘉峪关市"},{"label":"金昌市"},{"label":"白银市"},{"label":"天水市"},{"label":"武威市"},{"label":"张掖市"},{"label":"平凉市"},{"label":"酒泉市"},{"label":"庆阳市"},{"label":"定西市"},{"label":"陇南市"},{"label":"临夏回族自治州"},{"label":"甘南藏族自治州"}],"label":"甘肃省"},{"children":[{"label":"呼和浩特市"},{"label":"包头市"},{"label":"乌海市"},{"label":"赤峰市"},{"label":"通辽市"},{"label":"鄂尔多斯市"},{"label":"呼伦贝尔市"},{"label":"巴彦淖尔市"},{"label":"乌兰察布市"},{"label":"兴安盟"},{"label":"锡林郭勒盟"},{"label":"阿拉善盟"}],"label":"内蒙古自治区"},{"children":[{"label":"武汉市"},{"label":"黄石市"},{"label":"十堰市"},{"label":"宜昌市"},{"label":"襄阳市"},{"label":"鄂州市"},{"label":"荆门市"},{"label":"孝感市"},{"label":"荆州市"},{"label":"黄冈市"},{"label":"咸宁市"},{"label":"随州市"},{"label":"恩施土家族苗族自治州"}],"label":"湖北省"},{"children":[{"label":"福州市"},{"label":"厦门市"},{"label":"莆田市"},{"label":"三明市"},{"label":"泉州市"},{"label":"漳州市"},{"label":"南平市"},{"label":"龙岩市"},{"label":"宁德市"}],"label":"福建省"},{"children":[{"label":"石家庄市"},{"label":"唐山市"},{"label":"秦皇岛市"},{"label":"邯郸市"},{"label":"邢台市"},{"label":"保定市"},{"label":"张家口市"},{"label":"承德市"},{"label":"沧州市"},{"label":"廊坊市"},{"label":"衡水市"}],"label":"河北省"},{"children":[{"label":"杭州市"},{"label":"宁波市"},{"label":"温州市"},{"label":"嘉兴市"},{"label":"湖州市"},{"label":"绍兴市"},{"label":"金华市"},{"label":"衢州市"},{"label":"舟山市"},{"label":"台州市"},{"label":"丽水市"}],"label":"浙江省"},{"children":[{"label":"济南市"},{"label":"青岛市"},{"label":"淄博市"},{"label":"枣庄市"},{"label":"东营市"},{"label":"烟台市"},{"label":"潍坊市"},{"label":"济宁市"},{"label":"泰安市"},{"label":"威海市"},{"label":"日照市"},{"label":"莱芜市"},{"label":"临沂市"},{"label":"德州市"},{"label":"聊城市"},{"label":"滨州市"},{"label":"菏泽市"}],"label":"山东省"},{"children":[{"label":"重庆市"}],"label":"重庆"},{"children":[{"label":"西宁市"},{"label":"海东市"},{"label":"海北藏族自治州"},{"label":"黄南藏族自治州"},{"label":"海南藏族自治州"},{"label":"果洛藏族自治州"},{"label":"玉树藏族自治州"},{"label":"海西蒙古族藏族自治州"}],"label":"青海省"},{"children":[{"label":"银川市"},{"label":"石嘴山市"},{"label":"吴忠市"},{"label":"固原市"},{"label":"中卫市"}],"label":"宁夏回族自治区"},{"children":[{"label":"乌鲁木齐市"},{"label":"克拉玛依市"},{"label":"吐鲁番市"},{"label":"哈密地区"},{"label":"昌吉回族自治州"},{"label":"博尔塔拉蒙古自治州"},{"label":"巴音郭楞蒙古自治州"},{"label":"阿克苏地区"},{"label":"克孜勒苏柯尔克孜自治州"},{"label":"喀什地区"},{"label":"和田地区"},{"label":"伊犁哈萨克自治州"},{"label":"塔城地区"},{"label":"阿勒泰地区"}],"label":"新疆维吾尔自治区"},{"children":[{"label":"天津市"}],"label":"天津"}]
      ,currentCity:'吉林省'
    }
  },
  components:{
    heritageInheritor,
    heritageProject
  },
  watch:{
    'currentCity'(newValue) {
      let len = newValue.length
      if (len === 2) {
        this.currentCity = newValue[0] + '-' + newValue[1]
        len = null
      } else if (len === 1) {
        this.currentCity = newValue[0]
      }
    }
  }
}
</script>

<template>
  <el-container class="carousel-box">
    <el-main class="main-content">
      <el-tabs v-model="activeName"  type="border-card" style="
      border-bottom-right-radius: 20px;border-bottom-left-radius: 20px">
        <el-cascader placeholder="选择省市"
                     filterable
                     style="margin-bottom: 10px"
                     v-model="currentCity"
                     :options="options" :props="{checkStrictly: true, value: 'label', label: 'label' }"
                     clearable />
        <el-tab-pane label="非遗人" name="first">
          <heritage-inheritor :city="currentCity" :key="currentCity"/>
        </el-tab-pane>
        <el-tab-pane label="非遗项目" name="second">
          <heritage-project :city="currentCity" :key="currentCity"/>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>

<style scoped>
.carousel-box {
  height: 100vh;
  overflow: hidden;
  background-image: linear-gradient(to right bottom, #8e44ad, #f1c40f); /* 更改为较浅的紫色到黄色渐变 */
}
</style>