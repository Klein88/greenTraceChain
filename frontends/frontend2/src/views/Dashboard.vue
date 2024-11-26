<script setup>
import {
  ArrowRight,
} from '@element-plus/icons-vue'
import * as echarts from 'echarts';
import {onMounted} from "vue";

const createLine = () => {
  const myChart = echarts.init(document.getElementById('line'));
  myChart.setOption({
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985'
        }
      }
    },
    legend: {
      data: ['访问量', '注册量']
    },
    grid: {
      left: '1%',
      right: '2%',
      bottom: '1%',
      top: '15%',
      containLabel: true
    },
    xAxis: [
      {
        type: 'category',
        data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
      }
    ],
    yAxis: [
      {
        type: 'value'
      }
    ],
    series: [
      {
        name: '访问量',
        type: 'line',
        areaStyle: {},
        emphasis: {
          focus: 'series'
        },
        data: [10, 32, 41, 84, 45, 30, 10],
        smooth: true
      },
      {
        name: '注册量',
        type: 'line',
        areaStyle: {},
        emphasis: {
          focus: 'series'
        },
        data: [6, 82, 51, 34, 80, 30, 10],
        smooth: true
      }
    ]
  });
}

const createRadar = () => {
  const myChart = echarts.init(document.getElementById('radar'));
  myChart.setOption({
    tooltip: {
      trigger: 'item'
    },
    legend: {
      type: 'scroll',
      bottom: 0,
      data: (function () {
        const list = [];
        for (let i = 1; i <= 12; i++) {
          list.push(i + '月' + '');
        }
        return list;
      })()
    },
    visualMap: {
      top: 'middle',
      right: 0,
      color: ['red', 'yellow'],
      calculable: true
    },
    radar: {
      indicator: [
        { name: '图片' },
        { name: '文档' },
        { name: '表格' },
        { name: '压缩包' },
        { name: 'PDF' }
      ]
    },
    series: (function () {
      const series = [];
      for (let i = 1; i <= 12; i++) {
        series.push({
          type: 'radar',
          symbol: 'none',
          lineStyle: {
            width: 1
          },
          emphasis: {
            areaStyle: {
              color: 'rgba(0,250,0,0.3)'
            }
          },
          data: [
            {
              value: [
                (40 - i) * 10,
                (38 - i) * 4 + 60,
                i * 5 + 10,
                i * 9,
                (i * i) / 0.5
              ],
              name: i + '月' + ''
            }
          ]
        });
      }
      return series;
    })()
  });
}

