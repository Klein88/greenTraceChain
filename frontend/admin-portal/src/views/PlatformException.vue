<template>
  <div style="display: flex; flex-direction: column; margin-top: 20px; margin-left: 20px; margin-right: 20px">
    <DataTable :value="codes" show-gridlines table-style="min-width: 50rem">
      <Column field="code" header="HTTP状态码"></Column>
      <Column field="error" header="错误异常"></Column>
    </DataTable>
    <DataTable :value="exceptions" show-gridlines table-style="min-width: 50rem" style="margin-top: 30px">
      <Column field="code" header="JavaScript异常"></Column>
      <Column field="error" header="错误异常"></Column>
    </DataTable>
  </div>
</template>

<script setup>
import {ref} from "vue";

const codes = ref([
  {
    code: '404 Not Found',
    error: '客户端请求的资源在服务器上不存在。'
  },
  {
    code: '400 Bad Request',
    error: '服务器无法理解客户端的请求，通常是由于客户端发送的请求格式错误或包含无效参数。'
  },
  {
    code: '401 Unauthorized',
    error: '客户端请求需要身份验证，但未提供有效的认证信息。'
  },
  {
    code: '403 Forbidden',
    error: '服务器拒绝客户端的请求，通常是由于权限不足。'
  },
  {
    code: '500 Internal Server Error',
    error: '服务器在执行请求时发生了意外错误，无法完成请求。这可能是由于代码错误、数据库连接问题等引起的。'
  },
  {
    code: '502 Bad Gateway',
    error: '服务器作为网关或代理，从上游服务器接收到无效的响应。'
  },
  {
    code: '503 Service Unavailable',
    error: '服务器暂时无法处理请求，通常是由于服务器过载或维护而造成的。'
  },
  {
    code: '504 Gateway Timeout',
    error: '服务器作为网关或代理，未能及时从上游服务器接收到响应。'
  }
])
const exceptions = ref([
  {
    code: 'TypeError',
    error: '类型错误，通常是由于对一个值的类型进行了不支持的操作。'
  },
  {
    code: 'ReferenceError',
    error: '引用错误，通常是由于试图访问未定义的变量或函数。'
  },
  {
    code: 'SyntaxError',
    error: '语法错误，通常是由于 JavaScript 代码的语法错误导致的。'
  },
  {
    code: 'NetworkError',
    error: '网络错误，通常是由于网络连接问题或跨域请求受阻等原因引起的。'
  }
])
</script>

<style scoped>

</style>