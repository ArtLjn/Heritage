<script>
export default {
  data() {
    return{
      role:localStorage.getItem('role')
    }
  }
}
</script>

<template>
  <div class="link-box">
    <router-link to="/adminHome/index" class="link">首页</router-link>
    <router-link to="/adminHome/apply" class="link">申请</router-link>
    <router-link to="/adminHome/individual" class="link">个人</router-link>
    <router-link to="/adminHome/browse" class="link">Browse</router-link>
    <router-link to="/adminHome/order" class="link">Order</router-link>
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