const pathSymbols = {
  reindeer:
      'path://M-22.788,24.521c2.08-0.986,3.611-3.905,4.984-5.892 c-2.686,2.782-5.047,5.884-9.102,7.312c-0.992,0.005-0.25-2.016,0.34-2.362l1.852-0.41c0.564-0.218,0.785-0.842,0.902-1.347 c2.133-0.727,4.91-4.129,6.031-6.194c1.748-0.7,4.443-0.679,5.734-2.293c1.176-1.468,0.393-3.992,1.215-6.557 c0.24-0.754,0.574-1.581,1.008-2.293c-0.611,0.011-1.348-0.061-1.959-0.608c-1.391-1.245-0.785-2.086-1.297-3.313 c1.684,0.744,2.5,2.584,4.426,2.586C-8.46,3.012-8.255,2.901-8.04,2.824c6.031-1.952,15.182-0.165,19.498-3.937 c1.15-3.933-1.24-9.846-1.229-9.938c0.008-0.062-1.314-0.004-1.803-0.258c-1.119-0.771-6.531-3.75-0.17-3.33 c0.314-0.045,0.943,0.259,1.439,0.435c-0.289-1.694-0.92-0.144-3.311-1.946c0,0-1.1-0.855-1.764-1.98 c-0.836-1.09-2.01-2.825-2.992-4.031c-1.523-2.476,1.367,0.709,1.816,1.108c1.768,1.704,1.844,3.281,3.232,3.983 c0.195,0.203,1.453,0.164,0.926-0.468c-0.525-0.632-1.367-1.278-1.775-2.341c-0.293-0.703-1.311-2.326-1.566-2.711 c-0.256-0.384-0.959-1.718-1.67-2.351c-1.047-1.187-0.268-0.902,0.521-0.07c0.789,0.834,1.537,1.821,1.672,2.023 c0.135,0.203,1.584,2.521,1.725,2.387c0.102-0.259-0.035-0.428-0.158-0.852c-0.125-0.423-0.912-2.032-0.961-2.083 c-0.357-0.852-0.566-1.908-0.598-3.333c0.4-2.375,0.648-2.486,0.549-0.705c0.014,1.143,0.031,2.215,0.602,3.247 c0.807,1.496,1.764,4.064,1.836,4.474c0.561,3.176,2.904,1.749,2.281-0.126c-0.068-0.446-0.109-2.014-0.287-2.862 c-0.18-0.849-0.219-1.688-0.113-3.056c0.066-1.389,0.232-2.055,0.277-2.299c0.285-1.023,0.4-1.088,0.408,0.135 c-0.059,0.399-0.131,1.687-0.125,2.655c0.064,0.642-0.043,1.768,0.172,2.486c0.654,1.928-0.027,3.496,1,3.514 c1.805-0.424,2.428-1.218,2.428-2.346c-0.086-0.704-0.121-0.843-0.031-1.193c0.221-0.568,0.359-0.67,0.312-0.076 c-0.055,0.287,0.031,0.533,0.082,0.794c0.264,1.197,0.912,0.114,1.283-0.782c0.15-0.238,0.539-2.154,0.545-2.522 c-0.023-0.617,0.285-0.645,0.309,0.01c0.064,0.422-0.248,2.646-0.205,2.334c-0.338,1.24-1.105,3.402-3.379,4.712 c-0.389,0.12-1.186,1.286-3.328,2.178c0,0,1.729,0.321,3.156,0.246c1.102-0.19,3.707-0.027,4.654,0.269 c1.752,0.494,1.531-0.053,4.084,0.164c2.26-0.4,2.154,2.391-1.496,3.68c-2.549,1.405-3.107,1.475-2.293,2.984 c3.484,7.906,2.865,13.183,2.193,16.466c2.41,0.271,5.732-0.62,7.301,0.725c0.506,0.333,0.648,1.866-0.457,2.86 c-4.105,2.745-9.283,7.022-13.904,7.662c-0.977-0.194,0.156-2.025,0.803-2.247l1.898-0.03c0.596-0.101,0.936-0.669,1.152-1.139 c3.16-0.404,5.045-3.775,8.246-4.818c-4.035-0.718-9.588,3.981-12.162,1.051c-5.043,1.423-11.449,1.84-15.895,1.111 c-3.105,2.687-7.934,4.021-12.115,5.866c-3.271,3.511-5.188,8.086-9.967,10.414c-0.986,0.119-0.48-1.974,0.066-2.385l1.795-0.618 C-22.995,25.682-22.849,25.035-22.788,24.521z',
  plane:
      'path://M1.112,32.559l2.998,1.205l-2.882,2.268l-2.215-0.012L1.112,32.559z M37.803,23.96 c0.158-0.838,0.5-1.509,0.961-1.904c-0.096-0.037-0.205-0.071-0.344-0.071c-0.777-0.005-2.068-0.009-3.047-0.009 c-0.633,0-1.217,0.066-1.754,0.18l2.199,1.804H37.803z M39.738,23.036c-0.111,0-0.377,0.325-0.537,0.924h1.076 C40.115,23.361,39.854,23.036,39.738,23.036z M39.934,39.867c-0.166,0-0.674,0.705-0.674,1.986s0.506,1.986,0.674,1.986 s0.672-0.705,0.672-1.986S40.102,39.867,39.934,39.867z M38.963,38.889c-0.098-0.038-0.209-0.07-0.348-0.073 c-0.082,0-0.174,0-0.268-0.001l-7.127,4.671c0.879,0.821,2.42,1.417,4.348,1.417c0.979,0,2.27-0.006,3.047-0.01 c0.139,0,0.25-0.034,0.348-0.072c-0.646-0.555-1.07-1.643-1.07-2.967C37.891,40.529,38.316,39.441,38.963,38.889z M32.713,23.96 l-12.37-10.116l-4.693-0.004c0,0,4,8.222,4.827,10.121H32.713z M59.311,32.374c-0.248,2.104-5.305,3.172-8.018,3.172H39.629 l-25.325,16.61L9.607,52.16c0,0,6.687-8.479,7.95-10.207c1.17-1.6,3.019-3.699,3.027-6.407h-2.138 c-5.839,0-13.816-3.789-18.472-5.583c-2.818-1.085-2.396-4.04-0.031-4.04h0.039l-3.299-11.371h3.617c0,0,4.352,5.696,5.846,7.5 c2,2.416,4.503,3.678,8.228,3.87h30.727c2.17,0,4.311,0.417,6.252,1.046c3.49,1.175,5.863,2.7,7.199,4.027 C59.145,31.584,59.352,32.025,59.311,32.374z M22.069,30.408c0-0.815-0.661-1.475-1.469-1.475c-0.812,0-1.471,0.66-1.471,1.475 s0.658,1.475,1.471,1.475C21.408,31.883,22.069,31.224,22.069,30.408z M27.06,30.408c0-0.815-0.656-1.478-1.466-1.478 c-0.812,0-1.471,0.662-1.471,1.478s0.658,1.477,1.471,1.477C26.404,31.885,27.06,31.224,27.06,30.408z M32.055,30.408 c0-0.815-0.66-1.475-1.469-1.475c-0.808,0-1.466,0.66-1.466,1.475s0.658,1.475,1.466,1.475 C31.398,31.883,32.055,31.224,32.055,30.408z M37.049,30.408c0-0.815-0.658-1.478-1.467-1.478c-0.812,0-1.469,0.662-1.469,1.478 s0.656,1.477,1.469,1.477C36.389,31.885,37.049,31.224,37.049,30.408z M42.039,30.408c0-0.815-0.656-1.478-1.465-1.478 c-0.811,0-1.469,0.662-1.469,1.478s0.658,1.477,1.469,1.477C41.383,31.885,42.039,31.224,42.039,30.408z M55.479,30.565 c-0.701-0.436-1.568-0.896-2.627-1.347c-0.613,0.289-1.551,0.476-2.73,0.476c-1.527,0-1.639,2.263,0.164,2.316 C52.389,32.074,54.627,31.373,55.479,30.565z',
  rocket:
      'path://M-244.396,44.399c0,0,0.47-2.931-2.427-6.512c2.819-8.221,3.21-15.709,3.21-15.709s5.795,1.383,5.795,7.325C-237.818,39.679-244.396,44.399-244.396,44.399z M-260.371,40.827c0,0-3.881-12.946-3.881-18.319c0-2.416,0.262-4.566,0.669-6.517h17.684c0.411,1.952,0.675,4.104,0.675,6.519c0,5.291-3.87,18.317-3.87,18.317H-260.371z M-254.745,18.951c-1.99,0-3.603,1.676-3.603,3.744c0,2.068,1.612,3.744,3.603,3.744c1.988,0,3.602-1.676,3.602-3.744S-252.757,18.951-254.745,18.951z M-255.521,2.228v-5.098h1.402v4.969c1.603,1.213,5.941,5.069,7.901,12.5h-17.05C-261.373,7.373-257.245,3.558-255.521,2.228zM-265.07,44.399c0,0-6.577-4.721-6.577-14.896c0-5.942,5.794-7.325,5.794-7.325s0.393,7.488,3.211,15.708C-265.539,41.469-265.07,44.399-265.07,44.399z M-252.36,45.15l-1.176-1.22L-254.789,48l-1.487-4.069l-1.019,2.116l-1.488-3.826h8.067L-252.36,45.15z',
  train:
      'path://M67.335,33.596L67.335,33.596c-0.002-1.39-1.153-3.183-3.328-4.218h-9.096v-2.07h5.371 c-4.939-2.07-11.199-4.141-14.89-4.141H19.72v12.421v5.176h38.373c4.033,0,8.457-1.035,9.142-5.176h-0.027 c0.076-0.367,0.129-0.751,0.129-1.165L67.335,33.596L67.335,33.596z M27.999,30.413h-3.105v-4.141h3.105V30.413z M35.245,30.413 h-3.104v-4.141h3.104V30.413z M42.491,30.413h-3.104v-4.141h3.104V30.413z M49.736,30.413h-3.104v-4.141h3.104V30.413z  M14.544,40.764c1.143,0,2.07-0.927,2.07-2.07V35.59V25.237c0-1.145-0.928-2.07-2.07-2.07H-9.265c-1.143,0-2.068,0.926-2.068,2.07 v10.351v3.105c0,1.144,0.926,2.07,2.068,2.07H14.544L14.544,40.764z M8.333,26.272h3.105v4.141H8.333V26.272z M1.087,26.272h3.105 v4.141H1.087V26.272z M-6.159,26.272h3.105v4.141h-3.105V26.272z M-9.265,41.798h69.352v1.035H-9.265V41.798z'
};
const createPictorialBar = () => {
  const myChart = echarts.init(document.getElementById('pictorialBar'));
  myChart.setOption({
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'none'
      },
      formatter: function (params) {
        return params[0].name + ': ' + params[0].value;
      }
    },
    xAxis: {
      data: ['百度', '直接访问', '坐飞机', '坐高铁'],
      axisTick: { show: false },
      axisLine: { show: false },
      axisLabel: {
        color: '#e54035'
      }
    },
    yAxis: {
      splitLine: { show: false },
      axisTick: { show: false },
      axisLine: { show: false },
      axisLabel: { show: false }
    },
    color: ['#e54035'],
    series: [
      {
        name: 'hill',
        type: 'pictorialBar',
        barCategoryGap: '-130%',
        symbol: 'path://M0,10 L10,10 C5.5,10 5.5,5 5,0 C4.5,5 4.5,10 0,10 z',
        itemStyle: {
          opacity: 0.5
        },
        emphasis: {
          itemStyle: {
            opacity: 1
          }
        },
        data: [123, 60, 25, 80],
        z: 10
      },
      {
        name: 'glyph',
        type: 'pictorialBar',
        barGap: '-100%',
        symbolPosition: 'end',
        symbolSize: 50,
        symbolOffset: [0, '-120%'],
        data: [
          {
            value: 123,
            symbol: pathSymbols.reindeer,
            symbolSize: [60, 60]
          },
          {
            value: 60,
            symbol: pathSymbols.rocket,
            symbolSize: [50, 60]
          },
          {
            value: 25,
            symbol: pathSymbols.plane,
            symbolSize: [65, 35]
          },
          {
            value: 80,
            symbol: pathSymbols.train,
            symbolSize: [50, 30]
          }
        ]
      }
    ]
  });
}

