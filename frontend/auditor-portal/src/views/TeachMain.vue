<template>
  <el-button type="primary" @click="back()" style="margin-left: 1100px">返回</el-button>
  <el-card shadow="never" style="margin-top: 10px">
  <template #header>
    <text style="font-size: 40px;font-weight: bold">碳市场再次释放扩围预期：钢铁、石化等7个行业碳排放核算报告核查工作已开展</text>
    <div style="display: flex;margin-top: 5px;margin-left: 5px">
      <span>华夏日报</span>
      <span style="color: #999;margin-left: 30px">2024-03-21</span>
      <span style="color: #999;margin-left: 5px">17:21</span>
    </div>
  </template>
    <img src="@/assets/teach/xinwen1.jpeg"height="500" width="500" style="margin-left: 300px"/>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">本报（chinatimes.net.cn）记者刘诗萌 北京报道</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">2024年初，随着全国温室气体自愿减排交易市场上线和《碳排放权交易管理暂行条例》（下称《条例》）正式对外公布，全国碳市场目前只剩下一个问题待解：什么时候启动扩围？</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">“我们正在积极推动，争取尽快实现我国碳排放权交易市场的首次扩围。”2月26日，在国新办举行的《碳排放权交易管理暂行条例》政策例行吹风会上，生态环境部副部长赵英民指出，中国的碳排放主要集中在发电、钢铁、建材、有色、石化、化工、造纸、航空等重点行业，除了电力行业之外，其他7个行业虽然没有纳入配额管控，但是其碳排放核算报告核查工作已经开展起来了。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">2021年7月，全国碳排放权交易市场启动，首批被纳入的约2200家重点排放单位均属于电力行业。目前碳市场已经完成了两个履约周期，第一个履约周期是2019—2020年，第二个履约周期是2021、2022年，但尚未进行首次扩围。事实上，自从碳市场启动以来，何时开始扩围始终是一个备受外界关注的话题。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">伦敦证券交易所集团碳分析师宋雨彤此前接受《华夏时报》记者采访时表示，预计2024年全国碳市场很可能会出现扩围动作，水泥、电解铝、航空等重点行业都有可能被纳入。</span>
    <span style="display: inline-block;font-size: 20px;font-weight: bold;margin-top: 20px">减碳成本降低约350亿元</span>
    <span style="display: inline-block;font-size: 20px">作为全球覆盖温室气体排放量最大的碳市场，中国的碳排放权交易市场在过去的两个履约周期中表现如何？</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">《华夏时报》记者从生态环境部获悉，目前全国碳排放权交易市场覆盖年二氧化碳排放量约51亿吨，占到了全国二氧化碳排放总量的40%以上，纳入重点排放单位2257家。市场活跃度方面，第二个履约周期有明显提升：截至2023年底，全国碳排放权交易市场累计成交量达到4.4亿吨，成交额约249亿元。第二个履约周期成交量比第一个履约周期增长了19%，成交额比第一个履约周期增长了89%。并且，碳价整体呈现平稳上涨态势。由启动时的每吨48元上涨至每吨80元左右，上涨66%左右。第二个履约周期企业参与交易的积极性明显提升，参与交易的企业占总数的82%，比第一个履约周期上涨了近50%。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">赵英民表示，全国强制碳市场启动两年半以来，总体运行平稳，制度规范日趋完善，市场活跃度逐步提升，碳排放数据质量全面改善，碳排放管理能力明显提升，价格发现机制作用日益显现。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">数据显示，碳市场确实降低了行业和全社会的减碳成本。据测算，这两个履约周期，全国电力行业总体减排成本降低了约350亿元。这意味着，通过碳排放配额交易，碳市场为企业履行减碳责任提供了更为灵活的选择，帮助行业实现了低成本的减碳。随着碳排放权交易市场覆盖行业范围不断扩大，碳排放资源在全国范围内不同行业间的优化配置将最终实现全国总的减排成本最小化。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">在促进企业减排温室气体、推动行业绿色低碳转型高质量发展的同时，全国碳市场也为全社会开展气候投融资、碳资产管理等碳定价活动锚定了基准价格，推动社会各界关注应对气候变化工作，积极参与降碳、减污、扩绿、增长，推动生产生活方式的低碳化、绿色化，从而推动全社会的绿色低碳发展。</span>
    <span style="display: inline-block;font-size: 20px;font-weight: bold;margin-top: 20px">2024年能否实现扩围？</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">在上述发布会上，赵英民也回应了碳市场扩围这一热点问题。他指出，发电、钢铁、建材、有色、石化、化工、造纸、航空这八个行业占我国二氧化碳排放75%左右，这些重点行业工业化程度高，有一定的人才、技术、管理基础，更容易实现对碳排放的量化控制管理和影响含碳产品和服务的价格。将上述高排放行业尽早纳入全国碳排放权交易市场，就抓住了全国75%的排放。充分发挥市场在碳排放资源配置中的决定作用，可以使全社会的降碳成本实现最优、最小化，从而助力实现我国的双碳目标，推动绿色低碳转型和美丽中国建设。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">记者了解到，过去几年间，生态环境部针对扩围主要开展了两项工作：一是在全国范围内对上述重点行业组织开展了年度的碳排放核算报告核查工作；二是开展扩围的专项研究，对重点行业的配额分配方法、核算报告方法、核算要求指南、扩围实施路径等，开展了专题研究评估论证，相关的技术文件起草工作已经基本完成。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">过去一年间，生态环境部环境规划院多次举办碳市场扩围相关的学术研讨。5月召开了全国碳市场扩大行业覆盖范围专项研究启动会，6月召开了石化行业纳入全国碳市场专项研究第一次工作会议，7月、9月举行了两期全国碳市场扩围关键技术学术沙龙，首期聚焦“间接排放是否应纳入全国碳市场”议题，第二期聚焦水泥、钢铁、电解铝等行业配额盈缺率、清缴履约上限、绿电在配额与核算中的体现、核算核查方法等问题。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px">“扩围工作将把握好节奏力度，科学合理确定不同行业的纳入时间，分阶段、有步骤地积极推动碳排放权交易市场覆盖碳排放重点行业，从而构建更加有效、更有活力、更具国际影响力的碳市场。”赵英民说。</span>
    <span style="display: inline-block;font-size: 20px;margin-top: 20px;margin-left: 800px">责任编辑：徐芸茜 主编：公培佳</span>
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
