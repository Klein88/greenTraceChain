<template>
  <div class="main2" style="display: flex; flex-direction: column">
    <el-card shadow="never" style="margin-top: 5px">
      <div style="display: flex; justify-content: center; margin-top: 40px">
        <el-steps :active="4" style="width: 90%" align-center>
          <el-step title="申报" icon="Document" description="42"></el-step>
          <el-step title="发布" icon="Upload" description="17"></el-step>
          <el-step title="审核" icon="DocumentChecked" description="6"></el-step>
          <el-step title="完成" icon="Finished" description="31"></el-step>
        </el-steps>
      </div>
      <div style="display: flex; align-items: center; margin-left: 30px; margin-top: 40px">
        <text>公司编号：</text>
        <el-input v-model="searchQuery" placeholder="请输入公司编号" style="width: 20%; height: 38px; margin-left: 5px" />
        <el-button type="primary" icon="Search" @click="search" style="width: 100px; height: 37px; font-size: 16px; margin-left: 20px">搜索</el-button>
        <el-button type="success" icon="List" @click="refresh" style="width: 180px; height: 37px; font-size: 16px; margin-left: 80px">查询所有</el-button>
      </div>
      <div class="main-middle">
        <el-table :data="tableData" border style="width: 100%" height="440" table-layout="auto" :header-cell-style="{ background:'#f4f4f5' }">
          <el-table-column prop="Id" label="ID" align="center" />
          <el-table-column prop="CompanyId" label="公司ID" align="center" />
          <el-table-column prop="createAt" label="创建时间" align="center" sortable />
          <el-table-column prop="updateAt" label="更新时间" align="center" sortable />
          <el-table-column prop="State" label="审核状态" align="center"
                           :filters="stateOptions"
                           :filter-method="filterState"
                           filter-placement="bottom">
            <template #default="scope">
              <el-tag :type="stateTagType(scope.row.State)" :disable-transitions="false">
                {{ stateTagText(scope.row.State) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="Type" label="企业类型" align="center"
                           :filters="typeOptions"
                           :filter-method="filterType"
                           filter-placement="bottom">
            <template #default="scope">
              <el-tag :type="typeTagType(scope.row.Type)" :disable-transitions="false">
                {{ typeTagText(scope.row.Type) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center">
            <template #default="scope">
              <el-button size="small" type="danger" icon="Promotion" @click="traceSource(scope.$index, scope.row)">溯源</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import {getCurrentInstance, onMounted, ref, watch} from "vue";
import { useRouter } from "vue-router";

const router = useRouter()
const currentInstance = getCurrentInstance()
const { proxy } = currentInstance

const tableData = ref([])
const searchQuery = ref('')
const originalTableData = ref([])
const dialogVisible = ref(false)
const currentCompanyData = ref({});
const Type = ref('')
const PGERepairREC = ref([])
const PGERetireREC = ref([])

onMounted(async () => {
  tableData.value = await proxy.$axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetAllTransactionData', {})
      .then(function (response) {
        console.log(response)
        return response.data.data.map(item => {
          item.createAt = item.createAt.split('T')[0];
          item.updateAt = item.updateAt.split('T')[0];
          return item;
        })
      })
      .catch(function (error) {
        console.log(error)
      })
  originalTableData.value = [...tableData.value]
})

const stateTagType = (value) => {
  switch (value) {
    case 0:
      return 'warning'
    case 1:
      return 'success'
    case 2:
      return 'danger'
    default:
      return ''
  }
}

const stateTagText = (value) => {
  switch (value) {
    case 0:
      return '未审核'
    case 1:
      return '已通过'
    case 2:
      return '未通过'
    default:
      return ''
  }
}

const stateOptions = [
  { text: '未审核', value: 0 },
  { text: '已通过', value: 1 },
  { text: '未通过', value: 2 }
]

const filterState = (value, row) => {
  return row.State === value
}

const typeTagType = (value) => {
  switch (value) {
    case "0":
      return 'primary'
    case "1":
      return 'primary'
    default:
      return ''
  }
}

const typeTagText = (value) => {
  switch (value) {
    case "0":
      return '电网企业'
    case "1":
      return '化工企业'
    default:
      return ''
  }
}

const typeOptions = [
  { text: '电网企业', value: "0" },
  { text: '化工企业', value: "1" }
]

const filterType = (value, row) => {
  return row.Type === value
}

const refresh = async () => {
  tableData.value = await proxy.$axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetAllTransactionData', {})
      .then(function (response) {
        console.log(response)
        return response.data.data.map(item => {
          item.createAt = item.createAt.split('T')[0];
          item.updateAt = item.updateAt.split('T')[0];
          return item;
        })
      })
      .catch(function (error) {
        console.log(error)
      })
  originalTableData.value = [...tableData.value];
}

watch(searchQuery, () => {
  if (searchQuery.value) {
    tableData.value = originalTableData.value.filter(item => item.CompanyId.includes(searchQuery.value))
  } else {
    tableData.value = originalTableData.value
  }
})

// const handleInfo = async (index, row) => {
//   await proxy.$axios.post('http://47.97.176.174:8087/chainmaker?cmb=SearchPGEById', {
//     Id : row.Id
//   }).then(res=>{
//     console.log(res)
//     currentCompanyData.value = res.data.data
//     if (res.data.Type == '0'){
//       Type.value = '电网企业'
//     } else if (res.data.Type == '1'){
//       Type.value = '化工企业'
//     }
//     PGERepairREC.value = [...res.data.data.Data.PGERepairREC]
//     PGERetireREC.value = [...res.data.data.Data.PGERetireREC]
//   }).catch(function (error) {
//     console.log(error);
//   })
//   dialogVisible.value = true
// }

const traceSource = (index, row) => {
  router.push({
    path: '/traceSource1/'+row.Id,
    query: {
      Id: row.Id,
      State: row.State,
      CompanyId: row.CompanyId,
      Type: row.Type,
      createAt: row.createAt,
      updateAt: row.updateAt
    }
  })
}
</script>

<style scoped>
.main2 {
  margin: 15px 15px 0 15px;
}

.main-top {
  margin: 15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-middle {
  margin: 20px 10px;
}

:deep(.el-card__body) {
  padding: 0;
}
</style>
