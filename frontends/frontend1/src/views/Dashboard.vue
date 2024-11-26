<template>
  <div>
    <div style="display: flex">
      <span style="font-size: 20px;margin-left: 20px;font-weight: bold">碳排分析</span>
      <el-row style="display: flex;margin-left: 750px">
        <el-button icon="DocumentCopy"></el-button>
        <el-button icon="Download"></el-button>
        <el-button icon="DataLine"></el-button>
      </el-row>
      <el-select
          v-model="value"
          placeholder="Select"
          style="width: 140px;margin-left: 12px"
      >
        <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        />
      </el-select>
    </div>
    <div style="background-color: #FFFFFF;margin-top: 20px;margin-left: 20px;width: 1155px;height: 260px;border-radius:10px">
      <span style="font-size: 15px;margin-left: 30px;font-weight: bold;margin-top: 20px;display: inline-block">碳排概览</span>
      <div style="display: flex;height: 100%">
        <div style="margin-top: 20px;margin-left: 55px">
          <span style="display: block;font-size: 12px">碳排总量</span>
          <span style="display: block;font-size: 12px;color: #999;">/tCO₂e</span>
          <span style="display: block;font-size: 40px;">{{total.toFixed(2)}}</span>
          <span style="font-size: 12px">上年&nbsp;&nbsp;&nbsp;&nbsp;0&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;同比&nbsp;&nbsp;&nbsp;&nbsp;0%</span>
        </div>
        <div style="margin-left: 100px;margin-top: 25px">
          <el-divider direction="vertical" style="display:block;background-color: orange;height: 50px;" border-style="#009900"></el-divider>
          <el-divider direction="vertical" style="display:block;background-color: darkgreen;margin-top: 18px;height: 50px;" border-style="#009900"></el-divider>
        </div>
        <div style="margin-top: 25px;margin-left: 10px">
          <div>
            <span style="display: block;font-size: 14px">碳排放额度</span>
            <span style="display: block;font-size: 14px;margin-left: 5px">7600.00</span>
          </div>
          <div style="margin-top: 28px;margin-left: 2px">
            <span style="display: block;font-size: 14px">剩余额度</span>
            <span style="display: block;font-size: 14px;margin-left: 5px">{{cc}}</span>
          </div>
        </div>
        <div style="margin-left: 220px;margin-top: 20px">
          <span style="display: block;font-size: 25px;margin-left: 5px">碳币剩余</span>
          <span style="display: block;font-size: 45px;margin-left: 16px">{{cb.toFixed(2)}}</span>
          <span style="font-size: 12px">上年&nbsp;&nbsp;&nbsp;&nbsp;0&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;同比&nbsp;&nbsp;&nbsp;&nbsp;0%</span>
        </div>
        <div style="margin-left: 140px;margin-top: 20px">
          <span style="display: block;font-size: 12px">碳排水平</span>
          <span style="display: block;font-size: 12px;color: #999;">较国家碳排标准</span>
          <span style="font-size: 60px;color: forestgreen">低</span>
        </div>
      </div>
    </div>
  </div>

  <div style="width: 1155px;margin-left: 20px;margin-top: 20px">
    <el-tabs v-model="activeName" class="demo-tabs" tab-active-color="#008000">
      <el-tab-pane label="碳排详情" name="first">
        <el-tabs
            v-model="activeNamenew"
            type="card"
        >
          <el-tab-pane label="总量" name="firstnew">
              <div style="display: flex">
                <div>
                  <div style="width: 500px;background-color: #FFFFFF;height: 145px;border-radius: 10px">
                      <span style="display: inline-block;font-size: 12px;font-weight: bold;margin-top: 10px;margin-left: 30px">碳排总量</span>
                      <span style="display: block;font-size: 12px;color: #999;margin-left: 30px">/tCO₂e</span>
                      <span style="display: block;font-size: 40px;margin-left: 30px">{{total.toFixed(2)}}</span>
                      <span style="font-size: 12px;margin-left: 30px">上年&nbsp;&nbsp;&nbsp;&nbsp;0&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;同比&nbsp;&nbsp;&nbsp;&nbsp;0%</span>
                  </div>
                  <div style="display: flex;margin-top: 10px">
                      <div style="width: 270px;background-color: #FFFFFF;height: 170px;border-radius: 10px">
                        <span style="display: inline-block;font-size: 10px;font-weight: bold;margin-top: 15px;margin-left: 30px">碳排放量类别占比</span>
                        <div id="main" style="height: 150px;width: 250px;margin-bottom: 10px"></div>
                      </div>
                      <div style="width: 215px;background-color: #FFFFFF;height: 170px;border-radius: 10px;margin-left: 15px">
                        <span style="display: inline-block;font-size: 10px;font-weight: bold;margin-top: 15px;margin-left: 30px">温室气体排放量占比</span>
                        <div class="demo-progress" style="margin-left: 10px">
                          <span style="font-size: 10px;font-weight: bold">CO₂</span>
                          <el-progress :percentage="99"/>
                          <span style="font-size: 10px;font-weight: bold">NO₂</span>
                          <el-progress :percentage="1"/>
                          <span style="font-size: 10px;font-weight: bold">CH₄</span>
                          <el-progress :percentage="0"/>
                        </div>
                      </div>
                  </div>
                </div>
                <div style="width: 700px;background-color: #FFFFFF;height: 325px;border-radius: 10px;margin-left: 15px">
                    <span style="font-weight: bold;font-size: 18px;margin-top: 30px;margin-left: 40px;display: inline-block">趋势分析</span>
                      <div id="mainnew" style="width: 650px;height: 300px"></div>
                </div>
              </div>
          </el-tab-pane>
          <el-tab-pane label="电气" name="secondnew">Config</el-tab-pane>
          <el-tab-pane label="化工" name="thirdnew">Role</el-tab-pane>
        </el-tabs>
      </el-tab-pane>
      <el-tab-pane label="市场详情" name="second">Config</el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import * as echarts from 'echarts';

