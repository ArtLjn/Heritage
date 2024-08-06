<template>
  <div>
    <el-row :gutter="20" style="margin-bottom: 20px;">
      <el-col :span="8">
        <el-card>
          <div>
            <h3>总账号数</h3>
            <p style="font-size: 24px; color: #409EFF;">{{ totalAccounts }}</p>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <div>
            <h3>活跃账号数</h3>
            <p style="font-size: 24px; color: #67C23A;">{{ activeAccounts }}</p>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <div>
            <h3>非活跃账号数</h3>
            <p style="font-size: 24px; color: #F56C6C;">{{ inactiveAccounts }}</p>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row style="margin-bottom: 20px;">
      <el-col :span="24">
        <el-card>
          <el-input
              v-model="searchQuery"
              placeholder="搜索省份或城市"
              @input="filterAccounts"
              clearable
          ></el-input>
        </el-card>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="24">
        <el-card>
          <div id="chart" style="width: 100%; height: 400px;"></div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import * as echarts from 'echarts';
import { GetAllAccount } from "@/api/auth";

export default {
  name: "HomePage",
  data() {
    return {
      accounts: [],
      filteredAccounts: [],
      provinceData: {},
      totalAccounts: 0,
      activeAccounts: 0,
      inactiveAccounts: 0,
      searchQuery: ""
    };
  },
  mounted() {
    this.fetchAccountData();
  },
  methods: {
    async fetchAccountData() {
      try {
        const res = await GetAllAccount();
        this.accounts = res.data.data[0];
        this.filteredAccounts = this.accounts;
        this.totalAccounts = this.accounts.length;
        this.processData();
        this.initChart();
      } catch (error) {
        console.error('Failed to fetch account data', error);
      }
    },
    processData() {
      this.provinceData = this.accounts.reduce((acc, account) => {
        const province = account.city.split('-')[0]; // 提取省份信息
        if (!acc[province]) {
          acc[province] = 0;
        }
        acc[province]++;
        if (account.level === 1) {
          this.activeAccounts++;
        } else {
          this.inactiveAccounts++;
        }
        return acc;
      }, {});
    },
    initChart() {
      const chartDom = document.getElementById('chart');
      const myChart = echarts.init(chartDom);

      const option = {
        title: {
          text: 'Account Distribution by Province',
          left: 'center'
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'shadow'
          }
        },
        xAxis: {
          type: 'category',
          data: Object.keys(this.provinceData),
          axisLabel: {
            interval: 0,
            rotate: 30
          }
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            name: 'Accounts',
            type: 'bar',
            data: Object.values(this.provinceData),
            itemStyle: {
              color: '#3398DB'
            }
          }
        ]
      };

      myChart.setOption(option);
    },
    filterAccounts() {
      const query = this.searchQuery.toLowerCase();
      this.filteredAccounts = this.accounts.filter(account => {
        return (
            account.city.toLowerCase().includes(query)
        );
      });
      this.processData();
      this.initChart();
    }
  }
};
</script>

<style scoped>
#chart {
  width: 100%;
  height: 400px;
}
</style>
