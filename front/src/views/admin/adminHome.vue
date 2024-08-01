<script>
import Aside from "@/components/admin/Aside.vue";
export default {
  components:{
    Aside,
  }
}
</script>

<template>
<el-container class="el-container">
  <el-aside class="el-aside">
    非物质文化遗产
    <Aside></Aside>
  </el-aside>
  <el-container>
    <el-main class="el-main">
      <RouterView></RouterView>
    </el-main>
  </el-container>
</el-container>
</template>

<style scoped>
.el-main{
  background-color: #f5f5f5;
  height: auto;
  min-height: 100vh;
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
  background-color: #282C34;
  position: fixed;
}
.el-container{
  /* background-color: #f1724c; */
  min-height: 100vh;
  height: 100%;
}

</style>