<template>
  <div>
    <div><carousel :items="carouselItems" :interval="5000" style="width: 100%;margin-top: 5px"></carousel></div>
  </div>
    <el-tabs type="border-card">
      <el-tab-pane label="每日资讯">
        <el-card shadow="never" >
          <template #header>
            <text style="font-size: 30px">每日资讯</text>
          </template>
          <div class="click" style="display:flex;background-color: #FFFFFF;width: 100%;height: 220px;" @click="navigateToDocumentPage">
            <img src="@/assets/teach/xinwen1.jpeg"height="220" width="220"/>
            <div>
              <p style="font-weight: bold;font-size: 20px;margin-top: 5px;margin-left: 5px">紧急！钢铁、石化等7个行业碳排放核算报告核查工作已开展</p>
              <p style="margin-top: 15px;margin-left: 5px">2024年初，随着全国温室气体自愿减排交易市场上线和《碳排放权交易管理暂行条例》正式对外公布，全国碳市场目前只剩下一个问题待解：什么时候启动扩围？</p>
              <div style="display: flex;margin-top: 60px;margin-left: 5px">
                <p style="color: #999;">发布时间：2024-03-21</p>
                <img src="@/assets/teach/yanjing.svg" height="25" width="25" style="margin-left: 20px">
                <p style="margin-left: 5px;color: #999">13</p>
              </div>
            </div>
          </div>
          <el-divider />
          <div style="display:flex;background-color: #FFFFFF;width: 100%;height: 220px;">
            <img src="@/assets/teach/xinwen2.jpeg"height="220" width="220"/>
            <div>
              <p style="font-weight: bold;font-size: 20px;margin-top: 5px;margin-left: 5px">从“漂绿”到“绿色沉默”，碳核算制度如何助力行业走出两难困境</p>
              <p style="margin-top: 15px;margin-left: 5px">3月15日是国际消费者权益日，每年的今天世界各地有关组织都会通过多种形式，探访消费市场中存在的各种违法侵权行为，加强消费者权益保护。其实，在“双碳”领域，有些企业也会存在虚假的绿色营销等行为，影响消费者体验和正当权益。</p>
              <div style="display: flex;margin-top: 60px;margin-left: 5px">
                <p style="color: #999;">发布时间：2024-03-20</p>
                <img src="@/assets/teach/yanjing.svg" height="25" width="25" style="margin-left: 20px">
                <p style="margin-left: 5px;color: #999">120</p>
              </div>
            </div>
          </div>
          <el-divider />
          <div style="display:flex;background-color: #FFFFFF;width: 100%;height: 220px;">
            <img src="@/assets/teach/xinwen3.jpeg"height="220" width="220"/>
            <div>
              <p style="font-weight: bold;font-size: 20px;margin-top: 5px;margin-left: 5px">碳市场扩围紧锣密鼓开展，相关核算核查技术指南征求意见</p>
              <p style="margin-top: 15px;margin-left: 5px">全国碳市场扩大覆盖范围的准备工作正紧锣密鼓地开展，眼下，钢铁、有色金属、建材等已纳入全国碳市场序列的重点排放行业正开展碳排放报告和核查等准备，相关核算与报告指南、核查技术指南也已开始征求各方意见。
              </p>
              <div style="display: flex;margin-top: 60px;margin-left: 5px">
                <p style="color: #999;">发布时间：2024-03-20</p>
                <img src="@/assets/teach/yanjing.svg" height="25" width="25" style="margin-left: 20px">
                <p style="margin-left: 5px;color: #999">10</p>
              </div>
            </div>
          </div>
          <el-divider />
          <div style="display:flex;background-color: #FFFFFF;width: 100%;height: 220px;">
            <img src="@/assets/teach/xinwen4.jpeg"height="220" width="220"/>
            <div>
              <p style="font-weight: bold;font-size: 20px;margin-top: 5px;margin-left: 5px">生态环境分区管控助力高质量发展—零碳研究院碳报</p>
              <p style="margin-top: 15px;margin-left: 5px">2020年9月，我国明确提出碳达峰、碳中和目标。2022年6月15日，新京报成立零碳研究院，如何准确理解和把握双碳政策趋势？碳中和背景下，不同产业和企业面临哪些机遇和挑战？研究院于2022年6月起推出《碳报》，研究最新双碳政策、权威声音、低碳样本等重点内容，并进行分析解读。
              </p>
              <div style="display: flex;margin-top: 60px;margin-left: 5px">
                <p style="color: #999;">发布时间：2024-03-19</p>
                <img src="@/assets/teach/yanjing.svg" height="25" width="25" style="margin-left: 20px">
                <p style="margin-left: 5px;color: #999">1452</p>
              </div>
            </div>
          </div>
          <el-divider />
          <div style="display:flex;background-color: #FFFFFF;width: 100%;height: 220px;">
            <img src="@/assets/teach/xinwen5.jpeg"height="220" width="220"/>
            <div>
              <p style="font-weight: bold;font-size: 20px;margin-top: 5px;margin-left: 5px">做好“预算”管理 积极稳妥推进“双碳”</p>
              <p style="margin-top: 15px;margin-left: 5px">随着我国构建完成碳达峰碳中和“1+N”政策体系，“双碳”工作已经由顶层设计阶段步入了实质性推进阶段，有必要加紧研究部署具体的制度落实举措。碳预算制度是落实中共中央部署的推进能耗“双控”向碳排放总量和强度“双控”转变的必要创新性制度，碳预算管理是预算制度在碳排放管理上的应用。制定碳排放年度目标和碳预算方案，有助于实现对地区和重点领域碳排放的灵活管控和高效调控，有助于更好统筹短期与中长期、发展与减排的关系，为积极稳妥推进碳达峰碳中和目标提供有力支撑。
              </p>
              <div style="display: flex;margin-top: 60px;margin-left: 5px">
                <p style="color: #999;">发布时间：2024-03-19</p>
                <img src="@/assets/teach/yanjing.svg" height="25" width="25" style="margin-left: 20px">
                <p style="margin-left: 5px;color: #999">2380</p>
              </div>
            </div>
          </div>
          <el-divider />
          <div style="display:flex;background-color: #FFFFFF;width: 100%;height: 220px;">
            <img src="@/assets/teach/xinwen6.jpeg"height="220" width="220"/>
            <div>
              <p style="font-weight: bold;font-size: 20px;margin-top: 5px;margin-left: 5px">强化顶层设计，建立统一 规范的碳排放统计核算体系</p>
              <p style="margin-top: 15px;margin-left: 5px">当前，世界百年未有之大变局加速演进，气候变化带来的不仅是海平面上升和更多极端天气等环境问题，也带来了大国博弈的政治和经济问题。建设碳核算体系不仅关乎环境治理，也关乎大国博弈下的经济社会发展以及能源转型。
              </p>
              <div style="display: flex;margin-top: 60px;margin-left: 5px">
                <p style="color: #999;">发布时间：2024-03-19</p>
                <img src="@/assets/teach/yanjing.svg" height="25" width="25" style="margin-left: 20px">
                <p style="margin-left: 5px;color: #999">3578</p>
              </div>
            </div>
          </div>
          <el-divider />
          <div style="display:flex;background-color: #FFFFFF;width: 100%;height: 220px;">
            <img src="@/assets/teach/xinwen7.jpeg"height="220" width="220"/>
            <div>
              <p style="font-weight: bold;font-size: 20px;margin-top: 5px;margin-left: 5px">报告解读｜2024政府工作报告释放的十大“双碳”信号</p>
              <p style="margin-top: 15px;margin-left: 5px">关于积极稳妥推进碳达峰碳中和，报告称，积极稳妥推进碳达峰碳中和。扎实开展“碳达峰十大行动”。提升碳排放统计核算核查能力，建立碳足迹管理体系，扩大全国碳市场行业覆盖范围。深入推进能源革命，控制化石能源消费，加快建设新型能源体系。加强大型风电光伏基地和外送通道建设，推动分布式能源开发利用，发展新型储能，促进绿电使用和国际互认，发挥煤炭、煤电兜底作用，确保经济社会发展用能需求。
              </p>
              <div style="display: flex;margin-top: 60px;margin-left: 5px">
                <p style="color: #999;">发布时间：2024-03-19</p>
                <img src="@/assets/teach/yanjing.svg" height="25" width="25" style="margin-left: 20px">
                <p style="margin-left: 5px;color: #999">5876</p>
              </div>
            </div>
          </div>
          <el-divider />
          <div style="display:flex;background-color: #FFFFFF;width: 100%;height: 220px;">
            <img src="@/assets/teach/xinwen8.jpeg"height="220" width="220"/>
            <div>
              <p style="font-weight: bold;font-size: 20px;margin-top: 5px;margin-left: 5px">欧盟新电池法引行业热议：碳足迹核算体系需国际互认 电池护照实施将面临三大挑战</p>
              <p style="margin-top: 15px;margin-left: 5px">欧洲是我国电池产品出口的主要市场。据海关总署公布数据，今年一季度，中国锂电池出口额排名前五的国家依次是美国、德国、韩国、荷兰和越南，占总出口额的62.6%。其中，中国出口至德国和荷兰两个欧洲国家的总金额为272.97亿元，超过了出口至美国的224.7亿元。
              </p>
              <div style="display: flex;margin-top: 60px;margin-left: 5px">

                <p style="color: #999;">发布时间：2024-03-18</p>
                <img src="@/assets/teach/yanjing.svg" height="25" width="25" style="margin-left: 20px">
                <p style="margin-left: 5px;color: #999">10593</p>
              </div>
            </div>
          </div>


        </el-card>
      </el-tab-pane>
      <el-tab-pane label="教学视频">
        <el-card shadow="never" >
          <template #header>
            <text style="font-size: 30px">教学视频</text>
          </template>
         <div style="height: 280px">
           <div style="background-color: rgba(255, 255, 255, 0.5);display: flex;margin-top: 20px">
             <div style="background-color: #FFFFFF;border-radius: 10px;border: 1px solid #000000;height: 230px;width: 340px" @click="navigateToDocumentPage2">
               <div class="video-preview" style="margin-left: 12px">
                 <img src="@/assets/teach/pit1.png"height="160" width="300" style="margin-top: 10px;margin-left: 4px"/>
               </div>
               <div style="background-color: black;border: 1px solid black"></div>
               <span style="font-size: 14px;font-weight: bold;margin-left: 5px;margin-top: 5px">科普动画丨碳排放是如何进行核算的，企业的碳排放</span><br>
               <span style="font-size: 14px;font-weight: bold;margin-left: 5px">又是怎么计算的呢?</span>
             </div>
             <div style="background-color: #FFFFFF;border-radius: 10px;border: 1px solid #000000;height: 230px;width: 340px;margin-left: 60px">
               <div class="video-preview" style="margin-left: 12px">
                 <img src="@/assets/teach/pit2.png"height="160" width="300" style="margin-top: 10px;margin-left: 4px"/>
               </div>
               <div style="background-color: black;border: 1px solid black"></div>
               <span style="font-size: 14px;font-weight: bold;margin-left: 15px;margin-top: 5px;display: inline-block">电气模型教学</span><br>
             </div>
             <div style="background-color: #FFFFFF;border-radius: 10px;border: 1px solid #000000;height: 230px;width: 340px;margin-left: 60px">
               <div class="video-preview" style="margin-left: 12px">
                 <img src="@/assets/teach/pit3.png"height="160" width="300" style="margin-top: 10px;margin-left: 4px"/>
               </div>
               <div style="background-color: black;border: 1px solid black"></div>
               <span style="font-size: 14px;font-weight: bold;margin-left: 15px;margin-top: 5px;display: inline-block">化工模型教学</span><br>
             </div>
           </div>
         </div>
        </el-card>
      </el-tab-pane>

    </el-tabs>

