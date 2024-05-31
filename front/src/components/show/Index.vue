<script>

import {GetCateGory} from "@/api/heritage";

export default {
  data() {
    return{
      cateGoryList:[]
    }
  },
  mounted() {
    GetCateGory().then((res) => {
      if (res.data.code === 200) {
        const list = res.data.data[0];
        Object.keys(list).forEach(key => {
          this.cateGoryList.push(key)
        })
      }
    })
  }
}
</script>

<template>
  <el-container class="carousel-box">
  </el-container>
</template>

<style scoped>
.carousel-box{
  background-color: red;
  width: 100%;
  height: 100%;
}
</style>