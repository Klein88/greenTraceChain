<template>
  <el-dialog v-model="loading"width="40%">
    <div style="height: 150px">
      <span style="font-size: 30px;margin-top: 50px;display: inline-block;margin-left: 30px">数据处理中，请耐心等待...</span>
    </div>
  </el-dialog>
  <div>
    <span style="font-size: 30px;margin-left: 20px;margin-top: 20px">碳核算申报</span>
    <div style="display: flex;margin-left: 40px">
      <span style="margin-top: 3px">核算模型:</span>
      <el-select v-model="selectedModel" placeholder="请选择模型" style="width: 14%;margin-left: 5px">
        <el-option
            v-for="model in models"
            :key="model.id"
            :label="model.name"
            :value="model.id"
        ></el-option>
      </el-select>
    </div>
    <div style="display: flex">
      <div class="table" style="width: 74%">
        <!-- 根据选择显示表格 -->
        <div v-if="selectedModel === 'electrical'" style="margin-top: 15px">
          <div>
            <el-steps :active="currentIndex" finish-status="success" align-center>
              <el-step v-for="(step, index) in steps" :key="index" :title="step.title"></el-step>
            </el-steps>
            <!-- 显示当前步骤内容 -->
            <component :is="currentStepComponent" @stepComplete="handleStepComplete" />
            <div style="text-align: center;margin-bottom: 20px;margin-top: 20px">
              <!-- 上一步和下一步按钮 -->
              <el-button v-if="currentIndex > 0" @click="prevStep">上一步</el-button>
              <el-button v-if="currentIndex === steps.length - 1" @click="submit">下一步</el-button>
              <el-button v-else-if="currentIndex < steps.length - 1" @click="nextStep">下一步</el-button>
            </div>
          </div>
        </div>
        <div v-else-if="selectedModel === 'chemicalindustry'">
          <div>
            <el-steps :active="currentIndexnew" finish-status="success" align-center>
              <el-step v-for="(step, index) in stepsnew" :key="index" :title="step.title"></el-step>
            </el-steps>
            <!-- 显示当前步骤内容 -->
            <component :is="currentStepComponentnew" @stepComplete="handleStepCompletenew" />
            <div style="text-align: center;margin-bottom: 20px;margin-top: 20px">
              <!-- 上一步和下一步按钮 -->
              <el-button v-if="currentIndexnew > 0" @click="prevStepnew">上一步</el-button>
              <el-button v-if="currentIndexnew === stepsnew.length - 1" @click="submitnew">下一步</el-button>
              <el-button v-else-if="currentIndexnew < stepsnew.length - 1" @click="nextStepnew">下一步</el-button>
            </div>
          </div>
        </div>
      </div>
      <div class="video" style="width: 25.5%;height: 100%">
        <div v-if="selectedModel == 'electrical'" style="background-color: #FFFFFF;width: 100%;padding: 10px;margin-bottom: 87px">
          <span style="font-size: 20px">核算教程</span><br>
          <span style="font-size: 16px">教学视频</span>
          <div class="video-preview">
            <video src="../assets/teach/dianqivideo.mp4" width="280" height="200" controls></video>
          </div>
          <span style="display: inline-block;margin-top: 5px">电网公式:</span>
          <span style="display: inline-block;margin-top: 2px">E = (总和（REC容量i - REC回收i）*23.9) + (EL上网 + EL输入 - EL输出 - EL售电) * EF电网</span>
          <span style="display: inline-block;margin-top: 5px">输配电损耗的电量计算公式：</span>
          <span style="display: inline-block;margin-left: 10px">AD网损=EL供电 - EL售电</span><br>
          <span style="display: inline-block;margin-top: 5px">供电量计算公式：</span>
          <span style="display: inline-block;margin-left: 10px">EL供电=EL上网 + EL输入 - EL输出</span><br>
          <el-divider></el-divider>

          <span style="display: inline-block;font-size: 13px">REC容量i ——退役设备i的六氟化硫容量</span>
          <span style="display: inline-block;font-size: 13px">REC回收i ——退役设备i的六氟化硫实际回收量</span>
          <span style="display: inline-block;font-size: 13px">REC容量j ——修理设备j的六氟化硫容量</span>
          <span style="display: inline-block;font-size: 13px">REC回收j ——修理设备j的六氟化硫实际回收量</span>
          <span style="display: inline-block;font-size: 13px">&nbsp&nbsp&nbsp&nbspAD网损——输配电损耗的电量</span>
          <span style="display: inline-block;font-size: 13px">&nbsp&nbsp&nbsp&nbsp&nbspEF电网——区域电网年平均供电排放因子</span>
          <span style="display: inline-block;font-size: 13px">&nbsp&nbsp&nbsp&nbsp&nbspEL供电——供电量</span>
          <span style="display: inline-block;font-size: 13px">&nbsp&nbsp&nbsp&nbsp&nbspEL售电——售电量，即终端用户用电量</span>
          <span style="display: inline-block;font-size: 13px">&nbsp&nbsp&nbsp&nbsp&nbspEL上网——电厂上网电量</span>
          <span style="display: inline-block;font-size: 13px">&nbsp&nbsp&nbsp&nbsp&nbspEL输入——自外省输入电量</span>
          <span style="display: inline-block;font-size: 13px">&nbsp&nbsp&nbsp&nbsp&nbspEL输出——自外省输出电量</span>
        </div>
        <div v-if="selectedModel == 'chemicalindustry'" style="background-color: #FFFFFF;width: 100%;padding: 10px;margin-bottom: 87px">
          <span style="font-size: 20px">核算教程</span><br>
          <span style="font-size: 16px">教学视频</span>
          <div class="video-preview">
            <video src="../assets/teach/huagongvideo.mp4" width="280" height="200" controls></video>
          </div>
          <span style="display: inline-block;margin-top: 5px">化工公式:</span>
          <span style="display: inline-block;margin-top: 2px">Eghg =(总和 (ADi * (NCVi  * EFi) * OFi * 44 /12)) + ( ( ( (总和ADr * CCr  -  总和ADp*CCp - 总和ADw*CCw) * 44 /12 ) + (总和ADi*EFi*PURi) )  + ( (En2o硝酸总和ADj*EFj*(1 - nk*jk)/1000 )   +  (En2o总和ADj*EFj*(1 - nk*jk)/1000 )  )* 310 ) - (Q * PURco2 *197.7) + (AD电力*EF电力) + (AD热力*EF热力)</span>
          <el-divider></el-divider>
          <span style="display: inline-block;font-size: 13px">附表二</span><br>
          <span style="display: inline-block;font-size: 13px">ADi&nbsp燃烧量&nbsp&nbsp&nbsp&nbsp&nbsp&nbspCCi&nbsp含碳量&nbsp&nbsp&nbsp&nbspNCVi&nbsp低位发热量 </span>
          <span style="display: inline-block;font-size: 13px">EFi&nbsp单位热值含碳量&nbsp&nbsp&nbsp&nbspOFi&nbsp碳氧化率</span><br>
          <span style="display: inline-block;font-size: 13px">附表三</span><br>
          <span style="display: inline-block;font-size: 13px">ADr&nbsp碳输入活动水平数据&nbsp&nbsp&nbsp&nbsp&nbspCCr&nbsp碳输入含碳量</span>
          <span style="display: inline-block;font-size: 13px">ADp&nbspADw&nbsp碳输出活动水平数据</span>
          <span style="display: inline-block;font-size: 13px">CCp&nbspCCw&nbsp碳输出含碳量</span><br>
          <span style="display: inline-block;font-size: 13px">附表四</span><br>
          <span style="display: inline-block;font-size: 13px">ADi&nbsp消耗量&nbsp&nbsp&nbspEFi&nbspCo2排放因子</span><br>
          <span style="display: inline-block;font-size: 13px">附表五</span><br>
          <span style="display: inline-block;font-size: 13px">ADj&nbsp硝酸产量&nbsp&nbsp&nbspEFj&nbspN2o生成因子&nbsp&nbsp&nbspnk&nbspN2o去除率</span><br>
          <span style="display: inline-block;font-size: 13px">uk&nbsp尾气处理设备使用率</span><br>
          <span style="display: inline-block;font-size: 13px">附表六</span><br>
          <span style="display: inline-block;font-size: 13px">ADj&nbsp己二酸产量&nbsp&nbsp&nbspEFj&nbspN2o生成因子&nbspnk&nbspN2o去除率</span><br>
          <span style="display: inline-block;font-size: 13px">uk&nbsp尾气处理设备使用率</span><br>
          <span style="display: inline-block;font-size: 13px">附表七</span><br>
          <span style="display: inline-block;font-size: 13px">AD电力&nbsp&nbsp净购入量&nbsp&nbsp&nbsp&nbsp&nbsp&nbspEF电力&nbsp&nbspCo2排放因子</span><br>
          <span style="display: inline-block;font-size: 13px">AD热力&nbsp&nbsp净购入量&nbsp&nbsp&nbsp&nbsp&nbsp&nbspEF热力&nbsp&nbspCo2排放因子</span><br>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Electrical_6FS from "@/views/Electrical/Electrical_6FS.vue";
