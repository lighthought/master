<template>
  <div class="course-learning-page">
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="10" animated />
    </div>
    
    <div v-else-if="course" class="learning-content">
      <!-- 学习头部 -->
      <div class="learning-header">
        <div class="course-info">
          <h1 class="course-title">{{ course.title }}</h1>
          <div class="course-meta">
            <span class="mentor">大师：{{ course.mentorName }}</span>
            <span class="progress">学习进度：{{ course.progress || 0 }}%</span>
          </div>
        </div>
        
        <div class="learning-actions">
          <el-button @click="showNotes = true" type="info" size="large">
            <el-icon><EditPen /></el-icon>
            我的笔记
          </el-button>
          <el-button @click="showDiscussion = true" type="warning" size="large">
            <el-icon><ChatDotRound /></el-icon>
            课程讨论
          </el-button>
        </div>
      </div>
      
      <div class="learning-layout">
        <!-- 左侧：视频播放区域 -->
        <div class="video-section">
          <div class="video-player">
            <div v-if="currentLesson" class="video-container">
              <div class="video-placeholder">
                <el-icon size="80" color="#409eff">
                  <VideoPlay />
                </el-icon>
                <h3>{{ currentLesson.title }}</h3>
                <p>视频播放区域</p>
                <el-button @click="togglePlay" type="primary" size="large">
                  <el-icon><VideoPlay v-if="!isPlaying" /><VideoPause v-else /></el-icon>
                  {{ isPlaying ? '暂停' : '播放' }}
                </el-button>
              </div>
              
              <!-- 视频控制栏 -->
              <div class="video-controls">
                <div class="progress-bar">
                  <el-slider 
                    v-model="videoProgress" 
                    :max="100" 
                    @input="seekVideo"
                    :show-tooltip="false"
                  />
                </div>
                <div class="control-buttons">
                  <el-button @click="togglePlay" type="text" size="small">
                    <el-icon><VideoPlay v-if="!isPlaying" /><VideoPause v-else /></el-icon>
                  </el-button>
                  <span class="time-display">{{ formatTime(currentTime) }} / {{ formatTime(duration) }}</span>
                  <el-button @click="toggleFullscreen" type="text" size="small">
                    <el-icon><FullScreen /></el-icon>
                  </el-button>
                </div>
              </div>
            </div>
            
            <div v-else class="no-lesson">
              <el-empty description="请选择要学习的课程">
                <el-button type="primary" @click="selectFirstLesson">
                  开始学习
                </el-button>
              </el-empty>
            </div>
          </div>
          
          <!-- 课程信息 -->
          <div class="lesson-info" v-if="currentLesson">
            <h3>{{ currentLesson.title }}</h3>
            <p>{{ currentLesson.description || '暂无描述' }}</p>
            
            <div class="lesson-actions">
              <el-button @click="markAsCompleted" :type="currentLesson.completed ? 'success' : 'default'" size="small">
                <el-icon><Check v-if="currentLesson.completed" /><CircleCheck v-else /></el-icon>
                {{ currentLesson.completed ? '已完成' : '标记完成' }}
              </el-button>
              
              <el-button @click="showNotes = true" type="info" size="small">
                <el-icon><EditPen /></el-icon>
                记笔记
              </el-button>
              
              <el-button @click="showDiscussion = true" type="warning" size="small">
                <el-icon><ChatDotRound /></el-icon>
                提问
              </el-button>
            </div>
          </div>
        </div>
        
        <!-- 右侧：课程大纲 -->
        <div class="outline-section">
          <div class="outline-header">
            <h3>课程大纲</h3>
            <div class="outline-stats">
              <span>{{ completedLessons }}/{{ totalLessons }} 已完成</span>
              <el-progress :percentage="course.progress || 0" :stroke-width="6" />
            </div>
          </div>
          
          <div class="outline-content">
            <div v-for="(section, sectionIndex) in course.outline" :key="sectionIndex" class="outline-section">
              <div class="section-header" @click="toggleSection(sectionIndex)">
                <h4 class="section-title">
                  <span class="section-number">{{ sectionIndex + 1 }}</span>
                  {{ section.title }}
                </h4>
                <div class="section-meta">
                  <span class="lesson-count">{{ section.lessons.length }} 课时</span>
                  <span class="completed-count">{{ getCompletedCount(section) }}/{{ section.lessons.length }}</span>
                  <el-icon class="toggle-icon" :class="{ 'is-expanded': expandedSections.includes(sectionIndex) }">
                    <ArrowDown />
                  </el-icon>
                </div>
              </div>
              
              <div class="lessons-list" v-show="expandedSections.includes(sectionIndex)">
                <div 
                  v-for="(lesson, lessonIndex) in section.lessons" 
                  :key="lessonIndex" 
                  class="lesson-item"
                  :class="{ 
                    'active': currentLesson && currentLesson.id === lesson.id,
                    'completed': lesson.completed,
                    'locked': !lesson.available
                  }"
                  @click="selectLesson(lesson)"
                >
                  <div class="lesson-info">
                    <el-icon class="lesson-icon">
                      <VideoPlay v-if="lesson.type === 'video'" />
                      <Document v-else />
                    </el-icon>
                    <span class="lesson-title">{{ lesson.title }}</span>
                    <el-tag v-if="lesson.isFree" type="success" size="small" effect="light">免费</el-tag>
                  </div>
                  <div class="lesson-meta">
                    <span class="lesson-duration">{{ lesson.duration }}</span>
                    <el-icon v-if="lesson.completed" class="completed-icon" color="#67c23a">
                      <Check />
                    </el-icon>
                    <el-icon v-else-if="!lesson.available" class="locked-icon" color="#909399">
                      <Lock />
                    </el-icon>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 笔记对话框 -->
    <el-dialog v-model="showNotes" title="我的笔记" width="80%" max-width="800px">
      <div class="notes-content">
        <div class="notes-header">
          <h4>{{ currentLesson?.title }} - 学习笔记</h4>
          <el-button @click="addNote" type="primary" size="small">
            <el-icon><Plus /></el-icon>
            添加笔记
          </el-button>
        </div>
        
        <div class="notes-list">
          <div v-for="note in currentNotes" :key="note.id" class="note-item">
            <div class="note-header">
              <span class="note-time">{{ formatDate(note.createdAt) }}</span>
              <div class="note-actions">
                <el-button @click="editNote(note)" type="text" size="small">
                  <el-icon><EditPen /></el-icon>
                </el-button>
                <el-button @click="deleteNote(note.id)" type="text" size="small" color="#f56c6c">
                  <el-icon><Delete /></el-icon>
                </el-button>
              </div>
            </div>
            <div class="note-content">{{ note.content }}</div>
          </div>
        </div>
        
        <div v-if="currentNotes.length === 0" class="no-notes">
          <el-empty description="暂无笔记">
            <el-button @click="addNote" type="primary">
              添加第一条笔记
            </el-button>
          </el-empty>
        </div>
      </div>
    </el-dialog>
    
    <!-- 讨论对话框 -->
    <el-dialog v-model="showDiscussion" title="课程讨论" width="80%" max-width="800px">
      <div class="discussion-content">
        <div class="discussion-header">
          <h4>{{ currentLesson?.title }} - 课程讨论</h4>
          <el-button @click="showAskQuestion = true" type="primary" size="small">
            <el-icon><Plus /></el-icon>
            提问
          </el-button>
        </div>
        
        <div class="discussion-list">
          <div v-for="question in currentQuestions" :key="question.id" class="question-item">
            <div class="question-header">
              <div class="user-info">
                <el-avatar :size="32" :src="question.userAvatar" />
                <span class="user-name">{{ question.userName }}</span>
                <span class="question-time">{{ formatDate(question.createdAt) }}</span>
              </div>
              <el-tag v-if="question.status === 'answered'" type="success" size="small">已解答</el-tag>
              <el-tag v-else type="warning" size="small">待解答</el-tag>
            </div>
            <div class="question-content">{{ question.content }}</div>
            <div class="question-actions">
              <el-button @click="showAnswers(question)" type="text" size="small">
                <el-icon><ChatDotRound /></el-icon>
                {{ question.answerCount || 0 }} 个回答
              </el-button>
              <el-button @click="likeQuestion(question)" type="text" size="small">
                <el-icon><Star /></el-icon>
                {{ question.likeCount || 0 }}
              </el-button>
            </div>
          </div>
        </div>
        
        <div v-if="currentQuestions.length === 0" class="no-questions">
          <el-empty description="暂无讨论">
            <el-button @click="showAskQuestion = true" type="primary">
              发起第一个讨论
            </el-button>
          </el-empty>
        </div>
      </div>
    </el-dialog>
    
    <!-- 提问对话框 -->
    <el-dialog v-model="showAskQuestion" title="提出问题" width="60%" max-width="600px">
      <el-form :model="questionForm" label-width="80px">
        <el-form-item label="问题标题">
          <el-input v-model="questionForm.title" placeholder="请输入问题标题" />
        </el-form-item>
        <el-form-item label="问题内容">
          <el-input 
            v-model="questionForm.content" 
            type="textarea" 
            :rows="6"
            placeholder="请详细描述你的问题..."
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showAskQuestion = false">取消</el-button>
          <el-button type="primary" @click="submitQuestion">提交问题</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  VideoPlay, 
  VideoPause, 
  FullScreen, 
  EditPen, 
  ChatDotRound, 
  Check, 
  CircleCheck, 
  Document, 
  Lock, 
  ArrowDown, 
  Plus, 
  Delete, 
  Star 
} from '@element-plus/icons-vue'
import { ApiService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const course = ref<any>(null)
const currentLesson = ref<any>(null)
const loading = ref(true)
const isPlaying = ref(false)
const videoProgress = ref(0)
const currentTime = ref(0)
const duration = ref(0)
const expandedSections = ref<number[]>([0])
const showNotes = ref(false)
const showDiscussion = ref(false)
const showAskQuestion = ref(false)

const questionForm = ref({
  title: '',
  content: ''
})

// 加载课程详情
const loadCourseDetail = async () => {
  const courseId = route.params.id as string
  loading.value = true
  
  try {
    const response = await ApiService.courses.getCourseDetail(courseId)
    course.value = response.data
    
    // 检查是否已报名
    if (course.value.enrollmentStatus !== 'enrolled') {
      ElMessage.warning('请先报名课程才能学习')
      router.push(`/courses/${courseId}`)
      return
    }
    
    // 设置课程大纲的学习状态
    setupCourseOutline()
  } catch (error) {
    console.error('加载课程详情失败:', error)
    ElMessage.error('加载课程详情失败')
  } finally {
    loading.value = false
  }
}

// 设置课程大纲
const setupCourseOutline = async () => {
  if (!course.value?.outline) return
  
  try {
    // 获取用户学习记录
    const response = await ApiService.learning.getUserLearningRecords(
      authStore.user?.id || '1', 
      course.value.id
    )
    const learningRecords = response.data
    
    course.value.outline.forEach((section: any, sectionIndex: number) => {
      section.lessons.forEach((lesson: any, lessonIndex: number) => {
        // 设置课程ID
        lesson.id = `${sectionIndex}-${lessonIndex}`
        
        // 查找学习记录
        const record = learningRecords.find((r: any) => r.lessonId === lesson.id)
        
        // 设置完成状态
        lesson.completed = record?.completed || false
        lesson.progress = record?.progress || 0
        
        // 设置可用状态（第一个课程和免费课程可用，或者前置课程已完成）
        lesson.available = lessonIndex === 0 || lesson.isFree || 
          (sectionIndex > 0 && course.value.outline[sectionIndex - 1].lessons.every((l: any) => l.completed))
      })
    })
    
    // 更新课程进度
    updateProgress()
  } catch (error) {
    console.error('获取学习记录失败:', error)
  }
}

// 选择课程
const selectLesson = async (lesson: any) => {
  if (!lesson.available) {
    ElMessage.warning('请先完成前置课程')
    return
  }
  
  currentLesson.value = lesson
  isPlaying.value = false
  videoProgress.value = 0
  currentTime.value = 0
  duration.value = 300 // 5分钟视频
  
  // 加载笔记和讨论
  await loadNotes()
  await loadDiscussions()
  
  // 模拟加载视频
  ElMessage.info(`正在加载：${lesson.title}`)
}

// 选择第一个课程
const selectFirstLesson = () => {
  if (course.value?.outline?.[0]?.lessons?.[0]) {
    selectLesson(course.value.outline[0].lessons[0])
  }
}

// 切换播放状态
const togglePlay = () => {
  isPlaying.value = !isPlaying.value
  
  if (isPlaying.value) {
    // 模拟视频播放
    const interval = setInterval(async () => {
      if (currentTime.value < duration.value) {
        currentTime.value += 1
        videoProgress.value = (currentTime.value / duration.value) * 100
        
        // 每30秒更新一次学习进度
        if (currentTime.value % 30 === 0) {
          try {
            await ApiService.learning.updateLearningProgress({
              userId: authStore.user?.id || '1',
              courseId: course.value.id,
              lessonId: currentLesson.value.id,
              progress: videoProgress.value,
              studyTime: currentTime.value
            })
          } catch (error) {
            console.error('更新学习进度失败:', error)
          }
        }
      } else {
        clearInterval(interval)
        isPlaying.value = false
        markAsCompleted()
      }
    }, 1000)
  }
}

// 跳转视频
const seekVideo = (value: number) => {
  currentTime.value = (value / 100) * duration.value
  videoProgress.value = value
}

// 切换全屏
const toggleFullscreen = () => {
  ElMessage.info('切换全屏模式')
}

// 标记完成
const markAsCompleted = async () => {
  if (!currentLesson.value) return
  
  try {
    await ApiService.learning.markLessonCompleted({
      userId: authStore.user?.id || '1',
      courseId: course.value.id,
      lessonId: currentLesson.value.id
    })
    
    currentLesson.value.completed = true
    ElMessage.success('课程标记为已完成')
    
    // 更新学习进度
    updateProgress()
  } catch (error) {
    console.error('标记完成失败:', error)
    ElMessage.error('标记完成失败')
  }
}

// 更新学习进度
const updateProgress = () => {
  if (!course.value?.outline) return
  
  let totalLessons = 0
  let completedLessons = 0
  
  course.value.outline.forEach((section: any) => {
    section.lessons.forEach((lesson: any) => {
      totalLessons++
      if (lesson.completed) {
        completedLessons++
      }
    })
  })
  
  course.value.progress = Math.round((completedLessons / totalLessons) * 100)
}

// 切换章节展开
const toggleSection = (index: number) => {
  const sectionIndex = expandedSections.value.indexOf(index)
  if (sectionIndex > -1) {
    expandedSections.value.splice(sectionIndex, 1)
  } else {
    expandedSections.value.push(index)
  }
}

// 获取完成数量
const getCompletedCount = (section: any) => {
  return section.lessons.filter((lesson: any) => lesson.completed).length
}

// 格式化时间
const formatTime = (seconds: number) => {
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
}

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('zh-CN')
}

