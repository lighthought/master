<template>
  <div class="profile-page">
    <div class="container">
      <h1 class="page-title">个人中心</h1>
      <p class="page-description">管理个人信息和学习记录</p>
    </div>
  </div>
</template>

<script setup lang="ts">
// 个人中心页面组件
</script>

<style scoped lang="scss">
.profile-page {
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