import Transmissionanddistributionlosses from "@/views/Electrical/Transmissionanddistributionlosses.vue";
import ChemicalIndustry_GreenhouseGases from "@/views/ChemicalIndustry/ChemicalIndustry_GreenhouseGases.vue";
import ChemicalIndustry_FossilFuel from "@/views/ChemicalIndustry/ChemicalIndustry_FossilFuel.vue";
import ChemicalIndustry_Industry from "@/views/ChemicalIndustry/ChemicalIndustry_Industry.vue";
import ChemicalIndustry_M2CO3 from "@/views/ChemicalIndustry/ChemicalIndustry_M2CO3.vue";
import ChemicalIndustry_HNO3 from "@/views/ChemicalIndustry/ChemicalIndustry_HNO3.vue";
import ChemicalIndustry_Electrical from "@/views/ChemicalIndustry/ChemicalIndustry_Electrical.vue";
import ChemicalIndustry_HOOCCH2 from "@/views/ChemicalIndustry/ChemicalIndustry_HOOCCH2.vue";
import {ElMessage} from "element-plus";

export default {
  data() {
    return {
      loading:false,
      selectedModel: '',
      models: [
        { id: 'electrical', name: '电气' },
        { id: 'chemicalindustry', name: '化工' },
      ],
      currentIndex: 0, // 当前步骤索引
      currentIndexnew: 0, // 当前步骤索引
      steps: [
        { title: '六氟化硫回收' },
        { title: '输配电损失' },
      ],
      stepsnew: [
        { title: '温室气体' },
        { title: '化石燃料' },
        { title: '工业生产' },
        { title: '碳酸盐使用' },
        { title: '硝酸生产' },
        { title: '已二酸生产' },
        { title: '电力和热力' },
      ],
      stepData: {
        Electrical_6FS: {},
        Transmissionanddistributionlosses:{},
      }
    }
  },
  computed: {
    // 根据当前索引返回相应的步骤组件
    currentStepComponent() {
      switch (this.currentIndex) {
        case 0:
          return Electrical_6FS;
        case 1:
          return Transmissionanddistributionlosses;
        default:
          return null;
      }
    },
    currentStepComponentnew() {
      switch (this.currentIndexnew) {
        case 0:
          return ChemicalIndustry_GreenhouseGases;
        case 1:
          return ChemicalIndustry_FossilFuel;
        case 2:
          return ChemicalIndustry_Industry;
        case 3:
          return ChemicalIndustry_M2CO3;
        case 4:
          return ChemicalIndustry_HNO3;
        case 5:
          return ChemicalIndustry_HOOCCH2;
        case 6:
          return ChemicalIndustry_Electrical;
        default:
          return null;
      }
    }
  },
  methods: {
    handleStepComplete(data, componentName) {
      // this.$set(this.stepData, componentName, data);
      // 切换到下一个步骤
      if (this.currentIndex < this.steps.length - 1) {
        this.currentIndex++;
      }
    },
    handleStepCompletenew(data, componentName) {
      // this.$set(this.stepData, componentName, data);
      // 切换到下一个步骤
      if (this.currentIndexnew < this.stepsnew.length - 1) {
        this.currentIndexnew++;
      }
    },
    // 切换到上一步
    prevStep() {
      if (this.currentIndex > 0) {
        this.currentIndex--;
      }
    },
    // 切换到下一步
    nextStep() {
      if (this.currentIndex < this.steps.length - 1) {
        this.currentIndex++;
      }
    },
    // 切换到上一步
    prevStepnew() {
      if (this.currentIndexnew > 0) {
        this.currentIndexnew--;
      }
    },
    // 切换到下一步
    nextStepnew() {
      if (this.currentIndexnew < this.stepsnew.length - 1) {
        this.currentIndexnew++;
      }
    },
    // 提交
    submit() {
      this.loading = true;
      const ID=this.$store.state.companyData.CompanyID;
      const PGERetireREC=this.$store.state.PGERetireREC
      const formData2=this.$store.state.formData2
      const PGERepairREC=this.$store.state.PGERepairREC
      PGERepairREC.forEach(item => {
        item.RECV = parseFloat(item.RECV);
        item.RECR = parseFloat(item.RECR);
      });
      PGERetireREC.forEach(item => {
        item.RECV = parseFloat(item.RECV);
        item.RECR = parseFloat(item.RECR);
      });
      const upload={
        CompanyId: ID,
        PGERepairREC:PGERepairREC,
        PGERetireREC:PGERetireREC,
        ELSW: parseFloat(formData2.ELSW),
        ELSR: parseFloat(formData2.ELSR),
        ELSC: parseFloat(formData2.ELSC),
        ELSD: parseFloat(formData2.ELSD),
        ELSPDL: parseFloat(formData2.ELSPDL),
        EFDW: parseFloat(formData2.EFDW)
      }
      this.$http.post('http://47.97.176.174:8087/chainmaker?cmb=CalculationForPGE',upload).then((res)=>{
        this.loading=false
        const Total=res.data.data.Total
        const FSCO=res.data.data.FSCO
        const SPDCO=res.data.data.SPDCO
        console.log(Total,FSCO,SPDCO)
        this.$router.push({name:'carbonAccountingLast',query:{Total:Total,FSCO:FSCO,SPDCO:SPDCO}});
        ElMessage({
          message: '报告生成成功！',
          type: 'success',
        })
      }).catch((err) => {
        console.log(err);
      });
    },
    submitnew(){
      this.loading = true;
      const ID=this.$store.state.companyData.CompanyID;
      const greenhouse=this.$store.state.greenhouse;
      const FossilFuel=this.$store.state.FossilFuel
      const IndustryInput=this.$store.state.IndustryInput
      const IndustryOutput=this.$store.state.IndustryOutput
      const M2CO3=this.$store.state.M2CO3
      const HNO3=this.$store.state.HNO3
      const HOO=this.$store.state.HOO
      const Electrical=this.$store.state.Electrical
      greenhouse.forEach(item => {
        item.WS = parseFloat(item.WS);
        item.CO = parseFloat(item.CO);
      });
      FossilFuel.forEach(item => {
        item.AD = parseFloat(item.AD);
        item.CC = parseFloat(item.CC);
        item.CCDS = parseInt(item.CCDS);
        item.NCV = parseFloat(item.NCV);
        item.NCVDS = parseInt(item.NCVDS);
        item.EF = parseFloat(item.EF);
        item.OF = parseFloat(item.OF);
        item.OFDS = parseInt(item.OFDS);
      });
      IndustryInput.forEach(item => {
        item.AD = parseFloat(item.AD);
        item.CC = parseFloat(item.CC);
        item.CCDS = parseInt(item.CCDS);
      });
      IndustryOutput.forEach(item => {
        item.AD = parseFloat(item.AD);
        item.CC = parseFloat(item.CC);
        item.CCDS = parseInt(item.CCDS);
      });
      M2CO3.forEach(item => {
        item.AD = parseFloat(item.AD);
        item.EF = parseFloat(item.CC);
        item.EFDS = parseInt(item.CCDS);
      });
      HNO3.forEach(item => {
        item.AD = parseFloat(item.AD);
        item.EF = parseFloat(item.EF);
        item.EFDS = parseInt(item.EFDS);
        item.NK = parseFloat(item.NK);
        item.NKDS = parseInt(item.NKDS);
        item.UK = parseFloat(item.UK);
        item.UKDS = parseInt(item.UKDS);
      });
      HOO.forEach(item => {
        item.AD = parseFloat(item.AD);
        item.EF = parseFloat(item.EF);
        item.EFDS = parseInt(item.EFDS);
        item.NK = parseFloat(item.NK);
        item.NKDS = parseInt(item.NKDS);
        item.UK = parseFloat(item.UK);
        item.UKDS = parseInt(item.UKDS);
      });
      Electrical.forEach(item => {
        item.AD = parseFloat(item.AD);
        item.AD1 = parseFloat(item.AD1);
        item.AD2 = parseFloat(item.AD2);
        item.EF = parseFloat(item.EF);
      });
      let max=greenhouse[3].WS+greenhouse[3].CO;
      const upload={
        CompanyId:ID,
        COSummaryS:greenhouse,
        Ecors:FossilFuel,
        Eghg:{
          GYSCCORC:{
            GYSCCOR:IndustryInput,
            GYSCCOC:IndustryOutput
          },
          CACOS:M2CO3,
          XSCOS:HNO3,
          YRSCOS:HOO,
        },
        RCO:max,
        Drcos:Electrical
      }
      console.log(upload)
      this.$http.post('http://47.97.176.174:8087/chainmaker?cmb=CalculationForCE',upload).then((res)=>{
        this.loading=false
        console.log(res.data)
        this.$router.push({name:'CarbonAccoutingLastChemicallndustry'});
        ElMessage({
          message: '报告生成成功！',
          type: 'success',
        })
      }).catch((err) => {
        console.log(err);
      });
    }
  }
}
</script>
