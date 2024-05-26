<script>
  export default {
    data() {
      return {
        columns: [],
        data: []
      }
    }
  }
</script>

<template>
  <div class="order-box">
    <el-form inline style="float:left">
      <el-form-item label="订单号">
        <el-input  placeholder="输入订单号进行查询"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button  type="primary" >
          <el-icon style="vertical-align: middle">
            <Search />
          </el-icon>
          <span style="vertical-align: middle"> Search </span>
        </el-button>
        <el-button class="button-order" type="success" plain>新增订单</el-button>
      </el-form-item>
    </el-form>
  </div>

  <el-table border>
    <el-table-column type="selection" label="序号" width="80" >
      <template #default="{ row }">
        <el-checkbox v-model="row.selected"></el-checkbox>
      </template>
    </el-table-column>
    <el-table-column label="订单号" prop="orderNumber"></el-table-column>
    <el-table-column label="下单时间" prop="orderTime" sortable></el-table-column>
    <el-table-column label="顾客姓名"></el-table-column>
    <el-table-column label="订单状态"></el-table-column>
    <el-table-column label="物流公司名称"></el-table-column>
    <el-table-column label="发货日期"></el-table-column>
    <el-table-column label="收货地址"></el-table-column>
    <el-table-column label="联系电话"></el-table-column>
    <el-table-column label="备注"></el-table-column>
    <el-table-column label="操作">
<!--      <template #default="{ row }">-->
        <el-button type="text" size="small"><el-icon :size="18" :color="color"><edit></edit></el-icon></el-button>
        <el-button type="danger" size="small" icon="Delete" circle plain></el-button>
<!--      </template>-->
    </el-table-column>
    <template #empty>
      <el-empty :description="description"></el-empty>
    </template>
  </el-table>
</template>

<style scoped>
  .order-box {
    display: flex;
    flex-direction: row;
    margin-bottom: 10px;
  }
  .button-order {
    margin: 10px;
  }
</style>