const data = genData(20);
const createPie = () => {
  const myChart = echarts.init(document.getElementById('pie'));
  myChart.setOption({
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b} : {c} 吨 ({d}%)'
    },
    legend: {
      type: 'scroll',
      orient: 'vertical',
      right: 10,
      top: 20,
      bottom: 20,
      data: data.legendData
    },
    series: [
      {
        name: '碳排放量',
        type: 'pie',
        radius: '55%',
        center: ['40%', '50%'],
        data: data.seriesData,
        emphasis: {
          focus: 'self',
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.5)'
          }
        }
      }
    ]
  });
}
function genData(count) {
  const List = [
    '能源电力', '制造业', '交通运输', '建筑业', '农业', '矿业',
    '金融业', '信息技术', '零售业', '教育', '医疗健康', '媒体通讯',
    '服务业', '旅游休闲', '化工', '食品饮料', '纺织服装', '家具制造',
    '纸品印刷', '环保产业'
  ];

  const legendData = [];
  const seriesData = [];
  for (let i = 0; i < count; i++) {
    const name = List[i];
    legendData.push(name);
    seriesData.push({
      name: name,
      value: Math.round(Math.random() * 10000)
    });
  }
  return {
    legendData: legendData,
    seriesData: seriesData
  };
}

onMounted(() => {
  createLine()
  createRadar()
  createPictorialBar()
  createPie()
})
</script>

