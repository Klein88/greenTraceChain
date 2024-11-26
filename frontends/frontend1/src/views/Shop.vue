

<style scoped>
.app-container {
  background-image: url('../views/Shop/backgroundShop3.svg');
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
  height: 200vh;
}
</style>
<template>

<div class="app-container">

  <el-dialog v-model="this.viewDialogVisible" style="width: 900px">
      <div style="display: flex ; flex-direction: column;justify-content: center;align-items: center">
          <div style="display: flex;flex-direction: row;justify-content: left">
            <div style="font-weight: bold;font-size: 30px;color: cyan">
               -----------------交易详情-----------------
            </div>
          </div>
          <div style="display: flex ;flex-direction: row;justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black;margin-top: 30px">
            <div style="font-size: 25px ; font-weight: bold ; width: 300px ; background-color: #98F5FF">
               出售方 :
            </div>
            <div style="font-weight: bold;font-size: 13px">
               {{this.selecttrade.companyID}}
            </div>
          </div>

          <div v-if="this.selecttrade.TradeStatus == 'Closed'" style="display: flex ;flex-direction: row ; justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black">
            <div style="font-size: 25px ; font-weight: bold ; width: 300px ; background-color: #98F5FF">
              购买方 :
            </div>
            <div style="font-weight: bold;font-size: 13px">
              {{this.selecttrade.BuyerID}}
            </div>
          </div>

          <div v-if="this.selecttrade.TradeStatus == 'Closed'" style="display: flex ;flex-direction: row ; justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black">
            <div style="font-size: 25px ; font-weight: bold ; width: 300px ; background-color: #98F5FF">
              出售方出售碳额度 :
            </div>
            <div style="font-weight: bold;font-size: 13px">
              {{this.selecttrade.CarbonCredit}}
            </div>
          </div>
          <div style="display: flex ;flex-direction: row ; justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black">
            <div style="font-size: 25px ; font-weight: bold ;width: 300px ; background-color: #98F5FF">
              出售方需求碳币 :
            </div>
            <div style="font-weight: bold;font-size: 13px">
              {{this.selecttrade.CarbonCoin}}
            </div>
          </div>
          <div  style="display: flex ;flex-direction: row ; justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black">
            <div style="font-size: 25px ; font-weight: bold ;width: 300px ; background-color: #98F5FF">
              出售方需求碳币 :
            </div>
            <div style="font-weight: bold;font-size: 13px">
              {{this.selecttrade.CarbonCoin}}
            </div>
          </div>
          <div v-if="this.selecttrade.TradeStatus == 'Closed'" style="display: flex ;flex-direction: row ; justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black">
            <div style="font-size: 25px ; font-weight: bold ; width: 300px ; background-color: #98F5FF">
              交易ID :
            </div>
            <div style="font-weight: bold;font-size: 13px">
              {{this.selecttrade.TxID}}
            </div>
          </div>
          <div  style="display: flex ;flex-direction: row ; justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black">
            <div style="font-size: 25px ; font-weight: bold ; width: 300px ; background-color: #98F5FF">
              交易状态 :
            </div>
            <div style="font-weight: bold;font-size: 13px">
              {{this.selecttrade.TradeStatus}}
            </div>
          </div>
          <div  style="display: flex ;flex-direction: row ; justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black">
            <div style="font-size: 25px ; font-weight: bold ; width: 300px ; background-color: #98F5FF">
              创建时间 :
            </div>
            <div style="font-weight: bold;font-size: 13px">
              {{this.selecttrade.createAt}}
            </div>
          </div>

          <div v-if="this.selecttrade.TradeStatus == 'Closed'" style="display: flex ;flex-direction: row ; justify-content: left;align-items: center ; width: 850px ; height: 60px ; border: 1px solid black">
            <div style="font-size: 25px ; font-weight: bold ; width: 300px ; background-color: #98F5FF">
              交易时间 :
            </div>
            <div style="font-weight: bold;font-size: 13px">
              {{this.selecttrade.updateAt}}
            </div>
          </div>

        <button  style="width: 80px ; height: 40px ;background-color: orangered ; border-radius: 10px;margin-top: 30px" @click="this.viewDialogVisible = false">
            <div style="font-size: 20px ; font-weight: bold;color: white">
                关闭
            </div>
        </button>
      </div>
  </el-dialog>
  <el-dialog v-model="this.BuyVisible">
    <div style="display: flex ;flex-direction: column">
      <div style="display: flex;flex-direction: row;height: 100px">
          <div style="width: 70px">
          <img src="../views/Shop/addsell.svg">
          </div>
          <div style="font-size: 20px ; font-weight: bold;display: flex;justify-content: center;align-items: center">
            请输入需要出售的碳额度和期望的碳币。
          </div>
      </div>
      <div style="width: 98%;height: 2px;background-color: #72767b"/>

      <div style="display: flex;flex-direction: row;justify-content: left;align-items:center;margin-top: 40px">

        <div style="font-weight: bold;font-size: 20px">
           &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;您拥有的碳额度为 : {{this.cc}}
        </div>
        <div style="width: 40px">
          <img src="../views/Shop/tanedu48.png">
        </div>
      </div>
      <div style="display: flex;flex-direction: row;justify-content: left;align-items:center">
        <div style="font-weight: bold;font-size: 20px">
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;您拥有的碳币为 : {{this.cb}}
        </div>
        <div style="width: 40px">
          <img src="../views/Shop/tanbi32.png">
        </div>
      </div>
      <div style="background-color: #98F5FF ; border-radius: 20px  ; height: 150px ; display:flex;justify-content: center;align-items: center ; flex-direction: column;margin-top: 40px">
          <div style="display: flex;flex-direction: row ; justify-content: center ; align-items: center">
              <div style="font-size: 16px ; font-weight: bold">
              您所需要的碳币:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              </div>
             <el-input v-model="this.addtrade.carbonCoin" style="width: 100px ; margin-left: 40px"/>
              <div style="width: 40px;display: flex;justify-content: center;align-items: center;margin-left: 10px">
                <img src="../views/Shop/tanbi32.png">
              </div>
          </div>
          <div style="display: flex;flex-direction: row ; justify-content: center ; align-items: center">
            <div style="font-size: 16px ; font-weight: bold">
              您所需要出售的碳额度:
            </div>
            <el-input v-model="this.addtrade.carbonCredit" style="width: 100px ; margin-left: 40px"/>
            <div style="width: 40px;display: flex;justify-content: center;align-items: center;margin-left: 10px">
              <img src="../views/Shop/tanedu48.png">
            </div>
          </div>
      </div>

      <div style="display: flex;flex-direction: row ; justify-content: center;align-items: center;margin-top: 40px">
        <button @click="add" style="width: 100px ; height: 40px ; background-color: springgreen ; border-radius: 10px">确定出售</button>
        <div style="width: 150px"/>
        <button @click="this.BuyVisible = false" style="width: 100px ; height: 40px ; background-color: orangered ; border-radius: 10px">取消出售</button>
      </div>

    </div>
  </el-dialog>
  <div style="padding-left: 40px ; display: flex ; flex-direction: row;">
    <input type="text" placeholder="请输入需要搜索的卖方或卖方..." v-model="searchQuery" @input="handleSearch" style="width: 500px;height: 50px;background-color: white ; border-radius: 20px;padding-left: 40px">
    <img src="../views/Shop/sousuo2.svg" @click="handleSearch" style="padding-left: 20px">
  </div>
  <div  style="display: flex ;flex-direction: row;align-items: center">
    <div style="display: flex ; flex-direction: column ;">
      <div style="display: flex ; flex-direction: row">
        <div style="display: flex;justify-content: center ; align-items: center;width: 40px ; height: 100px;margin-left: 20px">
          <img src="../views/Shop/gouwuche.svg" />
        </div>
        <span style="font-weight: bold;font-size: 30px;margin-top: 30px;margin-left: 20px;display: inline-block">市场详情</span>
        <div style="width: 460px"/>
        <div style="width: 200px ; height: 60px ;border-radius: 10px;display: flex;background-color: white;box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.25);margin-top: 20px">
          <div style="display: flex ;justify-content: center;align-items: center;margin-left: 30px;">
            <div style="font-size: 20px ; font-weight: bold">
                月
            </div>
          </div>
          <div style="display: flex ; justify-content: center;align-items: center;margin-left: 100px">
            <img src="../views/Shop/shitaarrowsvg16.svg">
          </div>
        </div>
      </div>
      <div style="display: flex;padding-top: 20px ; ">
        <div style="width: 800px;background-color: #FFFFFF;height: 400px;margin-left: 15px;border-radius: 10px;display: flex;flex-direction: row ;background: linear-gradient(to bottom, #A0FF88 10%, white 15%);box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.25); ">
          <div id="BuyAndSellChart" style="width: 700px;height: 350px;padding-left: 20px;padding-top: 40px"></div>
          <div style="display: flex;flex-direction: column ; justify-content: center ; align-items: center">
            <div style="display: flex;flex-direction: row;justify-content: center;align-items: center">
              <div style="display: flex;width: 70px ; height: 40px ;background-color: #98F5FF;border-radius: 10px;justify-content: center;align-items: center">
                <img src="../views/Shop/chushousvg.svg">
              </div>

              <div style="display: flex;flex-direction: column;justify-content: center;align-items: center;width: 100px">
                <div style="font-size: 20px;color: #79bbff">
                      334
                </div>
                <div style="font-size: 10px;font-weight: bold">
                    本月交易数 / (件)
                </div>
              </div>
            </div>


            <div style="display: flex;flex-direction: row;justify-content: center;align-items: center">
              <div style="display: flex;width: 70px ; height: 40px ;background-color: #A0FF88;border-radius: 10px;justify-content: center;align-items: center">
                <img src="../views/Shop/shangjiasvg48.svg">
              </div>

              <div style="display: flex;flex-direction: column;justify-content: center;align-items: center;width: 100px">
                <div style="font-size: 20px;color: #A0FF88">
                  841
                </div>
                <div style="font-size: 10px;font-weight: bold">
                  本月上架数 / (件)
                </div>
              </div>
            </div>
          </div>


        </div>
      </div>
      <div style="width: 800px ;height: 180px ; margin-top: 6px;margin-left: 15px;display: flex;flex-direction: row; ">
          <div id="CarbonCreditInfo" style="width: 650px ; height: 180px ; background-color: white ; border-radius: 20px ; box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.25);"/>
          <div class="click-bg-image" style="margin-left: 20px ;box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.25);" @click="this.BuyVisible = true">
                <div style="font-size: 20px ; font-weight: bold ; color: black ; padding-left: 25px;padding-top: 140px">
                    发布交易
                </div>
          </div>
      </div>
    </div>

    <div style="width: 250px ;height: 700px;background-color: white;margin-left: 40px ; box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.25) ; border-radius: 10px">
        <div style="display: flex;flex-direction: row;justify-content: left">
          <div style="font-size: 25px ; font-weight: bold;color: #409eff;margin-left: 20px ;margin-top: 20px">
              用户信息
          </div>
        </div>
        <div style="width: 245px ; height: 2px ; background-color: gainsboro" />
        <div style="display: flex; flex-direction: column;justify-content: center;align-items: center;margin-top: 20px">
            <img src="../views/Shop/dianpu.svg" style="width: 130px ;height: 130px;border: 2px solid #000000;border-radius: 75px;background-color: #79bbff;">
            <div style="font-size: 23px ;font-weight: bold ;color: black">
              {{this.ID}}
            </div>
        </div>
        <div style="display: flex;flex-direction: row;justify-content: left">
          <div style="font-size: 20px ; font-weight: bold;color: #409eff;margin-left: 20px ;margin-top: 20px">
            资产详情 :
          </div>
        </div>

        <div style="display:flex;flex-direction: row;width: 200px ; height: 80px ;background-color: #BBFFFF ;border-radius: 10px;margin-left: 25px;justify-content: center;align-items: center">
              <div style="width: 80px;height: 50px;flex-direction: row">
                  <div style="display: flex;justify-content: center;align-items: center;width: 20px ; height: 20px  ;background-color: white ;border-radius: 5px">
                      <img src="../views/Shop/tanbi16.png">
                  </div>
                  <div style="display: flex;flex-direction: column ;justify-content: center;align-items: center">
                      <div style="font-size: 13px;font-weight: bold">
                          碳币:
                      </div>
                      <div style="font-size: 10px">
                          {{this.cb}}
                      </div>
                  </div>
              </div>
              <div style="width: 2px;height: 50px;background-color: #72767b;margin-left: 10px;margin-right: 10px"/>
              <div style="width: 80px;height: 50px">
                <div style="display: flex;justify-content: center;align-items: center;width: 20px ; height: 20px  ;background-color: white ;border-radius: 5px">
                  <img src="../views/Shop/tanedu16.png">
                </div>
                <div style="display: flex;flex-direction: column ;justify-content: center;align-items: center">
                  <div style="font-size: 13px;font-weight: bold">
                    碳额度:
                  </div>
                  <div style="font-size: 10px">
                    {{this.cc}}
                  </div>
                </div>
              </div>
        </div>

      <div style="display: flex ; justify-content: center ; align-items: center ; width: 230px ; height: 200px;margin-top: 50px;background-color: #F4C6F4 ; border-radius: 20px;margin-left: 10px">
          <img src="../views/Shop/online-shopping-1-20.svg">
      </div>
    </div>
  </div>


  <div style="width: 100% ;border-radius: 40px ; box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.25);margin-top: 50px">
    <div style="display: flex;flex-direction: row;justify-content: left;height: 200px;background : linear-gradient(to right, cyan, white,lightcyan), linear-gradient(to bottom, white, cyan,lightcyan);;border-top-left-radius: 20px ; border-top-right-radius: 20px">
        <img src="../views/Shop/shangcheng.svg">
      <div style="display: flex;justify-content: center;align-items: center">
          <div style="font-size: 27px ; font-weight: bolder;">
              交易市场
          </div>
      </div>
      <div style="width: 500px"/>
      <div style="display: flex;justify-content: center;align-items: center">
        <input type="text" placeholder="请输入需要搜索的卖方或卖方..." v-model="searchQuery" @input="handleSearch" style="width: 300px;height: 30px;background-color: white ; border-radius: 20px;padding-left: 40px">
        <img src="../views/Shop/sousuo32.svg">
      </div>
    </div>
    <div style="width: 100% ;height: 2px ; background-color: #72767b">
    </div>

    <div style="display: flex;flex-direction: row ;align-items: center;padding-top: 40px ; padding-left: 30px ; padding-bottom: 20px">
      <button @click="showAll"  style="background-color: darkgray ;width: 120px ; height: 60px;display: flex;justify-content: center;align-items: center;border-radius: 10px ; box-shadow: 3px 3px 3px 3px rgba(0,0,0,0.25)">
        <div style="display: flex;justify-content: center;align-items: center;font-weight: bold;font-size: 25px;color: white">
        显示全部
        </div>
      </button>

      <!-- 根据TradeStatus分类的按钮 -->
      <button @click="filterByStatus('Open')"  style="background-color: #f7ba2a;width: 120px ; height: 60px;display: flex;justify-content: center;align-items: center;margin-left: 40px ;border-radius: 10px ; box-shadow: 3px 3px 3px 3px rgba(0,0,0,0.25) ">
        <div style="display: flex;justify-content: center;align-items: center;font-weight: bold;font-size: 25px;color: white">
          可交易
        </div>
      </button>
      <button @click="filterByStatus('Closed')"  style="background-color: #67c23a ; width: 120px ; height: 60px;display: flex;justify-content: center;align-items: center;margin-left: 40px ;border-radius: 10px ; box-shadow: 3px 3px 3px 3px rgba(0,0,0,0.25)">
        <div style="display: flex;justify-content: center;align-items: center;font-weight: bold;font-size: 25px;color: white">
          已交易
        </div>
      </button>
      <div @click="filterByStatus('Canceled')" style="background-color: #72767b ; width: 120px ; height: 60px;display: flex;justify-content: center;align-items: center;margin-left: 40px ;border-radius: 10px ; box-shadow: 3px 3px 3px 3px rgba(0,0,0,0.25)">
        <div style="display: flex;justify-content: center;align-items: center;font-weight: bold;font-size: 25px;color: white">
          已取消
        </div>
      </div>


      <div style="width: 300px"/>

      <div style="width: 200px ; height: 60px ;border-radius: 10px;display: flex;background-color: white;box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.25);margin-top: 20px">
        <div style="display: flex ;justify-content: center;align-items: center;margin-left: 30px;">
          <div style="font-size: 20px ; font-weight: bold">
            按时间排序
          </div>
        </div>
        <div style="width: 40px"/>
        <div style="display: flex ; justify-content: center;align-items: center;">
          <img src="../views/Shop/shitaarrowsvg16.svg">
        </div>
      </div>

    </div>
    <div v-for="(rowDatas , Index) in chunkedAndFilteredData" key="Index" style="padding-left: 100px;padding-right: 100px;display: flex; flex-direction: row; justify-content: space-between;background-color: #DDFFFF">
    <div v-for="(datas , index) in rowDatas" key="index" style="display: flex; flex-direction: row; justify-content: space-between; padding-top: 40px ">

      <div style="width: 250px; height: 370px; background-color: #AAFFFF; border-radius: 20px; margin-right: 20px; justify-content: space-between;box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.25);">
          <div style="display: flex; flex-direction: row; justify-content: right;padding-top: 10px;padding-right: 10px">
              <div v-if="datas.TradeStatus == 'Open'" style="width:80px ;height: 40px;background-color: #f7ba2a;border-radius: 5px ;display: flex;justify-content:center;align-items: center ">
                <div style="font-size: 18px ; color: white ; font-weight: bold">
                    可交易
                </div>
              </div>
              <div v-if="datas.TradeStatus == 'Closed'" style="width:80px ;height: 40px;background-color: #67c23a;border-radius: 5px ;display: flex;justify-content:center;align-items: center ">
                <div style="font-size: 18px ; color: white ; font-weight: bold">
                  已交易
                </div>
              </div>
              <div v-if="datas.TradeStatus == 'Canceled'" style="width:80px ;height: 40px;background-color: #72767b;border-radius: 5px ;display: flex;justify-content:center;align-items: center ">
                <div style="font-size: 18px ; color: white ; font-weight: bold">
                  已取消
                </div>
              </div>
          </div>


          <div style="display: flex; flex-direction: row; justify-content: left;align-items: center;padding-top: 10px;padding-right: 10px;border: red;font-size: 15px ; color:black;font-weight: bold;padding-top: 50px">

            <img src="./Shop/tanedu48.png" style="padding-left: 20px ;padding-right: 20px">
            出售碳额度 : {{datas.CarbonCredit}}
          </div>


          <div style="display: flex; flex-direction: row; justify-content: left;align-items: center;padding-top: 10px;padding-right: 10px;border: red;font-size: 15px ; color:black;font-weight: bold;padding-bottom: 20px">

            <img src="./Shop/tanbi48.png" style="padding-left: 20px ;padding-right: 20px">
            需求碳币 : {{datas.CarbonCoin}}
          </div>
          <div style="width:250px ; height: 150px; background-color: #F0FDFD;border-radius: 20px; box-shadow: 2px 2px 2px 2px rgba(0,0,0,0.125)">
            <div style="display: flex;justify-content: center;align-items: center">
              <div style="font-size: 15px ; font-weight: bold ; color: #FFB444;margin-top: 30px">
                    卖家 : {{datas.companyID}}
              </div>
            </div>
            <div style="display: flex;justify-content: center;align-items: center">
              <div style="font-size: 20px ; font-weight: bolder;color: #67c23a">
                发布时间 : {{datas.createAt}}
              </div>
            </div>
            <div style="display: flex;justify-content: center;align-items: center">
                <div v-if="datas.TradeStatus == 'Open'" style="font-size: 20px;font-weight: bold;color: white ; width: 80px ; height: 30px ; background-color: red ; border-radius: 10px;display: flex;justify-content: center;align-items: center;box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.125);border: 2px solid white" @click="goToPage2(datas.CarbonCoin , datas.CarbonCredit , datas.companyID , this.ID ,datas.tradeID)">
                    交易
                </div>
                <div v-if="datas.TradeStatus == 'Open'" style=" width: 40px ; height: 30px ;">
                </div>
                <div style="font-size: 20px;font-weight: bold;color: white ; width: 80px ; height: 30px ; background-color: cyan ; border-radius: 10px;display: flex;justify-content: center;align-items: center;box-shadow: 5px 5px 5px 0px rgba(0,0,0,0.125);border: 2px solid white " @click="TradeSelectByID(datas.tradeID)">
                    详情
                </div>
            </div>
          </div>
      </div>
    </div>
  </div>
  </div>

