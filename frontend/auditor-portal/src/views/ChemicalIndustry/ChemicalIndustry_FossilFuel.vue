<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表二 化石燃料燃烧的活动水平和排放因子数据一览表</text>
    </template>
    <el-table :data="FossilFuel" border style="width:96%; margin-left: 8px; margin-top: 20px">
      <el-table-column label="燃料名称" width="120">
        <template #default="scope">
          <el-input v-model="scope.row.Breed" placeholder="请输入名称"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="燃烧量(吨或万Nm3)" width="160">
        <template #default="scope">
          <el-input v-model="scope.row.AD" placeholder="请输入数据"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="含碳量(tC/吨或tC/万Nm3)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.CC" placeholder="请输入数据"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="数据来源" width="170"  height="120">
        <template #default="scope">
          <el-radio-group v-model="scope.row.CCDS" size="small" style="display: flex;" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-left: -10px">计算值</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
      <el-table-column label="低位发热量*(GJ/吨或GJ/万Nm3)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.NCV" placeholder="请输入数据"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="数据来源" width="170"  height="120">
        <template #default="scope">
          <el-radio-group v-model="scope.row.NCVDS" size="small" style="display: flex;" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-left: -10px">计算值</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
      <el-table-column label="单位热值含碳量*(tC/GJ)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.EF" placeholder="请输入数据"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="碳氧化率(%)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.OF" placeholder="请输入数据"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="数据来源" width="170"  height="120">
        <template #default="scope">
          <el-radio-group v-model="scope.row.OFDS" size="small" style="display: flex;" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-left: -10px">计算值</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作">
        <template #default="{ row }">
          <el-button type="primary" @click="deleteRow1(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button type="primary" @click="addItem1" style="margin-top: 10px;margin-left: 340px">添加新燃料</el-button>
  </el-card>
</template>

<script>
export default {
  data(){
    return{
      FossilFuel:[
        {
          Breed: '无烟煤',
          AD: null,
          CC: null,
          CCDS: 0,
          NCV: null,
          NCVDS: 0,
          EF: null,
          OF: null,
          OFDS: 0
        },
        {
          Breed: '烟煤',
          AD: null,
          CC: null,
          CCDS: 0,
          NCV: null,
          NCVDS: 0,
          EF: null,
          OF: null,
          OFDS: 0
        },
        {
          Breed: '褐煤',
          AD: null,
          CC: null,
          CCDS: 0,
          NCV: null,
          NCVDS: 0,
          EF: null,
          OF: null,
          OFDS: 0
        },
        {
          Breed: '洗精煤',
          AD: null,
          CC: null,
          CCDS: 0,
          NCV: null,
          NCVDS: 0,
          EF: null,
          OF: null,
          OFDS: 0
        },
      ],
    }
  },
  methods:{
    deleteRow1(row){
      const index = this.FossilFuel.indexOf(row)
      if (index !== -1){
        this.FossilFuel.splice(index, 1)
      }
    },
    addItem1(){
      this.FossilFuel.push({
        Breed: '',
        AD: null,
        CC: null,
        CCDS: 0,
        NCV: null,
        NCVDS: 0,
        EF: null,
        OF: null,
        OFDS: 0
      });
    },
  },
  watch: {
    FossilFuel: {
      deep: true,
      handler(newFossilFuel) {
        this.$store.commit('updateFossilFuel', newFossilFuel);
      }
    },
  }
}
</script>