const chart = () => {
  const myChart = echarts.init(document.getElementById('main'));
  const option = {
    tooltip: {
      trigger: 'item'
    },

    series: [
      {
        type: 'pie',
        radius: ['40%', '70%'],
        center: ['52%', '43%'],
        avoidLabelOverlap: false,
        padAngle: 5,
        itemStyle: {
          borderRadius: 10
        },
        label: {
          show: false,
          position: 'center'
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 1048, name: '电气' },
          { value: 735, name: '化工' },
          { value: 580, name: '钢铁' },
          { value: 484, name: '平板玻璃' },
          { value: 300, name: '水泥' }
        ]
      }
    ]
  };
  myChart.setOption(option);
}

const chartnew =() =>{
  const myChart = echarts.init(document.getElementById('mainnew'));
  const option = {
    xAxis: {
      type: 'category',
      data: ['一月', '二月', '三月', '四月', '五月', '六月', '七月','八月','九月','十月','十一月','十二月']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: [756, 850, 901, 999, 1290, 1330, 1630, 1845, 2100, 2267, 2456, 2789],
        type: 'line',
        smooth: true
      }
    ],
    grid:[
      {
        top:'10%'
      }
    ]
  };
  myChart.setOption(option);
}

export default {
  data(){
    return{
      options:[
        {
          value: '2024',
          label: '2024',
        },
        {
          value: '2023',
          label: '2023',
        },
        {
          value: '2022',
          label: '2022',
        },
        {
          value: '2021',
          label: '2021',
        },
      ],
      value:'2024',
      total:6500.0,
      cc:'',
      cb:'',
      activeName:'first',
      activeNamenew:'firstnew',
    }
  },
  created() {
    this.cc = this.$store.state.companyData.CarbonCredit;
    this.cb=this.$store.state.companyData.CarbonCoin;
  },
  methods:{

  },
  mounted() {
    chart()
    chartnew()
  } ,
  computed: {
    // // 计算属性，确保cc只显示小数点后两位
    // formattedCc() {
    //   // 使用toFixed方法保留两位小数
    //   return parseFloat(this.cb).toFixed(2);
    // },
  },
}
</script>

