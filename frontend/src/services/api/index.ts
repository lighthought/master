// 导入所有API模块
import { authApi } from './auth'
import { userApi } from './user'
import { mentorApi } from './mentor'
import { appointmentApi } from './appointment'
import { courseApi } from './course'
import { learningApi } from './learning'
import { circlesApi } from './circles'
import { postsApi } from './posts'
import { incomeApi } from './income'
import { studentApi } from './student'

// 统一的API服务类
export class ApiService {
  // 认证相关API
  static auth = authApi

  // 用户管理API
  static user = userApi

  // 大师管理API
  static mentors = mentorApi

  // 预约管理API
  static appointments = appointmentApi

  // 课程管理API
  static courses = courseApi

  // 学习记录管理API
  static learning = learningApi

  // 圈子管理API
  static circles = circlesApi

  // 帖子管理API
  static posts = postsApi

  // 收入管理API
  static income = incomeApi

  // 学生管理API
  static students = studentApi
}

// 导出默认实例
export default ApiService 