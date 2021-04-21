import Vue from 'vue'
import App from '@/App.vue'
import router from '@/router'
// import elment
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import axios from 'axios'
import CSS  from '@/assets/css/myCss.css'

axios.defaults.headers.common['Authorization'] = window.localStorage.getItem("access_token")

axios.interceptors.response.use(
  response => {
      // 如果返回的状态码为200，说明接口请求成功，可以正常拿到数据,否则的话抛出错误
      if (response.status === 200) {
          return response
      } else {
          return Promise.reject(response);
      }
      // return response
  },
  // 错误处理
  error => {
      if (error.response.status) {
          switch (error.response.status) {
              // 401: 未登录
              // 未登录则跳转登录页面，并携带当前页面的路径
              // 在登录成功后返回当前页面，这一步需要在登录页操作。
              case 401:
                  router.replace({
                      path: '/login',
                      query: {
                          'redirect': router.currentRoute.fullPath
                      }
                  });
                  break;

              // 403 token过期，刷新token
              case 403:
                  axios({
                    method : 'post',
                    url : '/user/refreshToken',
                    headers :{
                      'Authorization' : localStorage.getItem('refresh_token')
                    }
                  }).then(res => {
                    // refresh_token 过期
                    console.log(res)
                    if (res.data.code == 4030) {
                        setTimeout(() => {
                          router.replace({
                              path: '/login',
                              query: {
                                  'redirect': router.currentRoute.fullPath
                              }
                          });
                        }, 1000);
                    }else{
                      localStorage.setItem("access_token",res.data.data)
                      router.go(0)
                    }
                  })
                  break;
              // 404请求不存在
              case 404:
                this.$message({
                    title:'404',
                    message:'请求页面不存在',
                    type:'error'
                })
                  break;
              // 其他错误，直接抛出错误提示
              default:
                this.$message({
                  title:'404',
                  message:'请求页面不存在',
                  type:'error'
              })
          }
          // return error.response
      }
      return error
  }
); 

// //禁止F12键盘事件
// document.addEventListener('keydown', function(event){
//   return 123 != event.keyCode || (event.returnValue = false)
// });
//禁止右键、选择、复制
// document.addEventListener('contextmenu', function(event){
//   return event.returnValue = false
// })

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(CSS)

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
