<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表四 碳酸盐使用的活动水平和排放因子数据一览表</text>
    </template>
    <el-table :data="M2CO3" border style="width:96%; margin-left: 8px; margin-top: 20px">
        <el-table-column label="碳酸盐种类" prop="Id" width="120">
          <template #default="scope">
            <el-input v-model="scope.row.Breed" placeholder="请输入名称"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="消耗量(单位：吨)" width="160">
          <template #default="scope">
            <el-input v-model="scope.row.AD" placeholder="请输入数据"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="CO2排放因子(吨CO2/吨碳酸盐)" width="140">
          <template #default="scope">
            <el-input v-model="scope.row.EF" placeholder="请输入数据"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="数据来源" width="260"  height="120">
          <template #default="scope">
            <el-radio-group v-model="scope.row.EFDS" size="small" style="display: flex;" >
              <el-radio :label="0">检测值</el-radio><br>
              <el-radio :label="1" style="margin-left: -10px">化学计算</el-radio>
              <el-radio :label="2" style="margin-left: -10px">缺省值</el-radio>
            </el-radio-group>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作">
          <template #default="{ row }">
            <el-button type="primary" @click="deleteRow1(row)">删除</el-button>
          </template>
        </el-table-column>
    </el-table>
    <el-button type="primary" @click="addItem1" style="margin-top: 10px;margin-left: 330px">添加新碳酸盐种类</el-button>
  </el-card>

</template>
<script>
export default {
  data(){
    return{
      M2CO3:[
        {
          Breed: '',
          AD: null,
          EF: null,
          EFDS: 0
        },
      ],
    }
  },
  methods:{
    deleteRow1(row){
      const index = this.M2CO3.indexOf(row)
      if (index !== -1){
        this.M2CO3.splice(index, 1)
      }
    },
    addItem1(){
      this.M2CO3.push({
        Breed: '',
        AD: null,
        EF: null,
        EFDS: 0
      });
    },
  },
  watch: {
    M2CO3: {
      deep: true,
      handler(newM2CO3) {
        this.$store.commit('updateM2CO3', newM2CO3);
      }
    },
  }
}
</script>
