<script>
import {HeritageProject, QueryHeritageByLocate} from "@/api/heritage";

export default {
  data() {
    return {
      heritageProjectList: {
        currentPage: 1,
        list:[],
        total:0,
        totalPages: 0
      },
      heritageProjectForm:{
        number: '',
        name: '',
        category: '',
        locate: '',
        rank: '',
        details: '',
      },
      showHeritageProject: false,
    }
  },
  props:{
    city:String
  },
  mounted() {
    this.queryHeritageProjectByLocate(this.$props.city)
  },
  methods:{
    handleProjectCurrentChange(val) {
      this.heritageProjectList.currentPage = val
      this.queryHeritageProjectByLocate(this.$props.city)
    },
    queryHeritageProjectByLocate(city) {
      QueryHeritageByLocate(
          this.heritageProjectList.currentPage,
          10,
          HeritageProject,
          city
      ).then(res => {
        const b = res.data.data[0]
        this.heritageProjectList.list = b.list
        this.heritageProjectList.currentPage = b.currentPage;
        this.heritageProjectList.total = b.total;
        this.heritageProjectList.totalPages = b.totalPages;
      })
    },
    handleView(row) {
      this.heritageProjectForm = row
      this.showHeritageProject = true
    }
  }
}
</script>

<template>
  <el-table border :data="heritageProjectList.list">
    <el-table-column type="selection" label="序号" width="80" >
      <template #default="{ row }">
        <el-checkbox v-model="row.selected"></el-checkbox>
      </template>
    </el-table-column>
    <el-table-column label="number" width="100" prop="number"></el-table-column>
    <el-table-column label="name"  width="200" prop="name"></el-table-column>
    <el-table-column label="category" width="100" prop="category"></el-table-column>
    <el-table-column label="locate"  width="200" prop="locate"></el-table-column>
    <el-table-column label="rank"  width="100" prop="rank"></el-table-column>
    <el-table-column label="内容" width="100">
      <template #default="{ row }">
        <el-button type="primary" @click="handleView(row)">查看</el-button>
      </template>
    </el-table-column>
    <template #empty>
      <el-empty :description="description"></el-empty>
    </template>
  </el-table>
  <el-pagination
      style="margin-top: 20px;"
      v-model:current-page="heritageProjectList.currentPage"
      :page-count="heritageProjectList.totalPages"
      :small=false
      :disabled=false
      :background=true
      layout="prev, pager, next, jumper"
      @current-change="handleProjectCurrentChange"
  />

  <el-card style="margin-top: 40px" v-if="showHeritageProject">
    <el-form>
      <el-form-item label="文物编号">
        {{heritageProjectForm.number}}
      </el-form-item>
      <el-form-item label="项目">
        {{heritageProjectForm.name}}
      </el-form-item>
      <el-form-item label="类别">
        {{heritageProjectForm.category}}
      </el-form-item>
      <el-form-item label="地点">
        {{heritageProjectForm.locate}}
      </el-form-item>
      <el-form-item label="级别">
        {{heritageProjectForm.rank}}
      </el-form-item>
      <el-form-item label="详情">
        <div v-html="heritageProjectForm.details"></div>
      </el-form-item>
      <el-button type="primary" @click="showHeritageProject = false">关闭</el-button>
    </el-form>
  </el-card>
</template>

<style scoped>

</style>