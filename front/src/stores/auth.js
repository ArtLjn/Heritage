import { defineStore } from 'pinia'
import {ref} from "vue";
import {GetAllAccount} from "@/api/auth";
import axios from "axios";


export const useAuthStore = defineStore('auth', () => {
  const accountList = ref([])

  function getAllAccount() {
    GetAllAccount().then(res => {
      accountList.value = res.data.data[0];
    })
  }

  function fetchSampleClients() {
    axios
      .get(`/api/account/getAllAccount`,{
        headers:{
          Authorization: localStorage.getItem("token")
        }
      })
      .then((result) => {
        accountList.value = result.data.data[0]
      })
      .catch((error) => {
        alert(error.message)
      })
  }
  return {
    getAllAccount,
    accountList,
    fetchSampleClients
  }
})
