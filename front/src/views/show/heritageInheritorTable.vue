<script>
import CardBoxModal from '@/components/CardBoxModal.vue'
import BaseLevel from '@/components/BaseLevel.vue'
import BaseButtons from '@/components/BaseButtons.vue'
import BaseButton from '@/components/BaseButton.vue'
import {mdiAccount, mdiAlert, mdiContentCopy, mdiEye, mdiOpenInNew} from '@mdi/js'
import {HeritageInheritor, QueryHeritageByLocate} from "@/api/heritage";
import router from "@/router";
import NotificationBar from "@/components/NotificationBar.vue";

export default {
  components:{
    NotificationBar,
    CardBoxModal,
    BaseLevel,
    BaseButtons,
    BaseButton,
  },
  props: {
    checkable: Boolean,
    city:String,
  },
  data() {
    return {
      isModalActive: false,
      isModalDangerActive: false,
      searchQuery: '',
      currentPage: 0,
      itemsPerPage: 10,
      heritageInheritorList: {
        currentPage: 1,
        list:[],
        total:0,
        totalPages: 0
      },
      showHeritageInheritor: false,
    };
  },
  computed: {
    // 计算属性生成页码列表
    totalInheritorPagesList() {
      const pages = [];
      // 这里确保总页数不为0时，生成页码
      if (this.heritageInheritorList.totalPages > 0) {
        for (let i = 1; i <= this.heritageInheritorList.totalPages; i++) {
          pages.push(i);
        }
      }
      return pages;
    },
  },
  mounted() {
    this.queryHeritageInheritorByLocate(this.$props.city)
  },
  methods: {
    mdiAlert() {
      return mdiAlert
    },
    mdiEye() {
      return mdiEye
    },
    mdiOpenInNew() {
      return mdiOpenInNew
    },
    mdiAccount() {
      return mdiAccount
    },
    mdiContentCopy() {
      return mdiContentCopy
    },
    // 更新当前页，并重新获取数据
    handleInheritorCurrentChange(page) {
      if (page >= 1 && page <= this.heritageInheritorList.totalPages) {
        this.heritageInheritorList.currentPage = page;
        this.queryHeritageInheritorByLocate(this.$props.city)
      }
    },
    queryHeritageInheritorByLocate(city) {
      QueryHeritageByLocate(
        this.heritageInheritorList.currentPage,
        10,
        HeritageInheritor,
        city
      ).then(res => {
        const b = res.data.data[0]
        this.heritageInheritorList.list = b.list
        this.heritageInheritorList.currentPage = b.currentPage;
        this.heritageInheritorList.total = b.total;
        this.heritageInheritorList.totalPages = b.totalPages;
      })
    },
    handleSeeInheritor(data) {
      router.push({
        path: "/seeInheritor",
        query: data
      });
    }
  }
};
</script>

<template>
  <NotificationBar color="warning" :icon="mdiAlert()">
    <b>注意：</b> 该表格为非遗传承人数据表
  </NotificationBar>
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
      <th>Number</th>
      <th>Inheritor</th>
      <th>Project</th>
      <th>CateGory</th>
      <th>Locate</th>
      <th>Rank</th>
      <th />
    </tr>
    </thead>
    <tbody>
    <tr v-for="client in heritageInheritorList.list" :key="client.id">
      <td data-label="Number">{{ client.number }}</td>
      <td data-label="Inheritor">{{ client.inheritor }}</td>
      <td data-label="Project">{{ client.project }}</td>
      <td data-label="CateGory">{{ client.category }}</td>
      <td data-label="Locate">{{ client.locate }}</td>
      <td data-label="Rank">{{ client.rank }}</td>
      <td class="before:hidden lg:w-1 whitespace-nowrap">
        <BaseButtons type="justify-start lg:justify-end" no-wrap>
          <BaseButton
            color="info"
            :icon="mdiEye()"
            small
            @click="handleSeeInheritor(client)"
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
          v-for="page in totalInheritorPagesList"
          :key="page"
          :active="page === heritageInheritorList.currentPage"
          :label="page === '...' ? '...' : page"
          :color="page === heritageInheritorList.currentPage ? 'lightDark' : 'whiteDark'"
          small
          @click="handleInheritorCurrentChange(page)"
        />
      </BaseButtons>
      <small>Page {{  heritageInheritorList.currentPage  }} of {{ heritageInheritorList.totalPages  }}</small>
    </BaseLevel>
  </div>
</template>
