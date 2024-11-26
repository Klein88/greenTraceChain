<template>
  <div style="display: flex">
    <div>
      <el-card style="margin-top: 10px; margin-left: 20px; margin-right: 20px">
        <template #header>
          <text style="font-size: 30px">自定义模型</text>
        </template>
        <div style="display: flex;margin-left: 20px;margin-top: 20px">
          <span style="font-size: 25px">公式：</span>
          <el-input v-model="formulaall" style="width: 400px" disabled></el-input>
        </div>
        <el-button type="primary" class="mt-4" @click="addItem()" style="margin-top: 10px; margin-left: 30px">新增参数</el-button>
        <div style="display: flex;margin-left: 100px">
          <el-table :data="newmodelcanshu" border style="width: 450px; margin-left: 8px; margin-top: 20px">
            <el-table-column label="参数" width="200">
              <template #default="scope">
                <el-input v-model="scope.row.canshu" placeholder="请输入参数" disabled></el-input>
              </template>
            </el-table-column>
            <el-table-column label="值" width="250">
              <template #default="scope">
                <el-input v-model="scope.row.zhi" placeholder="请输入值"></el-input>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div style="display: flex;margin-left: 100px;margin-top: 20px">
          <el-button type="primary" size="large"  @click="calculate">计算</el-button>
          <div style="display: flex;height: 30px;margin-left: 50px">
            <span style="font-size: 25px">计算结果：</span>
            <el-input v-model="result" style="width: 400px"></el-input>
          </div>
        </div>
      </el-card>
    </div>
    <el-card shadow="never" style="margin-top: 10px; margin-left: 20px; margin-right: 20px;width: 300px">
      <template #header>
        <text style="font-size: 30px">上传新模型资料</text>
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
      <el-button type="primary" @click="customUpload" style="margin-top: 10px">提交材料</el-button>

    </el-card>
  </div>
</template>
<script>
import {ElMessage} from "element-plus";

export default {
  data() {
    return {
      newmodelcanshu: [],
      formulaall: '',
      result:'',
      fileList:[]
    }
  },
  created() {
    this.formulaall = this.$route.query.Total
    this.extractParameters(this.formulaall);
  },
  methods: {
    addItem() {
      this.newmodelcanshu.push({
        canshu: '',
        zhi: ''
      });
    },
    calculate() {
      const res1 = this.combineFormulas(this.newmodelcanshu);
      const values = {};
      res1.forEach((value, key) => {
        values[key] = value;
      });
      this.$http.post('http://47.97.176.174:8087/chainmaker?cmb=Calculate', {
        companyId: this.$store.state.companyData.CompanyID,
        expression: this.formulaall,
        values: values
      }).then((res) => {
        console.log(res.data);
        this.result=res.data.data
        ElMessage({
          message: '计算成功！',
          type: 'success',
        });
      }).catch((err) => {
        console.log(err);
      });
    },
    extractParameters(expression) {
      const regex = /[a-zA-Z]+/g;
      const matches = expression.match(regex);
      const uniqueParams = [...new Set(matches)]; // 过滤掉重复的参数
      uniqueParams.forEach(param => {
        this.newmodelcanshu.push({
          canshu: param,
          zhi: ''
        });
      });
    },
    combineFormulas(array) {
      const values = new Map();
      array.forEach(item => {
        const floatValue = parseFloat(item.zhi);
        if (!isNaN(floatValue)) { // 确保是有效的数值
          values.set(item.canshu, floatValue);
        }
      });
      return values;
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
  }
}
</script>
