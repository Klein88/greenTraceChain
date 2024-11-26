<template>
  <div style="display: flex; flex-direction: column">
    <div style="display: flex; flex-direction: row; justify-content: space-around; align-items: center; margin-top: 40px">
      <el-card shadow="never" style="width: 550px; border-radius: 5px; background: linear-gradient(to right, #94c8ea, #e5f5e0)">
        <div style="display: flex; flex-direction: row; justify-content: space-around; align-items: center">
          <div style="display: flex; flex-direction: column">
            <span style="font-size: 21px; font-weight: bold; color: white">今日审核</span>
            <div style="display: flex; flex-direction: row; align-items: center; margin-top: 10px">
              <span style="font-size: 18px; font-weight: bold; color: white">{{ sum }}</span>
              <el-icon style="margin-left: 30px; color: #67c23a"><ArrowUpBold /></el-icon>
              <span style="margin-left: 5px; color: #67c23a">+2.0%</span>
            </div>
          </div>
          <div style="width: 70px; height: 70px; background: white; border-radius: 8px; display: flex; justify-content: center; align-items: center">
<!--            <el-icon color="white" style="font-size: 28px"><InfoFilled /></el-icon>-->
            <svg height="70" node-id="1" sillyvg="true" template-height="60" template-width="60" version="1.1" viewBox="0 0 60 60" width="70" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><defs node-id="16"></defs><path d="M 38.80 27.20 L 34.40 31.50 L 34.40 27.20 L 32.10 27.20 C 30.30 27.20 28.90 25.80 28.90 24.00 L 28.90 21.20 C 21.80 21.30 16.10 27.10 16.10 34.20 C 16.10 41.40 21.90 47.20 29.10 47.20 C 36.30 47.20 42.10 41.40 42.10 34.20 C 42.10 31.60 41.30 29.20 40.00 27.20 L 38.80 27.20 Z" fill="#5d8ef9" fill-rule="nonzero" node-id="20" stroke="none" target-height="26" target-width="25.999998" target-x="16.1" target-y="21.2"/><path d="M 34.40 31.50 L 34.40 27.20 L 32.10 27.20 C 30.50 27.20 29.20 26.00 28.90 24.50 C 23.60 24.60 19.30 28.90 19.30 34.30 C 19.30 39.70 23.70 44.10 29.10 44.10 C 34.50 44.10 38.90 39.70 38.90 34.30 C 38.90 32.30 38.30 30.40 37.20 28.80 L 34.40 31.50 Z" fill="#8ac9f9" fill-rule="nonzero" node-id="22" stroke="none" target-height="19.599998" target-width="19.600002" target-x="19.3" target-y="24.5"/><path d="M 29.10 36.80 L 25.10 33.80 L 29.10 27.70 L 33.10 33.80 Z" fill="#ffffff" fill-rule="nonzero" node-id="24" stroke="none" target-height="9.099998" target-width="7.999998" target-x="25.1" target-y="27.7"/><path d="M 29.10 40.80 L 25.10 34.50 L 29.10 37.70 L 33.10 34.50 Z" fill="#ffffff" fill-rule="nonzero" node-id="26" stroke="none" target-height="6.299999" target-width="7.999998" target-x="25.1" target-y="34.5"/><path d="M 18.90 46.10 L 19.20 46.20 L 18.90 46.30 C 17.00 46.80 15.60 48.20 15.10 50.10 L 15.00 50.40 L 15.00 50.10 C 14.50 48.20 13.10 46.80 11.20 46.30 L 10.90 46.20 L 11.20 46.10 C 13.10 45.60 14.50 44.20 15.00 42.30 L 15.10 42.00 L 15.20 42.30 C 15.60 44.10 17.10 45.60 18.90 46.10 L 18.90 46.10 Z" fill="#5d8ef9" fill-rule="nonzero" node-id="28" stroke="none" target-height="8.400002" target-width="8.300001" target-x="10.9" target-y="42"/><path d="M 24.80 14.70 L 25.20 14.80 L 24.80 14.90 C 22.70 15.40 21.00 17.10 20.50 19.20 L 20.40 19.60 L 20.30 19.20 C 19.80 17.10 18.10 15.40 16.00 14.90 L 15.60 14.80 L 16.00 14.70 C 18.10 14.20 19.80 12.50 20.30 10.40 L 20.40 10.00 L 20.50 10.40 C 21.00 12.50 22.60 14.20 24.80 14.70 L 24.80 14.70 Z" fill="#ffbe1b" fill-rule="nonzero" node-id="30" stroke="none" target-height="9.6" target-width="9.6" target-x="15.6" target-y="10"/><path d="M 15.90 21.70 L 16.10 21.80 L 15.90 21.90 C 14.50 22.30 13.40 23.40 13.00 24.80 L 12.90 25.00 L 12.80 24.80 C 12.40 23.40 11.30 22.30 9.90 21.90 L 9.70 21.80 L 9.90 21.70 C 11.30 21.30 12.40 20.20 12.80 18.80 L 12.90 18.60 L 13.00 18.80 C 13.30 20.30 14.40 21.40 15.90 21.70 L 15.90 21.70 Z" fill="#ffbe1b" fill-rule="nonzero" node-id="32" stroke="none" target-height="6.3999996" target-width="6.4000006" target-x="9.7" target-y="18.6"/><path d="M 46.10 13.20 L 32.10 13.20 C 30.30 13.20 28.90 14.60 28.90 16.40 L 28.90 24.00 C 28.90 25.80 30.30 27.20 32.10 27.20 L 34.40 27.20 L 34.40 31.50 L 38.80 27.20 L 46.10 27.20 C 47.90 27.20 49.30 25.80 49.30 24.00 L 49.30 16.40 C 49.30 14.60 47.90 13.20 46.10 13.20 Z" fill="#330d84" fill-rule="nonzero" node-id="34" stroke="none" target-height="18.3" target-width="20.4" target-x="28.9" target-y="13.2"/><path d="M 38.00 24.50 C 37.70 24.50 37.30 24.40 37.10 24.10 L 34.40 21.50 C 33.90 21.00 33.90 20.10 34.40 19.60 C 34.90 19.10 35.80 19.10 36.30 19.60 L 38.00 21.30 L 42.20 16.30 C 42.70 15.70 43.50 15.70 44.10 16.10 C 44.70 16.60 44.70 17.40 44.30 18.00 L 39.00 24.10 C 38.80 24.30 38.40 24.50 38.00 24.50 C 38.10 24.50 38.00 24.50 38.00 24.50 Z" fill="#8ac9f9" fill-rule="nonzero" node-id="36" stroke="none" target-height="8.8" target-width="10.799999" target-x="33.9" target-y="15.7"/><path d="M 50.80 37.10 L 51.00 37.20 L 50.80 37.30 C 49.30 37.70 48.20 38.80 47.80 40.30 L 47.70 40.50 L 47.60 40.30 C 47.20 38.80 46.10 37.70 44.60 37.30 L 44.40 37.20 L 44.60 37.10 C 46.10 36.70 47.20 35.60 47.60 34.10 L 47.70 33.90 L 47.80 34.10 C 48.20 35.60 49.30 36.80 50.80 37.10 L 50.80 37.10 Z" fill="#330d84" fill-rule="nonzero" node-id="38" stroke="none" target-height="6.5999985" target-width="6.5999985" target-x="44.4" target-y="33.9"/></svg>
          </div>
        </div>
      </el-card>
      <el-card shadow="never" style="width: 550px; border-radius: 5px; background: linear-gradient(to right, #FFB5B4, #FFCEBF)">
        <div style="display: flex; flex-direction: row; justify-content: space-around; align-items: center">
          <div style="display: flex; flex-direction: column">
            <span style="font-size: 21px; font-weight: bold; color: white">审核未通过</span>
            <div style="display: flex; flex-direction: row; align-items: center; margin-top: 10px">
              <span style="font-size: 18px; font-weight: bold; color: white">{{ fail }}</span>
              <el-icon style="margin-left: 30px; color: red"><ArrowDownBold /></el-icon>
              <span style="margin-left: 5px; color: red">-3.2%</span>
            </div>
          </div>
          <div style="width: 70px; height: 70px; background: white; border-radius: 8px; display: flex; justify-content: center; align-items: center">
