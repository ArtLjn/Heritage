<script>
export default {
  data() {
    return{
      role:localStorage.getItem('role'),
      city:localStorage.getItem("city")
    }
  }
}
</script>

<template>
  <div class="link-box">
    <router-link to="/adminHome/index" class="link" v-if="city === 'admin'">首页</router-link>
    <router-link to="/adminHome/apply" v-if="city !== 'admin'" class="link">申请</router-link>
    <router-link to="/adminHome/heritage" v-if="city !== 'admin'" class="link">遗产</router-link>
    <router-link to="/adminHome/individual"  class="link">个人</router-link>
    <router-link to="/adminHome/manager" class="link" v-if="city === 'admin'">管理</router-link>
  </div>
</template>
<style scoped>
.link-box{
  display: flex;
  margin-top: 40px;
  flex-direction: column;
}
.link {
  color: white;
  text-decoration: none;
  font-size: 15px;
  padding: 10px 40px;
  margin: 10px 10px;
}
.link:hover{
  background-color: rgb(100, 100, 100);
}
</style>