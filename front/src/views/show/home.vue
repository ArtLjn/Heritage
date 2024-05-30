<script>

import Header from "@/components/show/Header";
import Footer from "@/components/show/Footer";
export default {
  components:{
    Header,
    Footer
  },
}
</script>

<template>
    <el-container class="top-container">
      <el-header class="header">
        <Header/>
      </el-header>
      <el-main class="main">
        <router-view></router-view>
      </el-main>
      <el-footer class="footer">
        <Footer/>
      </el-footer>
    </el-container>
</template>

<style scoped>
@import "../../assets/css/common.css";

.top-container{
  background-color: #ffffff;
}

</style>