// 计算属性
const completedLessons = computed(() => {
  if (!course.value?.outline) return 0
  let count = 0
  course.value.outline.forEach((section: any) => {
    section.lessons.forEach((lesson: any) => {
      if (lesson.completed) count++
    })
  })
  return count
})

const totalLessons = computed(() => {
  if (!course.value?.outline) return 0
  let count = 0
  course.value.outline.forEach((section: any) => {
    count += section.lessons.length
  })
  return count
})

const currentNotes = ref<any[]>([])
const currentQuestions = ref<any[]>([])

// 加载笔记
const loadNotes = async () => {
  if (!currentLesson.value) return
  
  try {
    const response = await ApiService.learning.getCourseNotes(
      authStore.user?.id || '1',
      course.value.id,
      currentLesson.value.id
    )
    currentNotes.value = response.data
  } catch (error) {
    console.error('加载笔记失败:', error)
  }
}

// 加载讨论
const loadDiscussions = async () => {
  if (!currentLesson.value) return
  
  try {
    const response = await ApiService.learning.getCourseDiscussions(
      course.value.id,
      currentLesson.value.id
    )
    currentQuestions.value = response.data
  } catch (error) {
    console.error('加载讨论失败:', error)
  }
}

// 笔记相关方法
const addNote = async () => {
  if (!currentLesson.value) return
  
  try {
    const noteContent = prompt('请输入笔记内容：')
    if (!noteContent) return
    
    await ApiService.learning.addNote({
      userId: authStore.user?.id || '1',
      courseId: course.value.id,
      lessonId: currentLesson.value.id,
      content: noteContent
    })
    
    ElMessage.success('笔记添加成功')
    loadNotes()
  } catch (error) {
    console.error('添加笔记失败:', error)
    ElMessage.error('添加笔记失败')
  }
}

