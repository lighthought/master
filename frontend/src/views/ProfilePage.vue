<template>
  <div class="profile-page">
    <div class="container">
      <!-- 页面头部 -->
      <div class="page-header">
        <h1 class="page-title">个人中心</h1>
        <p class="page-description">管理个人信息和学习记录</p>
      </div>

      <!-- 加载状态 -->
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="10" animated />
      </div>

      <div v-else class="profile-content">
        <!-- 用户基本信息卡片 -->
        <div class="profile-card">
          <div class="profile-header">
            <div class="avatar-section">
              <el-avatar :size="80" :src="userInfo.avatar" />
              <div class="user-info">
                <h2 class="user-name">{{ userInfo.name }}</h2>
                <p class="user-email">{{ userInfo.email }}</p>
                <div class="user-status">
                  <el-tag :type="userInfo.status === 'active' ? 'success' : 'warning'">
                    {{ userInfo.status === 'active' ? '正常' : '待激活' }}
                  </el-tag>
                </div>
              </div>
            </div>
            <div class="profile-actions">
              <el-button type="primary" @click="showEditDialog = true">
                <el-icon><Edit /></el-icon>
                编辑资料
              </el-button>
            </div>
          </div>
        </div>

        <!-- 身份信息 -->
        <div class="profile-card">
          <div class="card-header">
            <h3 class="card-title">
              <el-icon><User /></el-icon>
              身份信息
            </h3>
            <el-button type="text" @click="showAddIdentityDialog = true">
              <el-icon><Plus /></el-icon>
              添加身份
            </el-button>
          </div>
          
          <div class="identities-list">
            <div 
              v-for="identity in userInfo.identities" 
              :key="identity.id"
              class="identity-item"
              :class="{ 'active': identity.id === userInfo.currentIdentityId }"
            >
              <div class="identity-info">
                <div class="identity-avatar">
                  <el-avatar :size="50" :src="identity.avatar" />
                  <el-tag 
                    :type="identity.type === 'master' ? 'success' : 'warning'"
                    size="small"
                    class="identity-type"
                  >
                    {{ identity.type === 'master' ? '大师' : '学徒' }}
                  </el-tag>
                </div>
                <div class="identity-details">
                  <h4 class="identity-name">{{ identity.name }}</h4>
                  <p class="identity-domain">{{ identity.domain }}</p>
                  <div class="identity-meta">
                    <span class="meta-item">
                      <el-icon><Calendar /></el-icon>
                      创建于 {{ formatDate(identity.createdAt) }}
                    </span>
                    <span v-if="identity.verified" class="meta-item verified">
                      <el-icon><Check /></el-icon>
                      已认证
                    </span>
                  </div>
                </div>
              </div>
              <div class="identity-actions">
                <el-button 
                  v-if="identity.id !== userInfo.currentIdentityId"
                  type="primary" 
                  size="small"
                  @click="switchIdentity(identity.id)"
                >
                  切换到此身份
                </el-button>
                <el-button 
                  v-else
                  type="success" 
                  size="small"
                  disabled
                >
                  当前身份
                </el-button>
                <el-button 
                  type="text" 
                  size="small"
                  @click="editIdentity(identity)"
                >
                  编辑
                </el-button>
              </div>
            </div>
          </div>
        </div>

        <!-- 个人详细信息 -->
        <div class="profile-card">
          <div class="card-header">
            <h3 class="card-title">
              <el-icon><InfoFilled /></el-icon>
              详细信息
            </h3>
          </div>
          
          <div class="details-grid">
            <div class="detail-item">
              <label class="detail-label">真实姓名</label>
              <span class="detail-value">{{ userInfo.realName || '未设置' }}</span>
            </div>
            <div class="detail-item">
              <label class="detail-label">手机号码</label>
              <span class="detail-value">{{ userInfo.phone || '未设置' }}</span>
            </div>
            <div class="detail-item">
              <label class="detail-label">性别</label>
              <span class="detail-value">{{ getGenderText(userInfo.gender) }}</span>
            </div>
            <div class="detail-item">
              <label class="detail-label">生日</label>
              <span class="detail-value">{{ userInfo.birthday || '未设置' }}</span>
            </div>
            <div class="detail-item">
              <label class="detail-label">所在地区</label>
              <span class="detail-value">{{ userInfo.location || '未设置' }}</span>
            </div>
            <div class="detail-item">
              <label class="detail-label">个人简介</label>
              <span class="detail-value">{{ userInfo.bio || '未设置' }}</span>
            </div>
          </div>
        </div>

        <!-- 账号安全 -->
        <div class="profile-card">
          <div class="card-header">
            <h3 class="card-title">
              <el-icon><Lock /></el-icon>
              账号安全
            </h3>
          </div>
          
          <div class="security-list">
            <div class="security-item">
              <div class="security-info">
                <div class="security-icon">
                  <el-icon><Key /></el-icon>
                </div>
                <div class="security-details">
                  <h4>登录密码</h4>
                  <p>定期更换密码可以保护账号安全</p>
                </div>
              </div>
              <el-button type="primary" size="small" @click="showChangePasswordDialog = true">
                修改密码
              </el-button>
            </div>
            
            <div class="security-item">
              <div class="security-info">
                <div class="security-icon">
                  <el-icon><Phone /></el-icon>
                </div>
                <div class="security-details">
                  <h4>手机绑定</h4>
                  <p>{{ userInfo.phone ? '已绑定手机：' + maskPhone(userInfo.phone) : '未绑定手机' }}</p>
                </div>
              </div>
              <el-button 
                :type="userInfo.phone ? 'default' : 'primary'" 
                size="small" 
                @click="showBindPhoneDialog = true"
              >
                {{ userInfo.phone ? '更换手机' : '绑定手机' }}
              </el-button>
            </div>
            
            <div class="security-item">
              <div class="security-info">
                <div class="security-icon">
                  <el-icon><Message /></el-icon>
                </div>
                <div class="security-details">
                  <h4>邮箱绑定</h4>
                  <p>{{ userInfo.email ? '已绑定邮箱：' + maskEmail(userInfo.email) : '未绑定邮箱' }}</p>
                </div>
              </div>
              <el-button 
                :type="userInfo.email ? 'default' : 'primary'" 
                size="small" 
                @click="showBindEmailDialog = true"
              >
                {{ userInfo.email ? '更换邮箱' : '绑定邮箱' }}
              </el-button>
            </div>
          </div>
        </div>

        <!-- 学习统计（学徒身份） -->
        <div v-if="hasApprenticeIdentity" class="profile-card">
          <div class="card-header">
            <h3 class="card-title">
              <el-icon><Reading /></el-icon>
              学习统计
            </h3>
            <el-button type="text" @click="$router.push('/learning-records')">
              <el-icon><ArrowRight /></el-icon>
              查看学习记录
            </el-button>
          </div>
          
          <div class="stats-grid">
            <div class="stat-item">
              <div class="stat-number">{{ learningStats.totalCourses }}</div>
              <div class="stat-label">已报名课程</div>
            </div>
            <div class="stat-item">
              <div class="stat-number">{{ learningStats.completedCourses }}</div>
              <div class="stat-label">已完成课程</div>
            </div>
            <div class="stat-item">
              <div class="stat-number">{{ learningStats.totalStudyTime }}</div>
              <div class="stat-label">总学习时长(小时)</div>
            </div>
            <div class="stat-item">
              <div class="stat-number">{{ learningStats.averageRating }}</div>
              <div class="stat-label">平均评分</div>
            </div>
          </div>
        </div>
        
        <!-- 教学统计 -->
        <div v-if="hasMasterIdentity" class="profile-card">
          <div class="card-header">
            <h3 class="card-title">
              <el-icon><UserFilled /></el-icon>
              教学统计
            </h3>
            <el-button type="text" @click="$router.push('/student-management')">
              <el-icon><ArrowRight /></el-icon>
              学生管理
            </el-button>
            <el-button type="text" @click="$router.push('/income-stats')">
              <el-icon><ArrowRight /></el-icon>
              收入统计
            </el-button>
          </div>
          
          <div class="stats-grid">
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><VideoPlay /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ learningStats.totalCourses }}</div>
                <div class="stat-label">已报名课程</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><Clock /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ learningStats.totalStudyTime }}</div>
                <div class="stat-label">总学习时长(小时)</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><Trophy /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ learningStats.completedCourses }}</div>
                <div class="stat-label">已完成课程</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><Star /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ learningStats.averageRating }}</div>
                <div class="stat-label">平均评分</div>
              </div>
            </div>
          </div>
        </div>

        <!-- 教学统计（大师身份） -->
        <div v-if="hasMasterIdentity" class="profile-card">
          <div class="card-header">
            <h3 class="card-title">
              <el-icon><UserFilled /></el-icon>
              教学统计
            </h3>
          </div>
          
          <div class="stats-grid">
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><User /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ teachingStats.totalStudents }}</div>
                <div class="stat-label">总学员数</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><VideoPlay /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ teachingStats.totalCourses }}</div>
                <div class="stat-label">发布课程</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><Money /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">¥{{ teachingStats.totalIncome }}</div>
                <div class="stat-label">总收入</div>
              </div>
            </div>
            
            <div class="stat-card">
              <div class="stat-icon">
                <el-icon><Star /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ teachingStats.averageRating }}</div>
                <div class="stat-label">平均评分</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 编辑个人资料对话框 -->
    <el-dialog
      v-model="showEditDialog"
      title="编辑个人资料"
      width="90%"
      max-width="600px"
      :close-on-click-modal="false"
    >
      <el-form :model="editForm" label-width="100px">
        <el-form-item label="头像">
          <el-upload
            action="#"
            :auto-upload="false"
            :show-file-list="false"
            accept="image/*"
            @change="handleAvatarUpload"
          >
            <el-avatar :size="80" :src="editForm.avatar" />
            <div class="upload-tip">点击更换头像</div>
          </el-upload>
        </el-form-item>
        
        <el-form-item label="真实姓名">
          <el-input v-model="editForm.realName" placeholder="请输入真实姓名" />
        </el-form-item>
        
        <el-form-item label="手机号码">
          <el-input v-model="editForm.phone" placeholder="请输入手机号码" />
        </el-form-item>
        
        <el-form-item label="性别">
          <el-radio-group v-model="editForm.gender">
            <el-radio label="male">男</el-radio>
            <el-radio label="female">女</el-radio>
            <el-radio label="other">其他</el-radio>
          </el-radio-group>
        </el-form-item>
        
        <el-form-item label="生日">
          <el-date-picker
            v-model="editForm.birthday"
            type="date"
            placeholder="选择生日"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        
        <el-form-item label="所在地区">
          <el-input v-model="editForm.location" placeholder="请输入所在地区" />
        </el-form-item>
        
        <el-form-item label="个人简介">
          <el-input
            v-model="editForm.bio"
            type="textarea"
            :rows="4"
            placeholder="请输入个人简介"
            maxlength="200"
            show-word-limit
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showEditDialog = false">取消</el-button>
          <el-button type="primary" @click="saveProfile" :loading="saving">
            保存
          </el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 修改密码对话框 -->
    <el-dialog
      v-model="showChangePasswordDialog"
      title="修改密码"
      width="90%"
      max-width="500px"
      :close-on-click-modal="false"
    >
      <el-form :model="passwordForm" label-width="100px">
        <el-form-item label="当前密码">
          <el-input
            v-model="passwordForm.currentPassword"
            type="password"
            placeholder="请输入当前密码"
            show-password
          />
        </el-form-item>
        
        <el-form-item label="新密码">
          <el-input
            v-model="passwordForm.newPassword"
            type="password"
            placeholder="请输入新密码"
            show-password
          />
        </el-form-item>
        
        <el-form-item label="确认密码">
          <el-input
            v-model="passwordForm.confirmPassword"
            type="password"
            placeholder="请再次输入新密码"
            show-password
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showChangePasswordDialog = false">取消</el-button>
          <el-button type="primary" @click="changePassword" :loading="changingPassword">
            确认修改
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { ElMessage } from 'element-plus'
import {
  Edit, User, Plus, Calendar, Check, InfoFilled, Lock, Key, Phone, Message,
  Reading, VideoPlay, Clock, Trophy, Star, UserFilled, Money, ArrowRight
} from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'
import { ApiService } from '@/services/api'

