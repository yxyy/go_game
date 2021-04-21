// vue.config.js 配置说明
//官方vue.config.js 参考文档 https://cli.vuejs.org/zh/config/#css-loaderoptions
// 这里只列一部分，具体配置参考文档
module.exports = {
    devServer: {
        proxy: {//配置跨域
            '/admin': {
                target: 'http://127.0.0.1:8090/',// 后台
                ws: true,
                changOrigin: true//允许跨域
            },
            '/user': {
                target: 'http://localhost:8070/',// 账号中心
                ws: true,
                changOrigin: true//允许跨域
            }
            
        }
    }
    
  };