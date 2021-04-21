<template>
    <el-container class="home-container">
        <!-- 菜单栏 -->
        <el-aside width="auto"> 
            <app-aside class="aside-menu" /> 
        </el-aside>
        <el-container>
            <el-header>
                <app-header class="header" /> 
            </el-header>
            <el-main class="main">
                <router-view></router-view>
            </el-main>

        </el-container>
    </el-container>
</template>

<script>
import AppAside from './components/aside';
import AppHeader from './components/header';

export default {
    name: "Layout",
    components : {
        AppAside,
        AppHeader
    },
    methods: {
        handleOpen(key, keyPath) {
            console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
            console.log(key, keyPath);
        }
    }
};
</script>

<style lang="less" scoped>
    .home-container {
        position: fixed;
        left: 0;
        top: 0;
        right: 0;
        bottom: 0;
    }
    .el-header {
        color: #333;
        text-align: center;
        line-height: 60px;
        background-color: #e9eef3;
    }
    .header{
        background-color: #fff;
    }
    .el-main {
        background-color: #e9eef3;
        color: #333;
        text-align: center;
        line-height: 160px;
        float: left;
    }
    .main{
        line-height: 50px;
    }


  

</style>