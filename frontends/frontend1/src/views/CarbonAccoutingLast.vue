<template>
  <span style="font-size: 30px;margin-left: 20px">核算结果：</span>
  <el-card shadow="never" style="margin-top: 10px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表1 报告主体()年二氧化碳排放量报告</text>
    </template>
    <el-form :model="receivedData" label-width="340px">
      <el-form-item label="企业二氧化碳排放总量(吨/二氧化碳)">
        <el-input v-model="receivedData.Total" disabled></el-input>
      </el-form-item>
      <el-form-item label="使用六氟化硫设备修理与退役过程产生的排放(吨/二氧化碳)">
        <el-input v-model="receivedData.FSCO" disabled></el-input>
      </el-form-item>
      <el-form-item label="输配电引起的二氧化碳排放(吨/二氧化碳)">
        <el-input v-model="receivedData.SPDCO" disabled></el-input>
      </el-form-item>
    </el-form>
  </el-card>
  <div style="display: flex">
    <el-table :data="PGERepairREC" border style="width: 500px; margin-left: 50px; margin-top: 20px">
      <el-table-column label="六氟化硫回收">
        <el-table-column label="修理设备" prop="Id" width="100"></el-table-column>
        <el-table-column label="设备容量(千克)" width="200">
          <template #default="scope">
            <el-input v-model="scope.row.RECV" disabled></el-input>
          </template>
        </el-table-column>
        <el-table-column label="实际回收量(千克)" width="200">
          <template #default="scope">
            <el-input v-model="scope.row.RECR" disabled></el-input>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
    <el-table :data="PGERetireREC" border style="width:500px;margin-left: 50px; margin-top: 20px">
      <el-table-column label="六氟化硫回收">
        <el-table-column label="退役设备" prop="Id" width="100"></el-table-column>
        <el-table-column label="设备容量(千克)" width="200">
          <template #default="scope">
            <el-input v-model="scope.row.RECV" disabled></el-input>
          </template>
        </el-table-column>
        <el-table-column label="实际回收量(千克)" width="200">
          <template #default="scope">
            <el-input v-model="scope.row.RECR" disabled></el-input>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
  </div>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>输配电损失</text>
    </template>
    <el-form :model="SPDData" label-width="180px">
      <el-form-item label="电厂上网电量(兆瓦时)">
        <el-input v-model="SPDData.ELSW" disabled></el-input>
      </el-form-item>
      <el-form-item label="自外省输入电量(兆瓦时)">
        <el-input v-model="SPDData.ELSR" disabled></el-input>
      </el-form-item>
      <el-form-item label="向外省输出电量(兆瓦时)">
        <el-input v-model="SPDData.ELSC" disabled></el-input>
      </el-form-item>
      <el-form-item label="售电量(兆瓦时)">
        <el-input v-model="SPDData.ELSD" disabled></el-input>
      </el-form-item>
      <el-form-item label="输配电量(兆瓦时)">
        <el-input v-model="SPDData.ELSPDL" disabled></el-input>
      </el-form-item>
      <el-form-item label="区域电网年平均供电排放因子(吨二氧化碳/兆瓦时)">
        <el-input v-model="SPDData.EFDW" disabled></el-input>
      </el-form-item>
      <el-form-item>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>上传电气模型证明资料</text>
    </template>
    <el-upload
        ref="upload"
        class="upload-demo"
        action="https://jsonplaceholder.typicode.com/posts/"
        :on-preview="handlePreview"
        :on-remove="handleRemove"
        :file-list="fileList"
        list-type="picture-card"
        :auto-upload="false"
        @change="handleFileChange">
    <i class="el-icon-plus"></i>
    </el-upload>

    <!-- 添加一个上传按钮 -->
    <el-button type="primary" @click="customUpload" style="margin-top: 10px">提交所有数据</el-button>

  </el-card>
</template>
<script>
export default {
  data(){
    return{
      receivedData:{
        Total:'',
        FSCO:'',
        SPDCO:''
      },
      SPDData:{},
      PGERetireREC:[],
      PGERepairREC:[],
      fileList:[]
    }
  },
  created() {
    this.receivedData.Total = this.$route.query.Total
    this.receivedData.SPDCO=this.$route.query.SPDCO
    this.receivedData.FSCO=this.$route.query.FSCO
    this.dataload()
  },
  methods:{
    dataload(){
      const PGERetireREC=this.$store.state.PGERetireREC
      const formData2=this.$store.state.formData2
      const PGERepairREC=this.$store.state.PGERepairREC
      this.SPDData=formData2
      this.PGERepairREC=PGERepairREC
      this.PGERetireREC=PGERetireREC
      console.log(this.receivedData)
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      console.log(file);
    },
    handleFileChange(file, fileList) {
      this.fileList = fileList;
      console.log(this.fileList); // 检查fileList是否正确更新
    },
    async customUpload() {
      // 确保有文件被选中
      if (this.fileList.length === 0) {
        alert("请先选择文件！");
        return;
      }

      // 使用FormData来构建上传的数据
      const formData = new FormData();
      this.fileList.forEach((fileItem) => {
        formData.append("files[]", fileItem.raw); // "files" 是你的后端期望的字段名
        formData.append("CompanyId",this.$store.state.companyData.CompanyID)
      });

      try {
        this.$http.post("http://47.97.176.174:8087/chainmaker?cmb=TransmissionPGEFile", formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }).then((res)=>{
          this.$router.push({name:'SuccessUpload'})
        });
      } catch (error) {
        console.error("上传失败:", error);
        // 上传失败的处理逻辑...
      }
    }
  },
}
</script>