const editNote = async (note: any) => {
  try {
    const newContent = prompt('编辑笔记内容：', note.content)
    if (!newContent || newContent === note.content) return
    
    await ApiService.learning.updateNote(note.id, newContent)
    ElMessage.success('笔记更新成功')
    loadNotes()
  } catch (error) {
    console.error('更新笔记失败:', error)
    ElMessage.error('更新笔记失败')
  }
}

const deleteNote = async (noteId: string) => {
  try {
    await ApiService.learning.deleteNote(noteId)
    ElMessage.success('笔记删除成功')
    loadNotes()
  } catch (error) {
    console.error('删除笔记失败:', error)
    ElMessage.error('删除笔记失败')
  }
}

// 讨论相关方法
const showAnswers = (question: any) => {
  ElMessage.info('查看回答功能')
}

const likeQuestion = async (question: any) => {
  try {
    await ApiService.learning.likeDiscussion(question.id)
    ElMessage.success('点赞成功')
    loadDiscussions()
  } catch (error) {
    console.error('点赞失败:', error)
    ElMessage.error('点赞失败')
  }
}

const submitQuestion = async () => {
  if (!questionForm.value.title || !questionForm.value.content) {
    ElMessage.warning('请填写完整的问题信息')
    return
  }
  
  try {
    await ApiService.learning.addDiscussion({
      courseId: course.value.id,
      lessonId: currentLesson.value?.id || '',
      userId: authStore.user?.id || '1',
      userName: authStore.user?.identities?.[0]?.name || '学员',
      userAvatar: authStore.user?.avatar || 'https://via.placeholder.com/32x32/4CAF50/FFFFFF?text=U',
      title: questionForm.value.title,
      content: questionForm.value.content
    })
    
    ElMessage.success('问题提交成功')
    showAskQuestion.value = false
    questionForm.value = { title: '', content: '' }
    loadDiscussions()
  } catch (error) {
    console.error('提交问题失败:', error)
    ElMessage.error('提交问题失败')
  }
}

