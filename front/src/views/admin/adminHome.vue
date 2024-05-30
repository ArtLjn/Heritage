<script>
import Aside from "@/components/admin/Aside.vue";
import Header from "@/components/admin/Header.vue";
export default {
  components:{
    Aside,
    Header,
  }
}
</script>

<template>
<el-container class="el-container">
  <el-aside class="el-aside">
    测试平台系统
    <Aside></Aside>
  </el-aside>
  <el-container>
    <el-header class="el-header">
      <Header></Header>
    </el-header>
    <el-main class="el-main">
      <RouterView></RouterView>
    </el-main>
  </el-container>
</el-container>
</template>

<style scoped>
.el-header {
  width: 100%;
  padding-top: 15px;
  background-color: #334157;
  border-left: 2px solid #333;
  position: fixed;
  margin-left: 180px;
  z-index: 999;
  display: flex; /* 将header设置为弹性容器 */
  justify-content: space-between; /* 按钮和左侧元素之间留出相等的空间 */
}
.el-main{
  background-color: #f5f5f5;
  height: auto;
  min-height: 100vh;
  margin-top: 60px;
  width: calc(100% - 180px);
  margin-left: 180px;
}
.el-aside{
  padding-top: 20px;
  width: 180px;
  color: white;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #334157;
  position: fixed;
}
.el-container{
  /* background-color: #f1724c; */
  min-height: 100vh;
  height: 100%;
}

</style>