<template>
  <div style="border-radius: 10px;background-color: #FFFFFF;width: 1150px;margin-left: 20px;">
    <div style="display: flex;margin-top: 20px;display: inline-block">
      <el-button icon="Back" size="large" @click="back" text/>
      <span style="margin-top: 3px;font-size: 23px;font-weight: bold">交易信息</span>
    </div>
    <div style="width: 100% ;height: 2px ;background-color: #72767b;margin-top: 15px" />
    <div style="margin-left: 20px;margin-top: 20px">
      <span style="font-weight: bold;font-size: 20px">所属区块信息</span>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">区块高度:</span>
        <span style="margin-left: 10px;color: #409eff">{{block_height}}</span>
      </div>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">区块时间戳:</span>
        <span style="margin-left: 10px;color: #409eff">{{block_timestamp}}</span>
      </div>

    </div>

    <div style="margin-left: 20px;margin-top: 30px">
      <span style="font-weight: bold;font-size: 20px">交易信息</span>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">交易ID:</span>
        <span style="margin-left: 10px">{{txid}}</span>
      </div>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">交易类型:</span>
        <span style="margin-left: 10px">INVOKE_CONTRACT</span>
      </div>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">交易状态:</span>
        <span style="margin-left: 10px">SUCCESS</span>
      </div>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">交易发起时间:</span>
        <span style="margin-left: 10px">2024-03-28 23:21:20</span>
      </div>

    </div>
    <div style="margin-left: 20px;margin-top: 30px">
      <span style="font-weight: bold;font-size: 20px">合约执行信息</span>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">目标合约:</span>
        <span style="margin-left: 10px;color: #409eff">{{contract_name}}</span>
      </div>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">合约执行信息:</span>
        <span style="margin-left: 10px">Success</span>
      </div>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">合约调用方法:</span>
        <span style="margin-left: 10px">{{method}}</span>
      </div>
      <div style="display: flex;margin-left: 30px;margin-top: 10px">
        <span style="font-size: 14px">合约调用入参数</span>
        <el-table :data="tableData" stripe style="width: 88%;margin-left: 10px">
          <el-table-column prop="Id" label="#" width="120" />
          <el-table-column prop="key" label="Key" width="250" />
          <el-table-column prop="value" label="Value" width="595" />
        </el-table>
      </div>

    </div>
  </div>
</template>
<script>
export default {
  data(){
    return{
      txid:'',
      block_timestamp:'',
      block_height:'',
      contract_name:'',
      method:'',
      tableData:[]
    }
  },
  created() {
    this.txid=this.$route.query.txid;
    this.getdata()
  },
  methods:{
    getdata(){
      this.$http.post("http://47.97.176.174:8087/chainmaker?cmb=TxIdToSearchTx",{
        txid:this.txid
      }).then((res) => {
        this.block_timestamp=res.data.data.block_timestamp;
        this.block_height=res.data.data.block_height;
        this.contract_name=res.data.data.transaction.payload.contract_name;
        this.method=res.data.data.transaction.payload.method;
        const parameters = res.data.data.transaction.payload.parameters;
        // 为每个参数对象添加 Id，并解码 value
        parameters.forEach((param, index) => {
          param.Id = index + 1; // 添加 Id，从1开始
          param.value = atob(param.value); // 对 value 进行 Base64 解码
        });
        this.tableData=parameters
      }).catch((err) => {
        console.log(err);
      });
    },
    back(){
      this.$router.push('/Traceability');
    }
  }
}
</script>
