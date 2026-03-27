<template>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>输配电损失</text>
    </template>
    <el-form :model="formData2" label-width="180px">
      <el-form-item label="电厂上网电量(兆瓦时)">
        <el-input v-model="formData2.ELSW" placeholder="请输入数据"></el-input>
      </el-form-item>
      <el-form-item label="自外省输入电量(兆瓦时)">
        <el-input v-model="formData2.ELSR" placeholder="请输入数据"></el-input>
      </el-form-item>
      <el-form-item label="向外省输出电量(兆瓦时)">
        <el-input v-model="formData2.ELSC" placeholder="请输入数据"></el-input>
      </el-form-item>
      <el-form-item label="售电量(兆瓦时)">
        <el-input v-model="formData2.ELSD" placeholder="请输入数据"></el-input>
      </el-form-item>
      <el-form-item label="输配电量(兆瓦时)">
        <el-input v-model="formData2.ELSPDL" placeholder="请输入数据"></el-input>
      </el-form-item>
      <el-form-item label="区域电网年平均供电排放因子(吨二氧化碳/兆瓦时)">
        <el-input v-model="formData2.EFDW" placeholder="请输入数据"></el-input>
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
      formData2:{},
    }
  },
  methods:{

  },
  watch: {
    formData2: {
      deep: true,
      handler(newFormData) {
        // 将 formData2 传递到全局变量
        this.$store.commit('updateFormData2', newFormData); // 使用正确的 mutation 名称
      }
    },
  }
}
</script>
