<template>
    <div class="header">
        <el-button type="danger" class="logout" @click="logout">Logout</el-button>
    </div>
</template>

<script>
import {LogOut} from "@/api/auth";

export default {
    methods:{
        logout() {
            LogOut();
        }
    }
}
</script>

<style scoped>
.logout {
  float: right;
  margin-top: 13px;
}
.header{
    /* background-color: white; */
    width: 85%;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
}
</style>