</div>
</template>

<style>
.click-bg-image {
  width: 130px ;
  height: 180px ;
  border-radius: 20px;
  background-image: url("../views/Shop/clicktoshow.svg");
  background-size: cover;
  background-color: #D6FFCC;
}

</style>


<!--<>---------------->
<script>
import {ElMessage} from "element-plus";
import * as echarts from 'echarts';


var BuyAndSellChart = null
function BuyAndSellChartWays(){
  BuyAndSellChart = echarts.init(document.getElementById("BuyAndSellChart"))
  const option = {
    color: ['#DEFEFE', '#AFFC9D'],
    title: {
      text: '交易(/3月)'
    },
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
      data: ['上架数', '交易数']
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: [
      {
        type: 'category',
        boundaryGap: false,
        data: ['3月24日', '3月25日', '3月26日', '3月27日', '3月28日', '3月29日', '3月30日']
      }
    ],
    yAxis: [
      {
        type: 'value'
      }
    ],
    series: [
      {
        name: '上架数',
        type: 'line',
        stack: 'Total',
        smooth: true,
        lineStyle: {
          width: 0
        },
        showSymbol: false,
        areaStyle: {
          opacity: 0.8,
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: 'rgb(128, 255, 165)'
            },
            {
              offset: 1,
              color: 'rgb(1, 191, 236)'
            }
          ])
        },
        emphasis: {
          focus: 'series'
        },
        data: [140, 232, 101, 264, 90, 340, 250]
      },
      {
        name: '交易数',
        type: 'line',
        stack: 'Total',
        smooth: true,
        lineStyle: {
          width: 0
        },
        showSymbol: false,
        areaStyle: {
          opacity: 0.8,
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: 'rgb(0, 221, 255)'
            },
            {
              offset: 1,
              color: 'rgb(77, 119, 255)'
            }
          ])
        },
        emphasis: {
          focus: 'series'
        },
        data: [120, 282, 111, 234, 220, 340, 310]
      }
    ]
  };

  BuyAndSellChart.setOption(option );

}


