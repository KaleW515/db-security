module.exports = {
  assetsDir: 'static',
  productionSourceMap: false,
  devServer: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8080/',
        ws: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
}
