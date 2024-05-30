const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave:false,
  assetsDir:'static',
  productionSourceMap:false,
  devServer: {
    port:8081,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
      }
    },
    historyApiFallback: true //增加这个选项
  },
})