const CarbonCreditInfoWays = () => {
  const CarbonCreditInfo = echarts.init(document.getElementById("CarbonCreditInfo"))

  const option = {
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    legend: {
      data: ['电气碳额度交易总量', '化工碳额度交易总量']
    },
    series: [
      {
        name: '3月份',
        type: 'pie',
        selectedMode: 'single',
        radius: [0, '30%'],
        label: {
          position: 'inner',
          fontSize: 14
        },
        labelLine: {
          show: false
        },
        data: [{ value: 3203045, name: '碳额度交易总量' }]
      },
      {
        name: '3月份',
        type: 'pie',
        radius: ['45%', '60%'],
        labelLine: {
          length: 30
        },
        label: {
          formatter: '{a|{a}}{abg|}\n{hr|}\n  {b|{b}：}{c}  {per|{d}%}  ',
          backgroundColor: '#F6F8FC',
          borderColor: '#8C8D8E',
          borderWidth: 1,
          borderRadius: 4,
          rich: {
            a: {
              color: '#6E7079',
              lineHeight: 22,
              align: 'center'
            },
            hr: {
              borderColor: '#8C8D8E',
              width: '100%',
              borderWidth: 1,
              height: 0
            },
            b: {
              color: '#4C5058',
              fontSize: 14,
              fontWeight: 'bold',
              lineHeight: 33
            },
            per: {
              color: '#fff',
              backgroundColor: '#4C5058',
              padding: [3, 4],
              borderRadius: 4
            }
          }
        },
        data: [
          { value: 1326613, name: '电气碳额度交易总量' },
          { value: 1876432, name: '化工碳额度交易总量' }
        ]
      }
    ]
  };

  CarbonCreditInfo.setOption(option)

}




