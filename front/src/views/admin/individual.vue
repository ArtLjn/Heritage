<script>
export default {
  name: "individual",
  data() {
    return {
      individual:{
        city:"",
        rank: ""
      }
    }
  },
  mounted() {
    this.individual.city = localStorage.getItem("city")
    const rank = localStorage.getItem("rank")
    switch (rank) {
      case '1':
        this.individual.rank = "province"
        break
      case '2':
        this.individual.rank = "city"
        break
    }
  }
}
</script>

<template>
 <el-form>
   <el-form-item label="city:">
     <span>{{individual.city}}</span>
   </el-form-item>
   <el-form-item label="rank:">
     <span>{{individual.rank}}</span>
   </el-form-item>
 </el-form>
</template>

<style scoped>

</style>