<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表六 己二酸生产过程的活动水平和N2O排放因子数据一览表</text>
    </template>
    <el-form :model="HOO" label-width="180px">
      <div style="display: flex; align-items: center;">
        <el-text style="margin-left: 0px">己二酸生产工艺类型</el-text>
        <el-text style="margin-left: 15px">己二酸产量</el-text>
        <el-text style="margin-left: 15px">N2O生成因子</el-text>
        <el-text style="margin-left: 25px">数据来源</el-text>
        <el-text style="margin-left: 25px">N2O去除率</el-text>
        <el-text style="margin-left: 25px">数据来源</el-text>
        <el-text style="margin-left: 25px">尾气处理设备使用率</el-text>
        <el-text style="margin-left: 25px">数据来源</el-text>
      </div>
      <div>
        <el-text style="margin-left: 161px">(吨)</el-text>
        <el-text style="margin-left: 32px">(kgN2O/吨硝酸)</el-text>
        <el-text style="margin-left: 125px">(%)</el-text>
        <el-text style="margin-left: 181px">(%)</el-text>
      </div>
      <el-divider></el-divider>
      <el-form-item label="硝酸氧化" style="margin-left: -88px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HOO[0].AD" placeholder="输入数据" style="width:85px;margin-left: 30px"></el-input>
          <el-input v-model="HOO[0].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px"></el-input>
          <el-radio-group v-model="HOO[0].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HOO[0].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"></el-input>
          <el-radio-group v-model="HOO[0].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HOO[0].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px"></el-input>
          <el-radio-group v-model="HOO[0].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="其他" style="margin-left: -88px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HOO[1].AD" placeholder="输入数据" style="width:85px;margin-left: 30px"></el-input>
          <el-input v-model="HOO[1].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px"></el-input>
          <el-radio-group v-model="HOO[1].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HOO[1].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"></el-input>
          <el-radio-group v-model="HOO[1].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HOO[1].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px"></el-input>
          <el-radio-group v-model="HOO[1].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
export default {
  data(){
    return{
      HOO:[
        {
          Breed: '硝酸氧化',
          AD: null,
          EF: null,
          EFDS: 0,
          NK: null,
          NKDS: 0,
          UK: null,
          UKDS: 0
        },
        {
          Breed: '其他',
          AD: null,
          EF: null,
          EFDS: 0,
          NK: null,
          NKDS: 0,
          UK: null,
          UKDS: 0
        },
      ],
    }
  },
  methods:{
  },
  watch: {
    HOO: {
      deep: true,
      handler(newHOO) {
        this.$store.commit('updateHOO', newHOO);
      }
    },
  }
}
</script>

