<template>
  <div class="embedded-iframe">
    <iframe ref="iframeRef" :src="externalUrl" frameborder="0" width="100%" height="100%" style="overflow: hidden"></iframe>
  </div>
</template>

<script setup>
import {onMounted, ref} from "vue";
import {useRoute} from "vue-router";

const route = useRoute()

const externalUrl = ref('')
const iframeRef = ref(null)

// 根据路由参数获取外部链接地址
const fetchExternalUal = () => {
  const { params } = route
  if (params.url){
    externalUrl.value = decodeURIComponent(params.url)
  }
}

onMounted(()=>{
  fetchExternalUal()
})
</script>

<style scoped>
.embedded-iframe {
  width: 1600px; /* iframe宽度填满容器 */
  height: 950px; /* iframe高度填满容器 */
  position: absolute; /* 脱离文档流以应用transform */
  margin: 20px 10px 0px 15px;
  transform-origin: 0 0; /* 设置缩放的基准点 */
  transform: scale(0.71); /* 缩小到原来尺寸的0.8 */
}
</style>