<!--            <el-icon color="white" style="font-size: 28px"><CircleCloseFilled /></el-icon>-->
            <svg height="70" node-id="1" sillyvg="true" template-height="65" template-width="65" version="1.1" viewBox="0 0 60 60" width="70" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><defs node-id="16"></defs><path d="M 18.30 45.20 L 18.60 45.30 L 18.30 45.40 C 16.50 45.90 15.10 47.30 14.60 49.10 L 14.50 49.40 L 14.40 49.00 C 13.90 47.20 12.50 45.70 10.70 45.30 L 10.40 45.20 L 10.70 45.10 C 12.50 44.60 14.00 43.20 14.40 41.40 L 14.50 41.10 L 14.60 41.40 C 15.00 43.30 16.50 44.70 18.30 45.20 L 18.30 45.20 Z" fill="#ffbe1b" fill-rule="nonzero" node-id="20" stroke="none" target-height="8.300003" target-width="8.200001" target-x="10.4" target-y="41.1"/><path d="M 15.30 17.60 L 15.50 17.70 L 15.30 17.80 C 13.90 18.10 12.90 19.20 12.50 20.60 L 12.40 20.80 L 12.30 20.60 C 12.00 19.20 10.90 18.20 9.50 17.80 L 9.30 17.70 L 9.50 17.60 C 10.90 17.30 11.90 16.20 12.30 14.80 L 12.40 14.60 L 12.50 14.80 C 12.90 16.20 13.90 17.20 15.30 17.60 L 15.30 17.60 Z" fill="#330d84" fill-rule="nonzero" node-id="22" stroke="none" target-height="6.199999" target-width="6.2" target-x="9.3" target-y="14.6"/><path d="M 28.00 13.70 C 19.30 13.70 12.30 20.70 12.30 29.40 L 43.40 26.50 C 42.10 19.20 35.70 13.70 28.00 13.70 Z" fill="#ffbe1b" fill-rule="nonzero" node-id="24" stroke="none" target-height="15.7" target-width="31.100002" target-x="12.3" target-y="13.7"/><path d="M 12.40 29.40 C 12.40 38.10 19.40 45.10 28.10 45.10 C 35.80 45.10 42.30 39.50 43.50 32.10 L 12.40 29.40 Z" fill="#ffbe1b" fill-rule="nonzero" node-id="26" stroke="none" target-height="15.699999" target-width="31.1" target-x="12.4" target-y="29.4"/><path d="M 16.30 29.70 C 16.50 36.00 21.70 41.10 28.10 41.10 C 33.80 41.10 38.50 37.10 39.60 31.70 L 16.30 29.70 Z" fill="#ffe37b" fill-rule="nonzero" node-id="28" stroke="none" target-height="11.399998" target-width="23.3" target-x="16.3" target-y="29.7"/><path d="M 39.50 26.80 C 38.30 21.50 33.60 17.60 28.00 17.60 C 21.60 17.60 16.40 22.70 16.20 29.00 L 39.50 26.80 Z" fill="#ffe37b" fill-rule="nonzero" node-id="30" stroke="none" target-height="11.4" target-width="23.3" target-x="16.2" target-y="17.6"/><path d="M 30.60 31.00 C 30.70 31.20 30.70 31.40 30.70 31.60 C 30.70 32.10 30.50 32.50 30.20 32.70 C 29.80 32.90 29.30 33.00 28.60 33.00 L 26.90 33.00 L 26.90 30.70 L 24.40 30.50 L 24.40 33.00 L 22.80 33.00 L 22.80 35.20 L 24.40 35.20 L 25.00 35.20 L 25.00 36.50 L 27.00 36.50 L 27.00 35.20 L 28.30 35.20 L 28.30 36.50 L 30.30 36.50 L 30.30 35.00 C 31.40 34.90 32.20 34.50 32.70 34.00 C 33.20 33.40 33.50 32.70 33.50 31.90 C 33.50 31.60 33.50 31.40 33.40 31.20 L 30.60 31.00 Z" fill="#ffffff" fill-rule="nonzero" node-id="32" stroke="none" target-height="6" target-width="10.700001" target-x="22.8" target-y="30.5"/><path d="M 24.30 28.30 L 26.80 28.10 L 26.80 25.90 L 28.00 25.90 C 28.80 25.90 29.30 26.00 29.60 26.10 C 29.90 26.30 30.10 26.60 30.10 27.10 C 30.10 27.40 30.10 27.60 30.00 27.70 L 32.70 27.40 C 32.80 27.20 32.80 26.90 32.80 26.60 C 32.80 25.90 32.50 25.20 32.00 24.60 C 31.70 24.30 31.30 24.00 30.70 23.80 C 30.50 23.70 30.40 23.70 30.20 23.70 L 30.20 22.20 L 28.20 22.20 L 28.20 23.50 L 26.90 23.50 L 26.90 22.20 L 24.90 22.20 L 24.90 23.50 L 24.30 23.50 L 22.70 23.50 L 22.70 25.70 L 24.30 25.70 L 24.30 28.30 Z" fill="#ffffff" fill-rule="nonzero" node-id="34" stroke="none" target-height="6.0999985" target-width="10.099998" target-x="22.7" target-y="22.2"/><path d="M 49.20 36.50 L 49.40 36.60 L 49.20 36.70 C 47.80 37.10 46.70 38.20 46.30 39.60 L 46.20 39.80 L 46.10 39.60 C 45.70 38.20 44.60 37.10 43.20 36.70 L 43.00 36.60 L 43.20 36.50 C 44.60 36.10 45.70 35.00 46.10 33.60 L 46.20 33.40 L 46.30 33.60 C 46.60 35.00 47.70 36.10 49.20 36.50 L 49.20 36.50 Z" fill="#5d8ef9" fill-rule="nonzero" node-id="36" stroke="none" target-height="6.3999977" target-width="6.4000015" target-x="43" target-y="33.4"/><path d="M 53.60 28.70 L 54.00 28.80 L 53.60 28.90 C 51.60 29.40 50.00 31.00 49.40 33.10 L 49.30 33.40 L 49.30 33.00 C 48.80 31.00 47.20 29.40 45.10 28.80 L 44.80 28.70 L 45.10 28.60 C 47.10 28.10 48.70 26.50 49.30 24.40 L 49.40 24.10 L 49.50 24.40 C 50.00 26.60 51.60 28.20 53.60 28.70 L 53.60 28.70 Z" fill="#5d8ef9" fill-rule="nonzero" node-id="38" stroke="none" target-height="9.300001" target-width="9.200001" target-x="44.8" target-y="24.1"/></svg>
          </div>
        </div>
      </el-card>
    </div>
    <el-card shadow="never" style="margin: 30px 15px 0px 15px">
      <div style="display: flex; align-items: center; margin-top: 10px;">
        <text style="font-size: 16px; margin-left: 20px">公司编号：</text>
        <el-input v-model="companyid" placeholder="请输入公司编号" style="width: 180px; height: 37px; font-size: 16px; margin-left: 5px"></el-input>
        <el-button type="danger" icon="CircleClose" @click="nopass" style="width: 180px; height: 37px; font-size: 16px; margin-left: 20px">审核不通过</el-button>
        <el-button type="warning" icon="List" @click="getTableData" style="width: 180px; height: 37px; font-size: 16px; margin-left: 80px">查询所有</el-button>
      </div>
      <el-table :data="tableData" border height="420" :header-cell-style="{ background:'#f4f4f5' }" style="width: 98%; margin-left: 10px; margin-top: 20px">
        <el-table-column prop="ID" label="编号" align="center" width="70"></el-table-column>
        <el-table-column prop="CompanyID" label="用户编号" align="center" width="150"></el-table-column>
        <el-table-column prop="CompanyName" label="用户名称" align="center" width="180"></el-table-column>
        <el-table-column prop="Type" label="用户类型" align="center" width="130">
          <template #default="{ row }">
            <el-tag :type="userType(row.Type)">
              {{ userText(row.Type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="CarbonCoin" label="碳币" align="center" width="130"></el-table-column>
        <el-table-column prop="CarbonCredit" label="碳额度" align="center" width="130"></el-table-column>
        <el-table-column prop="Examine" label="账户审核情况" align="center" width="140"
                         :filters="examineOptions"
                         :filter-method="filterExamine"
                         filter-placement="bottom">
          <template #default="{ row }">
            <el-tag :type="examineType(row.Examine)">
              {{ examineText(row.Examine) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" width="143">
          <template #default="scope">
            <el-button type="success" icon="CircleCheck" @click="pass(scope.row)">审核通过</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
  <el-dialog v-model="dialogVisible" title="分发碳币碳额度" style="display: flex; flex-direction: column; align-items: center">
    <el-form :inline="true" style="margin-top: 20px">
      <el-form-item label="公司编号" style="width: 260px">
        <el-input v-model="CompanyID" placeholder="请输入公司编号"></el-input>
      </el-form-item>
      <el-form-item label="分发碳币" style="width: 260px">
        <el-input v-model="CarbonCoin" :disabled="disableInput" placeholder="请输入分发的碳币数量"></el-input>
      </el-form-item>
      <el-form-item label="分发碳额度" style="width: 260px">
        <el-input v-model="CarbonCredit" :disabled="disableInput" placeholder="请输入分发的碳币额度"></el-input>
      </el-form-item>
    </el-form>
    <el-button type="danger" @click="determine" style="width: 150px; margin-left: 210px; margin-top: 20px">确定分发</el-button>
  </el-dialog>
</template>

<script setup>
import {computed, getCurrentInstance, onMounted, ref} from "vue";
import {ElMessage} from "element-plus";

const currentInstance = getCurrentInstance()
const { proxy } = currentInstance

const tableData = ref([])
const sum = ref(0)
const fail = ref(0)

onMounted(() => {
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=GetCompanyList").then(res=>{
    console.log(res)
    tableData.value = res.data.data
    sum.value = res.data.data.length
    for (let i = 0; i < res.data.data.length; i++){
      if (res.data.data[i].Examine == 3){
        fail.value += 1
      }
    }
  }).catch(err=>{
    console.log(err)
  })
})

const examineOptions = [
  { text: '未审核', value: 0 },
  { text: '审核中', value: 1 },
  { text: '已通过', value: 2 },
  { text: '未通过', value: 3 }
]

const filterExamine = (value, row) => {
  return row.Examine === value
}

const examineType = (Type) => {
  switch (Type) {
    case 0:
      return 'default';
    case 1:
      return 'warning';
    case 2:
      return 'success';
    case 3:
      return 'danger';
    default:
      return '';
  }
}

const examineText = (Type) => {
  switch (Type) {
    case 0:
      return '未审核';
    case 1:
      return '审核中';
    case 2:
      return '已通过';
    case 3:
      return '未通过';
    default:
      return '';
  }
}

const userType = (Type) => {
  switch (Type) {
    case 0:
      return 'primary';
    case 1:
      return 'success';
    case 2:
      return 'warning';
    case 3:
      return 'danger';
    default:
      return '';
  }
}

const userText = (Type) => {
  switch (Type) {
    case 0:
      return '企业';
    case 1:
      return '数据审核员';
    case 2:
      return '第三方监管机构';
    case 3:
      return '管理员';
    default:
      return '';
  }
}

const getTableData = () => {
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=GetCompanyList").then(res=>{
    console.log(res)
    tableData.value = res.data.data
  }).catch(err=>{
    console.log(err)
  })
}

const dialogVisible = ref(false)
const CompanyID = ref('')
const CarbonCoin = ref('')
const CarbonCredit = ref('')
const Type = ref(-1)
const companyid = ref('')

const pass = (row) => {
  dialogVisible.value = true
  CompanyID.value = row.CompanyID
  Type.value = row.Type
}

const determine = () => {
  if (Type.value == 0){
    // proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=APDistributeCoin",{
    //   companyID: row.CompanyID,
    //   carbonCoin: row.CarbonCoin.toString(),
    //   carbonCredit: row.CarbonCredit.toString()
    // }).then(res=>{
    //   row.Examine = 2
    //   ElMessage.success("审核通过")
    //   console.log(res)
    // })
    proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=APDistributeCoin",{
      companyID: CompanyID.value,
      carbonCoin: CarbonCoin.value,
      carbonCredit: CarbonCredit.value
    }).then(res=>{
      console.log(res)
      dialogVisible.value = false
      ElMessage.success("审核通过，分配成功")
      getTableData()
      CompanyID.value = ''
      CarbonCoin.value = ''
      CarbonCredit.value = ''
      Type.value = -1
    }).catch(err=>{
      ElMessage.error("该公司不存在")
      console.log(err)
    })
  } else {
    proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=APRegisiterForDAPAndTPRP",{
      CompanyID: CompanyID.value
    }).then(res=>{
      dialogVisible.value = false
      ElMessage.success("审核通过")
      getTableData()
      CompanyID.value = ''
      Type.value = -1
      console.log(res)
    })
  }
}

const nopass = () => {
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=APReviewFailed",{
    CompanyID: companyid.value
  }).then(res=>{
    ElMessage.success("审核不通过，操作成功")
    getTableData()
    companyid.value = ''
    console.log(res)
  })
}

const disableInput = computed(()=>{
  if (Type.value == 0){
    return false
  } else {
    return true
  }
})
</script>

<style scoped>

</style>