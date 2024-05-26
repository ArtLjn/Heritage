<template>
    <div class="background-image">
      <div class="box">
        <form  class="form" @submit.prevent="logins">
          <h1>注册</h1>
          <div class="input-box">
            <label for="username">Username:</label>
            <input type="text" id="username" name="username" v-model="loginForm.name" required>
    
            <label for="password">Password:</label>
            <input type="password" id="password" name="password" v-model="loginForm.password" required>
          </div>
          <button type="submit">Register</button>
          <button type="button" @click="handleLogin">Log In</button>
        </form>
      </div>
    </div>
    </template>
    
    <script>
    import router from "@/router";

    export default {
        data() {
        return {
          loginForm: {
            name: '',
            password: '',
          },
        };
      },
      methods: {
        handleLogin() {
          router.push('/');
        },
      },
    };
    </script>
    
    <style scoped>
    .background-image {
      position: absolute;
      position: fixed;
      top: 0px;
      left: 0px;
      height: 100%;
      width: 100%;
      z-index: 1;
    }
    
    .box {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      height: 100%;
      animation: slide-in-blurred 0.5s both;
    }
    .form {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      width: 400px;
      height: auto;
      background-color: rgba(179, 179, 179, 0.7);
      padding: 100px 50px;
      border-radius: 30px;
    }
    
    form h1 {
      margin: 10px;
      font-size: 28px;
      text-align: center;
      color: #f1724c;
    }
    
    .input-box {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      width: 100%;
    }
    
    input {
      margin: 10px;
      padding: 12px 20px;
      width: 100%;
      border: none;
      border-radius: 25px;
      box-sizing: border-box;
      background-color: #f7f7f7;
      font-size: 16px;
      color: #333333;
    }
    button[type="submit"] {
      margin: 10px;
      margin-top: 20px;
      padding: 12px 20px;
      width: 100%;
      border: none;
      border-radius: 25px;
      box-sizing: border-box;
      background-color: #f1724c;
      font-size: 16px;
      color: #ffffff;
      cursor: pointer;
      transition: background-color 0.3s ease-out;
      }
      
      button[type="submit"]:hover {
         background-color: #f14c4c;
      }
      
      button[type="button"] {
        margin: 10px;
        margin-top: 0px;
        padding: 12px 20px;
        width: 100%;
        border: none;
        border-radius: 25px;
        box-sizing: border-box;
        background-color: #f7f7f7;
        font-size: 16px;
        color: #333333;
        cursor: pointer;
        transition: background-color 0.3s ease-out;
      }
    </style>