<template>
  <el-scrollbar>
    <div class="main">
      <div class="main-top">
        <el-card shadow="never">
          <div class="main-top-illustration">
            <img src="../assets/dashboard/illustration.svg" alt="illustration">
            <div class="main-top-illustration-text">
              <h1>基于区块链的碳核算和碳交易系统</h1>
              <span>本选题是利用区块链技术解决碳排放数据的真实性和可靠性，以及打通核算与交易流程中的各个环节，消除信息孤岛，实现信息共享，提升碳交易流程的透明性。</span>
            </div>
          </div>
        </el-card>
        <div class="main-top-coffee">
          <img src="@/assets/dashboard/coffee.svg" alt="coffee">
          <span>咖啡，让每一天都充满活力。</span>
        </div>
      </div>
      <div class="main-middle">
        <div class="main-middle-top">
          <el-card shadow="never">
            <div class="main-middle-top-outer">
              <span>企业注册量</span>
              <div class="main-middle-top-inner">
                <div>
                  <svg t="1711658157625" class="icon" viewBox="0 0 1184 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="114415" width="25" height="25"><path d="M1170.272 877.728l0 73.152-1170.272 0 0-877.728 73.152 0 0 804.576 1097.152 0zM1097.152 164.576l0 248.576q0 12-11.136 16.864t-20.288-4.288l-69.152-69.152-361.728 361.728q-5.728 5.728-13.152 5.728t-13.152-5.728l-133.152-133.152-237.728 237.728-109.728-109.728 334.272-334.272q5.728-5.728 13.152-5.728t13.152 5.728l133.152 133.152 265.152-265.152-69.152-69.152q-9.152-9.152-4.288-20.288t16.864-11.136l248.576 0q8 0 13.152 5.152t5.152 13.152z" p-id="114416" fill="#8595f4"></path></svg>
                  <span>56</span>
                </div>
                <span>+14%</span>
              </div>
            </div>
          </el-card>
          <el-card shadow="never">
            <div class="main-middle-top-outer">
              <span>附表上传量</span>
              <div class="main-middle-top-inner">
                <div>
                  <svg t="1711658436438" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="116345" width="25" height="25"><path d="M890.49536 272q8 8 16 20.576l-269.728 0 0-269.728q12.576 8 20.576 16zM618.49536 365.728l310.848 0 0 603.424q0 22.848-16 38.848t-38.848 16l-768 0q-22.848 0-38.848-16t-16-38.848l0-914.272q0-22.848 16-38.848t38.848-16l457.152 0 0 310.848q0 22.848 16 38.848t38.848 16zM709.91936 786.272l0-36.576q0-8-5.152-13.152t-13.152-5.152l-402.272 0q-8 0-13.152 5.152t-5.152 13.152l0 36.576q0 8 5.152 13.152t13.152 5.152l402.272 0q8 0 13.152-5.152t5.152-13.152zM709.91936 640l0-36.576q0-8-5.152-13.152t-13.152-5.152l-402.272 0q-8 0-13.152 5.152t-5.152 13.152l0 36.576q0 8 5.152 13.152t13.152 5.152l402.272 0q8 0 13.152-5.152t5.152-13.152zM709.91936 493.728l0-36.576q0-8-5.152-13.152t-13.152-5.152l-402.272 0q-8 0-13.152 5.152t-5.152 13.152l0 36.576q0 8 5.152 13.152t13.152 5.152l402.272 0q8 0 13.152-5.152t5.152-13.152z" p-id="116346" fill="#ad85f4"></path></svg>
                  <span>34</span>
                </div>
                <span>+50%</span>
              </div>
            </div>
          </el-card>
          <el-card shadow="never">
            <div class="main-middle-top-outer">
              <span>企业总数</span>
              <div class="main-middle-top-inner">
                <div>
                  <svg t="1711658001467" class="icon" viewBox="0 0 1097 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="104465" width="25" height="25"><path d="M338.857143 512q-92.571429 2.857143-151.428572 73.142857H110.857143q-46.857143 0-78.857143-23.142857T0 494.285714q0-201.714286 70.857143-201.714285 3.428571 0 24.857143 12t55.714285 24.285714T219.428571 341.142857q38.285714 0 76-13.142857-2.857143 21.142857-2.857142 37.714286 0 79.428571 46.285714 146.285714z m612 364q0 68.571429-41.714286 108.285714t-110.857143 39.714286H298.857143q-69.142857 0-110.857143-39.714286T146.285714 876q0-30.285714 2-59.142857t8-62.285714T171.428571 692.571429t24.571429-55.714286 35.428571-46.285714 48.857143-30.571429T344 548.571429q5.714286 0 24.571429 12.285714t41.714285 27.428571 61.142857 27.428572 77.142858 12.285714 77.142857-12.285714 61.142857-27.428572 41.714286-27.428571 24.571428-12.285714q34.857143 0 63.714286 11.428571t48.857143 30.571429 35.428571 46.285714 24.571429 55.714286 15.142857 62 8 62.285714 2 59.142857zM365.714286 146.285714q0 60.571429-42.857143 103.428572t-103.428572 42.857143-103.428571-42.857143-42.857143-103.428572 42.857143-103.428571T219.428571 0t103.428572 42.857143 42.857143 103.428571z m402.285714 219.428572q0 90.857143-64.285714 155.142857T548.571429 585.142857 393.428571 520.857143 329.142857 365.714286t64.285714-155.142857T548.571429 146.285714t155.142857 64.285715T768 365.714286z m329.142857 128.571428q0 44.571429-32 67.714286t-78.857143 23.142857h-76.571428q-58.857143-70.285714-151.428572-73.142857 46.285714-66.857143 46.285715-146.285714 0-16.571429-2.857143-37.714286 37.714286 13.142857 76 13.142857 33.714286 0 68-12.285714t55.714285-24.285714 24.857143-12q70.857143 0 70.857143 201.714285z m-73.142857-348q0 60.571429-42.857143 103.428572t-103.428571 42.857143-103.428572-42.857143-42.857143-103.428572 42.857143-103.428571 103.428572-42.857143 103.428571 42.857143 42.857143 103.428571z" p-id="104466" fill="#74a8b5"></path></svg>
                  <span>86</span>
                </div>
                <span>+28%</span>
              </div>
            </div>
          </el-card>
          <el-card shadow="never">
            <div class="main-middle-top-outer">
              <span>已使用模型数</span>
              <div class="main-middle-top-inner">
                <div>
                  <svg t="1711657970526" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="103398" width="25" height="25"><path d="M933.973614 270.237771l0 482.756979 60.344367 0 0 181.034123L813.283858 934.028872l0-60.344367L209.838146 873.684506l0 60.344367L28.804023 934.028872 28.804023 752.99475l60.344367 0L89.148389 270.237771l-60.344367 0L28.804023 89.203648l181.034123 0 0 60.344367 603.445712 0 0-60.344367 181.034123 0L994.31798 270.237771 933.973614 270.237771zM89.149413 209.893404l60.344367 0 0-60.344367-60.344367 0L89.149413 209.893404zM149.493779 813.339116l-60.344367 0 0 60.344367 60.344367 0L149.493779 813.339116zM813.283858 752.99475l60.344367 0L873.628224 270.237771l-60.344367 0 0-60.344367L209.838146 209.893404l0 60.344367-60.344367 0 0 482.756979 60.344367 0 0 60.344367 603.445712 0L813.283858 752.99475zM813.283858 390.926504l0 362.067222L390.872269 752.993726l0-120.688733L209.838146 632.304993 209.838146 270.237771l422.412612 0 0 120.688733L813.283858 390.926504zM571.905368 571.960627 571.905368 330.582137 270.182512 330.582137l0 241.378489L571.905368 571.960627zM752.939491 451.27087l-120.688733 0 0 181.034123L451.216635 632.304993l0 60.344367 301.722856 0L752.939491 451.27087zM873.628224 209.893404l60.344367 0 0-60.344367-60.344367 0L873.628224 209.893404zM933.973614 813.339116l-60.344367 0 0 60.344367 60.344367 0L933.973614 813.339116z" fill="#f48595" p-id="103399"></path></svg>
                  <span>5</span>
                </div>
                <span>+88%</span>
              </div>
            </div>
          </el-card>
        </div>
        <div class="main-middle-middle">
          <div class="main-middle-middle-left">
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <span>企业增长情况</span>
                </div>
              </template>
              <div id="line" style="width: 380px; height:254px;"></div>
            </el-card>
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <span>附表增长情况</span>
                </div>
              </template>
              <div id="radar" style="width: 380px; height:254px;"></div>
            </el-card>
          </div>
          <div class="main-middle-middle-right">
            <el-card shadow="hover">
              <template #header>
                <div class="card-header">
                  <span>刚刚加入的企业</span>
                </div>
              </template>
              <div class="card-body">
                <el-scrollbar height="274px">
                  <el-card shadow="never">
                    <div class="item">
                      <img src="@/assets/dashboard/companyLogo/银河电力.webp" alt="companyLogo" height="50px" style="border-radius: 50%">
                      <div>
                        <h1>银河电力有限公司</h1>
                        <span>1分钟前加入了我们</span>
                      </div>
                      <el-button
                          type="primary"
                          tag="a"
                          href="https://github.com"
                          target="_blank"
                          :icon="ArrowRight"
                          size="small"
                          text
                      >
                      </el-button>
                    </div>
                  </el-card>
                  <el-card shadow="never">
                    <div class="item">
                      <img src="../assets/dashboard/companyLogo/绿盛化工.webp" alt="companyLogo" height="50px" style="border-radius: 50%">
                      <div>
                        <h1>绿盛化工有限公司</h1>
                        <span>7分钟前加入了我们</span>
                      </div>
                      <el-button
                          type="primary"
                          tag="a"
                          href="https://github.com"
                          target="_blank"
                          :icon="ArrowRight"
                          size="small"
                          text
                      >
                      </el-button>
                    </div>
                  </el-card>
                  <el-card shadow="never">
                    <div class="item">
                      <img src="@/assets/dashboard/companyLogo/星辉能源.webp" alt="companyLogo" height="50px" style="border-radius: 50%">
                      <div>
                        <h1>星辉能源有限公司</h1>
                        <span>13分钟前加入了我们</span>
                      </div>
                      <el-button
                          type="primary"
                          tag="a"
                          href="https://github.com"
                          target="_blank"
                          :icon="ArrowRight"
                          size="small"
                          text
                      >
                      </el-button>
                    </div>
                  </el-card>
                  <el-card shadow="never">
                    <div class="item">
                      <img src="@/assets/dashboard/companyLogo/金瑞科技.webp" alt="companyLogo" height="50px" style="border-radius: 50%">
                      <div>
                        <h1>金瑞科技有限公司</h1>
                        <span>22分钟前加入了我们</span>
                      </div>
                      <el-button
                          type="primary"
                          tag="a"
                          href="https://github.com"
                          target="_blank"
                          :icon="ArrowRight"
                          size="small"
                          text
                      >
                      </el-button>
                    </div>
                  </el-card>
                  <el-card shadow="never">
                    <div class="item">
                      <img src="@/assets/dashboard/companyLogo/星河电能.webp" alt="companyLogo" height="50px" style="border-radius: 50%">
                      <div>
                        <h1>星河电能有限公司</h1>
                        <span>27分钟前加入了我们</span>
                      </div>
                      <el-button
                          type="primary"
                          tag="a"
                          href="https://github.com"
                          target="_blank"
                          :icon="ArrowRight"
                          size="small"
                          text
                      >
                      </el-button>
                    </div>
                  </el-card>
                </el-scrollbar>
              </div>
            </el-card>
          </div>
        </div>
        <div class="main-middle-bottom">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span>企业来源</span>
              </div>
            </template>
            <div id="pictorialBar" style="width: 520px; height:399px;"></div>
          </el-card>
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <span>碳排放数据</span>
              </div>
            </template>
            <div id="pie" style="width: 520px; height:399px;"></div>
          </el-card>
        </div>
      </div>
    </div>
  </el-scrollbar>
