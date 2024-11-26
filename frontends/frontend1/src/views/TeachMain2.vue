<template>
  <el-button type="primary" @click="back()" style="margin-left: 1100px">返回</el-button>
  <el-card shadow="never" style="margin-top: 10px">
    <template #header>
      <text style="font-size: 35px;font-weight: bold;margin-left: 20px">科普动画丨碳排放是如何进行核算的，企业的碳排放又是怎么计算的呢？</text>
      <div style="display: flex;margin-top: 5px;margin-left: 5px">
        <span>国家大气污染防治攻关联合中心</span>
        <span style="color: #999;margin-left: 30px">2022-12-04</span>
        <span style="color: #999;margin-left: 5px">17:37</span>
      </div>
    </template>
    <div class="video-preview" style="margin-left: 130px">
      <video src="../assets/teach/video1.mp4" width="880" height="600" controls></video>
    </div>
    <img src="@/assets/teach/v1.png"height="100" width="300" style="margin-left: 400px"/><br>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;padding: 20px">碳排放的核算主要是通过收集历史碳排放数据确定基准值，结合未来发展计划，比如产量、投资等因素，再由此测算未来碳排放的潜力。</span><br>
    <img src="@/assets/teach/v2.png"height="360" width="600" style="margin-left: 280px"/><br>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;padding: 20px">碳排放核算方法主要有三钟：包括排放因子法、质量平衡法、实测法。其中，排放因子法是目前应用最为广泛、普遍的方法。</span>
    <img src="@/assets/teach/v3.png"height="360" width="600" style="margin-left: 280px"/><br>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;padding: 20px">“碳排放核算途径通常可分为自上而下和自下而上两大类。所谓自上而下，指的是国家或政府层面的宏观测量。</span>
    <img src="@/assets/teach/v4.png"height="360" width="600" style="margin-left: 280px"/><br>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;padding: 20px">那反过来，自下而上则是下级单位的自行测算后向上级单位披露与汇总统计，包括企业的自测与披露，地方对中央的汇报汇总及各国对国际社会提交反馈。</span>
    <img src="@/assets/teach/v5.png"height="360" width="600" style="margin-left: 280px"/><br>
    <img src="@/assets/teach/v10.png"height="100" width="300" style="margin-left: 400px"/><br>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;padding: 20px">要做好企业碳排放核算，首先要了解企业层级温室气体排放的三种核算边界，即范围1、范围2、范围3三种类型。范围1是直接排放，范围2和3是间接排放。</span>
    <img src="@/assets/teach/v6.png"height="360" width="600" style="margin-left: 280px"/><br>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;padding: 20px">范围1排放指的是在企业实体控制范围之内，直接控制或拥有的排放源所产生的直接排放，包括静止燃烧、移动燃烧、化学或生产过程或无组织逸散等活动。</span>
    <img src="@/assets/teach/v7.png"height="360" width="600" style="margin-left: 280px"/><br>
    <span style="display: inline-block;font-size: 20px;padding: 20px">范围2排放指的是企业自用的外购电力和热力所产生的间接排放，也包括使用电力、蒸汽、热水、制冷器等。</span>
    <img src="@/assets/teach/v8.png"height="360" width="700" style="margin-left: 220px"/><br>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;padding: 20px">范围3排放是指企业在范围2以外的间接排放，包括企业供应链及企业上下游可能产生的所有排放，比如原材料的采掘、生产和运输、产品的使用及回收等。</span>
    <img src="@/assets/teach/v9.png"height="360" width="600" style="margin-left: 280px"/><br>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;margin-left: 900px">来源：环保科普365</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;margin-left: 900px">编辑：卫雅琦</span>
  </el-card>
</template>
<script>
export default {
  methods:{
    back(){
      this.$router.push('/Teach');
    }
  }
}
</script>
