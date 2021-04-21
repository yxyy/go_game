<template>
    <el-card class="box-card">
        <div slot="header" class="clearfix">
            <div style="margin-top: 15px;">
               <el-row>
                   <el-col :span="2"><el-button type="success" icon="el-icon-plus" circle @click="dialogFormVisible = true"></el-button></el-col>
                    <el-col :span="8" class="search-box">
                        <el-input placeholder="请输入内容" v-model="search" class="input-with-select" autosize>
                            <el-button slot="append" icon="el-icon-search"></el-button>
                        </el-input>
                    </el-col>
                </el-row>
               
            </div>
        </div>
        <div class="text item">
            <el-table :data="tableData" style="width: 100%" max-height="600">
                <el-table-column fixed  prop="date"  label="日期" ></el-table-column>
                <el-table-column  prop="name"  label="姓名" ></el-table-column>
                <el-table-column  prop="province"  label="省份" ></el-table-column>
                <el-table-column prop="city" label="市区"></el-table-column>
                <el-table-column prop="address" label="地址"></el-table-column>
                <el-table-column prop="zip" label="邮编"></el-table-column>
                <el-table-column fixed="right" label="操作">
                    <template slot-scope="scope">
                        <el-button type="primary" icon="el-icon-edit" circle @click="dialogFormVisible = true"></el-button>
                        <el-button type="danger" icon="el-icon-delete" circle @click.native.prevent="deleteRow(scope.$index, tableData)"></el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div v-show="total > 0">
                <span  class="total">共 {{total}} 条记录</span>
                <el-pagination class="pager"  background  layout="prev, pager, next"  :total="total"></el-pagination>
            </div>
            
        </div>

        <!-- 增加和编辑 -->
        <el-dialog :title="title" :visible.sync="dialogFormVisible" :modal-append-to-body='false' class="dialog">
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
                            <el-select v-model="form.parent" placeholder="请选择父菜单" class="slect-box">
                                <el-option label="数据分析" value="1"></el-option>
                                <el-option label="系统设置" value="2"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row> 

                <el-row>
                    <el-col :span="span" :offset="offset" >
                        <el-form-item label="图标" :label-width="formLabelWidth">
                            <el-select v-model="form.icon" placeholder="请选择图标" class="slect-box">
                                <el-option label="区域一" value="shanghai"></el-option>
                                <el-option label="区域二" value="beijing"></el-option>
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>                
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
            </div>
        </el-dialog>

    </el-card>

</template>

<script>
export default {
    name : 'Home',
    data() {
      return {
        total : 1000,
        search : '',
        tableData: [{
          date: '2016-05-03',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-02',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-04',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-01',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-08',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-06',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-07',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-08',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }, {
          date: '2016-05-09',
          name: '王小虎',
          province: '上海',
          city: '普陀区',
          address: '上海市普陀区金沙江路 1518 弄',
          zip: 200333
        }],
        title : '添加菜单',
        dialogFormVisible: false,
        form: {
          name: '',
          mark: '',
          url: '',
          parent: '',
          icon: ''
        },
        formLabelWidth: '170px',
        span : 14,
        offset :2
      }
    },
    methods: {
      deleteRow(index, rows) {
        rows.splice(index, 1);
      }
    }
    // created() {
    //     console.log("狗子函数");
    // }
}
</script>

<style scoped>

    .dialog{
        width: 80%;
        margin-left: 10%;
    }

    .search-box{
        float: right;
    }
    .total{
        float: left;
        margin-top: 20px;
    }
    .pager {
        float: right;
        margin-top: 20px;
    }
    .slect-box{
        width: 100%;
    }
    
</style>