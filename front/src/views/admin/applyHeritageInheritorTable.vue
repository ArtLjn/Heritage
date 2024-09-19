<script>
import CardBoxModal from '@/components/CardBoxModal.vue'
import BaseLevel from '@/components/BaseLevel.vue'
import BaseButtons from '@/components/BaseButtons.vue'
import BaseButton from '@/components/BaseButton.vue'
import {mdiAlert, mdiEye} from '@mdi/js'
import {QueryHeritageTask} from "@/api/heritage";
import NotificationBar from "@/components/NotificationBar.vue";
import router from "@/router";

export default {
  components:{
    NotificationBar,
    CardBoxModal,
    BaseLevel,
    BaseButtons,
    BaseButton,
  },
  props: {
    checkable: Boolean
  },
  data() {
    return {
      isModalActive: false,
      isModalDangerActive: false,
      heritageInheritorTask:{
        currentPage: 1,
        list:[],
        total:0,
        totalPages: 0
      },
    };
  },
  computed: {
    // 计算属性生成页码列表
    totalInheritorPagesList() {
      const pages = [];
      // 这里确保总页数不为0时，生成页码
      if (this.heritageInheritorTask.totalPages > 0) {
        for (let i = 1; i <= this.heritageInheritorTask.totalPages; i++) {
          pages.push(i);
        }
      }
      return pages;
    },
  },
  mounted() {
    this.queryHeritageInheritorTask();
  },
  methods: {
    mdiAlert() {
      return mdiAlert
    },
    mdiEye() {
      return mdiEye
    },
    queryHeritageInheritorTask() {
      QueryHeritageTask(
        this.heritageInheritorTask.currentPage,
        10,
        1,
        localStorage.getItem("city")
      ).then(res => {
        const data = res.data.data[0];
        this.heritageInheritorTask.list = data.list;
        this.heritageInheritorTask.total = data.total;
        this.heritageInheritorTask.totalPages = data.totalPages;
        this.heritageInheritorTask.currentPage = data.currentPage;
      })
    },
    // 更新当前页，并重新获取数据
    handleInheritorCurrentChange(page) {
      if (page >= 1 && page <= this.heritageInheritorTask.totalPages) {
        this.heritageInheritorTask.currentPage = page;
        this.queryHeritageInheritorTask(); // 每次点击分页时重新获取数据
      }
    },
    handleSeeInheritor(data) {
      router.push({
        path: "/watchInheritor",
        query: data
      });
    }
  }
};
</script>

<template>
  <NotificationBar color="warning" :icon="mdiAlert()">
    <b>注意：</b> 该表格为非遗传承人的申请任务表
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
      <th>ID</th>
      <th>UUID</th>
      <th>Type</th>
      <th>Locate</th>
      <th>ApplyLevel</th>
      <th>PassLevel</th>
      <th>CreateTime</th>
      <th />
    </tr>
    </thead>
    <tbody>
    <tr v-for="client in heritageInheritorTask.list" :key="client.id">
      <td data-label="ID">{{ client.id }}</td>
      <td data-label="UUID">{{ client.uuid }}</td>
      <td data-label="Type">{{ client.type }}</td>
      <td data-label="Locate">{{ client.locate }}</td>
      <td data-label="ApplyLevel">{{ client.apply_level }}</td>
      <td data-label="PassLevel">{{ client.pass_level }}</td>
      <td data-label="CreateTime" class="lg:w-1 whitespace-nowrap">
        <small class="text-gray-500 dark:text-slate-400" :title="client.create_time">{{
            client.create_time
          }}</small>
      </td>
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
          :active="page === heritageInheritorTask.currentPage"
          :label="page === '...' ? '...' : page"
          :color="page === heritageInheritorTask.currentPage ? 'lightDark' : 'whiteDark'"
          small
          @click="handleInheritorCurrentChange(page)"
        />
      </BaseButtons>
      <small>Page {{  heritageInheritorTask.currentPage  }} of {{ heritageInheritorTask.totalPages  }}</small>
    </BaseLevel>
  </div>

</template>