</template>

<style scoped>
.main {
  margin: 15px;
}

.main-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-top :deep(.el-card__body) {
  padding: 0;
}

.main-top .el-card {
  transition: all .3s ease;
}

.main-top .el-card:hover {
  transform: translateY(-4px) scale(1.02);
}

.main-top-illustration {
  height: 130px;
  width: 850px;
  background-color: #E1EAF9;
  display: flex;
  justify-content: left;
  align-items: center;
}

.main-top-illustration img {
  height: 100px;
  width: 150px;
  margin: 0 15px;
}

.main-top-illustration-text h1 {
  font-size: 20px;
  color: #3f6ad8;
}

.main-top-coffee {
  height: 130px;
  width: 270px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
}

.main-top-coffee img {
  height: 80px;
  width: 80px;
  transition: all .3s ease;
}

.main-top-coffee img:hover {
  transform: translateY(-4px) scale(1.02);
}

.main-top-coffee span {
  transition: all .3s ease;
}

.main-top-coffee span:hover {
  transform: translateY(-4px) scale(1.02);
}

.main-middle {
  margin-top: 15px;
}

.main-middle-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-middle-top .el-card {
  height: 125px;
  width: 275px;
  background-color: #e9edf2;
  transition: all .3s ease;
}

.main-middle-top .el-card:hover {
  transform: translateY(-4px) scale(1.02);
  box-shadow: 0 14px 24px #0003;
}

