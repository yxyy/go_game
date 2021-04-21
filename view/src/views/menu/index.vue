<template>
    <el-card class="box-card">
        <div slot="header" class="clearfix">
            <div style="margin-top: 15px;">
               <el-row>
                   <el-col :span="2"><el-button type="success" icon="el-icon-plus" circle  @click="openDiaolog('添加菜单',-1)"></el-button></el-col>
                    <el-col :span="8" class="search-box">
                        <el-input placeholder="请输入内容" v-model="form.name" class="input-with-select" autosize>
                            <el-button slot="append" icon="el-icon-search" @click="serach"></el-button>
                        </el-input>
                    </el-col>
                </el-row>
               
            </div>
        </div>
        <div class="text item">
            <el-table  :data="tableData"     max-height="650"      row-key="id"  border  
                        :tree-props="{children: 'children', hasChildren: 'hasChildren'}">
                <el-table-column  prop="name"  label="名称"    sortable ></el-table-column>
                <el-table-column  prop="mark"  label="标识"    sortable ></el-table-column> 
                <el-table-column  prop="url"   label="接口url"     sortable ></el-table-column>
                <el-table-column  prop="sort"   label="排序"     sortable ></el-table-column>
                <el-table-column  prop=""  label="图标"    sortable >
                    <template slot-scope="scope">
                        <i :class="scope.row.icon"></i>
                    </template>
                </el-table-column>
                <el-table-column   label="状态" >
                    <template slot-scope="scope">
                        <el-switch  v-model="scope.row.status" 
                                    :active-value="1" 
                                    :inactive-value="0" 
                                    active-color="#13ce66"  
                                    inactive-color="#ff4949"
                                    @change="changeSwitch(scope.row)">
                        </el-switch>
                    </template>
                </el-table-column>

                <el-table-column fixed="right" label="操作" sortable>
                    <template slot-scope="scope">
                        <el-button type="primary" icon="el-icon-edit" circle @click="openDiaolog('编辑菜单',scope.row.id)"></el-button>
                        <el-button type="danger" icon="el-icon-delete" circle @click.native.prevent="deleteRow(scope.row.id, tableData)"></el-button>
                    </template>
                </el-table-column>
            </el-table>

            <div v-show="total > 0">
                <span  class="total">共 {{total}} 条记录</span>
                <el-pagination class="pager"  background  layout="prev, pager, next"  :total="total"></el-pagination>
            </div>
            
        </div>

        <!-- 增加和编辑 -->
        <el-dialog :title="title" :visible.sync="dialogFormVisible" :modal-append-to-body='false' class="dialog" @close="close">
            <el-form :model="form">
                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="菜单名称：" :label-width="formLabelWidth" >
                            <el-input v-model="form.name" autocomplete="off" placeholder="请输入菜单名称"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="标识：" :label-width="formLabelWidth" >
                            <el-input v-model="form.mark" autocomplete="off" placeholder="请输入标识"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="接口url:" :label-width="formLabelWidth" >
                            <el-input v-model="form.url" autocomplete="off" placeholder="请输入url"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="排序：" :label-width="formLabelWidth">
                            <el-input v-model="form.sort" autocomplete="off" placeholder="数字也大越靠后"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>

                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="父菜单" :label-width="formLabelWidth">
                            <el-select filterable v-model="form.parent" placeholder="请选择父菜单" class="slect-box" no-data-tex="一级菜单">
                                <el-option label="一级菜单" :value="form.parent" ></el-option>
                                <el-option v-for="item in parent" :label="item.name" :value="item.id" :key="item.id"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row> 

                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="图标" :label-width="formLabelWidth">
                            <el-select filterable v-model="form.icon" placeholder="请选择图标" class="slect-box">
                                <el-option label="" :value="form.parent" ></el-option>
                                <el-option v-for="item in icons" :value="item" :key="item" >
                                    <i :class="item"></i>
                                </el-option>
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
    import http from 'axios'
    import  qs from 'qs'
    export default {
        name : 'Menu',
        data() {
            return {
                total : 1000,
                search : '',
                tableData : [],
                title : '',
                dialogFormVisible: false,
                form: {
                    id : '',
                    name: '',
                    mark: '',
                    url: '',
                    parent: '',
                    icon: '',
                    sort : ''
                },
                parent : [],
                formLabelWidth: '170px',
                span : 14,
                offset :2,
                icons : [
                    'el-icon-s-home',
                    'el-icon-user',
                    'el-icon-setting',
                    'el-icon-more',
                    'el-icon-star-off',
                    'el-icon-star-on',
                    'el-icon-circle-plus',
                    'el-icon-circle-plus-outline',
                    'el-icon-remove-outline',
                    'el-icon-picture-outline',
                    'el-icon-bell',
                    'el-icon-s-order',
                    'el-icon-s-platform',
                    'el-icon-s-cooperation',
                    'el-icon-menu',
                    'el-icon-loading',
                    'el-icon-folder-opened',
                    'el-icon-document-remove',
                    'el-icon-search',
                    'el-icon-s-promotion',
                    'el-icon-s-data'
                ]
            }
        },
        methods: {
            deleteRow(index, rows) {
                console.log(index, rows)
                rows.splice(index, 1);
            },
            openDiaolog(title,index){
                this.dialogFormVisible = true //打开模态框
                this.title = title            //设置标题
                if (index < 0) {
                    return
                }
                this.getOne(index)
                
            },
            getpParentMenu(p){
                let _this = this
                http({
                    method :'GET',
                    url : '/admin/getParentMenu/'+p
                }).then(function(res){
                    if (res.data.code == 200) {
                        _this.parent =  res.data.menuList
                    }
                }).catch(error => {
                    _this.$message.error(error)
                })
            },
            getOne(index){
                let _this = this
                http({
                    method :'POST',
                    url : '/admin/menu',
                    params : {"id":index}
                }).then(function(res){
                    console.log(res.data)
                    if (res.data.code == 200) {
                        _this.form =  res.data.data
                    }
                }).catch(error => {
                    _this.$message.error(error)
                })
            },
            updateData(){
                let _this = this
                http({
                    method: 'post',
                    url : '/admin/menu/UpdataMenu',
                    data : qs.stringify(this.form)
                }).then(function(res){
                    _this.$message({
                        title:'成功',
                        message:res.data.msg,
                        type:'success'
                    })
                }).catch(error => {
                    _this.$message.error(error)
                })
            },
            getList(){
                let _this = this
                http({
                    method :'GET',
                    url : '/admin/menu/index'
                }).then(function(res){
                    _this.tableData =  res.data.data
                    _this.total = _this.tableData.length
                    // _this.tableData =  {}
                }).catch(error => {
                    console.log(error);
                    _this.$message.error('网络出错')
                })
            },
            send(){
                this.dialogFormVisible = false
                this.updateData()
                this.getList()
                this.form = {} 
            },
            close(){
                this.form = {} 
            },
            serach(){
                let _this = this
                http({
                    method :'GET',
                    url : '/admin/menu/index',
                    params : {name : _this.form.name}
                }).then(function(res){
                    // _this.tableData =  res.data.menuList
                    // _this.total = _this.tableData.length
                    // _this.tableData =  {}

                    console.log(res.data)
                }).catch(error => {
                    console.log(error);
                    _this.$message.error('网络出错')
                })
            }
        },
        created() {
           this.getpParentMenu(0)
           this.getList()
        }
    }
</script>