onMounted(() => {
  loadCourseDetail()
})
</script>

<style scoped lang="scss">
.course-learning-page {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: var(--bg-primary);
}

.loading-container {
  padding: var(--spacing-xl);
}

.learning-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.learning-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-lg) var(--spacing-xl);
  background: var(--bg-card);
  border-bottom: 1px solid var(--border-color-light);
  
  .course-info {
    .course-title {
      font-size: var(--font-size-h3);
      font-weight: var(--font-weight-bold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-sm) 0;
    }
    
    .course-meta {
      display: flex;
      gap: var(--spacing-lg);
      font-size: var(--font-size-medium);
      color: var(--text-secondary);
      
      .progress {
        color: var(--primary-color);
        font-weight: var(--font-weight-medium);
      }
    }
  }
  
  .learning-actions {
    display: flex;
    gap: var(--spacing-md);
  }
}

.learning-layout {
  flex: 1;
  display: grid;
  grid-template-columns: 1fr 350px;
  gap: 0;
  overflow: hidden;
  
  @media (max-width: 1024px) {
    grid-template-columns: 1fr;
  }
}

.video-section {
  display: flex;
  flex-direction: column;
  background: var(--bg-card);
  
  .video-player {
    flex: 1;
    display: flex;
    flex-direction: column;
    
    .video-container {
      flex: 1;
      display: flex;
      flex-direction: column;
      background: #000;
    }
    
    .video-placeholder {
      flex: 1;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      color: white;
      text-align: center;
      
      h3 {
        margin: var(--spacing-lg) 0 var(--spacing-md) 0;
        font-size: var(--font-size-h4);
      }
      
      p {
        margin: 0 0 var(--spacing-lg) 0;
        opacity: 0.8;
      }
    }
    
    .video-controls {
      background: rgba(0, 0, 0, 0.8);
      padding: var(--spacing-md);
      
      .progress-bar {
        margin-bottom: var(--spacing-sm);
      }
      
      .control-buttons {
        display: flex;
        align-items: center;
        gap: var(--spacing-md);
        
        .time-display {
          color: white;
          font-size: var(--font-size-small);
          margin-left: auto;
        }
      }
    }
    
    .no-lesson {
      flex: 1;
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
  
  .lesson-info {
    padding: var(--spacing-lg);
    border-top: 1px solid var(--border-color-light);
    
    h3 {
      font-size: var(--font-size-h4);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-sm) 0;
    }
    
    p {
      font-size: var(--font-size-medium);
      color: var(--text-secondary);
      margin: 0 0 var(--spacing-lg) 0;
      line-height: 1.6;
    }
    
    .lesson-actions {
      display: flex;
      gap: var(--spacing-md);
    }
  }
}

.outline-section {
  background: var(--bg-secondary);
  border-left: 1px solid var(--border-color-light);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  
  .outline-header {
    padding: var(--spacing-lg);
    border-bottom: 1px solid var(--border-color-light);
    
    h3 {
      font-size: var(--font-size-h4);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0 0 var(--spacing-md) 0;
    }
    
    .outline-stats {
      display: flex;
      flex-direction: column;
      gap: var(--spacing-sm);
      
      span {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
      }
    }
  }
  
  .outline-content {
    flex: 1;
    overflow-y: auto;
    padding: var(--spacing-md);
  }
}

.outline-section {
  margin-bottom: var(--spacing-lg);
  
  &:last-child {
    margin-bottom: 0;
  }
  
  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: var(--spacing-md);
    background: var(--bg-card);
    border-radius: var(--border-radius-medium);
    cursor: pointer;
    transition: all var(--transition-normal);
    
    &:hover {
      background: var(--bg-tertiary);
    }
    
    .section-title {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
      font-size: var(--font-size-medium);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0;
      
      .section-number {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 24px;
        height: 24px;
        background: var(--primary-color);
        color: white;
        border-radius: 50%;
        font-size: var(--font-size-small);
        font-weight: var(--font-weight-bold);
      }
    }
    
    .section-meta {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
      font-size: var(--font-size-small);
      color: var(--text-secondary);
      
      .completed-count {
        color: var(--primary-color);
        font-weight: var(--font-weight-medium);
      }
      
      .toggle-icon {
        transition: transform var(--transition-normal);
        
        &.is-expanded {
          transform: rotate(180deg);
        }
      }
    }
  }
  
  .lessons-list {
    margin-top: var(--spacing-sm);
    display: flex;
    flex-direction: column;
    gap: 1px;
  }
  
  .lesson-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: var(--spacing-md);
    background: var(--bg-card);
    cursor: pointer;
    transition: all var(--transition-normal);
    
    &:hover {
      background: var(--bg-tertiary);
    }
    
    &.active {
      background: var(--primary-color);
      color: white;
      
      .lesson-title {
        color: white;
      }
      
      .lesson-duration {
        color: rgba(255, 255, 255, 0.8);
      }
    }
    
    &.completed {
      border-left: 3px solid #67c23a;
    }
    
    &.locked {
      opacity: 0.6;
      cursor: not-allowed;
      
      &:hover {
        background: var(--bg-card);
      }
    }
    
    .lesson-info {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
      flex: 1;
      
      .lesson-icon {
        color: var(--primary-color);
      }
      
      .lesson-title {
        font-size: var(--font-size-medium);
        color: var(--text-primary);
        flex: 1;
      }
    }
    
    .lesson-meta {
      display: flex;
      align-items: center;
      gap: var(--spacing-sm);
      
      .lesson-duration {
        font-size: var(--font-size-small);
        color: var(--text-secondary);
        min-width: 50px;
        text-align: right;
      }
    }
  }
}

