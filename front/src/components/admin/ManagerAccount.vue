<script>
import {GetAllAccount} from "@/api/auth";
import axios from "axios";

export default {
  name: "ManagerAccount",
  data() {
    return {
      Accounts:[],
      currentCity:'',
      AccountCopy:[]
    }
  },
  mounted() {
    GetAllAccount().then(res => {
      this.Accounts = res.data.data[0]
      this.AccountCopy = res.data.data[0]
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
  },
  watch:{
    'currentCity'(newValue) {
      const keywords = newValue.toLowerCase().split(' '); // 拆分用户输入的关键字并转换为小写
      this.Accounts = this.AccountCopy.map(acc => {
        let matches = 0;
        const regex = new RegExp(keywords.join('|'), 'gi'); // 构建正则表达式，匹配所有关键字，忽略大小写
        const matchContent = `${acc.id} ${acc.city} ${acc.level}`.toLowerCase(); // 将博客标题、内容、前言拼接并转换为小写
        matchContent.replace(regex, () => {
          matches++;
          return ''; // 使用空字符串替换匹配到的内容
        });
        return { acc, matches };
      }).filter(item => item.matches > 0)
          .sort((a, b) => b.matches - a.matches) // 根据匹配度降序排序
          .map(item => item.acc); // 仅保留博客对象
      if (newValue === '' || newValue === null) {
        this.Accounts = this.AccountCopy;
      }
    }
  }
}
</script>

<template>
  <el-form inline style="float:left">
    <el-form-item >
      <el-input v-model="currentCity"  placeholder="输入相关城市信息"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="warning"
                 plain  @click="downloadFile">导出Excel</el-button>
    </el-form-item>
  </el-form>

  <el-table :data="Accounts" class="custom-table"  stripe :lazy="true" border height="700" >
    <el-table-column label="ID" width="100" prop="id"></el-table-column>
    <el-table-column label="Address" width="400" prop="address"></el-table-column>
    <el-table-column label="Password" width="120" prop="password"></el-table-column>
    <el-table-column label="City" width="200" prop="city"></el-table-column>
    <el-table-column label="Level" width="100" prop="level"></el-table-column>
  </el-table>
</template>

<style scoped>

</style>