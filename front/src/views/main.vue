<script>
import {GetCateGory} from "@/api/heritage";

export default {
  name:"main",
  data() {
    return {
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
  <el-container class="top-container">
    <el-header class="header">
      <div class="top-header">
        <h1 class="title">区块链非物质文化遗产</h1>
        <input class="search" placeholder="站内搜索">
      </div>
      <div class="bottom-header">
        <router-link to="" class="el-link">首页</router-link>
        <router-link to="" class="el-link">非遗名录</router-link>
        <router-link to="" class="el-link">首页</router-link>
        <router-link to="" class="el-link">首页</router-link>
      </div>
    </el-header>
    <el-main class="main">
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
    </el-main>
    <el-footer class="footer">
     <span>主管部门：文化和旅游厅</span>
    </el-footer>
  </el-container>
</template>

<style scoped>
.heritage-view-aside-item-box{
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100px;
  width: 100%;
  background-color: #baab7c;
  flex-direction: column;
}
.top-container{
  background-color: #ffffff;
}
.footer{
  height: 100px;
  background-color: #eee7d2;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 600;
}
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
.title{
  font-size: 40px;
  font-weight: 600;
  text-align: center;
  padding: 20px 0;
}
.search{
  width: 300px;
  height: 30px;
  border-radius: 20px;
  border: 1px solid #eee7d2;
  padding: 0 20px;
  font-size: 16px;
}
.header{
  height:200px;
  padding: 0;
}
.main{
  height: 100%;
  min-height: 100vh;
}
.top-header {
  height: 140px;
  display: flex; /* 启用Flexbox布局 */
  align-items: center; /* 垂直居中所有子元素 */
  justify-content: space-between; /* 子元素之间的空间平均分布 */
  flex-wrap: wrap; /* 如果需要，子元素可以换行显示 */
}
.bottom-header {
  height: 60px;
  background-color: #eee7d2;
  display: flex; /* 使用Flexbox布局 */
  justify-content: center; /* 水平居中子元素 */
  align-items: center; /* 垂直居中子元素 */
}
.el-link {
  color: black;
  font-size: 16px;
  font-weight: 600;
  padding: 0 10px;
  text-decoration: none;
}
</style>