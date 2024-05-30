<script>

import {GetCateGory} from "@/api/heritage";

export default {
  data() {
    return{
      imgList:[
        "https://www.jsfybh.com/intangiblemodule/mnt/sdc/img/heirImg/2024-05/2.jpg",
        "https://www.jsfybh.com/intangiblemodule/mnt/sdc/img/heirImg/2024-05/3.png",
        "https://www.jsfybh.com/intangiblemodule/mnt/sdc/img/heirImg/2024-05/4.png",
        "https://www.jsfybh.com/intangiblemodule/mnt/sdc/img/heirImg/2024-05/1.jpg"
      ],
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
  <div class="carousel-box">
    <el-carousel trigger="click" class="carousel" >
      <el-carousel-item v-for="item in imgList" :key="item" class="el-carousel-item">
        <img :src="item" alt="" class="el-empty__image">
      </el-carousel-item>
    </el-carousel>
  </div>
  <el-container class="heritage-view">
    <el-aside  class="heritage-view-aside">
      {{cateGoryList}}
      <router-link to=""></router-link>
    </el-aside>
    <el-container class="heritage-view-main">
      shabi
    </el-container>
    <el-aside class="heritage-view-aside-right">
      <span>非遗</span>
    </el-aside>
  </el-container>
</template>

<style scoped>
.heritage-view-aside-right {
  width: 550px;
  background-color: #baab7c;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
}
.heritage-view-main {
  width: 100%;
  height: 100%;
  background-color: #eee7d2;
}
.heritage-view-aside {
  width: 200px;
  background-color: #baab7c;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
}
.heritage-view {
  margin-top: 10px;
  background-color: #bc982c;
  height: 500px;
}
.el-carousel-item {
  width: 100%;
  height: 100%;
}
.el-empty__image {
  width: 100%; /* 图片宽度占满轮播图 */
  height: 100%; /* 图片高度占满轮播图 */
  min-height: 400px;
  object-fit: cover; /* 保持图片比例，裁剪超出部分 */
}
.carousel-box{
  height: 400px;
  background-color: #baab7c;
  padding: 0;
  margin: 0;
  width: 100%; /* 占据所有宽度 */
  overflow: hidden; /* 隐藏溢出的内容 */
}
.carousel{
  height: 400px;
  padding: 0;
}
</style>