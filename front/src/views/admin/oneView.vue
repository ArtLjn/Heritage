<template>
  <el-calendar ref="calendar">
    <template v-slot:header="{ date }">
      <span>Custom header content</span>
      <span>{{ date }}</span>
      <el-button-group>
        <el-button size="small" @click="selectDate('prev-year')">Previous Year</el-button>
        <el-button size="small" @click="selectDate('prev-month')">Previous Month</el-button>
        <el-button size="small" @click="selectDate('today')">Today</el-button>
        <el-button size="small" @click="selectDate('next-month')">Next Month</el-button>
        <el-button size="small" @click="selectDate('next-year')">Next Year</el-button>
      </el-button-group>
    </template>
  </el-calendar>
</template>

<script>
import {ref} from 'vue'
import 'element-plus/theme-chalk/index.css'
import {ElCalendar, ElButton, ElButtonGroup} from 'element-plus'

export default {
  components: {
    ElCalendar,
    ElButton,
    ElButtonGroup
  },
  data() {
    return {
      calendar: ref(null)
    }
  },
  methods: {
    selectDate(command) {
      console.log(command)
      // if (this.calendar.value) {
      //   this.calendar.value[command]()
      // }
    }
  },
}
</script>

<style>
.box-card {
  width: 480px;
  height: 200px;
}
</style>