<template>
    <el-row>
        <el-col :span="1" class="header-col"><i class="el-icon-menu"></i></el-col >
        <el-col :span="1" class="header-col"> <span>后台管理</span></el-col >
        <el-col :span="3" :offset="1" class="header-col" > 
            <el-breadcrumb separator-class="el-icon-arrow-right head-breadcrumb">
                <el-breadcrumb-item :to="{ path: '/' }" class="head-breadcrumb">首页</el-breadcrumb-item>
                <el-breadcrumb-item class="head-breadcrumb">{{routercommand}}</el-breadcrumb-item>
            </el-breadcrumb>
        </el-col >
        <el-col :span="2" class="header-col user">
            <el-dropdown   @command="userInfo">
            <el-image :src="head_img"  class="head-img"></el-image><i class="el-icon-arrow-down el-icon--right head-icon-img"></i>
            <el-dropdown-menu slot="dropdown" >
                <el-dropdown-item  command="/info">{{nickname}}</el-dropdown-item>
                <el-dropdown-item command="/logout">注销</el-dropdown-item>
            </el-dropdown-menu>
            </el-dropdown>
        </el-col >
    </el-row>
</template>

<script>
export default {
    name : 'Header',
    data() {
        return {
            nickname : '',
            head_img : '',
            routercommand : '',
            // routers : {
            //     "user":{
            //         "name":
            //         "管理员管理",
            //         "path":"/user"
            //     },
            //     "userInfo":{
            //         "name":
            //         "个人详情",
            //         "path":"/info"
            //     }
            // },
        }
    },
     methods: {
        userInfo(router){
            this.$router.push(router)
        }
    },
    created () {
        let userInfo = JSON.parse(localStorage.getItem('userInfo'))
        this.nickname = userInfo.nickname
        this.head_img = userInfo.head_img
        this.routercommand = this.$route.name ? this.$route.name : '..'
    }
}
</script>
