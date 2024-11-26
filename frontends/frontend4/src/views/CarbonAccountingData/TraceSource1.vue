<template>
  <div class="traceSource1" style="display: flex; flex-direction: column">
    <div style="display: flex; justify-content: center; align-items: center; margin: 40px 20px 0 20px; padding: 20px 0 20px 0; background: white; border-radius: 3px">
      <el-button type="primary" icon="ArrowLeftBold" @click="back" style="width: 120px; margin-left: 30px">返回</el-button>
      <el-steps :active="6" style="width: 90%" align-center>
        <el-step title="评估中" icon="Document"></el-step>
        <el-step title="待发布" icon="DocumentCopy"></el-step>
        <el-step title="发布中" icon="Upload"></el-step>
        <el-step title="签约中" icon="Memo"></el-step>
        <el-step title="认证中" icon="DocumentChecked"></el-step>
        <el-step title="已完成" icon="Finished"></el-step>
      </el-steps>
    </div>
    <div style="display: flex; justify-content: space-around; align-items: center">
      <el-card style="width: 250px; height: 420px; margin-top: 18px">
        <div style="display: flex; flex-direction: column; justify-content: center; align-items: center; margin-top: 20px">
          <span style="font-size: 20px; font-weight: bold">公司详情</span>
          <span style="margin-top: 15px">状态：{{ data.State }}</span>
          <!--          <el-button type="primary" text style="margin-left: 200px">评估结果</el-button>-->
        </div>
        <div style="display: flex; flex-direction: column; justify-content: center; height: 250px">
          <el-card style="height: 250px; margin-top: 50px; display: flex; flex-direction: column; justify-content: center">
            <el-descriptions :column="1">
              <el-descriptions-item label="公司编号">{{ data.CompanyId }}</el-descriptions-item>
              <el-descriptions-item label="企业类型">{{ data.Type }}</el-descriptions-item>
              <el-descriptions-item label="创建时间">{{ data.createAt }}</el-descriptions-item>
              <el-descriptions-item label="更新时间">{{ data.updateAt }}</el-descriptions-item>
            </el-descriptions>
          </el-card>
        </div>
      </el-card>

      <el-card style="width: 750px; margin-top: 50px; margin-bottom: 35px; margin-right: 30px">
        <div style="display: flex; align-items: center; justify-content: space-between">
          <span style="font-size: 20px; font-weight: bold">发布详情</span>
          <el-button type="primary" text @click="more(data)">
            更多
            <el-icon class="el-icon--right"><ArrowRightBold /></el-icon>
          </el-button>
        </div>
        <el-card style="margin-top: 20px">
          <div style="display: flex; justify-content: space-between; align-items: center; margin-left: 10px; margin-right: 20px">
            <span style="font-size: 17px; font-weight: bold">最新记录</span>
            <el-button type="primary" text style="margin-left: 20px" @click="traceSource(data)">溯源</el-button>
          </div>
          <el-descriptions border :column="2" style="margin-top: 15px">
            <el-descriptions-item width="120">
              <template #label>
                <span class="custom-label">区块高度</span>
              </template>
              {{ data.block_height }}
            </el-descriptions-item>
            <el-descriptions-item width="120">
              <template #label>
                <span class="custom-label">区块时间戳</span>
              </template>
              {{ data.block_timestamp }}
            </el-descriptions-item>
            <el-descriptions-item label="区块链ID">{{ data.chain_id }}</el-descriptions-item>
            <el-descriptions-item label="交易发起组织">{{ data.org_id }}</el-descriptions-item>
            <el-descriptions-item label="合约名称">{{ data.contract_name }}</el-descriptions-item>
            <el-descriptions-item label="合约调用方法">{{ data.method }}</el-descriptions-item>
            <el-descriptions-item label="合约调用参数" :span="2">{{ data.parameters }}</el-descriptions-item>
            <el-descriptions-item label="合约执行结果">{{ data.contract_result_message }}</el-descriptions-item>
            <el-descriptions-item label="消耗GAS量">{{ data.gas_used }}</el-descriptions-item>
            <el-descriptions-item label="交易ID" :span="2">{{ data.tx_id }}</el-descriptions-item>
          </el-descriptions>
        </el-card>
      </el-card>
    </div>
  </div>
