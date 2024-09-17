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
      heritageProjectTaskList: {
        currentPage: 1,
        list:[],
        total:0,
        totalPages:0
      },
    };
  },
  computed: {
    // 计算属性生成页码列表
    totalInheritorPagesList() {
      const pages = [];
      // 这里确保总页数不为0时，生成页码
      if (this.heritageProjectTaskList.totalPages > 0) {
        for (let i = 1; i <= this.heritageProjectTaskList.totalPages; i++) {
          pages.push(i);
        }
      }
      return pages;
    },
  },
  mounted() {
    this.queryHeritageProjectTask();
  },
  methods: {
    mdiAlert() {
      return mdiAlert
    },
    mdiEye() {
      return mdiEye
    },
    queryHeritageProjectTask() {
      QueryHeritageTask(
        this.heritageProjectTaskList.currentPage,
        10,
        2,
        localStorage.getItem("city")
      ).then(res => {
        const data = res.data.data[0];
        this.heritageProjectTaskList.list = data.list;
        this.heritageProjectTaskList.total = data.total;
        this.heritageProjectTaskList.totalPages = data.totalPages;
        this.heritageProjectTaskList.currentPage = data.currentPage;
      })
    },
    // 更新当前页，并重新获取数据
    handleProjectCurrentChange(page) {
      if (page >= 1 && page <= this.heritageProjectTaskList.totalPages) {
        this.heritageProjectTaskList.currentPage = page;
        this.queryHeritageProjectTask(); // 每次点击分页时重新获取数据
      }
    },
    handleSeeProject(data) {
      router.push({
        path: "/watchProject",
        query: data
      });
    }
  }
};
</script>

<template>
  <NotificationBar color="warning" :icon="mdiAlert()">
    <b>注意：</b> 该表格为<b>非遗传承项目</b>的申请任务表
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
      <th>Level</th>
      <th>CreateTime</th>
      <th />
    </tr>
    </thead>
    <tbody>
    <tr v-for="client in heritageProjectTaskList.list" :key="client.id">
      <td data-label="ID">{{ client.id }}</td>
      <td data-label="UUID">{{ client.uuid }}</td>
      <td data-label="Type">{{ client.type }}</td>
      <td data-label="Locate">{{ client.locate }}</td>
      <td data-label="Level">{{ client.level }}</td>
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
            @click="handleSeeProject(client)"
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
          :active="page === heritageProjectTaskList.currentPage"
          :label="page === '...' ? '...' : page"
          :color="page === heritageProjectTaskList.currentPage ? 'lightDark' : 'whiteDark'"
          small
          @click="handleProjectCurrentChange(page)"
        />
      </BaseButtons>
      <small>Page {{  heritageProjectTaskList.currentPage  }} of {{ heritageProjectTaskList.totalPages  }}</small>
    </BaseLevel>
  </div>

</template>
