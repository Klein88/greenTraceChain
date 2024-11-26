<template>
  <div class="traceSource4">
    <div style="display: flex; flex-direction: row; margin-top: 30px">
      <el-button type="primary" icon="ArrowLeft" @click="back(data)" style="width: 120px; margin-left: 20px">返回</el-button>
      <el-button type="default" icon="DArrowLeft" @click="quit" style="width: 120px; margin-left: 20px">退出</el-button>
    </div>
    <el-tabs type="border-card" style="margin: 30px 20px 30px 20px; height: 900px">
      <el-tab-pane label="附表1：报告主体 2024 年二氧化碳排放量报告">
        <el-descriptions :column="1" border style="margin: 20px 10px 20px 10px">
          <el-descriptions-item label="使用六氟化硫设备修理与退役过程产生的排放(tCO2)" width="200">{{ currentCompanyData.FSCO }}</el-descriptions-item>
          <el-descriptions-item label="输配电引起的二氧化碳排放(tCO2)">{{ currentCompanyData.SPDCO }}</el-descriptions-item>
          <el-descriptions-item label="企业二氧化碳排放总量(tCO2)">{{ currentCompanyData.Total }}</el-descriptions-item>
        </el-descriptions>
      </el-tab-pane>

      <el-tab-pane label="附表2：报告主体活动水平数据">
        <p style="text-align: left; margin-left: 20px; font-size: 18px; font-weight: bold">六氟化硫回收</p>
        <el-table :data="PGERepairREC" border :header-cell-style="{ background:'#f4f4f5' }" style="width: 96%; margin-left: 15px; margin-top: 20px">
          <el-table-column label="修理设备" prop="Id"></el-table-column>
          <el-table-column label="设备容量(千克)" prop="RECV"></el-table-column>
          <el-table-column label="实际回收量(千克)" prop="RECR"></el-table-column>
        </el-table>
        <el-table :data="PGERetireREC" border :header-cell-style="{ background:'#f4f4f5' }" style="width: 96%; margin-left: 15px; margin-top: 20px">
          <el-table-column label="退役设备" prop="Id"></el-table-column>
          <el-table-column label="设备容量(千克)" prop="RECV"></el-table-column>
          <el-table-column label="实际回收量(千克)" prop="RECR"></el-table-column>
        </el-table>
        <p style="text-align: left; margin-top: 50px; margin-left: 20px; font-size: 18px; font-weight: bold">输配电损失</p>
        <el-descriptions :column="1" border style="margin: 20px 10px 20px 10px">
          <el-descriptions-item label="电厂上网电量(兆瓦时)" width="200">{{ ELSW }}</el-descriptions-item>
          <el-descriptions-item label="自外省输入电量(兆瓦时)">{{ ELSR }}</el-descriptions-item>
          <el-descriptions-item label="向外省输出电量(兆瓦时)">{{ ELSC }}</el-descriptions-item>
          <el-descriptions-item label="售电量(兆瓦时)">{{ ELSD }}</el-descriptions-item>
          <el-descriptions-item label="输配电量(兆瓦时)">{{ ELSPDL }}</el-descriptions-item>
        </el-descriptions>
      </el-tab-pane>

      <el-tab-pane label="附表3：报告主体排放因子">
        <el-descriptions :column="3" border direction="vertical" style="margin: 20px 10px 20px 10px">
          <el-descriptions-item label="名称" width="200">电力</el-descriptions-item>
          <el-descriptions-item label="排放因子单位">吨二氧化碳/兆瓦时</el-descriptions-item>
          <el-descriptions-item label="二氧化碳排放因子">{{ EFDW }}</el-descriptions-item>
        </el-descriptions>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>
import {useRoute, useRouter} from "vue-router";
import {getCurrentInstance, onMounted, ref} from "vue";

const route = useRoute()
const router = useRouter()
const currentInstance = getCurrentInstance()
const { proxy } = currentInstance

const Id = route.query.Id
const CompanyId = route.query.CompanyId
const data = ref({})

const currentCompanyData = ref({});
const PGERepairREC = ref([])
const PGERetireREC = ref([])
const ELSW = ref(0)
const ELSR = ref(0)
const ELSC = ref(0)
const ELSD = ref(0)
const ELSPDL = ref(0)

const EFDW = ref(0)

onMounted( async ()=>{
  data.value.Id = Id
  console.log(Id)
  data.value.CompanyId = CompanyId

  await proxy.$axios.post('http://47.97.176.174:8087/chainmaker?cmb=SearchPGEById', {
    Id : parseInt(data.value.Id)
  }).then(res=>{
    console.log(res)
    currentCompanyData.value = res.data.data

    PGERepairREC.value = [...res.data.data.Data.PGERepairREC]
    PGERetireREC.value = [...res.data.data.Data.PGERetireREC]

    ELSW.value = res.data.data.Data.ELSW
    ELSR.value = res.data.data.Data.ELSR
    ELSC.value = res.data.data.Data.ELSC
    ELSD.value = res.data.data.Data.ELSD
    ELSPDL.value = res.data.data.Data.ELSPDL
    EFDW.value = res.data.data.Data.EFDW

    data.value.currentCompanyData = currentCompanyData
    data.value.PGERepairREC = PGERepairREC
    data.value.PGERetireREC = PGERetireREC
    data.value.ELSW = ELSW
    data.value.ELSR = ELSR
    data.value.ELSC = ELSC
    data.value.ELSD = ELSD
    data.value.ELSPDL = ELSPDL
    data.value.EFDW = EFDW
  }).catch(function (error) {
    console.log(error);
  })

  // 在页面加载时从 sessionStorage 中获取数据并显示在页面上
  const savedData = getDataFromSessionStorage();
  if (savedData) {
    data.value = savedData;
  }
})

const back = (data) => {
  console.log(data)
  // 在页面跳转时保存数据到 sessionStorage 中
  sessionStorage.setItem('myData4', JSON.stringify(data));
  router.push('/traceSource3/'+data.CompanyId)
}

const quit = () => {
  sessionStorage.removeItem('myData1');
  sessionStorage.removeItem('myData2');
  sessionStorage.removeItem('myData3');
  sessionStorage.removeItem('myData4');
  router.push('/carbonAccountingData')
}

// 在页面加载时从 sessionStorage 中获取数据
const getDataFromSessionStorage = () => {
  const data = sessionStorage.getItem('myData4');
  return data ? JSON.parse(data) : null;
}
</script>

<style scoped>
/* 添加带有透明度的背景颜色 */
.traceSource4::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 950px;
  background-image: url('@/assets/traceSource4.png');
  background-size: cover;
  z-index: -1;
  opacity: 0.3;
}
</style>