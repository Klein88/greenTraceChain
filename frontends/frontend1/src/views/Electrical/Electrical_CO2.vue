<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表1 报告主体()年二氧化碳排放量报告</text>
    </template>
    <el-form :model="formData" label-width="340px">
      <el-form-item label="企业二氧化碳排放总量(tCO2)">
        <el-input v-model="formData.text1" placeholder="请输入数据"></el-input>
      </el-form-item>
      <el-form-item label="使用六氟化硫设备修理与退役过程产生的排放(tCO2)">
        <el-input v-model="formData.text2" placeholder="请输入数据"></el-input>
      </el-form-item>
      <el-form-item label="输配电引起的二氧化碳排放(tCO2)">
        <el-input v-model="formData.text3" placeholder="请输入数据"></el-input>
      </el-form-item>
    </el-form>
  </el-card>

</template>
<script>
export default {
  data(){
    return{
      formData:{},
    }
  },
  methods:{
  },
  watch: {
    formData: {
      deep: true,
      handler(newFormData) {
        // 将 formData 传递到全局变量
        this.$store.commit('updateFormData', newFormData);
      }
    },
  }
}
</script>