</template>

<script setup>
import {useRoute, useRouter} from "vue-router";
import {getCurrentInstance, onMounted, ref} from "vue";

const router = useRouter()
const route = useRoute()
const currentInstance = getCurrentInstance()
const { proxy } = currentInstance

const Id = route.query.Id
let State = ref('')
const CompanyId = route.query.CompanyId
let Type = ref('')
const createAt = route.query.createAt
const updateAt = route.query.updateAt
let data = ref({})

const tracesourceData = ref([])
const Txid = ref('')
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

const stateTagText = (value) => {
  switch (value) {
    case '0':
      return '未审核'
    case '1':
      return '已审核通过'
    case '2':
      return '未审核通过'
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

onMounted(async ()=>{
  State.value = stateTagText(route.query.State)
  Type.value = typeTagText(route.query.Type)

  data.value.Id = Id
  data.value.State = State.value
  data.value.CompanyId = CompanyId
  data.value.Type = Type.value
  data.value.createAt = createAt
  data.value.updateAt = updateAt

  await proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=SearchAllTransactionDataByCompanyId",{
    CompanyId: data.value.CompanyId
  }).then(res=>{
    console.log(res)
    tracesourceData.value = [...res.data.data]
  })

  data.value.tracesourceData = tracesourceData.value

  await proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=TxIdToSearchTx",{
    Txid: data.value.tracesourceData[data.value.tracesourceData.length-1].Txid
  }).then(res=>{
    console.log(res)
    Txid.value = data.value.tracesourceData[data.value.tracesourceData.length-1].Txid
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
  })

  data.value.Txid = Txid.value
  data.value.block_height = block_height.value
  data.value.block_timestamp = block_timestamp.value
  data.value.chain_id = chain_id.value
  data.value.contract_name = contract_name.value
  data.value.method = method.value
  data.value.parameters = parameters.value
  data.value.tx_id = tx_id.value
  data.value.contract_result_message = contract_result_message.value
  data.value.gas_used = gas_used.value
  data.value.org_id = org_id.value

  // 在页面加载时从 sessionStorage 中获取数据并显示在页面上
  const savedData = getDataFromSessionStorage();
  if (savedData) {
    data.value = savedData;
  }
})

const back = () => {
  sessionStorage.removeItem('myData1');
  sessionStorage.removeItem('myData2');
  sessionStorage.removeItem('myData3');
  sessionStorage.removeItem('myData4');
  router.push('/carbonAccountingData')
}

const traceSource = (data) => {
  // 在页面跳转时保存数据到 sessionStorage 中
  sessionStorage.setItem('myData1', JSON.stringify(data));
  console.log(data.CompanyId)
  router.push({
    path: '/traceSource3/'+data.CompanyId,
    query: {
      Id: data.Id,
      CompanyId: data.CompanyId,
      Txid: data.Txid
    }
  })
}

const more = (data) => {
  // 在页面跳转时保存数据到 sessionStorage 中
  sessionStorage.setItem('myData1', JSON.stringify(data));
  router.push({
    path: '/traceSource2/'+data.CompanyId,
    query: {
      Id: data.Id,
      CompanyId: data.CompanyId
    }
  })
}

// 在页面加载时从 sessionStorage 中获取数据
const getDataFromSessionStorage = () => {
  const data = sessionStorage.getItem('myData1');
  return data ? JSON.parse(data) : null;
}
</script>

<style scoped>
/* 添加带有透明度的背景颜色 */
.traceSource1::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 830px;
  background-image: url('@/assets/traceSource1.png');
  background-size: cover;
  z-index: -1;
  opacity: 0.3;
}

.custom-label {
  width: 100px; /* 设置 label 的宽度为 120px */
  display: inline-block; /* 使宽度生效 */
}
</style>