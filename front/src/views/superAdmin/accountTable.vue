<script>
import CardBoxModal from '@/components/CardBoxModal.vue'
import BaseLevel from '@/components/BaseLevel.vue'
import BaseButtons from '@/components/BaseButtons.vue'
import BaseButton from '@/components/BaseButton.vue'
import { GetAllAccount } from "@/api/auth";
import {mdiAccount, mdiContentCopy, mdiOpenInNew} from '@mdi/js'
import FormControl from "@/components/FormControl.vue"
import FormField from "@/components/FormField.vue";
import axios from "axios";
import UserAvatar from "@/components/UserAvatar.vue";

export default {
  components:{
    UserAvatar,
    FormField,
    CardBoxModal,
    BaseLevel,
    BaseButtons,
    BaseButton,
    FormControl,
  },
  props: {
    checkable: Boolean
  },
  data() {
    return {
      isModalActive: false,
      isModalDangerActive: false,
      searchQuery: '',
      accountList: [],  // 当前展示的账户数据
      AccountCopy: [],  // 原始账户数据副本
      currentPage: 0,
      itemsPerPage: 10
    };
  },
  computed: {
    numPages() {
      return Math.ceil(this.filteredAccounts.length / this.itemsPerPage);
    },
    filteredAccounts() {
      return this.accountList.filter(client => {
        return (
          client.id.toString().includes(this.searchQuery) ||
          client.address.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
          client.password.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
          client.city.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
          client.level.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
      });
    },
    itemsPaginated() {
      const start = this.currentPage * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.filteredAccounts.slice(start, end);
    },
    pagesList() {
      const totalPages = this.numPages;
      const maxPagesToShow = 5;
      const delta = 2;
      let pages = [];

      for (let i = 0; i < totalPages; i++) {
        if (totalPages <= maxPagesToShow) {
          pages.push(i);
        } else {
          if (i === 0 || i === totalPages - 1 || (i >= this.currentPage - delta && i <= this.currentPage + delta)) {
            pages.push(i);
          } else if (pages[pages.length - 1] !== '...') {
            pages.push('...');
          }
        }
      }

      return pages;
    },
    currentPageHuman() {
      return this.currentPage + 1;
    }
  },
  watch: {
    searchQuery(newValue) {
      // 如果搜索框为空，则还原完整列表
      if (newValue === '' || newValue === null) {
        this.accountList = this.AccountCopy;
        return;
      }

      const keywords = newValue.toLowerCase().split(' '); // 将用户输入转换为小写并按空格拆分为关键字数组
      const regex = new RegExp(keywords.join('|'), 'gi'); // 构建正则表达式，匹配所有关键字

      // 筛选符合条件的账户数据
      this.accountList = this.AccountCopy.filter(acc => {
        const matchContent = `${acc.id} ${acc.city} ${acc.level}`.toLowerCase(); // 拼接搜索内容并转换为小写
        return regex.test(matchContent); // 测试是否匹配关键字
      });
    }
  },
  mounted() {
    GetAllAccount().then(res => {
      this.accountList = res.data.data[0];
      this.AccountCopy = [...this.accountList];  // 保存原始账户数据
    });
  },
  methods: {
    mdiOpenInNew() {
      return mdiOpenInNew
    },
    mdiAccount() {
      return mdiAccount
    },
    mdiContentCopy() {
      return mdiContentCopy
    },
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
    },
    copyToClipboard(client) {
      // 复制当前行的 JSON 数据
      const clientData = JSON.stringify(client, null, 2); // 将数据转化为格式化的 JSON 字符串
      navigator.clipboard.writeText(clientData).then(() => {
        this.$notify.success('JSON 数据已复制到剪贴板');
      }).catch(err => {
        console.error('复制失败', err);
      });
    }
  }
};
</script>

<template>
  <!-- 搜索框 -->
  <div class="p-4">
    <FormField>
      <FormControl v-model="searchQuery" :icon="mdiAccount()" placeholder="请输入账号信息"/>
      <BaseButtons>
        <BaseButton
          color="warning"
          label="导出Excel"
          :icon="mdiOpenInNew()"
          @click="downloadFile"
        />
      </BaseButtons>
    </FormField>
  </div>

  <!-- 模态框 -->
  <CardBoxModal v-model="isModalActive" title="Sample modal">
    <p>Lorem ipsum dolor sit amet <b>adipiscing elit</b></p>
    <p>This is a sample modal</p>
  </CardBoxModal>

  <CardBoxModal v-model="isModalDangerActive" title="Please confirm" button="danger" has-cancel>
    <p>Lorem ipsum dolor sit amet <b>adipiscing elit</b></p>
    <p>This is a sample modal</p>
  </CardBoxModal>

  <!-- 表格 -->
  <table>
    <thead>
    <tr>
      <th>ID</th>
      <th>Address</th>
      <th>Password</th>
      <th>City</th>
      <th>Level</th>
      <th />
    </tr>
    </thead>
    <tbody>
    <tr v-for="client in itemsPaginated" :key="client.id">
      <td class="border-b-0 lg:w-6 before:hidden">
        <UserAvatar :username="client.address" class="w-24 h-24 mx-auto lg:w-6 lg:h-6" />
      </td>
      <td data-label="ID">{{ client.id }}</td>
      <td data-label="Address">{{ client.address }}</td>
      <td data-label="Password">{{ client.password }}</td>
      <td data-label="City">{{ client.city }}</td>
      <td data-label="Level">{{ client.level }}</td>
      <td class="before:hidden lg:w-1 whitespace-nowrap">
        <BaseButtons type="justify-start lg:justify-end" no-wrap>
          <BaseButton
            color="info"
            :icon="mdiContentCopy()"
            small
            @click="copyToClipboard(client)"
          />
        </BaseButtons>
      </td>
    </tr>
    </tbody>
  </table>

  <!-- 分页 -->
  <div class="p-3 lg:px-6 border-t border-gray-100 dark:border-slate-800">
    <BaseLevel>
      <BaseButtons>
        <BaseButton
          v-for="page in pagesList"
          :key="page"
          :active="page === currentPage"
          :label="page === '...' ? '...' : page + 1"
          :color="page === currentPage ? 'lightDark' : 'whiteDark'"
          small
          @click="page !== '...' && (currentPage = page)"
        />
      </BaseButtons>
      <small>Page {{ currentPageHuman }} of {{ numPages }}</small>
    </BaseLevel>
  </div>
</template>
