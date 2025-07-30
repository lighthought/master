<template>
  <div class="home-page">
    <!-- 欢迎区域 -->
    <section class="welcome-section">
      <div class="container">
        <div class="welcome-content">
          <h1 class="welcome-title">
            <span class="title-main">Master Guide</span>
            <span class="title-sub">大师指导平台</span>
          </h1>
          <p class="welcome-description">
            连接各领域大师与学习者，提供专业技艺培养计划、教学课程和实时线上指导服务
          </p>
          <div class="welcome-actions">
            <el-button type="primary" size="large" class="action-btn">
              <el-icon><User /></el-icon>
              开始学习
            </el-button>
            <el-button size="large" class="action-btn secondary">
              <el-icon><Star /></el-icon>
              成为大师
            </el-button>
          </div>
        </div>
      </div>
    </section>

    <!-- 学习路径区域 -->
    <section class="learning-paths-section">
      <div class="container">
        <h2 class="section-title">选择你的学习路径</h2>
        <div class="learning-paths-grid">
          <div class="path-card" v-for="path in learningPaths" :key="path.id">
            <div class="path-icon">
              <el-icon :size="32">
                <component :is="path.icon" />
              </el-icon>
            </div>
            <h3 class="path-title">{{ path.title }}</h3>
            <p class="path-description">{{ path.description }}</p>
            <el-button type="text" class="path-action">
              了解更多 <el-icon><ArrowRight /></el-icon>
            </el-button>
          </div>
        </div>
      </div>
    </section>

    <!-- 推荐大师区域 -->
    <section class="mentors-section">
      <div class="container">
        <h2 class="section-title">推荐大师</h2>
        <div class="mentors-grid">
          <div class="mentor-card" v-for="mentor in recommendedMentors" :key="mentor.id">
            <div class="mentor-avatar">
              <el-avatar :size="60" :src="mentor.avatar" />
              <div class="online-status" :class="{ online: mentor.isOnline }"></div>
            </div>
            <div class="mentor-info">
              <h3 class="mentor-name">{{ mentor.name }}</h3>
              <p class="mentor-domain">{{ mentor.domain }}</p>
              <div class="mentor-rating">
                <el-rate v-model="mentor.rating" disabled />
                <span class="student-count">{{ mentor.studentCount }} 学生</span>
              </div>
              <p class="mentor-price">¥{{ mentor.price }}/小时</p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// 学习路径数据
const learningPaths = ref([
  {
    id: 1,
    title: '1对1指导',
    description: '直接预约大师进行一对一专业指导',
    icon: 'User'
  },
  {
    id: 2,
    title: '结构化学习',
    description: '报名大师设计的系统化课程',
    icon: 'Reading'
  },
  {
    id: 3,
    title: '大师浏览',
    description: '浏览和筛选平台上的大师',
    icon: 'Search'
  },
  {
    id: 4,
    title: '其他学习方式',
    description: '探索更多学习途径',
    icon: 'More'
  }
])

// 推荐大师数据
const recommendedMentors = ref([
  {
    id: 1,
    name: '张大师',
    domain: '软件开发',
    avatar: '',
    rating: 4.8,
    studentCount: 156,
    price: 200,
    isOnline: true
  },
  {
    id: 2,
    name: '李大师',
    domain: 'UI设计',
    avatar: '',
    rating: 4.9,
    studentCount: 89,
    price: 180,
    isOnline: false
  },
  {
    id: 3,
    name: '王大师',
    domain: '数字营销',
    avatar: '',
    rating: 4.7,
    studentCount: 234,
    price: 150,
    isOnline: true
  }
])
</script>

<style scoped lang="scss">
.home-page {
  min-height: 100vh;
  background: var(--bg-primary);
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 var(--spacing-lg);
}

// 欢迎区域
.welcome-section {
  padding: var(--spacing-xxl) 0;
  background: linear-gradient(135deg, var(--bg-secondary) 0%, var(--bg-primary) 100%);
  text-align: center;
}

.welcome-content {
  max-width: 800px;
  margin: 0 auto;
}

