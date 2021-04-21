<template>
    <el-card class="box-card">
        <div slot="header" class="clearfix">
            <div style="margin-top: 15px;">
               <el-row>
                   <el-col :span="2">
                        <el-button type="success" icon="el-icon-plus" circle  @click="openDiaolog('添加管理员',-1)"></el-button>
                    </el-col>
                    <el-col :span="6" class="search-box">
                        <el-input placeholder="请输入内容" v-model="form.name" class="input-with-select" autosize v-on:keyup.enter.native="search">
                            <el-button slot="append" icon="el-icon-search" @click="search"></el-button>
                        </el-input>
                    </el-col>
                </el-row>
               
            </div>
        </div>
        <div class="text item">
            <el-table  :data="tableData"  max-height="650"  border >
                <el-table-column  prop="id"  label="id"  sortable ></el-table-column> 
                <el-table-column  prop="img"  label="头像">
                    <template slot-scope="scope">
                        <el-image :src="scope.row.head_img"  :preview-src-list=[scope.row.head_img] class="user-img"></el-image>
                    </template>
                </el-table-column>
                <el-table-column  prop="account"  label="账号"></el-table-column>
                <el-table-column  prop="nickname"   label="昵称"></el-table-column>
                <el-table-column  prop=""  label="状态" >
                    <template slot-scope="scope">
                        <el-switch  v-model="scope.row.status" 
                                    :active-value="0" 
                                    :inactive-value="1" 
                                    active-color="#13ce66"  
                                    inactive-color="#ff4949"
                                    @change="changeSwitch(scope.row)">
                        </el-switch>
                    </template>
                </el-table-column>
                <el-table-column  prop="group_id"   label="权限组"  sortable ></el-table-column>
                <el-table-column  prop="phone"  label="电话"></el-table-column>
                <el-table-column  prop="wechat"  label="微信" ></el-table-column>
                <el-table-column  prop="create_time"  label="创建时间" :formatter="dateFormat"></el-table-column>
                <el-table-column fixed="right" label="操作" sortable>
                    <template slot-scope="scope">
                        <el-button type="primary" icon="el-icon-edit" circle @click="openDiaolog('编辑管理员',scope.row.id)"></el-button>
                        <el-button type="danger" icon="el-icon-delete" circle @click.native.prevent="deleteRow(scope.row.id, tableData)"></el-button>
                    </template>
                </el-table-column>
            </el-table>

            <div class="block pager">
                <el-pagination
                    @prev-click="changePage"
                    @next-click="changePage"
                    @size-change="changePageLength"
                    @current-change="changePage"
                    :page-sizes="[10, 20, 30, 50]"
                    :page-size="pageSize"
                     layout="total, sizes, prev, pager, next, jumper"
                    :total="total">
                </el-pagination>
              </div>
            
        </div>

        <!-- 增加和编辑 -->
        <el-dialog :title="title" :visible.sync="dialogFormVisible" :modal-append-to-body='false' class="dialog" @close="close">
            <el-form :model="form">


                <el-row>
                    <el-col :span="4" :offset="9" >
                        <el-image :src="form.head_img" class= "user-img" fit="cover"></el-image>
                    </el-col>
                </el-row>

                 <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="账号" :label-width="formLabelWidth"  >
                            <el-input v-model="form.account" autocomplete="off" placeholder="请输入管理员账号" :disabled="form.id > 0"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="名称：" :label-width="formLabelWidth" >
                            <el-input v-model="form.nickname" autocomplete="off" placeholder="请输入管理员名称"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="电话:" :label-width="formLabelWidth" >
                            <el-input v-model="form.phone" autocomplete="off" placeholder="请输入电话"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="微信：" :label-width="formLabelWidth">
                            <el-input v-model="form.wechat" autocomplete="off" placeholder="请输入微信"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="权限组" :label-width="formLabelWidth">
                            <el-select filterable v-model="form.group_id" placeholder="请选择权限组" class="slect-box">
                                <el-option label="超级管理员" :value="group_id" ></el-option>
                                <el-option v-for="item in group_id" :label="item.name" :value="item.id" :key="item.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row> 

            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="send">确 定</el-button>
            </div>
        </el-dialog>

    </el-card>

</template>

<script>
    import axios from 'axios'
    import  qs from 'qs'
    export default {
        name : 'User',
        data() {
            return {
                pageSize: 2,
                total : 1000,
                page : '',
                length : 10,
                tableData : [],
                title : '',
                imgSrc : 'http://127.0.0.1:8090/static/head/default_head.jpg',
                dialogFormVisible: false,
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
            deleteRow(index, rows) {
                console.log(index, rows)
                rows.splice(index, 1);
            },
            openDiaolog(title,index){
                this.dialogFormVisible = true //打开模态框
                this.title = title            //设置标题
                if (index < 0) {
                    this.form.img = this.imgSrc
                    return
                }
                this.getOne(index)
                
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
            getList(){
                let _this = this
                axios.get('/user',{
                    params : {
                        page:this.page,
                        length:this.length
                    }
                }).then(function(res){
                    _this.tableData =  res.data.data
                    _this.total = res.data.total
                }).catch(error => {
                    _this.$message.error(error)
                })
            },
            send(){
                this.updateData()
            },
            close(){
                this.form = {} 
            },
            search(){
                let _this = this
                axios({
                    method :'GET',
                    url : '/user',
                    params : {search : _this.form.name}
                }).then(function(res){
                    _this.tableData =  res.data.data
                    _this.total = res.data.total
                }).catch(error => {
                    _this.$message.error(error)
                })
            },
            changeSwitch(index){
                console.log(index)
            },
            changePage(page){
                this.page = page
                this.getList()
            },
            changePageLength(length){
                this.length = length
                this.getList()
            }
        },
        created() {
           this.getList()
        }
    }
</script>