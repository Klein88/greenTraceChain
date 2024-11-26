<template>
  <el-card shadow="never" style="margin: 30px 15px 0 15px;">
    <div style="display: flex; align-items: center; margin-left: 30px; margin-top: 10px">
      <text>合约名称：</text>
      <el-input v-model="contractName" placeholder="请输入合约名称" style="width: 20%; height: 38px; margin-left: 5px" />
      <el-button type="primary" icon="Search" @click="search" style="width: 100px; height: 37px; font-size: 16px; margin-left: 20px">搜索</el-button>
      <el-button type="success" icon="List" @click="getTableData" style="width: 180px; height: 37px; font-size: 16px; margin-left: 80px">查询所有</el-button>
      <el-button type="danger" icon="FolderAdd" @click="add" style="width: 180px; height: 37px; font-size: 16px; margin-left: 80px">部署合约</el-button>
    </div>
    <el-table :data="tableData" border style="width: 100%; margin-top: 30px" height="550" :header-cell-style="{ background:'#f4f4f5' }">
      <el-table-column label="合约编号" prop="Id" width="85" align="center"></el-table-column>
      <el-table-column label="合约名称" prop="ContractName" width="150" align="center"></el-table-column>
      <el-table-column label="合约版本" prop="ContractVersion" width="90" align="center"></el-table-column>
      <el-table-column label="交易数量" prop="TxNum" width="90" align="center"></el-table-column>
      <el-table-column label="时间戳" prop="Timestamp" width="130" align="center"></el-table-column>
      <el-table-column label="地址" prop="Addr" width="370" align="center"></el-table-column>
      <el-table-column label="所属组织" prop="Sender" width="250" align="center"></el-table-column>
      <el-table-column label="操作" width="170">
        <template #default="scope">
          <el-button type="primary" text @click="freeze">冻结</el-button>
          <el-button type="primary" text @click="logout">注销</el-button>
          <el-button type="primary" text @click="upgrade">升级</el-button>
          <el-button type="primary" text @click="edit">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>

  <el-dialog v-model="visible1" title="部署合约">
    <el-form>
      <el-form-item label="合约名称">
        <el-input v-model="ContractName" style="margin-left: 30px"></el-input>
      </el-form-item>
      <el-form-item label="合约版本">
        <el-input v-model="ContractVersion" style="margin-left: 30px"></el-input>
      </el-form-item>
      <el-form-item label="虚拟机类型">
        <el-select v-model="VirtualMachineType" clearable placeholder="请选择虚拟机类型" style="margin-left: 15px">
          <el-option v-for="item in types" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="合约文件">
        <el-upload
            ref="upload"
            class="upload-demo"
            action="https://jsonplaceholder.typicode.com/posts/"
            :on-preview="handlePreview"
            :on-remove="handleRemove"
            :file-list="fileList"
            list-type="picture-card"
            :auto-upload="false"
            @change="handleFileChange"
            style="margin-left: 30px">
          <el-icon><Plus /></el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="部署理由">
        <el-input v-model="Reason" :rows="3" type="textarea" placeholder="部署合约需其他组织同意，请输入部署理由" style="margin-left: 30px"></el-input>
      </el-form-item>
      <el-form-item label="额外信息">
        <el-input v-model="key" style="width: 225px; margin-left: 30px" placeholder="Key"></el-input>
        <el-input v-model="value" style="width: 225px; margin-left: 20px" placeholder="Value"></el-input>
        <el-button type="primary" style="width: 80px; margin-left: 30px">增加</el-button>
      </el-form-item>
      <el-form-item label="合约调用方法">
        <el-input v-model="method" style="width: 150px; margin-left: 2px" placeholder="Method"></el-input>
        <el-select v-model="ff" clearable placeholder="请选择" style="width: 150px; margin-left: 10px">
          <el-option v-for="item in ffs" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
        <el-input v-model="param" style="width: 150px; margin-left: 10px" placeholder="Param"></el-input>
        <el-button type="primary" style="width: 80px; margin-left: 30px">增加</el-button>
      </el-form-item>
    </el-form>
    <div style="display: flex; justify-content: center; margin-top: 40px">
      <el-button type="primary" @click="determine1" style="width: 120px">确定</el-button>
      <el-button type="default" @click="cancle1" style="width: 120px; margin-left: 80px">取消</el-button>
    </div>
  </el-dialog>

  <el-dialog v-model="visible2" title="冻结合约">
    <p>冻结合约需经过其它组织投票同意，请输入冻结理由。</p>
    <el-form>
      <el-form-item label="理由">
        <el-input v-model="reason" :rows="5" type="textarea" placeholder="请输入冻结合约理由"></el-input>
      </el-form-item>
    </el-form>
    <div style="display: flex; justify-content: center; margin-top: 40px">
      <el-button type="primary" @click="determine2" style="width: 120px">确定</el-button>
      <el-button type="default" @click="cancle2" style="width: 120px; margin-left: 80px">取消</el-button>
    </div>
  </el-dialog>

  <el-dialog v-model="visible3" title="废止合约">
    <p>废止合约需经过其它组织投票同意，请输入废止理由。</p>
    <el-form>
      <el-form-item label="理由">
        <el-input v-model="reason2" :rows="5" type="textarea" placeholder="请输入废止合约理由"></el-input>
      </el-form-item>
    </el-form>
    <div style="display: flex; justify-content: center; margin-top: 40px">
      <el-button type="primary" @click="determine3" style="width: 120px">确定</el-button>
      <el-button type="default" @click="cancle3" style="width: 120px; margin-left: 80px">取消</el-button>
    </div>
  </el-dialog>

  <el-dialog v-model="visible4" title="升级合约">
    <el-form>
      <el-form-item label="合约名称">
        <el-input v-model="contractname" style="margin-left: 30px"></el-input>
      </el-form-item>
      <el-form-item label="升级版本号">
        <el-input v-model="contractversion" style="margin-left: 15px"></el-input>
      </el-form-item>
      <el-form-item label="虚拟机类型">
        <el-select v-model="virtualmachinetype" clearable placeholder="请选择虚拟机类型" style="margin-left: 15px">
          <el-option v-for="item in types" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="合约文件">
        <el-upload
            ref="upload"
            class="upload-demo"
            action="https://jsonplaceholder.typicode.com/posts/"
            :on-preview="handlePreview"
            :on-remove="handleRemove"
            :file-list="fileList"
            list-type="picture-card"
            :auto-upload="false"
            @change="handleFileChange"
            style="margin-left: 30px">
          <el-icon><Plus /></el-icon>
        </el-upload>
      </el-form-item>
      <el-form-item label="升级理由">
        <el-input v-model="reason3" :rows="3" type="textarea" placeholder="部署合约需其他组织同意，请输入部署理由" style="margin-left: 30px"></el-input>
      </el-form-item>
      <el-form-item label="额外信息">
        <el-input v-model="key" style="width: 225px; margin-left: 30px" placeholder="Key"></el-input>
        <el-input v-model="value" style="width: 225px; margin-left: 20px" placeholder="Value"></el-input>
        <el-button type="primary" style="width: 80px; margin-left: 30px">增加</el-button>
      </el-form-item>
      <el-form-item label="合约调用方法">
        <el-input v-model="method" style="width: 150px; margin-left: 2px" placeholder="Method"></el-input>
        <el-select v-model="ff" clearable placeholder="请选择" style="width: 150px; margin-left: 10px">
          <el-option v-for="item in ffs" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
        <el-input v-model="param" style="width: 150px; margin-left: 10px" placeholder="Param"></el-input>
        <el-button type="primary" style="width: 80px; margin-left: 30px">增加</el-button>
      </el-form-item>
    </el-form>
    <div style="display: flex; justify-content: center; margin-top: 40px">
      <el-button type="primary" @click="determine4" style="width: 120px">确定</el-button>
      <el-button type="default" @click="cancle4" style="width: 120px; margin-left: 80px">取消</el-button>
    </div>
  </el-dialog>
