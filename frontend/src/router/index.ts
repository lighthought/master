import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/HomePage.vue'),
    meta: { title: '首页' }
  },
  {
    path: '/mentors',
    name: 'Mentors',
    component: () => import('@/views/MentorsPage.vue'),
    meta: { title: '大师' }
  },
  {
    path: '/courses',
    name: 'Courses',
    component: () => import('@/views/CoursesPage.vue'),
    meta: { title: '课程' }
  },
  {
    path: '/community',
    name: 'Community',
    component: () => import('@/views/CommunityPage.vue'),
    meta: { title: '社群' }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/ProfilePage.vue'),
    meta: { title: '个人中心' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  // 设置页面标题
  document.title = to.meta.title ? `${to.meta.title} - Master Guide` : 'Master Guide'
  next()
})

export default router