// Store
const authStore = useAuthStore()

// 状态
const loading = ref(false)
const saving = ref(false)
const changingPassword = ref(false)
const showEditDialog = ref(false)
const showChangePasswordDialog = ref(false)
const showAddIdentityDialog = ref(false)
const showBindPhoneDialog = ref(false)
const showBindEmailDialog = ref(false)

// 数据
const userInfo = ref<any>({})
const learningStats = ref({
  totalCourses: 0,
  totalStudyTime: 0,
  completedCourses: 0,
  averageRating: 0
})
const teachingStats = ref({
  totalStudents: 0,
  totalCourses: 0,
  totalIncome: 0,
  averageRating: 0
})

// 表单
const editForm = ref({
  avatar: '',
  realName: '',
  phone: '',
  gender: '',
  birthday: '',
  location: '',
  bio: ''
})

const passwordForm = ref({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

// 计算属性
const hasApprenticeIdentity = computed(() => {
  return userInfo.value.identities?.some((identity: any) => identity.type === 'apprentice')
})

const hasMasterIdentity = computed(() => {
  return userInfo.value.identities?.some((identity: any) => identity.type === 'master')
})

// 方法
const loadUserInfo = async () => {
  loading.value = true
  try {
    const result = await ApiService.auth.getUserInfo(authStore.user!.id)
    userInfo.value = result.data
    
    // 初始化编辑表单
    editForm.value = {
      avatar: userInfo.value.avatar || '',
      realName: userInfo.value.realName || '',
      phone: userInfo.value.phone || '',
      gender: userInfo.value.gender || '',
      birthday: userInfo.value.birthday || '',
      location: userInfo.value.location || '',
      bio: userInfo.value.bio || ''
    }
    
    // 加载学习统计
    if (hasApprenticeIdentity.value) {
      const learningResult = await ApiService.userStats.getLearningStats(authStore.user!.id)
      learningStats.value = learningResult.data
    }
    
    // 加载教学统计
    if (hasMasterIdentity.value) {
      const teachingResult = await ApiService.userStats.getTeachingStats(authStore.user!.id)
      teachingStats.value = teachingResult.data
    }
  } catch (error) {
    ElMessage.error('加载用户信息失败')
  } finally {
    loading.value = false
  }
}

const switchIdentity = async (identityId: string) => {
  try {
    await ApiService.auth.switchIdentity(authStore.user!.id, identityId)
    await loadUserInfo()
    ElMessage.success('身份切换成功')
  } catch (error) {
    ElMessage.error('身份切换失败')
  }
}

const editIdentity = (identity: any) => {
  // TODO: 实现身份编辑功能
  ElMessage.info('身份编辑功能开发中')
}

const saveProfile = async () => {
  saving.value = true
  try {
    await ApiService.auth.updateIdentityInfo(
      authStore.user!.id,
      userInfo.value.currentIdentityId,
      editForm.value
    )
    await loadUserInfo()
    showEditDialog.value = false
    ElMessage.success('个人资料保存成功')
  } catch (error) {
    ElMessage.error('保存失败')
  } finally {
    saving.value = false
  }
}

const changePassword = async () => {
  if (passwordForm.value.newPassword !== passwordForm.value.confirmPassword) {
    ElMessage.error('两次输入的密码不一致')
    return
  }
  
  changingPassword.value = true
  try {
    await ApiService.auth.changePassword(
      passwordForm.value.currentPassword,
      passwordForm.value.newPassword
    )
    showChangePasswordDialog.value = false
    passwordForm.value = {
      currentPassword: '',
      newPassword: '',
      confirmPassword: ''
    }
    ElMessage.success('密码修改成功')
  } catch (error) {
    ElMessage.error('密码修改失败')
  } finally {
    changingPassword.value = false
  }
}

const handleAvatarUpload = (file: any) => {
  const reader = new FileReader()
  reader.onload = (e) => {
    editForm.value.avatar = e.target?.result as string
  }
  reader.readAsDataURL(file.raw)
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString()
}

const getGenderText = (gender: string) => {
  const genderMap: Record<string, string> = {
    male: '男',
    female: '女',
    other: '其他'
  }
  return genderMap[gender] || '未设置'
}

const maskPhone = (phone: string) => {
  return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
}

const maskEmail = (email: string) => {
  const [username, domain] = email.split('@')
  return username.substring(0, 3) + '***@' + domain
}

onMounted(() => {
  loadUserInfo()
})
</script>

<style scoped lang="scss">
.profile-page {
  padding: var(--spacing-xl) 0;
  background: var(--bg-page);
  min-height: 100vh;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 var(--spacing-lg);
}

.page-header {
  margin-bottom: var(--spacing-xl);
  
  .page-title {
    font-size: var(--font-size-h1);
    font-weight: var(--font-weight-bold);
    color: var(--text-primary);
    margin-bottom: var(--spacing-sm);
  }
  
  .page-description {
    font-size: var(--font-size-large);
    color: var(--text-secondary);
  }
}

.loading-container {
  padding: var(--spacing-xl) 0;
}

.profile-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.profile-card {
  background: var(--bg-card);
  border-radius: var(--border-radius-large);
  box-shadow: var(--shadow-card);
  padding: var(--spacing-xl);
  
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-lg);
    
    .card-title {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
      font-size: var(--font-size-h3);
      font-weight: var(--font-weight-medium);
      color: var(--text-primary);
      margin: 0;
    }
  }
}

