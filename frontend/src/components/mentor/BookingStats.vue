<template>
  <div class="booking-stats">
    <div class="stats-header">
      <h3>预约统计</h3>
      <el-button type="text" @click="refreshStats">
        <el-icon><Refresh /></el-icon>
        刷新
      </el-button>
    </div>
    
    <div class="stats-grid">
      <div class="stat-item">
        <div class="stat-icon pending">
          <el-icon><Clock /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-number">{{ stats.pending }}</div>
          <div class="stat-label">待确认</div>
          <div class="stat-trend" v-if="stats.pendingChange > 0">
            <el-icon><ArrowUp /></el-icon>
            +{{ stats.pendingChange }}
          </div>
        </div>
      </div>
      
      <div class="stat-item">
        <div class="stat-icon confirmed">
          <el-icon><Check /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-number">{{ stats.confirmed }}</div>
          <div class="stat-label">已确认</div>
          <div class="stat-trend" v-if="stats.confirmedChange > 0">
            <el-icon><ArrowUp /></el-icon>
            +{{ stats.confirmedChange }}
          </div>
        </div>
      </div>
      
      <div class="stat-item">
        <div class="stat-icon completed">
          <el-icon><Trophy /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-number">{{ stats.completed }}</div>
          <div class="stat-label">已完成</div>
          <div class="stat-trend" v-if="stats.completedChange > 0">
            <el-icon><ArrowUp /></el-icon>
            +{{ stats.completedChange }}
          </div>
        </div>
      </div>
      
      <div class="stat-item">
        <div class="stat-icon revenue">
          <el-icon><Money /></el-icon>
        </div>
        <div class="stat-content">
          <div class="stat-number">¥{{ stats.revenue }}</div>
          <div class="stat-label">本月收入</div>
          <div class="stat-trend" v-if="stats.revenueChange > 0">
            <el-icon><ArrowUp /></el-icon>
            +¥{{ stats.revenueChange }}
          </div>
        </div>
      </div>
    </div>
    
    <div class="stats-chart">
      <h4>预约趋势</h4>
      <div class="chart-placeholder">
        <el-empty description="图表功能开发中..." />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Refresh, Clock, Check, Trophy, Money, ArrowUp } from '@element-plus/icons-vue'

// Props
interface Props {
  mentorId?: string
}

const props = withDefaults(defineProps<Props>(), {
  mentorId: ''
})

// 状态
const loading = ref(false)

// 统计数据
const stats = ref({
  pending: 0,
  confirmed: 0,
  completed: 0,
  revenue: 0,
  pendingChange: 0,
  confirmedChange: 0,
  completedChange: 0,
  revenueChange: 0
})

// 加载统计数据
const loadStats = async () => {
  loading.value = true
  try {
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 500))
    
    // 模拟数据
    stats.value = {
      pending: 5,
      confirmed: 12,
      completed: 8,
      revenue: 2400,
      pendingChange: 2,
      confirmedChange: 3,
      completedChange: 1,
      revenueChange: 600
    }
  } catch (error) {
    console.error('加载统计数据失败:', error)
  } finally {
    loading.value = false
  }
}

// 刷新统计
const refreshStats = () => {
  loadStats()
}

// 组件挂载时加载数据
onMounted(() => {
  loadStats()
})
</script>

<style scoped lang="scss">
.booking-stats {
  background: var(--bg-card);
  border-radius: var(--border-radius-medium);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-light);
}

.stats-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
  
  h3 {
    font-size: var(--font-size-h4);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0;
  }
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-md);
  margin-bottom: var(--spacing-lg);
}

.stat-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  background: var(--bg-secondary);
  border-radius: var(--border-radius-small);
  transition: all var(--transition-normal);
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-light);
  }
}

.stat-icon {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: white;
  
  &.pending {
    background: linear-gradient(135deg, #e6a23c, #f56c6c);
  }
  
  &.confirmed {
    background: linear-gradient(135deg, #409eff, #67c23a);
  }
  
  &.completed {
    background: linear-gradient(135deg, #67c23a, #909399);
  }
  
  &.revenue {
    background: linear-gradient(135deg, #f56c6c, #e6a23c);
  }
}

.stat-content {
  flex: 1;
}

.stat-number {
  font-size: var(--font-size-h3);
  font-weight: var(--font-weight-bold);
  color: var(--text-primary);
  margin-bottom: var(--spacing-xs);
}

.stat-label {
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  margin-bottom: var(--spacing-xs);
}

.stat-trend {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--font-size-small);
  color: #67c23a;
  font-weight: var(--font-weight-medium);
}

.stats-chart {
  h4 {
    font-size: var(--font-size-h5);
    font-weight: var(--font-weight-semibold);
    color: var(--text-primary);
    margin: 0 0 var(--spacing-md) 0;
  }
}

.chart-placeholder {
  height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-secondary);
  border-radius: var(--border-radius-small);
}

// 响应式设计
@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .stat-item {
    flex-direction: column;
    text-align: center;
  }
  
  .stat-icon {
    width: 40px;
    height: 40px;
    font-size: 16px;
  }
  
  .stat-number {
    font-size: var(--font-size-h4);
  }
}
</style>