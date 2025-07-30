<template>
  <div class="home-page">
    <!-- 个性化欢迎区域 -->
    <PersonalizedWelcome />
    
    <!-- 学习路径选择区域（仅学徒身份显示） -->
    <LearningPathSelector v-if="authStore.isApprentice" />
    
    <!-- 推荐大师展示区域（仅学徒身份显示） -->
    <RecommendedMentors v-if="authStore.isApprentice" />
    
    <!-- 推荐课程展示区域（仅学徒身份显示） -->
    <RecommendedCourses v-if="authStore.isApprentice" />
           
    <!-- 推荐圈子展示区域 -->
    <RecommendedCircles />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import PersonalizedWelcome from '@/components/home/PersonalizedWelcome.vue'
import LearningPathSelector from '@/components/home/LearningPathSelector.vue'
import RecommendedMentors from '@/components/home/RecommendedMentors.vue'
import RecommendedCourses from '@/components/home/RecommendedCourses.vue'
import RecommendedCircles from '@/components/home/RecommendedCircles.vue'

const authStore = useAuthStore()
</script>

<style scoped lang="scss">
.home-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: var(--spacing-xl);
}

// 响应式设计
@media (max-width: 768px) {
  .home-page {
    padding: var(--spacing-lg);
  }
}
</style>