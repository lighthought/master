// 延迟函数
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟收入统计数据
const mockIncomeStats = {
  totalIncome: 15680,
  monthlyIncome: 3200,
  totalStudents: 24,
  averagePerStudent: 653,
  incomeGrowth: 15.6,
  topCourses: [
    {
      id: '1',
      title: 'Vue.js 基础入门',
      income: 4800,
      students: 12,
      averageRating: 4.8
    },
    {
      id: '2',
      title: 'Vue.js 进阶实战',
      income: 3600,
      students: 8,
      averageRating: 4.9
    },
    {
      id: '3',
      title: 'Vue.js 高级特性',
      income: 2800,
      students: 6,
      averageRating: 4.7
    }
  ],
  monthlyData: [
    { month: '2024-01', income: 2800, students: 5 },
    { month: '2024-02', income: 3200, students: 8 },
    { month: '2024-03', income: 3600, students: 12 },
    { month: '2024-04', income: 3200, students: 10 },
    { month: '2024-05', income: 3800, students: 15 },
    { month: '2024-06', income: 4200, students: 18 }
  ],
  recentTransactions: [
    {
      id: '1',
      studentName: '张学徒',
      courseTitle: 'Vue.js 基础入门',
      amount: 399,
      type: 'course_enrollment',
      status: 'completed',
      createdAt: '2024-03-20T14:30:00Z'
    },
    {
      id: '2',
      studentName: '李同学',
      courseTitle: 'Vue.js 进阶实战',
      amount: 599,
      type: 'course_enrollment',
      status: 'completed',
      createdAt: '2024-03-19T10:15:00Z'
    },
    {
      id: '3',
      studentName: '王学徒',
      courseTitle: '一对一指导',
      amount: 200,
      type: 'mentoring',
      status: 'completed',
      createdAt: '2024-03-18T16:45:00Z'
    },
    {
      id: '4',
      studentName: '赵同学',
      courseTitle: 'Vue.js 高级特性',
      amount: 799,
      type: 'course_enrollment',
      status: 'pending',
      createdAt: '2024-03-17T09:20:00Z'
    }
  ]
}

export const mockIncomeService = {
  // 获取收入统计数据
  async getIncomeStats(masterId: string, params: any = {}) {
    await delay(800)
    
    return {
      success: true,
      data: mockIncomeStats
    }
  },

  // 获取收入明细
  async getIncomeDetails(masterId: string, params: any = {}) {
    await delay(1000)
    
    let filteredTransactions = [...mockIncomeStats.recentTransactions]
    
    // 状态筛选
    if (params.status) {
      filteredTransactions = filteredTransactions.filter(transaction => 
        transaction.status === params.status
      )
    }
    
    // 类型筛选
    if (params.type) {
      filteredTransactions = filteredTransactions.filter(transaction => 
        transaction.type === params.type
      )
    }
    
    // 时间范围筛选
    if (params.startDate && params.endDate) {
      const startDate = new Date(params.startDate)
      const endDate = new Date(params.endDate)
      filteredTransactions = filteredTransactions.filter(transaction => {
        const transactionDate = new Date(transaction.createdAt)
        return transactionDate >= startDate && transactionDate <= endDate
      })
    }
    
    // 排序
    if (params.sortBy) {
      switch (params.sortBy) {
        case 'amount':
          filteredTransactions.sort((a, b) => b.amount - a.amount)
          break
        case 'date':
          filteredTransactions.sort((a, b) => 
            new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
          )
          break
        default:
          filteredTransactions.sort((a, b) => 
            new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
          )
      }
    }
    
    // 分页
    const page = params.page || 1
    const pageSize = params.pageSize || 10
    const start = (page - 1) * pageSize
    const end = start + pageSize
    const paginatedTransactions = filteredTransactions.slice(start, end)
    
    return {
      success: true,
      data: {
        transactions: paginatedTransactions,
        total: filteredTransactions.length,
        page,
        pageSize,
        totalPages: Math.ceil(filteredTransactions.length / pageSize)
      }
    }
  },

  // 导出收入报表
  async exportIncomeReport(masterId: string, params: any = {}) {
    await delay(2000)
    
    return {
      success: true,
      data: {
        downloadUrl: '/api/income/export/report.pdf',
        message: '报表导出成功'
      }
    }
  }
} 