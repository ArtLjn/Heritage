<script>
import {QueryHeritageTask} from "@/api/heritage";
import router from "@/router";

export default {
  name: "apply",
  data() {
    return {
      heritageInheritorTask:{
        currentPage: 1,
        list:[],
        total:0,
        totalPages: 0
      },
      heritageProjectTaskList: {
        currentPage: 1,
        list:[],
        total:0,
        totalPages:0
      },
      activeName:"first",
    }
  },
  mounted() {
    this.queryHeritageInheritorTask();
    this.queryHeritageProjectTask();
  },
  methods: {
    handleClick(event) {
      this.activeName = event.props.name
    },
    queryHeritageInheritorTask() {
      QueryHeritageTask(
          this.heritageInheritorTask.currentPage,
          10,
          1
      ).then(res => {
        const data = res.data.data[0];
        this.heritageInheritorTask.list = data.list;
        this.heritageInheritorTask.total = data.total;
        this.heritageInheritorTask.totalPages = data.totalPages;
        this.heritageInheritorTask.currentPage = data.currentPage;
      })
    },
    queryHeritageProjectTask() {
      QueryHeritageTask(
          this.heritageProjectTaskList.currentPage,
          10,
          2
      ).then(res => {
        const data = res.data.data[0];
        this.heritageProjectTaskList.list = data.list;
        this.heritageProjectTaskList.total = data.total;
        this.heritageProjectTaskList.totalPages = data.totalPages;
        this.heritageProjectTaskList.currentPage = data.currentPage;
      })
    },
    handleInheritorCurrentChange(val) {
      this.heritageInheritorTask.currentPage = val
      this.queryHeritageInheritorTask()
    },
    handleProjectCurrentChange(val) {
      this.heritageProjectTaskList.currentPage = val
      this.queryHeritageProjectTask()
    },
    handleView(row,raw) {
      router.push({
        path: '/adminHome/view',
        query: {
          field: row.field,
          raw: raw,
        }
      })
    },
    handleSizeChange(raw) {
      if(raw === 1) {
        this.queryHeritageInheritorTask()
      }else if(raw === 2) {
        this.queryHeritageProjectTask()
      }
    }
  }
}
</script>

<template>
  <el-tabs v-model="activeName"   @tab-click="handleClick" type="border-card">
    <el-tab-pane label="非遗人任务" name="first">
      <el-table border :data="heritageInheritorTask.list">
        <el-table-column type="selection" label="序号" width="80" >
          <template #default="{ row }">
            <el-checkbox v-model="row.selected"></el-checkbox>
          </template>
        </el-table-column>
        <el-table-column label="id" width="100" prop="id"></el-table-column>
        <el-table-column label="uuid"  width="100" prop="uuid"></el-table-column>
        <el-table-column label="create_time"  width="300" prop="create_time"></el-table-column>
        <el-table-column label="内容" width="100">
          <template #default="{ row }">
            <el-button type="primary" @click="handleView(row,1)">查看</el-button>
          </template>
        </el-table-column>
        <el-table-column label="操作">
        </el-table-column>
        <template #empty>
          <el-empty :description="description"></el-empty>
        </template>

      </el-table>
      <el-pagination
          v-model:current-page="heritageInheritorTask.currentPage"
          :page-count="heritageInheritorTask.totalPages"
          :small=false
          :disabled=false
          :background=true
          layout="prev, pager, next, jumper"
          @current-change="handleInheritorCurrentChange"
          @size-change="handleSizeChange(1)"
      />
    </el-tab-pane>
    <el-tab-pane label="非遗项目任务" name="second">
      <el-table border :data="heritageProjectTaskList.list">
        <el-table-column type="selection" label="序号" width="80" >
          <template #default="{ row }">
            <el-checkbox v-model="row.selected"></el-checkbox>
          </template>
        </el-table-column>
        <el-table-column label="id" width="100" prop="id"></el-table-column>
        <el-table-column label="uuid"  width="100" prop="uuid"></el-table-column>
        <el-table-column label="create_time"  width="300" prop="create_time"></el-table-column>
        <el-table-column label="内容" width="100">
          <template #default="{ row }">
            <el-button type="primary" @click="handleView(row,2)">查看</el-button>
          </template>
        </el-table-column>cd
        <el-table-column label="操作">
        </el-table-column>
        <template #empty>
          <el-empty :description="description"></el-empty>
        </template>

      </el-table>
      <el-pagination
          v-model:current-page="heritageProjectTaskList.currentPage"
          :page-count="heritageProjectTaskList.totalPages"
          :small=false
          :disabled=false
          :background=true
          layout="prev, pager, next, jumper"
          @current-change="handleProjectCurrentChange"
          @size-change="handleSizeChange(2)"
      />
    </el-tab-pane>
  </el-tabs>
  <el-main>
    <router-view></router-view>
  </el-main>
</template>
