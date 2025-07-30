# Master Guide 页面布局与组件设计

## 1. 整体布局架构

### 1.1 应用布局结构
```
┌─────────────────────────────────────────────────────────┐
│                    顶部导航栏 (Header)                    │
│  Logo | 身份切换器 | 搜索框 | 通知 | 用户菜单              │
├─────────────────────────────────────────────────────────┤
│                                                         │
│                    主要内容区域 (Main)                    │
│                                                         │
│  ┌─────────────────┐  ┌─────────────────┐              │
│  │   侧边栏导航     │  │    页面内容      │              │
│  │   (可选)        │  │                 │              │
│  └─────────────────┘  └─────────────────┘              │
│                                                         │
├─────────────────────────────────────────────────────────┤
│                    底部导航栏 (Bottom Nav)               │
│  首页 | 大师 | 课程 | 社群 | 我的                        │
└─────────────────────────────────────────────────────────┘
```

### 1.2 响应式布局策略
- **移动端**：单列布局，底部导航
- **平板端**：双列布局，侧边栏 + 主内容
- **桌面端**：三列布局，侧边栏 + 主内容 + 信息栏

## 2. 首页设计

### 2.1 首页布局结构
```
┌─────────────────────────────────────────────────────────┐
│                    个性化欢迎区域                        │
│  问候语 + 当前身份 + 快速切换 + 学习进度                 │
├─────────────────────────────────────────────────────────┤
│                    学习路径选择                          │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐      │
│  │1对1指导 │ │结构化学习│ │浏览大师 │ │其他方式 │      │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘      │
├─────────────────────────────────────────────────────────┤
│                    推荐大师展示                          │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐      │
│  │大师卡片1│ │大师卡片2│ │大师卡片3│ │大师卡片4│      │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘      │
├─────────────────────────────────────────────────────────┤
│                    推荐内容                              │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐      │
│  │课程卡片1│ │课程卡片2│ │动态卡片1│ │动态卡片2│      │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘      │
└─────────────────────────────────────────────────────────┘
```

### 2.2 个性化欢迎组件
```vue
<template>
  <div class="welcome-section">
    <div class="welcome-header">
      <h1 class="greeting">{{ greeting }}</h1>
      <div class="identity-info">
        <span class="current-identity">{{ currentIdentity.name }}</span>
        <el-tag :type="getIdentityType(currentIdentity.type)">
          {{ getIdentityLabel(currentIdentity.type) }}
        </el-tag>
      </div>
    </div>
    
    <div class="quick-actions">
      <el-button @click="switchIdentity" type="primary" size="small">
        切换身份
      </el-button>
      <el-button @click="viewProgress" type="info" size="small">
        查看进度
      </el-button>
    </div>
    
    <div class="progress-summary">
      <div class="progress-item">
        <span class="label">学习进度</span>
        <el-progress :percentage="learningProgress" />
      </div>
      <div class="progress-item">
        <span class="label">指导收入</span>
        <span class="value">¥{{ teachingIncome }}</span>
      </div>
    </div>
  </div>
</template>
```

### 2.3 学习路径选择组件
```vue
<template>
  <div class="learning-paths">
    <h2 class="section-title">选择你的学习路径</h2>
    <div class="path-grid">
      <div 
        v-for="path in learningPaths" 
        :key="path.id"
        class="path-card"
        @click="selectPath(path)"
      >
        <div class="path-icon">
          <el-icon :size="32">
            <component :is="path.icon" />
          </el-icon>
        </div>
        <h3 class="path-title">{{ path.title }}</h3>
        <p class="path-description">{{ path.description }}</p>
        <div class="path-meta">
          <span class="duration">{{ path.duration }}</span>
          <span class="price">{{ path.price }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
```

## 3. 大师页面设计

### 3.1 大师列表页面
```
┌─────────────────────────────────────────────────────────┐
│                    搜索和筛选区域                        │
│  搜索框 | 领域筛选 | 价格范围 | 评分筛选 | 在线状态       │
├─────────────────────────────────────────────────────────┤
│                    大师列表                              │
│  ┌─────────────────────────────────────────────────┐   │
│  │ 大师卡片1 - 头像 | 姓名 | 领域 | 评分 | 价格 | 预约 │   │
│  └─────────────────────────────────────────────────┘   │
│  ┌─────────────────────────────────────────────────┐   │
│  │ 大师卡片2 - 头像 | 姓名 | 领域 | 评分 | 价格 | 预约 │   │
│  └─────────────────────────────────────────────────┘   │
│  ┌─────────────────────────────────────────────────┐   │
│  │ 大师卡片3 - 头像 | 姓名 | 领域 | 评分 | 价格 | 预约 │   │
│  └─────────────────────────────────────────────────┘   │
├─────────────────────────────────────────────────────────┤
│                    分页控件                              │
└─────────────────────────────────────────────────────────┘
```

### 3.2 大师卡片组件
```vue
<template>
  <div class="mentor-card" @click="viewMentorDetail">
    <div class="mentor-header">
      <div class="mentor-avatar">
        <el-avatar :src="mentor.avatar" :size="60" />
        <div class="online-status" :class="{ online: mentor.isOnline }"></div>
      </div>
      <div class="mentor-info">
        <h3 class="mentor-name">{{ mentor.name }}</h3>
        <div class="mentor-badges">
          <el-tag v-if="mentor.isVerified" type="success" size="small">
            已认证
          </el-tag>
          <el-tag v-if="mentor.isOnline" type="primary" size="small">
            在线
          </el-tag>
        </div>
      </div>
      <div class="mentor-rating">
        <el-rate v-model="mentor.rating" disabled />
        <span class="rating-text">{{ mentor.rating }}分</span>
      </div>
    </div>
    
    <div class="mentor-details">
      <p class="mentor-domain">{{ mentor.domain }}</p>
      <p class="mentor-description">{{ mentor.description }}</p>
      
      <div class="mentor-stats">
        <span class="stat-item">
          <el-icon><User /></el-icon>
          {{ mentor.studentCount }} 学生
        </span>
        <span class="stat-item">
          <el-icon><Clock /></el-icon>
          {{ mentor.experience }} 年经验
        </span>
      </div>
    </div>
    
    <div class="mentor-footer">
      <div class="price-info">
        <span class="price">¥{{ mentor.price }}/小时</span>
      </div>
      <el-button type="primary" size="small" @click.stop="bookMentor">
        立即预约
      </el-button>
    </div>
  </div>
</template>
```

---

**文档版本**：v1.0.0  
**创建日期**：2024年12月  
**设计负责人**：Sally (UX Expert) 