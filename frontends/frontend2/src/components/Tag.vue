<script setup>
import { ref, watch, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const tags = ref([]);
const router = useRouter();
const route = useRoute();

watch(tags, (newTags) => {
  localStorage.setItem('tags', JSON.stringify(newTags));
}, { deep: true });

watch(() => route.path, (to) => {
  const tag = { name: route.name, label: route.meta.title };
  const isTagExist = tags.value.some((t) => t.name === tag.name);
  if (!isTagExist) {
    tags.value.push(tag);
  }
});

onMounted(() => {
  const loadedTags = JSON.parse(localStorage.getItem('tags') || '[]');
  if (loadedTags.length) {
    tags.value = loadedTags;
  } else {
    tags.value.push({ path: '/Home', name: 'home', label: '首页' });
  }
});

const handleClose = (tag, index) => {
  if (tag.name !== 'home') {
    tags.value.splice(index, 1);
  }
  if (tag.name !== route.name) {
    return;
  }
  const length = tags.value.length;
  if (index === length) {
    router.push({ name: tags.value[index - 1].name });
  } else {
    router.push({ name: tags.value[index].name });
  }
};

const changeMenu = (tag) => {
  if (tag.name === 'home') {
    if (route.path !== '/Home') {
      router.push({ path: '/Home' });
    }
    return;
  }
  if (tag.name !== 'home') {
    const toPath = router.resolve({ name: tag.name }).href;
    if (route.path !== toPath) {
      router.push({ name: tag.name });
    }
  }
};

const isActive = (tag) => {
  return tag.name === route.name;
};
</script>

<template>
  <div style="margin: 5px 15px">
    <el-scrollbar>
      <div style="height: 50px; width: 770px; display: flex; justify-content: left; align-items: center;">
        <el-tag
            v-for="(tag, index) in tags"
            :key="index"
            :closable="tag.name !== 'home'"
            :disable-transitions="false"
            @close="handleClose(tag, index)"
            @click="changeMenu(tag)"
            :effect="isActive(tag) ? 'dark' : 'plain'"
            size="large">
          {{ tag.label }}
        </el-tag>
      </div>
    </el-scrollbar>
  </div>
</template>

<style scoped>
.el-tag {
  margin-right: 10px;
  border: 0;
  cursor: pointer;
}
</style>
