<template>
  <div class="courses-page">
    <div class="container">
      <h1 class="page-title">课程页面</h1>
      <p class="page-description">浏览和报名专业课程</p>
    </div>
  </div>
</template>

<script setup lang="ts">
// 课程页面组件
</script>

<style scoped lang="scss">
.courses-page {
  padding: var(--spacing-xl) 0;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 var(--spacing-lg);
}

.page-title {
  font-size: var(--font-size-h1);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-md);
}

.page-description {
  font-size: var(--font-size-large);
  color: var(--text-secondary);
}
</style>