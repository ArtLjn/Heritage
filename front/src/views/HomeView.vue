<script setup>
import {ref, onMounted} from 'vue'
import {mdiCubeOutline, mdiAccountMultiple, mdiChartBar, mdiConsoleLine, mdiMapMarker} from '@mdi/js'
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
  accountAmount: 1000, // 审核账户数量
})

const participantDistribution = ref([
  {location: '北京', count: 300},
  {location: '上海', count: 200},
  {location: '广州', count: 150},
  {location: '深圳', count: 100},
])

const historicalTrends = ref({
  labels: ['2020', '2021', '2022', '2023'],
  datasets: [
    {
      label: '遗产保护项目',
      data: [10, 15, 25, 30],
      borderColor: '#42A5F5',
      backgroundColor: 'rgba(66, 165, 245, 0.5)',
    },
    {
      label: '传承人数量',
      data: [20, 25, 40, 45],
      borderColor: '#FFA726',
      backgroundColor: 'rgba(255, 167, 38, 0.5)',
    },
  ],
})

onMounted(() => {
  // 模拟获取数据
  setTimeout(() => {
    heritageData.value = {
      projects: 30,
      inheritors: 45,
      participation: 90,
      accountAmount: 1200,
    }
  }, 1000)
})

const learnMore = () => {
  router.push("/showHeritage")
}
</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <!-- 标题与简介 -->
      <SectionTitleLineWithButton title="区块链非遗保护平台概览" :icon="mdiCubeOutline" main>
        <BaseButton label="了解更多" color="contrast" small @click="learnMore"/>
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
          :number="heritageData.accountAmount"
          label="审核账户数量"
          color="text-yellow-500"
        />
        <CardBoxWidget
          :icon="mdiChartBar"
          :number="heritageData.participation"
          label="公众参与度"
          color="text-red-500"
        />
      </div>

      <!-- 参与者分布 -->
      <CardBox>
        <SectionTitleLineWithButton title="参与者分布" :icon="mdiMapMarker">
          <BaseButton label="查看更多" color="contrast" small/>
        </SectionTitleLineWithButton>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div v-for="(location, index) in participantDistribution" :key="index" class="p-4 bg-gray-100 rounded-md">
            <p class="text-lg font-semibold">{{ location.location }}</p>
            <p class="text-sm text-gray-600">参与人数: {{ location.count }}</p>
          </div>
        </div>
      </CardBox>
    </SectionMain>
  </LayoutAuthenticated>
</template>

<style scoped>
</style>
