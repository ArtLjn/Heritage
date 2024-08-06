<template>
  <admin-home-page  ></admin-home-page>
</template>
<script>
import adminHomePage from "@/components/admin/adminHomePage.vue";
export default {
  components:{
    adminHomePage
  },
  data()
  {
    return {
      city:localStorage.getItem("city")
    }
  },
}
</script>