// 对话框样式
.notes-content,
.discussion-content {
  .notes-header,
  .discussion-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-lg);
    padding-bottom: var(--spacing-md);
    border-bottom: 1px solid var(--border-color-light);
    
    h4 {
      font-size: var(--font-size-h5);
      font-weight: var(--font-weight-semibold);
      color: var(--text-primary);
      margin: 0;
    }
  }
}

.notes-list {
  .note-item {
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
    margin-bottom: var(--spacing-md);
    
    .note-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: var(--spacing-sm);
      
      .note-time {
        font-size: var(--font-size-small);
        color: var(--text-tertiary);
      }
      
      .note-actions {
        display: flex;
        gap: var(--spacing-xs);
      }
    }
    
    .note-content {
      font-size: var(--font-size-medium);
      color: var(--text-secondary);
      line-height: 1.6;
    }
  }
}

.discussion-list {
  .question-item {
    padding: var(--spacing-md);
    background: var(--bg-secondary);
    border-radius: var(--border-radius-medium);
    margin-bottom: var(--spacing-md);
    
    .question-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: var(--spacing-sm);
      
      .user-info {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        
        .user-name {
          font-size: var(--font-size-medium);
          font-weight: var(--font-weight-medium);
          color: var(--text-primary);
        }
        
        .question-time {
          font-size: var(--font-size-small);
          color: var(--text-tertiary);
        }
      }
    }
    
    .question-content {
      font-size: var(--font-size-medium);
      color: var(--text-secondary);
      line-height: 1.6;
      margin-bottom: var(--spacing-sm);
    }
    
    .question-actions {
      display: flex;
      gap: var(--spacing-md);
    }
  }
}

.no-notes,
.no-questions {
  text-align: center;
  padding: var(--spacing-xl) 0;
}

// 响应式设计
@media (max-width: 768px) {
  .learning-header {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: flex-start;
  }
  
  .learning-layout {
    grid-template-columns: 1fr;
  }
  
  .outline-section {
    border-left: none;
    border-top: 1px solid var(--border-color-light);
  }
}
</style> 