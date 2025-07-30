import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const routes: RouteRecordRaw[] = [
  {
    path: '/auth',
    name: 'Auth',
    component: () => import('@/views/AuthPage.vue'),
    meta: { title: '登录/注册', requiresGuest: true }
  },
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/HomePage.vue'),
    meta: { title: '首页', requiresAuth: true }
  },
  {
    path: '/mentors',
    name: 'Mentors',
    component: () => import('@/views/MentorsPage.vue'),
    meta: { title: '大师', requiresAuth: true }
  },
  {
    path: '/mentors/:id',
    name: 'MentorDetail',
    component: () => import('@/views/MentorDetailPage.vue'),
    meta: { title: '大师详情', requiresAuth: true }
  },
  {
    path: '/bookings',
    name: 'Bookings',
    component: () => import('@/views/BookingsPage.vue'),
    meta: { title: '我的预约', requiresAuth: true }
  },
  {
    path: '/booking-success/:bookingId',
    name: 'BookingSuccess',
    component: () => import('@/views/BookingSuccessPage.vue'),
    meta: { title: '预约成功', requiresAuth: true }
  },
  {
    path: '/mentor-bookings',
    name: 'MentorBookings',
    component: () => import('@/views/MentorBookingsPage.vue'),
    meta: { title: '预约管理', requiresAuth: true }
  },
  {
    path: '/courses',
    name: 'Courses',
    component: () => import('@/views/CoursesPage.vue'),
    meta: { title: '课程', requiresAuth: true }
  },
  {
    path: '/courses/:id',
    name: 'CourseDetail',
    component: () => import('@/views/CourseDetailPage.vue'),
    meta: { title: '课程详情', requiresAuth: true }
  },
  {
    path: '/courses/:id/enroll',
    name: 'CourseEnroll',
    component: () => import('@/views/CourseEnrollPage.vue'),
    meta: { title: '课程报名', requiresAuth: true }
  },
  {
    path: '/enroll-success/:id',
    name: 'EnrollSuccess',
    component: () => import('@/views/EnrollSuccessPage.vue'),
    meta: { title: '报名成功', requiresAuth: true }
  },
  {
    path: '/learning/:id',
    name: 'CourseLearning',
    component: () => import('@/views/CourseLearningPage.vue'),
    meta: { title: '课程学习', requiresAuth: true }
  },
  {
    path: '/circles',
    name: 'Circles',
    component: () => import('@/views/CirclesPage.vue'),
    meta: { title: '圈子', requiresAuth: true }
  },
  {
    path: '/posts',
    name: 'Posts',
    component: () => import('@/views/PostsPage.vue'),
    meta: { title: '动态', requiresAuth: true }
  },
  {
    path: '/community',
    name: 'Community',
    component: () => import('@/views/CommunityPage.vue'),
    meta: { title: '社群', requiresAuth: true }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/ProfilePage.vue'),
    meta: { title: '个人中心', requiresAuth: true }
  },
  {
    path: '/learning-records',
    name: 'LearningRecords',
    component: () => import('@/views/LearningRecordsPage.vue'),
    meta: { title: '学习记录', requiresAuth: true }
  },
  {
    path: '/student-management',
    name: 'StudentManagement',
    component: () => import('@/views/StudentManagementPage.vue'),
    meta: { title: '学生管理', requiresAuth: true }
  },
  {
    path: '/income-stats',
    name: 'IncomeStats',
    component: () => import('@/views/IncomeStatsPage.vue'),
    meta: { title: '收入统计', requiresAuth: true }
  },
  {
    path: '/identity',
    name: 'IdentityManagement',
    component: () => import('@/views/IdentityManagementPage.vue'),
    meta: { title: '身份管理', requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  // 设置页面标题
  document.title = to.meta.title ? `${to.meta.title} - Master Guide` : 'Master Guide'
  
  const authStore = useAuthStore()
  
  // 需要认证的页面
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/auth')
    return
  }
  
  // 需要访客的页面（登录/注册）
  if (to.meta.requiresGuest && authStore.isAuthenticated) {
    next('/')
    return
  }
  
  next()
})

export default router