<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表三 工业生产过程CO2排放的活动水平和排放因子数据一览表</text>
    </template>
      <el-table :data="IndustryInput" border style="width:96%; margin-left: 8px; margin-top: 20px">
        <el-table-column label="碳输入">
          <el-table-column label="物料名称" prop="Id" width="120">
            <template #default="scope">
              <el-input v-model="scope.row.Breed" placeholder="请输入名称"></el-input>
            </template>
          </el-table-column>
          <el-table-column label="活动水平数据(单位:吨或万Nm3)" width="160">
            <template #default="scope">
              <el-input v-model="scope.row.AD" placeholder="请输入数据"></el-input>
            </template>
          </el-table-column>
          <el-table-column label="含碳量(单位:tC/吨)" width="140">
            <template #default="scope">
              <el-input v-model="scope.row.CC" placeholder="请输入数据"></el-input>
            </template>
          </el-table-column>
          <el-table-column label="数据来源" width="260"  height="120">
            <template #default="scope">
              <el-radio-group v-model="scope.row.CCDS" size="small" style="display: flex;" >
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
        </el-table-column>
      </el-table>
        <el-button type="primary" @click="addItem1" style="margin-top: 10px;margin-left: 350px">添加新物料</el-button>
    <el-table :data="IndustryOutput" border style="width:96%; margin-left: 8px; margin-top: 20px">
      <el-table-column label="碳输出">
        <el-table-column label="物料名称" prop="Id" width="120">
          <template #default="scope">
            <el-input v-model="scope.row.Breed" placeholder="请输入名称"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="活动水平数据(单位:吨或万Nm3)" width="160">
          <template #default="scope">
            <el-input v-model="scope.row.AD" placeholder="请输入数据"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="含碳量(单位:tC/吨)" width="140">
          <template #default="scope">
            <el-input v-model="scope.row.CC" placeholder="请输入数据"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="数据来源" width="260"  height="120">
          <template #default="scope">
            <el-radio-group v-model="scope.row.CCDS" size="small" style="display: flex;" >
              <el-radio :label="0">检测值</el-radio><br>
              <el-radio :label="1" style="margin-left: -10px">化学计算</el-radio>
              <el-radio :label="2" style="margin-left: -10px">缺省值</el-radio>
            </el-radio-group>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作">
          <template #default="{ row }">
            <el-button type="primary" @click="deleteRow2(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
    <el-button type="primary" @click="addItem2" style="margin-top: 10px;margin-left: 350px">添加新物料</el-button>
  </el-card>

</template>
<script>
export default {
  data(){
    return{
      IndustryInput:[
        {
          Breed: '',
          AD: null,
          CC: null,
          CCDS: 0
        },
      ],
      IndustryOutput:[
        {
          Breed: '',
          AD: null,
          CC: null,
          CCDS: 0
        },
      ],
    }
  },
  methods:{
    deleteRow1(row){
      const index = this.IndustryInput.indexOf(row)
      if (index !== -1){
        this.IndustryInput.splice(index, 1)
      }
    },
    addItem1(){
      this.IndustryInput.push({
        Breed: '',
        AD: null,
        CC: null,
        CCDS: 0
      });
    },
    deleteRow2(row){
      const index = this.IndustryOutput.indexOf(row)
      if (index !== -1){
        this.IndustryOutput.splice(index, 1)
      }
    },
    addItem2(){
      this.IndustryOutput.push({
        Breed: '',
        AD: null,
        CC: null,
        CCDS: 0
      });
    },
  },
  watch: {
    IndustryInput: {
      deep: true,
      handler(newIndustryInput) {
        this.$store.commit('updateIndustryInput', newIndustryInput); // 使用正确的 mutation 名称
      }
    },
    IndustryOutput: {
      deep: true,
      handler(newIndustryOutput) {
        this.$store.commit('updateIndustryOutput', newIndustryOutput); // 使用正确的 mutation 名称
      }
    },
  }
}
</script>
