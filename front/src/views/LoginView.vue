<script setup>
import {onBeforeMount, reactive} from 'vue'
import { useRouter } from 'vue-router'
import { mdiAccount, mdiAsterisk } from '@mdi/js'
import SectionFullScreen from '@/components/SectionFullScreen.vue'
import CardBox from '@/components/CardBox.vue'
import FormCheckRadio from '@/components/FormCheckRadio.vue'
import FormField from '@/components/FormField.vue'
import FormControl from '@/components/FormControl.vue'
import BaseButton from '@/components/BaseButton.vue'
import BaseButtons from '@/components/BaseButtons.vue'
import LayoutGuest from '@/layouts/LayoutGuest.vue'
import {Login, VerifyToken} from "@/api/auth";
import {ElMessage, ElNotification} from "element-plus";
onBeforeMount(() => {
  VerifyToken()
})
const form = reactive({
  address:"",
  password:""
})

const router = useRouter()

const login = () => {
  if (form.address === '' || form.password === '') {
    ElMessage.warning("请填写完整表单")
    return
  }
  Login(form).then((res) => {
    if (res.data.code === 200) {
      const data = res.data.data[0];
      localStorage.setItem("city", data[0]);
      localStorage.setItem("rank", data[1]);
      localStorage.setItem("token", data[2]);
      localStorage.setItem("name", form.address);
      ElNotification.success("欢迎回来🥳")

      // Delay the redirect by 1 seconds (1000 milliseconds)
      setTimeout(() => {
        router.push("/dashboard");
      }, 1000);
    } else {
      ElNotification.error(res.data.msg)
    }
  })
}
</script>


<template>
  <LayoutGuest>
    <SectionFullScreen v-slot="{ cardClass }" bg="purplePink">
      <CardBox :class="cardClass" is-form @submit.prevent="login()">
        <FormField label="Login" help="Please enter your login">
          <FormControl
            v-model="form.address"
            :icon="mdiAccount"
            name="login"
            autocomplete="username"
            placeholder="请输入用户名"
          />
        </FormField>

        <FormField label="Password" help="Please enter your password">
          <FormControl
            v-model="form.password"
            :icon="mdiAsterisk"
            type="password"
            name="password"
            autocomplete="current-password"
            placeholder="请输入密码"
          />
        </FormField>

        <FormCheckRadio
          v-model="form.remember"
          name="remember"
          label="Remember"
          :input-value="true"
        />


        <template #footer>
          <BaseButtons>
            <BaseButton type="submit" color="info" label="Login" />
            <BaseButton to="/dashboard" color="info" outline label="Back" />
          </BaseButtons>
        </template>
      </CardBox>
    </SectionFullScreen>
  </LayoutGuest>
</template>
