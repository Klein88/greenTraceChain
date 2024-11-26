<template>
  <div style="display: flex; flex-direction: column; margin-top: 20px">
    <DataTable :value="codes" show-gridlines table-style="min-width: 50rem" style="margin-left: 20px; margin-right: 20px">
      <Column field="id" header="序号" style="width: 70px"></Column>
      <Column field="error" header="错误" style="width: 195px"></Column>
      <Column field="detail" header="详情"></Column>
    </DataTable>
  </div>
</template>

<script setup>
import {ref} from "vue";

const codes = ref([
  {
    id: 1,
    error: '发送资产至错误的地址',
    detail: '用户在转账或支付时，可能因为输入错误的接收地址导致资产发送到错误的地址，导致资产丢失。'
  },
  {
    id: 2,
    error: '忘记密码或私钥',
    detail: '用户如果忘记了钱包密码或私钥，将无法访问其资产，这可能会导致无法恢复资金的情况。'
  },
  {
    id: 3,
    error: '未备份钱包或私钥',
    detail: '如果用户未及时备份钱包或私钥，一旦设备损坏或丢失，将导致资产永久丢失。'
  },
  {
    id: 4,
    error: '交易延迟或被拒绝',
    detail: '由于网络拥堵或手续费设置不当，用户的交易可能会被延迟或被网络拒绝。'
  },
  {
    id: 5,
    error: '未及时升级客户端',
    detail: '如果用户未及时升级钱包客户端或节点软件，可能会因为安全漏洞或协议更新而受到攻击或无法与网络同步。'
  },
  {
    id: 6,
    error: '点击恶意链接',
    detail: '用户可能会收到钓鱼邮件或恶意链接，误点击后导致私钥泄露或恶意软件安装，从而造成资产损失。'
  },
  {
    id: 7,
    error: '未确认交易即关闭钱包',
    detail: '在交易尚未被确认时，用户可能关闭了钱包或离开了交易页面，导致交易未能成功完成。'
  },
  {
    id: 8,
    error: '忽略警告信息',
    detail: '在进行重要操作时，用户可能会忽略钱包或交易平台的警告信息，导致意外操作或资产丢失。'
  },
  {
    id: 9,
    error: '受到网络攻击',
    detail: '用户的钱包、交易所或其他相关服务可能会成为黑客攻击的目标，导致用户资产被盗或服务中断。'
  },
  {
    id: 10,
    error: '社会工程攻击',
    detail: '黑客可能会利用社会工程学手段，通过欺骗、诈骗等手段获取用户的私钥或敏感信息。'
  },
  {
    id: 11,
    error: '非法活动风险',
    detail: '用户参与的项目可能涉及非法或欺诈活动，导致用户资产受到损失。'
  },
  {
    id: 12,
    error: '跨链操作错误',
    detail: '在进行跨链操作时，用户可能由于操作失误导致资产丢失或被锁定。'
  },
  {
    id: 13,
    error: '监管风险',
    detail: '用户可能因为违反当地法规或监管要求而面临处罚或资产被冻结的风险。'
  }
])
</script>

<style scoped>

</style>