import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  base: "/heritage/",
  plugins: [vue()],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    open: false, // 是否自动在浏览器打开
    port: 8081, // 端口号
    host: "0.0.0.0",
    //这里的ip和端口是前端项目的;下面为需要跨域访问后端项目
    proxy: {
      '/api/': {   // '/api'是代理标识，用于告诉node，url前面是/api的就是使用代理的
        target: 'http://localhost:8080',//这里填入你要请求的接口的前缀
        ws: false,//代理websocked
        changeOrigin: true,  //是否跨域
        secure: true,  //是否https接口
      }
    }
  }
});
