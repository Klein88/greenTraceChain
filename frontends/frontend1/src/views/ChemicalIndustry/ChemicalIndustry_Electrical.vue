<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表七 净购入的电力和热力消费活动水平和排放因子数据一览表</text>
    </template>
    <el-form :model="Electrical" label-width="180px">
      <div>
        <el-text style="margin-left: 40px">类型</el-text>
        <el-text style="margin-left: 80px">净购入量</el-text>
        <el-text style="margin-left: 120px">购入量</el-text>
        <el-text style="margin-left: 120px">外供量</el-text>
        <el-text style="margin-left: 120px">CO2排放因子</el-text>
      </div>
      <div>
        <el-text style="margin-left: 140px">(MWh或GJ)</el-text>
        <el-text style="margin-left: 90px">(MWh或GJ)</el-text>
        <el-text style="margin-left: 90px">(MWh或GJ)</el-text>
        <el-text style="margin-left: 70px">(tCO2/MWh或tCO2/GJ)</el-text>
      </div>
      <el-divider></el-divider>
      <el-form-item label="电力" style="margin-left: -100px">
        <div style="display: flex">
          <el-input v-model="Electrical[0].AD" placeholder="请输入数据" style="width: 160px;margin-left: 10px"></el-input>
          <el-input v-model="Electrical[0].AD1" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
          <el-input v-model="Electrical[0].AD2" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
          <el-input v-model="Electrical[0].EF" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
      <el-form-item label="蒸汽" style="margin-left: -100px">
        <div style="display: flex">
          <el-input v-model="Electrical[1].AD" placeholder="请输入数据" style="width: 160px;margin-left: 10px"></el-input>
          <el-input v-model="Electrical[1].AD1" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
          <el-input v-model="Electrical[1].AD2" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
          <el-input v-model="Electrical[1].EF" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
      <el-form-item label="热水" style="margin-left: -100px">
        <div style="display: flex">
          <el-input v-model="Electrical[2].AD" placeholder="请输入数据" style="width: 160px;margin-left: 10px"></el-input>
          <el-input v-model="Electrical[2].AD1" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
          <el-input v-model="Electrical[2].AD2" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
          <el-input v-model="Electrical[2].EF" placeholder="请输入数据" style="width: 160px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
    </el-form>
  </el-card>

</template>
<script>
export default {
  data(){
    return{
      Electrical:[
        {
          Breed: "电力",
          AD: null,
          AD1: null,
          AD2:null,
          EF:null
        },
        {
          Breed: "蒸汽",
          AD: null,
          AD1: null,
          AD2:null,
          EF:null
        },
        {
          Breed: "热水",
          AD: null,
          AD1: null,
          AD2:null,
          EF:null
        },

      ],
    }
  },
  methods:{

  },
  watch: {
    Electrical: {
      deep: true,
      handler(newElectrical) {
        this.$store.commit('updateElectrical', newElectrical); // 使用正确的 mutation 名称
      }
    },
  }
}
</script>
