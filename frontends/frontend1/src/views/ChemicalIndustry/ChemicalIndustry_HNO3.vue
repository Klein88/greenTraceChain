<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表五 硝酸生产过程的活动水平和N2O排放因子数据一览表</text>
    </template>
    <el-form :model="HNO3" label-width="180px">
      <div style="display: flex; align-items: center;">
        <el-text style="margin-left: 0px">硝酸生产工艺类型</el-text>
        <el-text style="margin-left: 20px">硝酸产量</el-text>
        <el-text style="margin-left: 20px">N2O生成因子</el-text>
        <el-text style="margin-left: 25px">数据来源</el-text>
        <el-text style="margin-left: 25px">N2O去除率</el-text>
        <el-text style="margin-left: 25px">数据来源</el-text>
        <el-text style="margin-left: 25px">尾气处理设备使用率</el-text>
        <el-text style="margin-left: 25px">数据来源</el-text>
      </div>
      <div>
        <el-text style="margin-left: 148px">(吨)</el-text>
        <el-text style="margin-left: 25px">(kgN2O/吨硝酸)</el-text>
        <el-text style="margin-left: 125px">(%)</el-text>
        <el-text style="margin-left: 185px">(%)</el-text>
      </div>
      <el-divider></el-divider>
      <el-form-item label="高压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[0].AD" placeholder="输入数据" style="width:85px;margin-left: 30px"></el-input>
          <el-input v-model="HNO3[0].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px"></el-input>
          <el-radio-group v-model="HNO3[0].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[0].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"></el-input>
          <el-radio-group v-model="HNO3[0].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[0].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px"></el-input>
          <el-radio-group v-model="HNO3[0].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="中压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[1].AD" placeholder="输入数据" style="width:85px;margin-left: 30px"></el-input>
          <el-input v-model="HNO3[1].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px"></el-input>
          <el-radio-group v-model="HNO3[1].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[1].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"></el-input>
          <el-radio-group v-model="HNO3[1].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[1].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px"></el-input>
          <el-radio-group v-model="HNO3[1].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="常压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[2].AD" placeholder="输入数据" style="width:85px;margin-left: 30px"></el-input>
          <el-input v-model="HNO3[2].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px"></el-input>
          <el-radio-group v-model="HNO3[2].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[2].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"></el-input>
          <el-radio-group v-model="HNO3[2].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[2].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px"></el-input>
          <el-radio-group v-model="HNO3[2].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="双加压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[3].AD" placeholder="输入数据" style="width:85px;margin-left: 30px"></el-input>
          <el-input v-model="HNO3[3].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px"></el-input>
          <el-radio-group v-model="HNO3[3].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[3].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"></el-input>
          <el-radio-group v-model="HNO3[3].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[3].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px"></el-input>
          <el-radio-group v-model="HNO3[3].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="综合法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[4].AD" placeholder="输入数据" style="width:85px;margin-left: 30px"></el-input>
          <el-input v-model="HNO3[4].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px"></el-input>
          <el-radio-group v-model="HNO3[4].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[4].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"></el-input>
          <el-radio-group v-model="HNO3[4].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[4].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px"></el-input>
          <el-radio-group v-model="HNO3[4].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="低压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[5].AD" placeholder="输入数据" style="width:85px;margin-left: 30px"></el-input>
          <el-input v-model="HNO3[5].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px"></el-input>
          <el-radio-group v-model="HNO3[5].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[5].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"></el-input>
          <el-radio-group v-model="HNO3[5].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[5].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px"></el-input>
          <el-radio-group v-model="HNO3[5].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" >
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
      HNO3:[
        {
          Breed: '高压法',
          AD: null,
          EF: null,
          EFDS: 0,
          NK: null,
          NKDS: 0,
          UK: null,
          UKDS: 0
        },
        {
          Breed: '中压法',
          AD: null,
          EF: null,
          EFDS: 0,
          NK: null,
          NKDS: 0,
          UK: null,
          UKDS: 0
        },
        {
          Breed: '常压法',
          AD: null,
          EF: null,
          EFDS: 0,
          NK: null,
          NKDS: 0,
          UK: null,
          UKDS: 0
        },
        {
          Breed: '双加压法',
          AD: null,
          EF: null,
          EFDS: 0,
          NK: null,
          NKDS: 0,
          UK: null,
          UKDS: 0
        },
        {
          Breed: '综合法',
          AD: null,
          EF: null,
          EFDS: 0,
          NK: null,
          NKDS: 0,
          UK: null,
          UKDS: 0
        },
        {
          Breed: '低压法',
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
    HNO3: {
      deep: true,
      handler(newHNO3) {
        this.$store.commit('updateHNO3', newHNO3);
      }
    },
  }
}
</script>

