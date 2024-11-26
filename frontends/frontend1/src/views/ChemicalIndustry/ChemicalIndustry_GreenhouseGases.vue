<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表一 报告主体20__年温室气体排放量汇总</text>
    </template>
    <el-form :model="greenhouse" label-width="180px">
      <el-text style="margin-left: 80px">源类别</el-text>
        <el-text style="margin-left: 120px">温室气体本身质量(单位：吨)</el-text>
      <el-text style="margin-left: 120px">CO2当量（单位：吨CO2当量）</el-text>
      <el-divider></el-divider>
      <el-form-item label="化石燃料燃烧CO2排放">
        <div style="display: flex">
          <el-input v-model="greenhouse[0].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px"></el-input>
          <el-input v-model="greenhouse[0].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
      <el-form-item label="工业生产过程CO2排放">
        <div style="display: flex">
          <el-input v-model="greenhouse[1].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px"></el-input>
          <el-input v-model="greenhouse[1].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
      <el-form-item label="工业生产过程N2O排放">
        <div style="display: flex">
          <el-input v-model="greenhouse[2].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px"></el-input>
          <el-input v-model="greenhouse[2].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
      <el-form-item label="CO2回收利用量">
        <div style="display: flex">
          <el-input v-model="greenhouse[3].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px"></el-input>
          <el-input v-model="greenhouse[3].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
      <el-form-item label="企业净购入的电力和热力消费引起的CO2排放" style="height: 50px">
        <div style="display: flex">
          <el-input v-model="greenhouse[4].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px"></el-input>
          <el-input v-model="greenhouse[4].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
      <el-form-item label="企业温室气体排放总量（吨CO2当量）">
        <div style="display: flex">
          <el-input v-model="greenhouse[5].WS" placeholder="" style="width:300px;margin-left: 10px" disabled></el-input>
          <el-input v-model="greenhouse[5].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px"></el-input>
        </div>
      </el-form-item>
      <el-form-item>
      </el-form-item>
    </el-form>
  </el-card>

</template>
<script>
export default {
  data(){
    return{
      greenhouse:[
        {
          Breed: "化石燃料CO2",
          WS: null,
          CO: null
        },
        {
          Breed: "工业生产过程CO2",
          WS: null,
          CO: null
        },
        {
          Breed: "工业生产过程N2O",
          WS: null,
          CO: null
        },
        {
          Breed: "CO2回收利用量",
          WS: null,
          CO: null
        },
        {
          Breed: "企业净购入的电力和热力消费引起的CO2排放",
          WS: null,
          CO: null
        },
        {
          Breed: "企业温室气体排放总量",
          WS:0,
          CO: null
        },

      ],
    }
  },
  methods:{

  },
  watch: {
    greenhouse: {
      deep: true,
      handler(newgreenhouse) {
        this.$store.commit('updategreenhouse', newgreenhouse);
      }
    },
  }
}
</script>