.profile-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  
  .avatar-section {
    display: flex;
    align-items: center;
    gap: var(--spacing-lg);
    
    .user-info {
      .user-name {
        font-size: var(--font-size-h2);
        font-weight: var(--font-weight-bold);
        color: var(--text-primary);
        margin: 0 0 var(--spacing-xs) 0;
      }
      
      .user-email {
        color: var(--text-secondary);
        margin: 0 0 var(--spacing-sm) 0;
      }
    }
  }
}

.identities-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  
  .identity-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: var(--spacing-lg);
    border: 2px solid var(--border-color);
    border-radius: var(--border-radius-medium);
    transition: all 0.3s ease;
    
    &.active {
      border-color: var(--primary-color);
      background: var(--primary-color-light);
    }
    
    .identity-info {
      display: flex;
      align-items: center;
      gap: var(--spacing-md);
      
      .identity-avatar {
        position: relative;
        
        .identity-type {
          position: absolute;
          bottom: -5px;
          right: -5px;
        }
      }
      
      .identity-details {
        .identity-name {
          font-size: var(--font-size-large);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
          margin: 0 0 var(--spacing-xs) 0;
        }
        
        .identity-domain {
          color: var(--text-secondary);
          margin: 0 0 var(--spacing-xs) 0;
        }
        
        .identity-meta {
          display: flex;
          gap: var(--spacing-md);
          
          .meta-item {
            display: flex;
            align-items: center;
            gap: var(--spacing-xs);
            font-size: var(--font-size-small);
            color: var(--text-secondary);
            
            &.verified {
              color: var(--success-color);
            }
          }
        }
      }
    }
    
    .identity-actions {
      display: flex;
      gap: var(--spacing-sm);
    }
  }
}

