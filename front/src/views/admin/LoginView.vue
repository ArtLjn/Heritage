<template>
  <div id="login">
    <transition appear name="opacitytrans">
      <div class="container" id="container">
        <div class="form-container sign-in-container">
          <form action="#">
            <h1>非物质文化遗产</h1>
            <div v-if="showLogin">
              <input type="text" placeholder="账户" v-model="account" />
              <input type="password" placeholder="密码" v-model="password" />
              <div class="button" @click="handleLogin()" >登录</div>
            </div>
            
            <div v-if="showRegister">
              <input type="text" placeholder="名称" v-model="account" />
              <input type="password" placeholder="密码" v-model="password" />
              <div class="button" @click="login()" >注册</div>
              <div class="click-button" @click="onLogin">登录</div>
            </div>
          </form>
          
        </div>

        <div class="overlay-container">
          <div class="overlay">
            <div class="overlay-panel overlay-right">
              <img class="logo" src="@/assets/login.jpg" alt="" />
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>
<script>
import router from "@/router";
import {Login, VerifyToken} from "@/api/auth";

export default {
  data() {
    return {
      account: "",
      password: "",
      showLogin: true,
      showRegister:false
    };
  },
  mounted() {
    VerifyToken();
  },
  methods:{
    onRegister() {
      this.showLogin = false;
      this.showRegister = true;
    },
    onLogin() {
      this.showLogin = true;
      this.showRegister = false;
    },
    handleLogin() {
      if (this.account === "" || this.password === "") {
        this.$message.warning("请输入账号或密码");
        return;
      }
      const body = {
        "address":this.account,
        "password":this.password
      }
      Login(body).then((res) => {
        if (res.data.code === 200) {
          const data = res.data.data[0];
          localStorage.setItem("city",data[0])
          localStorage.setItem("rank",data[1])
          localStorage.setItem("token",data[2])
          localStorage.setItem("name",body.address)
          this.$message.success("登录成功")
          router.push("/adminHome")
        } else {
          this.$message.error(res.data.message)
        }
      }).catch((err) => {
        this.$message.error(err)
      })
    }
  }
};
</script>
<style  scoped>
#login {
  font-family: "Montserrat", sans-serif;
  background-image: linear-gradient(to right bottom, #b19bc8, #bdb287);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  /* margin: -20px 0 50px; */
  /* background-image: url("https://www.17sucai.com/preview/1749733/2019-06-22/%E7%99%BB%E5%BD%95/img/img1.png"); */
  background-size: cover;
}
.logo {
  width: 500px;
  margin-left: 90px;
  height: 570px;
}
h1 {
  font-weight: bold;
  margin: 0;
  color: beige;
}

p {
  font-size: 14px;
  font-weight: bold;
  line-height: 20px;
  letter-spacing: 0.5px;
  margin: 20px 0 30px;
}

span {
  font-size: 12px;
  color: beige;
}

a {
  color: #fff;
  font-size: 14px;
  text-decoration: none;
  margin: 15px 0;
}

.container {
  border-radius: 10px;
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
  position: absolute;
  overflow: hidden;
  width: 768px;
  max-width: 100%;
  min-height: 480px;
  opacity: 0.8;
}

.form-container form {
  background: rgba(45, 52, 54, 1);
  display: flex;
  flex-direction: column;
  padding: 0 50px;
  height: 100%;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.social-container {
  margin: 20px 0;
}

.social-container a {
  border: 1px solid #ddd;
  border-radius: 50%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  margin: 0 5px;
  height: 40px;
  width: 40px;
}

.form-container input {
  background: #eee;
  border: none;
  padding: 12px 15px;
  margin: 8px 0;
  width: 100%;
}

.button {
  cursor: pointer;
  border-radius: 20px;
  /* border: 1px solid #ff4b2b;
  background: #ff4b2b; */
  /* border: 1px solid #fa8817;
  background: #fa8817; */
    border: 1px solid #1BBFB4;
  background: #1BBFB4;
  color: #fff;
  font-size: 12px;
  font-weight: bold;
  padding: 12px 45px;
  margin-top: 20px;
  letter-spacing: 1px;
  text-transform: uppercase;
  transition: transform 80ms ease-in;
}

input[type="text"] {
  width: 240px;
  text-align: center;
  background: transparent;
  border: none;
  border-bottom: 1px solid #fff;
  font-family: "PLay", sans-serif;
  font-size: 16px;
  font-weight: 200px;
  padding: 10px 0;
  transition: border 0.5s;
  outline: none;
  color: #fff;
}

input[type="password"] {
  width: 240px;
  text-align: center;
  background: transparent;
  border: none;
  border-bottom: 1px solid #fff;
  font-family: "PLay", sans-serif;
  font-size: 16px;
  font-weight: bold;
  padding: 10px 0;
  transition: border 0.5s;
  outline: none;
  color: #fff;
}

input[type="email"] {
  width: 240px;
  text-align: center;
  background: transparent;
  border: none;
  border-bottom: 1px solid #fff;
  font-family: "PLay", sans-serif;
  font-size: 16px;
  font-weight: 200px;
  padding: 10px 0;
  transition: border 0.5s;
  outline: none;
  color: #fff;
}

.button:active {
  transform: scale(0.95);
}

.button:focus {
  outline: none;
}

.button.ghost {
  background: transparent;

  /* border-color: #fa8817;
  background-color: #fa8817; */
   border-color: #1BBFB4;
  background-color: #1BBFB4;
  margin: 0;
}

.form-container {
  position: absolute;
  top: 0;
  height: 100%;
  transition: all 0.6s ease-in-out;
}

.sign-in-container {
  left: 0;
  width: 50%;
  z-index: 2;
}

.sign-up-container {
  left: 0;
  width: 50%;
  z-index: 1;
  opacity: 0;
}

.overlay-container {
  position: absolute;
  top: 0;
  left: 50%;
  width: 50%;
  height: 100%;
  overflow: hidden;
  transition: transform 0.6s ease-in-out;
  z-index: 100;
}

.overlay {
  background: transparent;
  background: linear-gradient(to right, #ff4b2b, #ff416c) no repeat 0 0 / cover;
  color: #fff;
  position: absolute;
  left: -100%;
  height: 100%;
  width: 200%;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
}

.overlay-panel {
  position: absolute;
  top: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 0 40px;
  height: 100%;
  width: 50%;
  text-align: center;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
}

.overlay-right {
  right: 0;
  transform: translateX(0);
}

.click-button{
  color: white;
  margin-top: 10px;
  padding: 10px;
  background-color: rgb(232, 224, 224,0);
}
</style>
