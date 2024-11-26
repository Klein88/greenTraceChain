<template>
<!--  <el-table :data="tableData2" border style="width: 96%; margin-left: 20px; margin-top: 20px">-->
<!--    <el-table-column label="六氟化硫回收">-->
<!--      <el-table-column label="修理设备" prop="PGERepairREC.Id" width="90"></el-table-column>-->
<!--      <el-table-column label="设备容量(千克)" width="140">-->
<!--        <template #default="scope">-->
<!--          <el-input v-model="scope.row.PGERepairREC.RECV" placeholder="请输入数据"></el-input>-->
<!--        </template>-->
<!--      </el-table-column>-->
<!--      <el-table-column label="实际回收量(千克)" width="140">-->
<!--        <template #default="scope">-->
<!--          <el-input v-model="scope.row.PGERepairREC.RECR" placeholder="请输入数据"></el-input>-->
<!--        </template>-->
<!--      </el-table-column>-->
<!--      <el-table-column label="退役设备" prop="PGERepairREC.Id" width="90"></el-table-column>-->
<!--      <el-table-column label="设备容量(千克)" width="140">-->
<!--        <template #default="scope">-->
<!--          <el-input v-model="scope.row.PGERepairREC.RECV" placeholder="请输入数据"></el-input>-->
<!--        </template>-->
<!--      </el-table-column>-->
<!--      <el-table-column label="实际回收量(千克)" width="140">-->
<!--        <template #default="scope">-->
<!--          <el-input v-model="scope.row.PGERepairREC.RECR" placeholder="请输入数据"></el-input>-->
<!--        </template>-->
<!--      </el-table-column>-->
<!--      <el-table-column fixed="right" label="操作">-->
<!--        <template #default="{ row }">-->
<!--          <el-button type="primary" @click="deleteRow(row)">删除</el-button>-->
<!--        </template>-->
<!--      </el-table-column>-->
<!--    </el-table-column>-->
<!--  </el-table>-->
  <div style="display: flex">
    <el-table :data="PGERepairREC" border style="width: 96%; margin-left: 8px; margin-top: 20px">
      <el-table-column label="六氟化硫回收">
        <el-table-column label="修理设备" prop="Id" width="81"></el-table-column>
        <el-table-column label="设备容量(千克)" width="133">
          <template #default="scope">
            <el-input v-model="scope.row.RECV" placeholder="请输入数据"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="实际回收量(千克)" width="134">
          <template #default="scope">
            <el-input v-model="scope.row.RECR" placeholder="请输入数据"></el-input>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作">
          <template #default="{ row }">
            <el-button type="primary" @click="deleteRow1(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
    <el-table :data="PGERetireREC" border style="width: 96%;margin-left: 2px; margin-top: 20px">
      <el-table-column label="六氟化硫回收">
        <el-table-column label="退役设备" prop="Id" width="81"></el-table-column>
        <el-table-column label="设备容量(千克)" width="133">
          <template #default="scope">
            <el-input v-model="scope.row.RECV" placeholder="请输入数据"></el-input>
          </template>
        </el-table-column>
        <el-table-column label="实际回收量(千克)" width="134">
          <template #default="scope">
            <el-input v-model="scope.row.RECR" placeholder="请输入数据"></el-input>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作">
          <template #default="{ row }">
            <el-button type="primary" @click="deleteRow2(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
  </div>
  <el-button class="mt-4" @click="addItem()" style="margin-top: 10px; margin-left: 395px">新增设备</el-button>
</template>
<script>
export default {
  data(){
    return{
      PGERepairREC:[
        {
          Id:1,
          RECV:'',
          RECR:''
        }
      ],
      PGERetireREC:[
        {
          Id:1,
          RECV:'',
          RECR:''
        }
      ]
    }
  },
  methods:{
    deleteRow1(row){
      const index = this.PGERepairREC.indexOf(row)
      if (index !== -1){
        this.PGERepairREC.splice(index, 1)
      }
    },
    deleteRow2(row){
      const index = this.PGERetireREC.indexOf(row)
      if (index !== -1){
        this.PGERetireREC.splice(index, 1)
      }
    },
    addItem(){
      const repairequipment = this.PGERepairREC.length + 1
      const retiredequipment = this.PGERetireREC.length + 1
      this.PGERepairREC.push({
        Id: repairequipment,
      });
      this.PGERetireREC.push({
        Id: retiredequipment,
      });
    },
  },
  watch: {
    tableData2: {
      deep: true,
      handler(newTableData) {
        // 将 tableData2 传递到全局变量
        this.$store.commit('updateTableData2', newTableData); // 使用正确的 mutation 名称
      }
    },
    PGERepairREC:{
      deep: true,
      handler(newTableData) {
        // 将 tableData2 传递到全局变量
        this.$store.commit('updatePGERepairREC', newTableData); // 使用正确的 mutation 名称
      }
    },
    PGERetireREC:{
      deep: true,
      handler(newTableData) {
        // 将 tableData2 传递到全局变量
        this.$store.commit('updatePGERetireREC', newTableData); // 使用正确的 mutation 名称
      }
    }
  }
}
</script>
