<script>
import {useRoute} from "vue-router";
import router from "@/router";

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
      showProject:false
    }
  },
  mounted() {
    const route = useRoute();
    const raw  = route.query.raw.valueOf();
    const field = route.query.field;
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

  <el-button @click="back" type="primary">返回</el-button>
</template>

<style scoped>

</style>