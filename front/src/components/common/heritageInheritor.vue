<script>
import {HeritageInheritor, QueryHeritageByLocate} from "@/api/heritage";

export default {
  data() {
    return {
      heritageInheritorList: {
        currentPage: 1,
        list:[],
        total:0,
        totalPages: 0
      },
      heritageInheritorForm:{
        number: '',
        inheritor: '',
        project: '',
        category: '',
        locate: '',
        rank: '',
        details: '',
        inheritorImg:''
      },
      showHeritageInheritor: false,
    }
  },
  props:{
    city:String
  },
  mounted() {
    this.queryHeritageInheritorByLocate(this.$props.city)
  },
  methods:{
    handleInheritorCurrentChange(val) {
      this.heritageInheritorList.currentPage = val
      this.queryHeritageInheritorByLocate(this.$props.city)
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
    handleView(row) {
      this.heritageInheritorForm = row
      this.showHeritageInheritor = true
    }
  }
}
</script>

<template>
  <el-table border :data="heritageInheritorList.list">
    <el-table-column type="selection" label="序号" width="80" >
      <template #default="{ row }">
        <el-checkbox v-model="row.selected"></el-checkbox>
      </template>
    </el-table-column>
    <el-table-column label="number" width="100" prop="number"></el-table-column>
    <el-table-column label="inheritor"  width="100" prop="inheritor"></el-table-column>
    <el-table-column label="project"  width="200" prop="project"></el-table-column>
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
      v-model:current-page="heritageInheritorList.currentPage"
      :page-count="heritageInheritorList.totalPages"
      :small=false
      :disabled=false
      :background=true
      layout="prev, pager, next, jumper"
      @current-change="handleInheritorCurrentChange"
  />

  <el-card style="margin-top: 40px" v-if="showHeritageInheritor">
    <el-form>
      <el-form-item label="文物编号">
        {{heritageInheritorForm.number}}
      </el-form-item>
      <el-form-item label="传承人">
        {{heritageInheritorForm.inheritor}}
      </el-form-item>
      <el-form-item label="项目">
        {{heritageInheritorForm.project}}
      </el-form-item>
      <el-form-item label="类别">
        {{heritageInheritorForm.category}}
      </el-form-item>
      <el-form-item label="地点">
        {{heritageInheritorForm.locate}}
      </el-form-item>
      <el-form-item label="级别">
        {{heritageInheritorForm.rank}}
      </el-form-item>
      <el-form-item label="详情">
        <div v-html="heritageInheritorForm.details"></div>
      </el-form-item>
      <el-form-item label="图片">
        <el-image :src="heritageInheritorForm.inheritorImg" fit="cover" style="width: 200px; height: 200px"></el-image>
      </el-form-item>
      <el-button type="primary" @click="showHeritageInheritor = false">关闭</el-button>
    </el-form>
  </el-card>
</template>