</template>

<script>
import Carousel from './Carousel.vue';
import p1 from '@/assets/teach/p1.jpg'
import p2 from '@/assets/teach/p2.jpeg'
import p3 from '@/assets/teach/p3.jpeg'
import p4 from '@/assets/teach/p4.jpeg'
export default {
  name: 'App',
  components: {
    Carousel
  },
  data() {
    return {
      carouselItems: [
        { imageUrl: p1 },
        { imageUrl: p2 },
        { imageUrl: p3 },
        { imageUrl: p4 }
      ]
    };
  },
  methods: {
    navigateToDocumentPage() {
      this.$router.push('/TeachMain');
    },
    navigateToDocumentPage2() {
      this.$router.push('/TeachMain2');
    }
  }
};
</script>
<style>
.custom-link {
  text-decoration: none; /* 隐藏默认下划线 */
  position: relative; /* 使得伪元素的绝对定位相对于超链接 */
  color: #000; /* 设置超链接颜色 */
}

.custom-link::after {
  content: ""; /* 伪元素内容为空 */
  position: absolute; /* 绝对定位 */
  bottom: -3px; /* 距离超链接底部的距离 */
  left: 0; /* 与超链接左对齐 */
  width: 100%; /* 宽度与超链接相同 */
  height: 1px; /* 下划线高度 */
  background-color: #000; /* 下划线颜色 */
  transition: width 0.3s; /* 添加过渡效果 */
}

.custom-link:hover::after {
  width: 50%; /* 鼠标悬停时下划线宽度变为 50% */
  background-color: #FF0000; /* 鼠标悬停时下划线颜色变为红色 */
}
</style>
