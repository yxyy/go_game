<template>
    <el-menu default-active="2" class="el-menu-vertical-demo aside-menu" @open="handleOpen" @close="handleClose"
                background-color="#545c64" text-color="#fff" active-text-color="#ffd04b" :collapse="isCollapse" router>
        <el-submenu v-for="item in tableData" :index="item.url + '/' + item.name" :key="item.id">
            <template slot="title">
                <i :class="item.icon"></i>
                <span>{{item.name}}</span>
            </template>
            <el-menu-item-group v-for="child in item.children" :key="child.id">
                <el-menu-item :index="child.url"><i :class="child.icon"></i>{{child.name}}</el-menu-item>
            </el-menu-item-group>
        </el-submenu>
    </el-menu>
    
</template>

<script>
    import axios from 'axios'
    export default {
        name: "AppAside",
        data () {
            return {
                isCollapse : false,
                tableData : [],
            }
        },
        methods: {
            handleOpen(key, keyPath) {
                console.log(key, keyPath);
            },
            handleClose(key, keyPath) {
                console.log(key, keyPath);
            },
            Collapse(){
                this.isCollapse = !this.isCollapse
            }
        },
        created (){
            let _this = this
            // http({
            //     method :'GET',
            //     url : '/admin/menu/index'
            // }).then(function(res){
            //     _this.tableData =  res.data.data
            // }).catch(error => {
            //     _this.$message.error(error)
            // })

            axios.get('/admin/menu/index').then(function(res){
                 _this.tableData =  res.data.data
            }).catch(error => {
                _this.$message.error(error)
            })
        }
    }
</script>