export default {
  data(){
    return{
      tableData:[],
      viewDialogVisible: false,
      SelectIDView: false,
      BuyVisible:false,
      addtrade:{},
      cc:'',
      cb:'',
      selecttrade:{},
      ID:'',
      buytrade:{},
      searchQuery : '',
      it : ' ',
      currentFilter: null, // 当前过滤条件
      tradeStatuses: ['Open', 'Closed', 'Canceled'],
    }
  },
  mounted() {
    CarbonCreditInfoWays()
    if(BuyAndSellChart == null) {
      BuyAndSellChartWays()
    }
  },
  created() {
    this.getAllTrade();
    this.cc = this.$store.state.companyData.CarbonCredit;
    this.cb=this.$store.state.companyData.CarbonCoin;
    this.ID=this.$store.state.companyData.CompanyID;
  },
  computed:{
    // 根据过滤条件计算分块后的数据
    chunkedAndFilteredData() {
      if (this.currentFilter) {
        // 应用过滤条件
        const filteredData = this.tableData.filter(this.currentFilter);
        return this.chunkArray(filteredData);
      } else {
        // 显示全部数据
        return this.chunkArray(this.tableData);
      }
    },
  },
  methods:{
    showAll() {
      this.currentFilter = null;
    },

    // 根据TradeStatus过滤数据
    filterByStatus(status) {
      this.currentFilter = (item) => item.TradeStatus === status;
    },

    // 分块数组
    chunkArray(data) {
      const chunkSize = 3;
      return data.reduce((chunks, item, index) => {
        if (index % chunkSize === 0) {
          chunks.push([item]);
        } else {
          chunks[chunks.length - 1].push(item);
        }
        return chunks;
      }, []);
    },
      getAllTrade() {
      this.$http.post("http://47.97.176.174:8087/chainmaker?cmb=GetAllTransactionList").then((res) => {
        // 获取到后端数据后，遍历数据并处理订单生成日期
        res.data.data.forEach(item => {
          // 去除订单生成日期中的时区信息，并只保留日期部分
          item.createAt = this.fetchBackendDate(item.createAt);
        });
        // 逆序数组数据
        this.tableData = res.data.data.reverse();
        console.log("TableData : " , this.tableData)
      }).catch((err) => {
        console.log(err);
      });
    },
    fetchBackendDate(dateWithTimeZone) {
      // 解析带时区的日期字符串并提取年月日部分
      const backendDate = new Date(dateWithTimeZone);
      // 使用toISOString方法获取ISO格式的日期字符串，然后通过split方法获取年月日部分
      return backendDate.toISOString().split('T')[0]; // 返回格式化后的日期字符串（仅包含年月日）
    },
    add() {
      const ye=parseFloat(this.cc);
      const cs=parseFloat(this.addtrade.carbonCredit)
      if(this.addtrade.carbonCredit == null || this.addtrade.carbonCoin == null){
        ElMessage.error('请完整填写交易信息');
      }else {
        if (cs <= ye){
          try {
            this.$http.post(
                "http://47.97.176.174:8087/chainmaker?cmb=InitiateTransaction", {
                  companyID:this.ID,
                  carbonCredit:this.addtrade.carbonCredit,
                  carbonCoin:this.addtrade.carbonCoin,
                }
            ).then(response => {
              // 添加交易成功后更新表格数据
              this.getAllTrade();
              this.viewDialogVisible = false;
              this.addtrade = {};
              ElMessage({
                message: '发布交易成功',
                type: 'success',
              })
            });
          } catch (error) {
            console.error("Error adding new user:", error);
          }
        }else {
          ElMessage.error('发布失败，额度不足');
        }
      }
      this.BuyVisible = false
    },
    addnewtrade(){
      this.viewDialogVisible=true;
    },
    closeViewDialog() {
      this.viewDialogVisible = false;
      this.addtrade = {};
    },
    closeViewDialog2() {
      this.SelectIDView = false;
      this.selecttrade = {};
    },
    TradeSelectByID(id) {
      this.viewDialogVisible = true
      this.$http.post("http://47.97.176.174:8087/chainmaker?cmb=GetTransactionInfoByTradeId", {
        tradeId: id,
      }).then((res) => {
        const transactionInfo = res.data.data;
        // 处理时间字符串中的时区信息
       // transactionInfo.createAt = this.formatDate(transactionInfo.createAt);
        this.selecttrade = transactionInfo;
      });
    },
    formatDate(dateTimeString) {
      // 创建 Date 对象并获取年月日部分
      const date = new Date(dateTimeString);
      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, '0');
      const day = String(date.getDate()).padStart(2, '0');
      // 返回格式化后的日期字符串（年-月-日）
      return `${year}-${month}-${day}`;
    },
    getStatusColor(status) {
      // 根据状态返回相应的颜色
      if(status === 'Open'){
        return 'lightblue'
      }else if (status === 'Closed'){
        return 'green'
      }else if(status === 'Canceled'){
        return 'orange'
      }else {
        return 'gray';
      }
    },
    Select(id){
      this.TradeSelectByID(id);
      this.SelectIDView = true;
    },
    getText(text){
      if(text === 'Open'){
        return '可交易'
      }else if (text === 'Closed'){
        return '已交易'
      }else if(text === 'Canceled'){
        return '已取消'
      }else {
        return 'gray';
      }
    },
    BeforeBuy(row){
      if (row.companyID == this.ID){
        ElMessage.error("不可购买自己发布的交易！")
      }else {
        this.buytrade=row;
        this.BuyVisible=true;
      }
    },
    Buy(){
      if(this.buytrade.CarbonCoin > this.cb){
        ElMessage.error("余额不足！")
      }else {
        this.$http.post('http://47.97.176.174:8087/chainmaker?cmb=AToBTransaction',{
          tradeID:this.buytrade.tradeID,
          sellerID:this.buytrade.companyID,
          buyerID:this.ID,
          carbonCredit:this.buytrade.CarbonCredit.toString(),
          carbonCoin:this.buytrade.CarbonCoin.toString(),
        }).then((res)=>{
          this.getAllTrade();
          this.buytrade={};
          this.BuyVisible=false;
          ElMessage({
            message: '购买成功',
            type: 'success',
          })
        })
      }
    },
    cancelBuy(){
      this.BuyVisible=false;
    },
    handleSearch() {
      // 这里可以添加额外的搜索逻辑，比如发送请求到服务器搜索等
      console.log('搜索内容：', this.searchQuery);
    },
    goToPage2(CarbonCoin , CarbonCredit , Seller ,Buyer , TradeId) {

      if(this.cb < CarbonCoin){
        console.log(this.cb , "   " , CarbonCoin)
        ElMessage.error("用户余额不足，此交易不可进行")
      } else if (this.ID ==  Seller){
        ElMessage.error("不可对自己进行交易")
      }else {
        const jsonData = {
          CarbonCoin: CarbonCoin,
          CarbonCredit: CarbonCredit,
          Seller: Seller,
          Buyer: Buyer,
          TradeId: TradeId
        }; // 您的JSON数据
        const queryString = encodeURIComponent(JSON.stringify(jsonData)); // 将JSON对象转换为字符串并编码
        this.$router.push({name: 'TransactionDetail', query: {data: queryString}}); // 导航到第二个页面并传递查询参数
      }
    }
  }
}
</script>
