<script setup>
import { ref, onMounted } from 'vue'
import {mdiCubeOutline, mdiAccountMultiple, mdiChartBar, mdiConsoleLine} from '@mdi/js'
import CardBoxWidget from '@/components/CardBoxWidget.vue'
import SectionMain from '@/components/SectionMain.vue'
import BaseButton from '@/components/BaseButton.vue'
import LineChart from '@/components/Charts/LineChart.vue'
import SectionTitleLineWithButton from '@/components/SectionTitleLineWithButton.vue'
import CardBox from "@/components/CardBox.vue";
import LayoutAuthenticated from "@/layouts/LayoutAuthenticated.vue";
import router from "@/router";

const heritageData = ref({
  projects: 20,
  inheritors: 35,
  participation: 80, // 公众参与度
  rewards: 10000 // 总激励代币数
})

const chartData = ref(null)

const loadChartData = () => {
  chartData.value = {
    labels: ['Q1', 'Q2', 'Q3', 'Q4'],
    datasets: [
      {
        label: '公众参与度',
        data: [65, 59, 80, 81],
        borderColor: '#3b82f6',
        fill: false
      },
      {
        label: '奖励发放数',
        data: [1500, 3000, 4500, 6000],
        borderColor: '#f59e0b',
        fill: false
      }
    ]
  }
}

const learnMore = () => {
  router.push("/showHeritage")
}
onMounted(() => {
  loadChartData()
})

</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <!-- 标题与简介 -->
      <SectionTitleLineWithButton title="区块链非遗保护平台概览" :icon="mdiCubeOutline" main>
        <BaseButton label="了解更多" color="contrast" small  @click="learnMore"  />
      </SectionTitleLineWithButton>

      <!-- 数据展示卡片 -->
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-6 mb-6">
        <CardBoxWidget
          :icon="mdiAccountMultiple"
          :number="heritageData.projects"
          label="遗产保护项目"
          color="text-emerald-500"
        />
        <CardBoxWidget
          :icon="mdiAccountMultiple"
          :number="heritageData.inheritors"
          label="传承人"
          color="text-blue-500"
        />
        <CardBoxWidget
          :icon="mdiConsoleLine"
          :number="heritageData.rewards"
          label="奖励发放 (代币)"
          color="text-yellow-500"
        />
        <CardBoxWidget
          :icon="mdiChartBar"
          :number="heritageData.participation"
          label="公众参与度"
          color="text-red-500"
        />
      </div>

      <!-- 图表展示公众参与度与奖励发放 -->
      <SectionTitleLineWithButton title="参与情况与激励统计" :icon="mdiChartBar">
        <BaseButton label="刷新数据" color="whiteDark" @click="loadChartData" />
      </SectionTitleLineWithButton>

      <CardBox class="mb-6">
        <div v-if="chartData">
          <LineChart :data="chartData" class="h-96" />
        </div>
      </CardBox>
    </SectionMain>
  </LayoutAuthenticated>
</template>

<style scoped>
</style>