.details-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-lg);
  
  .detail-item {
    .detail-label {
      display: block;
      font-size: var(--font-size-small);
      color: var(--text-secondary);
      margin-bottom: var(--spacing-xs);
    }
    
    .detail-value {
      font-size: var(--font-size-medium);
      color: var(--text-primary);
    }
  }
}

.security-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
  
  .security-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: var(--spacing-lg);
    border: 1px solid var(--border-color);
    border-radius: var(--border-radius-medium);
    
    .security-info {
      display: flex;
      align-items: center;
      gap: var(--spacing-md);
      
      .security-icon {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        background: var(--primary-color-light);
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--primary-color);
      }
      
      .security-details {
        h4 {
          font-size: var(--font-size-medium);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
          margin: 0 0 var(--spacing-xs) 0;
        }
        
        p {
          font-size: var(--font-size-small);
          color: var(--text-secondary);
          margin: 0;
        }
      }
    }
  }
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-lg);
  
  .stat-card {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
    padding: var(--spacing-lg);
    background: var(--bg-light);
    border-radius: var(--border-radius-medium);
    
    .stat-icon {
      width: 50px;
      height: 50px;
      border-radius: 50%;
      background: var(--primary-color);
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-size: 20px;
    }
    
    .stat-content {
      .stat-number {
        font-size: var(--font-size-h3);
        font-weight: var(--font-weight-bold);
        color: var(--text-primary);
        margin-bottom: var(--spacing-xs);
      }
      
      .stat-label {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
      }
    }
  }
}

.upload-tip {
  text-align: center;
  font-size: var(--font-size-small);
  color: var(--text-secondary);
  margin-top: var(--spacing-xs);
}

.dialog-footer {
  text-align: right;
}

@media (max-width: 768px) {
  .profile-header {
    flex-direction: column;
    gap: var(--spacing-lg);
    text-align: center;
  }
  
  .identity-item {
    flex-direction: column;
    gap: var(--spacing-md);
    text-align: center;
  }
  
  .security-item {
    flex-direction: column;
    gap: var(--spacing-md);
    text-align: center;
  }
  
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>