<template>
  <span style="font-weight: bold;font-size: 30px;margin-left: 20px">溯源历史</span>
  <div class="card-container">
    <el-card v-for="(item, index) in jsonData" :key="index" class="card" shadow="hover" style="border-radius: 10px;width: 40%">
      <div style="display: flex">
        <span>交易编号:</span>
        <span>{{item.Id}}</span>
      </div>
      <div style="display: flex">
        <span>交易后碳币余额：</span>
        <span>{{item.CarbonCoin}}</span>
      </div>
      <div style="display: flex">
        <span>交易后排放量剩余：</span>
        <span>{{item.CarbonCredit}}</span>
      </div>
      <div>
        <span>交易编号：</span>
        <span>{{item.Txid}}</span>
      </div>
      <el-button @click="viewDetails(item.Txid)" style="margin-left: 350px" type="primary">详情</el-button>
    </el-card>
  </div>
</template>

<script>

import axios from "axios";

export default {
  data() {
    return {
      jsonData: [],
      CompanyID:'',
    };
  },
  mounted() {

  },
  created() {
    this.CompanyID = this.$route.query.CompanyID;
    this.getcompanydata()
  },
  methods:{
    getcompanydata(){
      axios.post("http://47.97.176.174:8087/chainmaker?cmb=GetCompanyTransactionByCompanyId",{
        CompanyId:this.CompanyID
      }).then((res) => {
        let number=1;
        res.data.data.forEach(item => {
          item.Id = number
          number++;
        });
        this.jsonData = res.data.data.reverse();
      }).catch((err) => {
        console.log(err);
      });
    },
    viewDetails(id){
      this.$router.push({ path: 'Traceabilitydetail', query: { txid: id } });
    }
  }
};
</script>

<style>
.card-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between; /* 水平方向上尽可能地平均分配空间 */
  margin-top: 10px;
  margin-left: 50px;
}

.card {
  margin-bottom: 20px;
  margin-left: 10px; /* 调整卡片之间的左边距 */
  margin-right: 100px; /* 调整卡片之间的右边距 */
  word-wrap: break-word;
}
</style>
