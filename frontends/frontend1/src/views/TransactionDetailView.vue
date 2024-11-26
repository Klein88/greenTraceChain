
<template>
  <div style="width: 100% ;height: 850px;background: linear-gradient(to bottom , #FF88FF 30% ,white 40%) ;display: flex ; flex-direction: column ; align-items: center">
    <div style="display: flex ; width: 100%;flex-direction: row ; ;justify-content: left ; align-items: center;padding-top: 100px ; margin-left: 100px">
      <button style="width: 100px ; height: 40px ; background-color: white;border-radius: 15px" @click="this.$router.go(-1)">
        <div style="display: flex;flex-direction: row ; justify-content: center ; align-items: center ">
          <div style="width: 20px ; height: 40px ;display: flex ; justify-content: center;align-items: center">
            <img src="../views/Shop/hidariarrow20.svg">
          </div>
          <div style="width: 10px" />
          <div style="font-weight: bold ; font-size: 20px">
            返回
          </div>
        </div>
      </button>
    </div>
    <div style="width: 660px ; height: 400px ;display: flex ; flex-direction: column;justify-content: center;align-items: center">
      <div style="width: 100% ; display: flex ; flex-direction: row ; justify-content: left ; align-items: center">
        <div style="font-weight: bold;font-size: 30px ; color: white">
          交易信息 :
        </div>
      </div>
      <el-descriptions border :column="1" style="width: 570px ; margin-top: 50px">
        <el-descriptions-item label="出售方">{{ this.receivedData.Seller }}</el-descriptions-item>
        <el-descriptions-item label="购买方">{{ this.receivedData.Buyer }}</el-descriptions-item>
        <el-descriptions-item label="出售碳币金额">{{ this.receivedData.CarbonCoin }}</el-descriptions-item>
        <el-descriptions-item label="获得碳额度">{{ this.receivedData.CarbonCredit }}</el-descriptions-item>
        <el-descriptions-item label="交易时间">{{ this.currentTime }}</el-descriptions-item>
      </el-descriptions>
      <div style="display: flex;flex-direction: row ; justify-content: center ; align-items: center;margin-top: 30px">
        <button @click="Buy" style="width: 100px ; height: 50px ; background-color: springgreen ; border-radius: 5px">
          确定交易
        </button>
        <div style="width: 100px" />
        <button @click="this.$router.go(-1)" style="width: 100px ; height: 50px ; background-color: deeppink ; border-radius: 5px">
          取消交易
        </button>
      </div>
    </div>
  </div>
  <div style=" display: flex ; height: 300px;flex-direction: row ; justify-content: left ; align-items: center">
    <div style="width: 300px ; height: 300px">
        <img src="../views/Shop/TransactionDetailsvg.svg">
    </div>
    <div style="font-size: 30px ; font-weight: bolder">
       ----请确保交易的安全
    </div>
  </div>
</template>

<script>
import {ElMessage} from "element-plus";

export default {
  data() {
    return {
      receivedData: null,
      currentTime: ''
    };
  },
  created() {
    const queryString = this.$route.query.data; // 获取查询参数中的字符串
    if (queryString) {
      try {
        const parsedData = JSON.parse(decodeURIComponent(queryString)); // 将字符串解码并解析为JSON对象
        this.receivedData = parsedData; // 将解析后的JSON对象保存到data中
        console.log(this.receivedData)
      } catch (error) {
        console.error('解析查询字符串时出错:', error);
      }
    }
    this.getCurrentTime();
    this.intervalId = setInterval(this.getCurrentTime, 1000); // 每秒更新一次时间
  },
  beforeDestroy() {
    clearInterval(this.intervalId); // 组件销毁前清除定时器
  },
  methods: {
    getCurrentTime() {
      const now = new Date();
      this.currentTime = now.toLocaleString(); // 转换为本地时间字符串
    },
    Buy(){
        this.$http.post('http://47.97.176.174:8087/chainmaker?cmb=AToBTransaction',{
          tradeID:this.receivedData.TradeId,
          sellerID:this.receivedData.Seller,
          buyerID:this.receivedData.Buyer,
          carbonCredit:this.receivedData.CarbonCredit.toString(),
          carbonCoin:this.receivedData.CarbonCoin.toString(),
        }).then((res)=>{
          ElMessage({
            message: '购买成功',
            type: 'success',
          })
          this.$router.go(-1)
        })
    },
  }
}
</script>