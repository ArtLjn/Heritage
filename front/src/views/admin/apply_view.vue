<script>
import {useRoute} from "vue-router";
import router from "@/router";
import {AuditHeritageTask} from "@/api/heritage";

export default {
  name: "apply_view",
  components: {},
  data() {
    return {
      heritageInheritor:{
        inheritor:"",
        inheritorImg:"",
        project:"",
        cateGory:0,
        level:0,
        details:"",
        locate:""
      },
      heritageProject:{
        name:"",
        category:0,
        level:0,
        locate:"",
        details:""
      },
      cateGory:{
        0:"民间文学",
        1:"传统音乐",
        2:"传统舞蹈",
        3:"传统戏剧",
        4:"曲艺",
        5:"传统体育",
        6:"传统美术",
        7:"传统技艺",
        8:"传统医药",
        9:"民俗"
      },
      rank:{
        0:"国家级",
        1:"省级",
        2:"市级",
        3:"人类非物质文化遗产"
      },
      showInheritor:true,
      showProject:false,
      id:''
    }
  },
  mounted() {
    const route = useRoute();
    const raw  = route.query.raw.valueOf();
    const field = route.query.field;
    this.id =  route.query.id;
    if (raw === '1') {
      this.showInheritor = true;
      this.showProject = false;
      this.heritageInheritor = JSON.parse(field);
    } else {
      this.showInheritor = false;
      this.showProject = true;
      this.heritageProject = JSON.parse(field);
    }
  },
  methods: {
    back() {
      router.back();
    },
    handleAudit() {
      this.$messageBox.confirm('是否确认审核', '提示', {
        confirmButtonText: 'pass',
        cancelButtonText: 'refuse',
        type:"success",
        beforeClose: (action, instance, done) => {
          let res = false;
          if (action === 'confirm') {
            res = true;
          }
          instance.confirmButtonLoading = true;
          instance.confirmButtonText = '执行中...';
          AuditHeritageTask(this.id,res).then(res => {
            instance.confirmButtonLoading = false;
            done();
            this.$message({
              type: 'success',
              message: '审核成功!'
            });
            router.back();
          }).catch(err => {
            instance.confirmButtonLoading = false;
            done();
            this.$message({
              type: 'error',
              message: '审核失败!'
            });
          })
        }
      }).catch(err => {

      })
    }
  },
}
</script>

<template>

  <el-form v-if="showInheritor">
    <el-form-item label="inheritor">
      <span>{{heritageInheritor.inheritor}}</span>
    </el-form-item>
    <el-form-item label="inheritorImg">
      <img :src="heritageInheritor.inheritorImg" style="max-width: 200px;height: 200px;    background-size: cover;
    background-repeat: no-repeat;" alt="img">
    </el-form-item>
    <el-form-item label="project">
      <span>{{heritageInheritor.project}}</span>
    </el-form-item>
    <el-form-item label="cateGory">
      <span>{{cateGory[heritageInheritor.cateGory]}}</span>
    </el-form-item>
    <el-form-item label="level">
      <span>{{rank[heritageInheritor.level]}}</span>
    </el-form-item>
    <el-form-item label="details">
      <div v-html="heritageInheritor.details"></div>
    </el-form-item>
    <el-form-item label="locate">
      <span>{{heritageInheritor.locate}}</span>
    </el-form-item>
  </el-form>

  <el-form v-if="showProject">
    <el-form-item label="name">
      <span>{{heritageProject.name}}</span>
    </el-form-item>
    <el-form-item label="category">
      <span>{{cateGory[heritageProject.category]}}</span>
    </el-form-item>
    <el-form-item label="level">
      <span>{{rank[heritageProject.level]}}</span>
    </el-form-item>
    <el-form-item label="locate">
      <span>{{heritageProject.locate}}</span>
    </el-form-item>
    <el-form-item label="details">
      <div v-html="heritageProject.details"></div>
    </el-form-item>
  </el-form>

  <el-button-group >
    <el-button @click="back" type="primary">返回</el-button>
    <el-button @click="handleAudit">审核</el-button>
  </el-button-group>

</template>

<style scoped>

</style>