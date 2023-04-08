const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    host: '0.0.0.0',
    port: 8080,
    https: false,
    proxy: {
      '/dev': {
        target:'http://193.203.215.185:8888',
        changeOrigin: true,
        pathRewrite: { '^/dev': '' },
        secure: false
      }
    }
  }
});
