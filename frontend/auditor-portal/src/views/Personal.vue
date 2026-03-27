<template>
  <div>
    <div style="display: flex">
      <el-button icon="Back" size="large" @click="back" text/>
      <span style="margin-top: 3px;font-size: 20px">用户信息</span>
    </div>
    <div style="background-color: #FFFFFF;width: 1100px;height: 200px;margin-left: 40px;margin-top: 10px;border-radius: 10px">
        <div style="display: flex">
          <span style="font-weight: bold;font-size: 30px;display: inline-block;margin-top: 20px;margin-left: 40px">{{ID}}</span>
          <el-button type="primary" style="margin-top: 20px;margin-left: 790px">编辑</el-button>
        </div>
      <div style="display: flex;margin-top: 35px">
        <div style="margin-left: 30px">
          <span>公司名:</span>
          <span style="margin-left: 10px">{{name}}</span>
        </div>
        <div style="margin-left: 200px">
          <span>邮箱:</span>
          <span style="margin-left: 10px">{{email}}</span>
        </div>
        <div style="margin-left: 280px">
          <span>电话号码:</span>
          <span style="margin-left: 10px">{{phone}}</span>
        </div>
      </div>
    </div>
    <div style="background-color: #FFFFFF;width: 1100px;height: 200px;margin-left: 40px;margin-top: 10px;border-radius: 10px">
      <div style="display: flex">
        <span style="font-size: 18px;display: inline-block;margin-top: 20px;margin-left: 25px">碳余额</span>
      </div>
      <div style="margin-top: 20px;margin-left: 20px">
        <div style="margin-left: 30px;margin-top: 5px">
          <span>碳币余额:</span>
          <span style="margin-left: 10px">{{cb}}</span>
        </div>
        <div style="margin-left: 30px;margin-top: 25px">
          <span>碳排放量余额:</span>
          <span style="margin-left: 10px">{{cc}}</span>
        </div>
      </div>
    </div>
    <div style="background-color: #FFFFFF;width: 1100px;height: 300px;margin-left: 40px;margin-top: 10px;border-radius: 10px">
      <div style="display: flex">
        <span style="font-size: 18px;display: inline-block;margin-top: 20px;margin-left: 25px">跟进日志</span>
      </div>
      <el-button type="primary" style="margin-top: 10px;margin-left: 20px">添加日志</el-button>
      <el-table :data="tableData" style="width: 96%;margin-left: 15px">
        <el-table-column prop="date" label="创建时间" width="200" />
        <el-table-column prop="name" label="跟进人" width="200" />
        <el-table-column prop="address" label="跟进备注" />
      </el-table>
    </div>
  </div>
</template>
<script>
export default {
  data(){
    return{
      ID:'',
      cc:'',
      cb:'',
      email:'',
      phone:'',
      name:'',
      tableData:[
        {
          date: '2024-03-29',
          name: '员工2082',
          address: '碳核算 电气模块核算成功',
        },
        {
          date: '2024-03-26',
          name: '员工2082',
          address: '交易成功，获得100碳币',
        },
        {
          date: '2024-03-25',
          name: '员工2082',
          address: '发布交易，希望有人购买',
        },
        {
          date: '2024-03-19',
          name: '员工2082',
          address: '账户注册成功！',
        },
      ]
    }
  },
  created() {
    this.ID=this.$store.state.companyData.CompanyID;
    this.cc = this.$store.state.companyData.CarbonCredit;
    this.cb=this.$store.state.companyData.CarbonCoin;
    this.email=this.$store.state.companyData.Email;
    this.phone=this.$store.state.companyData.PhoneNumber;
    this.name=this.$store.state.companyData.CompanyName;
  },
  methods:{
      back(){
        this.$router.push('/DashBoard');
      }
  }
}
</script>