</template>

<script setup>
import {getCurrentInstance, onMounted, ref} from "vue";
import { ElMessage } from 'element-plus'

const currentInstance = getCurrentInstance()
const { proxy } = currentInstance

const tableData = ref()
const visible1 = ref(false)
const visible2 = ref(false)
const visible3 = ref(false)
const visible4 = ref(false)

const ContractName = ref('')
const ContractVersion = ref('')
const VirtualMachineType = ref('')
const types = [
  {
    label: 'GASM',
    value: 'GASM'
  },
  {
    label: 'DOCKER_GO',
    value: 'DOCKER_GO'
  },
  {
    label: 'WASMER',
    value: 'WASMER'
  },
  {
    label: 'WXVM',
    value: 'WXVM'
  },
  {
    label: 'EVM',
    value: 'EVM'
  }
]
const Reason = ref('')
const key = ref('')
const value = ref('')
const method = ref('')
const ff = ref('')
const ffs = [
  {
    label: '调用',
    value: '调用'
  },
  {
    label: '查询',
    value: '查询'
  }
]
const param = ref('')

const reason = ref('')
const reason2 = ref('')

const contractname = ref('')
const contractversion = ref('')
const virtualmachinetype = ref('')
const reason3 = ref('')

onMounted(() => {
  getTableData()
})

const getTableData = () => {
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=GetContractList",{
    chainId: 'Chain01',
    contractName: '',
    pageNum: 0,
    pageSize: 100,
  }).then(res=>{
    tableData.value = res.data.data
    console.log(res)
  })
}

const add = () => {
  visible1.value = true
}

const determine1 = () => {
  visible1.value = false
  ElMessage.success("添加成功")
}

const cancle1 = () => {
  visible1.value = false
  ElMessage.info("已取消")
}

const freeze = () => {
  visible2.value = true
}

const determine2 = () => {
  visible2.value = false
  ElMessage.success("冻结成功")
}

const cancle2 = () => {
  visible2.value = false
  ElMessage.info("已取消")
}

const logout = () => {
  visible3.value = true
}

const determine3 = () => {
  visible3.value = false
  ElMessage.success("注销成功")
}

const cancle3 = () => {
  visible3.value = false
  ElMessage.info("已取消")
}

const upgrade = () => {
  visible4.value = true
}

const determine4 = () => {
  visible4.value = false
  ElMessage.success("升级成功")
}

const cancle4 = () => {
  visible4.value = false
  ElMessage.info("已取消")
}
</script>

<style scoped>

</style>