<template>
    <el-card class="box-card">
        <div slot="header" class="clearfix">
            <span>个人详情</span>
        </div>
        <div class="text item  box-card-main">
            <el-row>
                <el-col><el-image :src="form.head_img"  :preview-src-list=[form.head_img] class="user-img"></el-image></el-col>
            </el-row>
            <el-row :gutter="10">
                <el-col :offset="3" :span="5"><div class="grid-content bg-purple font-rigth">账号：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.account}}</div></el-col>
                <el-col  :span="5"><div class="grid-content bg-purple font-rigth" >昵称：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.nickname}}</div></el-col>
            </el-row>
            <el-row :gutter="10">
                <el-col :offset="3" :span="5"><div class="grid-content bg-purple font-rigth">电话：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.phone}}</div></el-col>
                <el-col  :span="5"><div class="grid-content bg-purple font-rigth" >微信：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.wechat}}</div></el-col>
            </el-row>
            <el-row :gutter="10">
                <el-col :offset="3" :span="5"><div class="grid-content bg-purple font-rigth">上次登陆时间：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.account}}</div></el-col>
                <el-col  :span="5"><div class="grid-content bg-purple font-rigth" >注册时间：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.nickname}}</div></el-col>
            </el-row>
            <el-row :gutter="10">
                <el-col :offset="3" :span="5"><div class="grid-content bg-purple font-rigth">上次登陆ip：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.account}}</div></el-col>
                <el-col  :span="5"><div class="grid-content bg-purple font-rigth" >本次登陆ip：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.nickname}}</div></el-col>
            </el-row>
            <el-row :gutter="10">
                <el-col :offset="3" :span="5"><div class="grid-content bg-purple font-rigth">权限组：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.account}}</div></el-col>
                <el-col  :span="5"><div class="grid-content bg-purple font-rigth" >本次登陆ip：</div></el-col>
                <el-col :span="5"><div class="grid-content bg-purple font-left">{{form.nickname}}</div></el-col>
            </el-row>
        </div>
    </el-card>


</template>

<script>
    import axios from 'axios'
    import  qs from 'qs'
    export default {
        name : 'UserInfo',
        data() {
            return {
                form: {
                    id : '',
                    head_img : '',
                    account: '',
                    nickname: '',
                    status: 0,
                    group_id: 1,
                    phone: '',
                    wechat: '',
                    create_time : ''
                },
                group_id: 1,
                group_ids : [],
                formLabelWidth: '170px',
                span : 14,
                offset :2,

            }
        },
        methods: {
            dateFormat(row){
                let date = new Date(row.create_time*1000);
                let Y = date.getFullYear() + '-';
                let M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
                let D = date.getDate() + ' ';
                let h = date.getHours() + ':';
                let m = date.getMinutes() + ':';
                let s = date.getSeconds(); 
                return Y+M+D+h+m+s
            },
            getOne(index){
                let _this = this
                axios({
                    method :'GET',
                    url : '/user/get',
                    params : {"id":index}
                }).then(function(res){
                    if (res.data.code == 200) {
                        _this.form =  res.data.data
                    }
                }).catch(error => {
                    _this.$message.error(error)
                })
            },
            updateData(){
                let _this = this
                axios({
                    method: 'post',
                    url : '/user/updateData',
                    data : qs.stringify(this.form)
                }).then(function(res){
                    if (res.data.code==200) {
                        _this.dialogFormVisible = false
                        _this.getList()
                        _this.$message({
                            title:'成功',
                            message:res.data.msg,
                            type:'success'
                        })
                    }else{
                       _this.$message.error(res.data.msg) 
                    }
                    
                }).catch(error => {
                    _this.$message.error(error)
                })
            },

            send(){
                this.updateData()
            }
        },
        created() {
           this.form = JSON.parse(localStorage.getItem("userInfo"))
        }
    }
</script>
<style>
  .el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }

</style>