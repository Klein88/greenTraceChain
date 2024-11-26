<template>
  <div class="traceSource2">
    <div style="display: flex; flex-direction: row; margin-top: 30px">
      <el-button type="primary" icon="ArrowLeft" @click="back(data)" style="width: 120px; margin-left: 20px">返回</el-button>
      <el-button type="default" icon="DArrowLeft" @click="quit" style="width: 120px; margin-left: 20px">退出</el-button>
    </div>
    <el-card style="margin: 30px 20px 30px 20px" v-for="(item, index) in tracesourceData.length" :key="index">
      <div style="display: flex; justify-content: space-between; align-items: center; margin-left: 10px; margin-right: 30px">
        <span style="font-size: 17px; font-weight: bold">记录 {{ index+1 }}</span>
        <el-button type="primary" text style="margin-left: 20px" @click="traceSource(data,index)">溯源</el-button>
      </div>
      <el-descriptions border :column="2" style="margin-top: 15px">
        <el-descriptions-item label="区块高度" width="120">{{ block_heights[index] }}</el-descriptions-item>
        <el-descriptions-item label="区块时间戳" width="100">{{ block_timestamps[index] }}</el-descriptions-item>
        <el-descriptions-item label="区块链ID">{{ chain_ids[index] }}</el-descriptions-item>
        <el-descriptions-item label="交易发起组织">{{ org_ids[index] }}</el-descriptions-item>
        <el-descriptions-item label="合约名称">{{ contract_names[index] }}</el-descriptions-item>
        <el-descriptions-item label="合约调用方法">{{ methods[index] }}</el-descriptions-item>
        <el-descriptions-item label="合约调用参数" :span="2">{{ parameterses[index] }}</el-descriptions-item>
        <el-descriptions-item label="合约执行结果">{{ contract_result_messages[index] }}</el-descriptions-item>
        <el-descriptions-item label="消耗GAS量">{{ gas_useds[index] }}</el-descriptions-item>
        <el-descriptions-item label="交易ID" :span="2">{{ tx_ids[index] }}</el-descriptions-item>
      </el-descriptions>
    </el-card>
  </div>
</template>

<script setup>
import {getCurrentInstance, onMounted, reactive, ref} from "vue";
import {useRoute, useRouter} from "vue-router";

const route = useRoute()
const router = useRouter()
const currentInstance = getCurrentInstance()
const { proxy } = currentInstance

const data = ref({})
const Id = route.query.Id
const CompanyId = route.query.CompanyId
const tracesourceData = ref([])
const block_height = ref(-1)
const block_timestamp = ref(-1)
const chain_id = ref('')
const contract_name = ref('')
const method = ref('')
const parameters = ref([])
const tx_id = ref('')
const contract_result_message = ref('')
const gas_used = ref(0)
const org_id = ref('')

const block_heights = ref([])
const block_timestamps = ref([])
const chain_ids = ref([])
const contract_names = ref([])
const methods = ref([])
const parameterses = ref([])
const tx_ids = ref([])
const contract_result_messages = ref([])
const gas_useds = ref([])
const org_ids = ref([])

onMounted(async () => {
  data.value.Id = Id
  data.value.CompanyId = CompanyId

  await proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=SearchAllTransactionDataByCompanyId",{
    CompanyId: data.value.CompanyId
  }).then(res=>{
    console.log(res)
    tracesourceData.value = [...res.data.data]
  })

  data.value.tracesourceData = tracesourceData.value

  for (let i = 0; i < data.value.tracesourceData.length; i++){
    await proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=TxIdToSearchTx",{
      Txid: data.value.tracesourceData[i].Txid
    }).then(res=>{
      console.log(res)
      block_height.value = res.data.data.block_height
      block_timestamp.value = res.data.data.block_timestamp
      chain_id.value = res.data.data.transaction.payload.chain_id
      contract_name.value = res.data.data.transaction.payload.contract_name
      method.value = res.data.data.transaction.payload.method
      const datas = []
      for (let i = 0; i < res.data.data.transaction.payload.parameters.length; i++){
        datas.push(res.data.data.transaction.payload.parameters[i].key)
      }
      parameters.value = [...datas]
      tx_id.value = res.data.data.transaction.payload.tx_id
      contract_result_message.value = res.data.data.transaction.result.contract_result.message
      gas_used.value = res.data.data.transaction.result.contract_result.gas_used
      org_id.value = res.data.data.transaction.sender.signer.org_id

      block_heights.value.push(block_height.value)
      block_timestamps.value.push(block_timestamp.value)
      chain_ids.value.push(chain_id.value)
      contract_names.value.push(contract_name.value)
      methods.value.push(method.value)
      parameterses.value.push(parameters.value)
      tx_ids.value.push(tx_id.value)
      contract_result_messages.value.push(contract_result_message.value)
      gas_useds.value.push(gas_used.value)
      org_ids.value.push(org_id.value)
    })
  }

  data.value.block_heights = block_heights.value
  data.value.block_timestamps = block_timestamps.value
  data.value.chain_ids = chain_ids.value
  data.value.contract_names = contract_names.value
  data.value.methods = methods.value
  data.value.parameterses = parameterses.value
  data.value.tx_ids = tx_ids.value
  data.value.contract_result_messages = contract_result_messages.value
  data.value.gas_useds = gas_useds.value
  data.value.org_ids = org_ids.value

  // 在页面加载时从 sessionStorage 中获取数据并显示在页面上
  const savedData = getDataFromSessionStorage();
  if (savedData) {
    data.value = savedData;
  }
})

const back = (data) => {
  // 在页面跳转时保存数据到 sessionStorage 中
  sessionStorage.setItem('myData2', JSON.stringify(data));
  router.push('/traceSource1/'+data.CompanyId)
}

const quit = () => {
  sessionStorage.removeItem('myData1');
  sessionStorage.removeItem('myData2');
  sessionStorage.removeItem('myData3');
  sessionStorage.removeItem('myData4');
  router.push('/carbonAccountingData')
}

const traceSource = (data, index) => {
  // 在页面跳转时保存数据到 sessionStorage 中
  sessionStorage.setItem('myData2', JSON.stringify(data));
  router.push({
    path: '/traceSource3/'+data.CompanyId,
    query: {
      Id: data.Id,
      CompanyId: data.CompanyId,
      Txid: data.tx_ids[index]
    }
  })
}

// 在页面加载时从 sessionStorage 中获取数据
const getDataFromSessionStorage = () => {
  const data = sessionStorage.getItem('myData2');
  return data ? JSON.parse(data) : null;
}
</script>

<style scoped>
/* 添加带有透明度的背景颜色 */
.traceSource2::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url('@/assets/traceSource2.png');
  background-size: cover;
  z-index: -1;
  opacity: 0.3;
}
</style>