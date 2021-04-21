<template>
  <div class="login" >
    <el-form class="login-form" :model="user"  :rules="rules" ref="ruleForm" id="qrcode">
        <el-form-item  prop="account">
            <el-input v-model="user.account" placeholder="请输入账号"></el-input>
        </el-form-item>
        <el-form-item  prop="password">
            <el-input v-model="user.password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item class="login-div">
            <el-button class="login-btn" type="primary" @click="submitForm('ruleForm')" :loading="loading">登陆</el-button>
            <el-button class="wechat-login-btn" type="primary" @click="wxlogin" :loading="loading">微信登陆</el-button>
        </el-form-item>
    </el-form>

  </div>
</template>
<script type='text/javascript' src="https://res.wx.qq.com/connect/zh_CN/htmledition/js/wxLogin.js"></script>
<script>
    import axios from 'axios'
    import  qs from 'qs'
    export default {
    name : 'Login',
    data() {
        return {
            loading : false,
            redirect : '/home',
            user: {
                account: '',
                password:''
            },
            rules: {
                name: [
                    { required: true, message: '请输入账号', trigger: 'blur' },
                    { min: 1, max: 10, message: '长度在 3 到 5 个字符', trigger: 'blur' }
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'change' },
                    { min: 1,  message: '长度大于6 个字符', trigger: 'blur' }
                ]
            }
      };
    },
    methods: {
        submitForm(formName) {
            this.loading = true
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    let _this = this
                    axios({
                        method :'POST',
                        url : '/user/login',
                        data : qs.stringify(this.user)
                    }).then(function(res){
                        if (res.data.code==200) {
                            console.log(res.data)
                            window.localStorage.setItem("userInfo",JSON.stringify(res.data.userInfo))
                            window.localStorage.setItem("access_token",res.data.data.access_token)
                            window.localStorage.setItem("refresh_token",res.data.data.refresh_token)
                            _this.$message({
                                title:'成功',
                                message:'登陆成功',
                                type:'success'
                            })
                            
                            _this.$router.push(_this.redirect)
                        }else{
                            _this.$message.error(res.data.msg)
                        }
                    }).catch(function(){
                        _this.$message.error('网络出错')
                    })
                }
            });
            this.loading=false
        },
        wxlogin(){
        //用于生成二维码的方法
            var obj=new WxLogin({
                self_redirect:false,
                id:'qrcode',//显示二维码的容器id
                appid:'wx89e240496b95918a',//在开放平台申请到的appid
                scope:'snsapi_login',
                redirect_uri:encodeURIComponent('http://120.78.167.190/wechat'),//扫码后跳转的链接，网址必须是在开放平台申请时填写的那个
                state:'123456',
                style:'black',
                href:'',
            })
        }
    },
    created() {
        this.redirect = this.$route.query.redirect
        if (!this.redirect) {
            this.redirect =  "/home"
        }
    }
}
</script>