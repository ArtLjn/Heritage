<script>
import {GetAllAccount,ExportAccount} from "@/api/auth";
import axios from "axios";

export default {
  name: "ManagerAccount",
  data() {
    return {
      Accounts:[]
    }
  },
  mounted() {
    GetAllAccount().then(res => {
      this.Accounts = res.data.data[0]
    })
  },
  methods:{
    async downloadFile() {
      try {
        const response = await axios({
          url: '/api/account/exportAccount', // 这里是你后端接口的URL
          method: 'GET',
          responseType: 'blob', // 确保接收的是二进制流
          headers:{
            Authorization:localStorage.getItem('token')
          }
        });

        // 创建一个URL链接到二进制数据
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'Account.xlsx'); // 设置下载的文件名
        document.body.appendChild(link);
        link.click();

        // 释放URL对象
        window.URL.revokeObjectURL(url);
        document.body.removeChild(link);
      } catch (error) {
        console.error('下载文件失败', error);
      }
    }
  }
}
</script>

<template>
  <el-button type="warning" plain style="margin-bottom: 20px;width: 100px" @click="downloadFile">导出Excel</el-button>
  <el-table :data="Accounts" class="custom-table"  stripe :lazy="true" border height="700">
    <el-table-column label="ID" width="100" prop="id"></el-table-column>
    <el-table-column label="Address" width="400" prop="address"></el-table-column>
    <el-table-column label="Password" width="120" prop="password"></el-table-column>
    <el-table-column label="City" width="200" prop="city"></el-table-column>
    <el-table-column label="Level" width="100" prop="level"></el-table-column>
  </el-table>
</template>

<style scoped>

</style>