.main-middle-top-outer {
  height: 85px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: start;
}

.main-middle-top-outer > span {
  color: #92969a;
  margin: 5px 0 0 5px;
}

.main-middle-top-inner {
  width: 235px;
  display: flex;
  justify-content: space-between;
  align-items: end;
}

.main-middle-top-inner > span {
  font-size: 18px;
  color: #2c3f5d;
  font-family: Helvetica Neue, Helvetica, PingFang SC, Hiragino Sans GB, Microsoft YaHei, SimSun, sans-serif;
  margin: 0 5px 0 0;
}

.main-middle-top-inner > div {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0 0 0 5px;
}

.main-middle-top-inner > div > span {
  font-size: 28px;
  color: #2c3f5d;
  font-family: Helvetica Neue, Helvetica, PingFang SC, Hiragino Sans GB, Microsoft YaHei, SimSun, sans-serif;
  margin: -5px 0 0 10px;
}

.main-middle-middle {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 15px 0;
}

.main-middle-middle-left {
  height: 355px;
  width: 850px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-middle-middle-left .el-card {
  height: 355px;
  width: 420px;
}

.main-middle-middle-right {
  height: 355px;
  width: 275px;
}

.main-middle-middle-right .el-card {
  height: 355px;
  width: 275px;
}

.main-middle-middle-right .el-card .card-body {
  margin: -10px;
}

.main-middle-middle-right .el-card .card-body .el-card {
  height: 88px;
  width: 255px;
  background-color: ghostwhite;
  margin-bottom: 10px;
}

.main-middle-middle-right .el-card .card-body .el-card .item{
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-middle-middle-right .el-card .card-body .el-card .item h1{
  font-size: 14px;
}

.main-middle-middle-right .el-card .card-body .el-card .item span{
  font-size: 13px;
}

.main-middle-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-middle-bottom .el-card {
  height: 500px;
  width: 560px;
}

h1 {
  margin: 0;
}
</style>
