<script>
import {LogOut} from "@/api/auth";

export default {
  name: "individual",
  methods: {LogOut},
  data() {
    return {
      individual:{
        accountName: localStorage.getItem("name"),
        city:"",
        rank: ""
      },
    }
  },
  mounted() {
    this.individual.city = localStorage.getItem("city")
    const rank = localStorage.getItem("rank")
    switch (rank) {
      case '1':
        this.individual.rank = "province"
        break
      case '2':
        this.individual.rank = "city"
        break
    }
  }
}
</script>


<template>
  <div class="user-info-container">
    <el-card class="user-info-card">
      <div slot="header" class="clearfix">
        <span>User Information</span>
        <el-button type="danger" class="logout" @click="LogOut()">Logout</el-button>
      </div>
      <el-form label-position="top" :model="individual" class="user-info-form">
        <el-form-item label="account">
          <el-input v-model="individual.accountName" disabled></el-input>
        </el-form-item>
        <el-form-item label="City">
          <el-input v-model="individual.city" disabled></el-input>
        </el-form-item>
        <el-form-item label="Rank" v-if="individual.city !== 'admin'">
          <el-tag type="success">
            {{ individual.rank }}
          </el-tag>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>


<style scoped>
.user-info-container {
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f0f2f5;
}

.user-info-card {
  width: 600px;
  background: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.logout {
  float: right;
}

.user-info-form {
  margin-top: 20px;
}
</style>
