<template>
  <el-carousel :interval="interval" arrow="always">
    <el-carousel-item v-for="(item, index) in items" :key="index">
      <img :src="item.imageUrl" style="width: 100%; height: 300px; object-fit: cover;" alt="">
    </el-carousel-item>
  </el-carousel>
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'Carousel',
  props: {
    items: {
      type: Array,
      required: true
    },
    interval: {
      type: Number,
      default: 3000
    }
  }
}
</script>
