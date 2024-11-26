<template>
  <span style="font-size: 30px;margin-left: 20px">核算结果：</span>
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
          <el-input v-model="greenhouse[0].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px" disabled></el-input>
          <el-input v-model="greenhouse[0].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
      <el-form-item label="工业生产过程CO2排放">
        <div style="display: flex">
          <el-input v-model="greenhouse[1].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px" disabled></el-input>
          <el-input v-model="greenhouse[1].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
      <el-form-item label="工业生产过程N2O排放">
        <div style="display: flex">
          <el-input v-model="greenhouse[2].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px" disabled></el-input>
          <el-input v-model="greenhouse[2].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
      <el-form-item label="CO2回收利用量">
        <div style="display: flex">
          <el-input v-model="greenhouse[3].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px" disabled></el-input>
          <el-input v-model="greenhouse[3].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
      <el-form-item label="企业净购入的电力和热力消费引起的CO2排放" style="height: 50px">
        <div style="display: flex">
          <el-input v-model="greenhouse[4].WS" placeholder="请输入数据" style="width:300px;margin-left: 10px" disabled></el-input>
          <el-input v-model="greenhouse[4].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
      <el-form-item label="企业温室气体排放总量（吨CO2当量）">
        <div style="display: flex">
          <el-input v-model="greenhouse[5].WS" placeholder="" style="width:300px;margin-left: 10px" disabled></el-input>
          <el-input v-model="greenhouse[5].CO" placeholder="请输入数据" style="width: 300px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
      <el-form-item>
      </el-form-item>
    </el-form>
  </el-card>
  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表二 化石燃料燃烧的活动水平和排放因子数据一览表</text>
    </template>
    <el-table :data="FossilFuel" border style="width:96%; margin-left: 8px; margin-top: 20px">
      <el-table-column label="燃料名称" width="120">
        <template #default="scope">
          <el-input v-model="scope.row.Breed" placeholder="请输入名称" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="燃烧量(吨或万Nm3)" width="160">
        <template #default="scope">
          <el-input v-model="scope.row.AD" placeholder="请输入数据" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="含碳量(tC/吨或tC/万Nm3)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.CC" placeholder="请输入数据" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="数据来源" width="170"  height="120">
        <template #default="scope">
          <el-radio-group v-model="scope.row.CCDS" size="small" style="display: flex;" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-left: -10px">计算值</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
      <el-table-column label="低位发热量*(GJ/吨或GJ/万Nm3)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.NCV" placeholder="请输入数据" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="数据来源" width="170"  height="120" >
        <template #default="scope">
          <el-radio-group v-model="scope.row.NCVDS" size="small" style="display: flex;" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-left: -10px">计算值</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
      <el-table-column label="单位热值含碳量*(tC/GJ)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.EF" placeholder="请输入数据" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="碳氧化率(%)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.OF" placeholder="请输入数据" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="数据来源" width="170"  height="120">
        <template #default="scope">
          <el-radio-group v-model="scope.row.OFDS" size="small" style="display: flex;" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-left: -10px">计算值</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
    </el-table>
  </el-card>

  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表三 工业生产过程CO2排放的活动水平和排放因子数据一览表</text>
    </template>
    <el-table :data="IndustryInput" border style="width:96%; margin-left: 8px; margin-top: 20px">
      <el-table-column label="碳输入">
        <el-table-column label="物料名称" prop="Id" width="120">
          <template #default="scope">
            <el-input v-model="scope.row.Breed" placeholder="请输入名称" disabled></el-input>
          </template>
        </el-table-column>
        <el-table-column label="活动水平数据(单位:吨或万Nm3)" width="160">
          <template #default="scope">
            <el-input v-model="scope.row.AD" placeholder="请输入数据" disabled></el-input>
          </template>
        </el-table-column>
        <el-table-column label="含碳量(单位:tC/吨)" width="140">
          <template #default="scope">
            <el-input v-model="scope.row.CC" placeholder="请输入数据" disabled></el-input>
          </template>
        </el-table-column>
        <el-table-column label="数据来源" width="260"  height="120">
          <template #default="scope">
            <el-radio-group v-model="scope.row.CCDS" size="small" style="display: flex;" disabled>
              <el-radio :label="0">检测值</el-radio><br>
              <el-radio :label="1" style="margin-left: -10px">化学计算</el-radio>
              <el-radio :label="2" style="margin-left: -10px">缺省值</el-radio>
            </el-radio-group>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
    <el-table :data="IndustryOutput" border style="width:96%; margin-left: 8px; margin-top: 20px">
      <el-table-column label="碳输出">
        <el-table-column label="物料名称" prop="Id" width="120">
          <template #default="scope">
            <el-input v-model="scope.row.Breed" placeholder="请输入名称" disabled></el-input>
          </template>
        </el-table-column>
        <el-table-column label="活动水平数据(单位:吨或万Nm3)" width="160">
          <template #default="scope">
            <el-input v-model="scope.row.AD" placeholder="请输入数据" disabled></el-input>
          </template>
        </el-table-column>
        <el-table-column label="含碳量(单位:tC/吨)" width="140">
          <template #default="scope">
            <el-input v-model="scope.row.CC" placeholder="请输入数据" disabled></el-input>
          </template>
        </el-table-column>
        <el-table-column label="数据来源" width="260"  height="120">
          <template #default="scope">
            <el-radio-group v-model="scope.row.CCDS" size="small" style="display: flex;" disabled>
              <el-radio :label="0">检测值</el-radio><br>
              <el-radio :label="1" style="margin-left: -10px">化学计算</el-radio>
              <el-radio :label="2" style="margin-left: -10px">缺省值</el-radio>
            </el-radio-group>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
  </el-card>


  <el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>附表四 碳酸盐使用的活动水平和排放因子数据一览表</text>
    </template>
    <el-table :data="M2CO3" border style="width:96%; margin-left: 8px; margin-top: 20px">
      <el-table-column label="碳酸盐种类" prop="Id" width="120">
        <template #default="scope">
          <el-input v-model="scope.row.Breed" placeholder="请输入名称" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="消耗量(单位：吨)" width="160">
        <template #default="scope">
          <el-input v-model="scope.row.AD" placeholder="请输入数据" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="CO2排放因子(吨CO2/吨碳酸盐)" width="140">
        <template #default="scope">
          <el-input v-model="scope.row.EF" placeholder="请输入数据" disabled></el-input>
        </template>
      </el-table-column>
      <el-table-column label="数据来源" width="260"  height="120">
        <template #default="scope">
          <el-radio-group v-model="scope.row.EFDS" size="small" style="display: flex;" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-left: -10px">化学计算</el-radio>
            <el-radio :label="2" style="margin-left: -10px">缺省值</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
    </el-table>
  </el-card>

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
          <el-input v-model="HNO3[0].AD" placeholder="输入数据" style="width:85px;margin-left: 30px" disabled></el-input>
          <el-input v-model="HNO3[0].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px" disabled></el-input>
          <el-radio-group v-model="HNO3[0].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[0].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px" disabled></el-input>
          <el-radio-group v-model="HNO3[0].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[0].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px" disabled></el-input>
          <el-radio-group v-model="HNO3[0].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="中压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[1].AD" placeholder="输入数据" style="width:85px;margin-left: 30px" disabled></el-input>
          <el-input v-model="HNO3[1].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px" disabled></el-input>
          <el-radio-group v-model="HNO3[1].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[1].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px" disabled></el-input>
          <el-radio-group v-model="HNO3[1].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[1].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px" disabled></el-input>
          <el-radio-group v-model="HNO3[1].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="常压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[2].AD" placeholder="输入数据" style="width:85px;margin-left: 30px" disabled></el-input>
          <el-input v-model="HNO3[2].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px" disabled></el-input>
          <el-radio-group v-model="HNO3[2].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[2].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px" disabled></el-input>
          <el-radio-group v-model="HNO3[2].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[2].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px" disabled></el-input>
          <el-radio-group v-model="HNO3[2].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="双加压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[3].AD" placeholder="输入数据" style="width:85px;margin-left: 30px" disabled></el-input>
          <el-input v-model="HNO3[3].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px" disabled></el-input>
          <el-radio-group v-model="HNO3[3].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[3].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px" disabled></el-input>
          <el-radio-group v-model="HNO3[3].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[3].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px" disabled></el-input>
          <el-radio-group v-model="HNO3[3].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="综合法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[4].AD" placeholder="输入数据" style="width:85px;margin-left: 30px" disabled></el-input>
          <el-input v-model="HNO3[4].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px" disabled></el-input>
          <el-radio-group v-model="HNO3[4].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[4].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px"disabled ></el-input>
          <el-radio-group v-model="HNO3[4].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[4].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px" disabled></el-input>
          <el-radio-group v-model="HNO3[4].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="低压法" style="margin-left: -100px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HNO3[5].AD" placeholder="输入数据" style="width:85px;margin-left: 30px" disabled></el-input>
          <el-input v-model="HNO3[5].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px" disabled></el-input>
          <el-radio-group v-model="HNO3[5].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px"disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[5].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px" disabled></el-input>
          <el-radio-group v-model="HNO3[5].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HNO3[5].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px" disabled></el-input>
          <el-radio-group v-model="HNO3[5].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
    </el-form>
  </el-card>

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
          <el-input v-model="HOO[0].AD" placeholder="输入数据" style="width:85px;margin-left: 30px" disabled></el-input>
          <el-input v-model="HOO[0].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px" disabled></el-input>
          <el-radio-group v-model="HOO[0].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HOO[0].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px" disabled></el-input>
          <el-radio-group v-model="HOO[0].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HOO[0].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px" disabled></el-input>
          <el-radio-group v-model="HOO[0].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item label="其他" style="margin-left: -88px;height: 30px;margin-top: 5px">
        <div style="display: flex; align-items: center">
          <el-input v-model="HOO[1].AD" placeholder="输入数据" style="width:85px;margin-left: 30px" disabled></el-input>
          <el-input v-model="HOO[1].EF" placeholder="请输入数据" style="width: 100px;margin-left: 10px" disabled></el-input>
          <el-radio-group v-model="HOO[1].EFDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HOO[1].NK" placeholder="请输入数据" style="width: 95px;margin-left: -10px" disabled></el-input>
          <el-radio-group v-model="HOO[1].NKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
          <el-input v-model="HOO[1].UK" placeholder="请输入数据" style="width: 140px;margin-left: -15px" disabled></el-input>
          <el-radio-group v-model="HOO[1].UKDS" size="small" style="margin-left: 15px;display: grid;margin-top: -10px" disabled>
            <el-radio :label="0">检测值</el-radio><br>
            <el-radio :label="1" style="margin-top: -35px">缺省值</el-radio>
          </el-radio-group>
        </div>
      </el-form-item>
    </el-form>
  </el-card>

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
          <el-input v-model="Electrical[0].AD" placeholder="请输入数据" style="width: 160px;margin-left: 10px" disabled></el-input>
          <el-input v-model="Electrical[0].AD1" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
          <el-input v-model="Electrical[0].AD2" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
          <el-input v-model="Electrical[0].EF" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
      <el-form-item label="蒸汽" style="margin-left: -100px">
        <div style="display: flex">
          <el-input v-model="Electrical[1].AD" placeholder="请输入数据" style="width: 160px;margin-left: 10px" disabled></el-input>
          <el-input v-model="Electrical[1].AD1" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
          <el-input v-model="Electrical[1].AD2" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
          <el-input v-model="Electrical[1].EF" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
      <el-form-item label="热水" style="margin-left: -100px">
        <div style="display: flex">
          <el-input v-model="Electrical[2].AD" placeholder="请输入数据" style="width: 160px;margin-left: 10px" disabled></el-input>
          <el-input v-model="Electrical[2].AD1" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
          <el-input v-model="Electrical[2].AD2" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
          <el-input v-model="Electrical[2].EF" placeholder="请输入数据" style="width: 160px;margin-left: 20px" disabled></el-input>
        </div>
      </el-form-item>
    </el-form>
  </el-card>


<el-card shadow="never" style="margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <template #header>
      <text>上传化工模型证明资料</text>
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
      greenhouse:{},
      fileList:[],
      FossilFuel:{},
      IndustryInput:{},
      IndustryOutput:{},
      M2CO3:{},
      HNO3:{},
      HOO:{},
      Electrical:{},
    }
  },
  created() {
    this.dataload()
  },
  methods:{
    dataload(){
      const greenhouse=this.$store.state.greenhouse
      const FossilFuel=this.$store.state.FossilFuel
      const IndustryInput=this.$store.state.IndustryInput
      const IndustryOutput=this.$store.state.IndustryOutput
      const M2CO3=this.$store.state.M2CO3
      const HNO3=this.$store.state.HNO3
      const HOO=this.$store.state.HOO
      const Electrical=this.$store.state.Electrical
      this.greenhouse=greenhouse
      this.FossilFuel=FossilFuel
      this.IndustryInput=IndustryInput
      this.IndustryOutput=IndustryOutput
      this.M2CO3=M2CO3
      this.HNO3=HNO3
      this.HOO=HOO
      this.Electrical=Electrical
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