.welcome-title {
  margin-bottom: var(--spacing-lg);
  
  .title-main {
    display: block;
    font-size: var(--font-size-h1);
    font-weight: var(--font-weight-bold);
    color: var(--primary-color);
    margin-bottom: var(--spacing-sm);
  }
  
  .title-sub {
    display: block;
    font-size: var(--font-size-h3);
    color: var(--text-secondary);
  }
}

.welcome-description {
  font-size: var(--font-size-large);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-xl);
  line-height: 1.6;
}

.welcome-actions {
  display: flex;
  gap: var(--spacing-md);
  justify-content: center;
  flex-wrap: wrap;
}

.action-btn {
  min-width: 160px;
  
  &.secondary {
    background: var(--bg-tertiary);
    border-color: var(--border-color);
    color: var(--text-primary);
    
    &:hover {
      background: var(--bg-card);
      border-color: var(--primary-color);
      color: var(--primary-color);
    }
  }
}

// 学习路径区域
.learning-paths-section {
  padding: var(--spacing-xxl) 0;
  background: var(--bg-primary);
}

.section-title {
  font-size: var(--font-size-h2);
  font-weight: var(--font-weight-semibold);
  text-align: center;
  margin-bottom: var(--spacing-xl);
  color: var(--text-primary);
}

.learning-paths-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: var(--spacing-lg);
}

.path-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  padding: var(--spacing-xl);
  text-align: center;
  transition: all var(--transition-normal);
  border: 1px solid var(--border-color);
  
  &:hover {
    transform: translateY(-4px);
    box-shadow: var(--shadow-medium);
    border-color: var(--primary-color);
  }
}

.path-icon {
  width: 80px;
  height: 80px;
  margin: 0 auto var(--spacing-lg);
  background: var(--primary-color);
  border-radius: var(--border-radius-full);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-primary);
}

.path-title {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  margin-bottom: var(--spacing-md);
  color: var(--text-primary);
}

.path-description {
  font-size: var(--font-size-medium);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-lg);
  line-height: 1.5;
}

.path-action {
  color: var(--primary-color);
  font-weight: var(--font-weight-medium);
  
  &:hover {
    color: var(--primary-light);
  }
}

// 推荐大师区域
.mentors-section {
  padding: var(--spacing-xxl) 0;
  background: var(--bg-secondary);
}

.mentors-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-lg);
}

.mentor-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  padding: var(--spacing-lg);
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  transition: all var(--transition-normal);
  border: 1px solid var(--border-color);
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-medium);
    border-color: var(--primary-color);
  }
}

.mentor-avatar {
  position: relative;
  flex-shrink: 0;
}

.online-status {
  position: absolute;
  bottom: 2px;
  right: 2px;
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: var(--text-tertiary);
  border: 2px solid var(--bg-card);
  
  &.online {
    background: var(--success-color);
  }
}

.mentor-info {
  flex: 1;
  min-width: 0;
}

.mentor-name {
  font-size: var(--font-size-h4);
  font-weight: var(--font-weight-semibold);
  margin-bottom: var(--spacing-xs);
  color: var(--text-primary);
}

.mentor-domain {
  font-size: var(--font-size-medium);
  color: var(--primary-color);
  margin-bottom: var(--spacing-sm);
  font-weight: var(--font-weight-medium);
}

.mentor-rating {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-sm);
}

.student-count {
  font-size: var(--font-size-small);
  color: var(--text-tertiary);
}

.mentor-price {
  font-size: var(--font-size-h5);
  font-weight: var(--font-weight-bold);
  color: var(--primary-color);
}

// 响应式设计
@media (max-width: 768px) {
  .container {
    padding: 0 var(--spacing-md);
  }
  
  .welcome-actions {
    flex-direction: column;
    align-items: center;
  }
  
  .action-btn {
    width: 100%;
    max-width: 280px;
  }
  
  .learning-paths-grid {
    grid-template-columns: 1fr;
  }
  
  .mentors-grid {
    grid-template-columns: 1fr;
  }
  
  .mentor-card {
    flex-direction: column;
    text-align: center;
  }
}
</style>