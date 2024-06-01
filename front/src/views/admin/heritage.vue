<script>
import heritageInheritor from "@/components/common/heritageInheritor.vue";
import heritageProject from "@/components/common/heritageProject.vue";
export default {
  data() {
    return {
      activeName: 'first',
      city: localStorage.getItem("city")
    }
  },
  components:{
    heritageInheritor,
    heritageProject
  }
}
</script>

<template>
  <el-tabs v-model="activeName"  type="border-card">
    <el-tab-pane label="非遗人" name="first">
      <heritage-inheritor :city="city"/>
    </el-tab-pane>
    <el-tab-pane label="非遗项目" name="second">
      <heritage-project :city="city"/>
    </el-tab-pane>
  </el-tabs>
</template>

<style scoped>

</style>