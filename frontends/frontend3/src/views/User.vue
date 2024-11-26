<script setup>
import {onMounted, ref, watch} from "vue";
import axios from "axios";
import {
  Refresh,
  Plus,
  Edit,
  Delete,
  Download,
  Search,
  InfoFilled
} from '@element-plus/icons-vue'
import * as XLSX from "xlsx";
import {ElMessage} from "element-plus";
import { useRouter } from 'vue-router';

const tableData = ref([])
const searchQuery = ref('')
const originalTableData = ref([])
const currentPage = ref(1)
const pageSize = ref(15)
const small = ref(false)
const background = ref(true)
const disabled = ref(false)
const router = useRouter();

onMounted(async () => {
  tableData.value = await axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetCompanyList', {})
      .then(function (response) {
        console.log(response)
        return response.data.data
      })
      .catch(function (error) {
        console.log(error)
      })
  originalTableData.value = [...tableData.value]
})

watch(searchQuery, () => {
  if (searchQuery.value) {
    tableData.value = originalTableData.value.filter(item => item.CompanyID.includes(searchQuery.value))
  } else {
    tableData.value = originalTableData.value
  }
})

const refresh = async () => {
  tableData.value = await axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetCompanyList', {})
      .then(function (response) {
        ElMessage.success('刷新成功')
        console.log(response)
        return response.data.data
      })
      .catch(function (error) {
        ElMessage.error('刷新失败')
        console.log(error)
      })
  originalTableData.value = [...tableData.value];
}

const exportToExcel = () => {
  const ws = XLSX.utils.json_to_sheet(tableData.value)
  const wb = XLSX.utils.book_new()
  XLSX.utils.book_append_sheet(wb, ws, "Sheet1")
  XLSX.writeFile(wb, "TableData.xlsx")
  ElMessage.success('导出成功');
}

const examineTagType = (value) => {
  switch (value) {
    case 0:
      return 'danger'
    case 1:
      return 'primary'
    case 2:
      return 'success'
    case 3:
      return 'warning'
    default:
      return ''
  }
}
const examineTagText = (value) => {
  switch (value) {
    case 0:
      return '未审核'
    case 1:
      return '审核中'
    case 2:
      return '已通过'
    case 3:
      return '未通过'
    default:
      return ''
  }
}
const examineOptions = [
  { text: '未审核', value: 0 },
  { text: '审核中', value: 1 },
  { text: '已通过', value: 2 },
  { text: '未通过', value: 3 }
]
const filterExamine = (value, row) => {
  return row.Examine === value
}

const typeTagType = (value) => {
  switch (value) {
    case 0:
      return 'primary'
    case 1:
      return 'primary'
    case 2:
      return 'primary'
    case 3:
      return 'primary'
    default:
      return ''
  }
}
const typeTagText = (value) => {
  switch (value) {
    case 0:
      return '企业'
    case 1:
      return '数据审核员'
    case 2:
      return '第三方监管机构'
    case 3:
      return '管理员'
    default:
      return ''
  }
}
const typeOptions = [
  { text: '企业', value: 0 },
  { text: '数据审核员', value: 1 },
  { text: '第三方监管机构', value: 2 },
  { text: '管理员', value: 3 }
]
const filterType = (value, row) => {
  return row.Type === value
}

const handleInfo = async (index, row) => {
  router.push({ path: '/Traceability', query: { CompanyID: row.CompanyID } });
  console.log(index, row)
}

const handleSizeChange = (val) => {
  console.log(`${val} items per page`)
}
const handleCurrentChange = (val) => {
  console.log(`current page: ${val}`)
}
</script>

<template>
  <el-scrollbar>
    <div class="main">
      <el-card shadow="never">
        <div class="main-top">
          <div>
            <el-tooltip content="刷新" placement="top">
              <el-button color="#40485b" :icon="Refresh" @click="refresh">刷新</el-button>
            </el-tooltip>
            <el-button type="success" :icon="Plus" disabled>添加</el-button>
            <el-button type="primary" :icon="Edit" disabled>修改</el-button>
            <el-button type="danger" :icon="Delete" disabled>删除</el-button>
            <el-tooltip content="导出为Excel" placement="top">
              <el-button type="primary" :icon="Download" @click="exportToExcel">导出</el-button>
            </el-tooltip>
          </div>
          <div>
            <el-button :icon="Search" circle></el-button>
            <el-input v-model="searchQuery" placeholder="请输入公司ID" style="width: 200px; margin-left: 10px;" clearable></el-input>
          </div>
        </div>
        <div class="main-middle">
          <el-table :data="tableData.slice((currentPage - 1) * pageSize, currentPage * pageSize)" border style="width: 100%" table-layout="auto" :header-cell-style="{ background:'#f4f4f5' }">
            <el-table-column prop="ID" label="ID" align="center" />
            <el-table-column prop="CompanyID" label="公司ID" align="center" />
            <el-table-column prop="CompanyName" label="公司名称" align="center" />
            <el-table-column prop="PhoneNumber" label="电话号码" align="center" />
            <el-table-column prop="Email" label="邮箱" align="center" />
            <el-table-column prop="Examine" label="审核状态" align="center"
                             :filters="examineOptions"
                             :filter-method="filterExamine"
                             filter-placement="bottom"
            >
              <template #default="scope">
                <el-tag :type="examineTagType(scope.row.Examine)" :disable-transitions="false">
                  {{ examineTagText(scope.row.Examine) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="Type" label="用户类型" align="center"
                             :filters="typeOptions"
                             :filter-method="filterType"
                             filter-placement="bottom"
            >
              <template #default="scope">
                <el-tag :type="typeTagType(scope.row.Type)" :disable-transitions="false">
                  {{ typeTagText(scope.row.Type) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
              <template #default="scope">
                <el-button size="small" type="primary" @click="handleInfo(scope.$index, scope.row)" :icon="InfoFilled">详情</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <el-divider style="margin: 10px auto" />
        <el-pagination
            v-model:current-page="currentPage"
            v-model:page-size="pageSize"
            :page-sizes="[15, 20, 25, 30]"
            :small="small"
            :disabled="disabled"
            :background="background"
            layout="total, sizes, ->, prev, pager, next, jumper"
            :total="tableData.length"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </el-card>
    </div>
  </el-scrollbar>
</template>

<style scoped>
.main {
  margin: 15px;
}

.main-top {
  margin: 10px 5px 10px 5px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-middle {
  margin: 5px;
}

:deep(.el-card__body) {
  padding: 0;
}

.el-pagination {
  margin: 10px;
}
</style>
