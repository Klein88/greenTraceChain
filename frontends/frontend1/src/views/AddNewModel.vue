<template>
  <el-card style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text style="font-size: 30px">自定义模型</text>
    </template>
    <el-button type="primary" class="mt-4" @click="addItem()" style="margin-top: 10px; margin-left: 30px">新增公式</el-button>
    <div style="display: flex;margin-left: 100px">
      <el-table :data="newmodel" border style="width: 500px; margin-left: 8px; margin-top: 20px">
        <el-table-column label="公式编号" width="150">
          <template #default="scope">
            <el-input v-model="scope.row.Id" placeholder="请输入编号"></el-input>
          </template>
        </el-table-column>
          <el-table-column label="公式" width="250">
            <template #default="scope">
              <el-input v-model="scope.row.formula" placeholder="请输入公式"></el-input>
            </template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="100">
            <template #default="{ row }">
              <el-button type="primary" @click="deleteRow1(row)">删除</el-button>
            </template>
          </el-table-column>
      </el-table>
    </div>
    <el-button type="primary" size="large" style="margin-top: 20px;margin-left: 280px" @click="hebing">合并公式</el-button>
    <div style="display: flex;margin-left: 20px;margin-top: 20px">
      <span style="font-size: 25px">公式合并为：</span>
      <el-input v-model="formulaall" style="width: 400px"></el-input>
    </div>
    <el-button type="primary" size="large" style="margin-top: 20px;margin-left: 280px" @click="nextstep">下一步</el-button>
  </el-card>
</template>
<script>
import {ElMessage} from "element-plus";

export default {
  data(){
    return{
      newmodel:[
        {
          Id:'',
          formula:''
        },
      ],
      formulaall:'',
    }
  },
  methods:{
    addItem(){
      this.newmodel.push({
        Id: '',
        formula:''
      });
    },
    deleteRow1(row){
      const index = this.newmodel.indexOf(row)
      if (index !== -1){
        this.newmodel.splice(index, 1)
      }
    },
    nextstep(){
      this.$router.push({name:'NewModel',query:{Total:this.formulaall}});
    },
    hebing(){
      const  { hebing, lastKey } = this.combineFormulas(this.newmodel);
      const formulas = {};
      hebing.forEach((value, key) => {
        formulas[key] = value;
      });
      this.$http.post('http://47.97.176.174:8087/chainmaker?cmb=MergeFormulas',{
        formulas: formulas,
        target: lastKey
      }).then((res)=>{
        console.log(res)
        this.formulaall=res.data.data
        ElMessage({
          message: '合并成功！',
          type: 'success',
        })
      }).catch((err) => {
        console.log(err);
      });
    },
    combineFormulas(array) {
      const hebing = new Map();
      let lastKey = '';
      array.forEach(item => {
        hebing.set(item.Id, item.formula);
        lastKey=item.Id
      });
      return { hebing, lastKey };
    }
  }
}
</script>
