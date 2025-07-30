# Master Guide UI/UX 设计规范

## 1. 设计理念

### 1.1 设计原则
- **用户至上**：以用户需求为中心，提供直观易用的界面
- **专业可信**：深色主题营造专业感，建立用户信任
- **技艺传承**：通过视觉设计体现传统与现代技艺的结合
- **双重身份**：清晰的身份标识和切换机制
- **移动优先**：响应式设计，优先考虑移动端体验

### 1.2 设计目标
- 打造专业、现代的技艺传承平台
- 支持用户在不同身份间无缝切换
- 提供流畅的学习和指导体验
- 建立活跃的社群互动环境

## 2. 色彩系统

### 2.1 主色调
```scss
// 主色调 - 体现专业和活力
--primary-color: #FF6B35;        // 橙色 - 主要操作按钮
--primary-light: #FF8A65;        // 浅橙色 - 次要操作
--primary-dark: #E55A2B;         // 深橙色 - 悬停状态

// 辅助色 - 体现技艺传承
--secondary-color: #FFD93D;      // 金黄色 - 强调和奖励
--secondary-light: #FFE082;      // 浅金色 - 背景强调
--secondary-dark: #FBC02D;       // 深金色 - 重要信息
```

### 2.2 背景色系
```scss
// 深色主题背景
--bg-primary: #1A1A1A;           // 主背景色
--bg-secondary: #2D2D2D;         // 次级背景色
--bg-tertiary: #404040;          // 三级背景色
--bg-card: #333333;              // 卡片背景色
--bg-overlay: rgba(0, 0, 0, 0.7); // 遮罩层
```

### 2.3 文字色系
```scss
// 文字颜色
--text-primary: #FFFFFF;         // 主要文字
--text-secondary: #CCCCCC;       // 次要文字
--text-tertiary: #999999;        // 辅助文字
--text-disabled: #666666;        // 禁用文字
--text-inverse: #1A1A1A;         // 反色文字
```

### 2.4 功能色系
```scss
// 状态颜色
--success-color: #4CAF50;        // 成功状态
--warning-color: #FF9800;        // 警告状态
--danger-color: #F44336;         // 错误状态
--info-color: #2196F3;           // 信息状态

// 大师/学徒身份色
--master-color: #FF6B35;         // 大师身份标识
--apprentice-color: #4CAF50;     // 学徒身份标识
--dual-color: #FFD93D;           // 双重身份标识
```

## 3. 字体规范

### 3.1 字体族
```scss
// 中文字体
--font-family-cn: 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif;

// 英文字体
--font-family-en: 'SF Pro Display', 'Segoe UI', 'Roboto', sans-serif;

// 数字字体
--font-family-number: 'SF Mono', 'Monaco', 'Inconsolata', monospace;
```

### 3.2 字体大小
```scss
// 标题字体
--font-size-h1: 32px;            // 页面主标题
--font-size-h2: 24px;            // 模块标题
--font-size-h3: 20px;            // 卡片标题
--font-size-h4: 18px;            // 小节标题
--font-size-h5: 16px;            // 小标题

// 正文字体
--font-size-large: 16px;         // 大号正文
--font-size-medium: 14px;        // 中号正文
--font-size-small: 12px;         // 小号正文
--font-size-mini: 10px;          // 迷你文字
```

### 3.3 字重规范
```scss
--font-weight-light: 300;        // 细体
--font-weight-normal: 400;       // 常规
--font-weight-medium: 500;       // 中等
--font-weight-semibold: 600;     // 半粗
--font-weight-bold: 700;         // 粗体
```

## 4. 间距系统

### 4.1 基础间距
```scss
--spacing-xs: 4px;               // 超小间距
--spacing-sm: 8px;               // 小间距
--spacing-md: 16px;              // 中间距
--spacing-lg: 24px;              // 大间距
--spacing-xl: 32px;              // 超大间距
--spacing-xxl: 48px;             // 特大间距
```

### 4.2 组件间距
```scss
// 卡片间距
--card-padding: 16px;
--card-margin: 12px;
--card-border-radius: 12px;

// 按钮间距
--button-padding-x: 16px;
--button-padding-y: 8px;
--button-border-radius: 8px;

// 输入框间距
--input-padding: 12px;
--input-border-radius: 8px;
```

## 5. 阴影系统

### 5.1 阴影层级
```scss
// 阴影定义
--shadow-light: 0 2px 8px rgba(0, 0, 0, 0.1);
--shadow-medium: 0 4px 16px rgba(0, 0, 0, 0.15);
--shadow-heavy: 0 8px 32px rgba(0, 0, 0, 0.2);
--shadow-card: 0 2px 12px rgba(0, 0, 0, 0.1);
--shadow-button: 0 2px 4px rgba(0, 0, 0, 0.1);
```

## 6. 圆角规范

### 6.1 圆角大小
```scss
--border-radius-small: 4px;      // 小圆角
--border-radius-medium: 8px;     // 中圆角
--border-radius-large: 12px;     // 大圆角
--border-radius-xl: 16px;        // 超大圆角
--border-radius-full: 50%;       // 圆形
```

## 7. 动画规范

### 7.1 过渡时间
```scss
--transition-fast: 0.15s;        // 快速过渡
--transition-normal: 0.25s;      // 正常过渡
--transition-slow: 0.35s;        // 慢速过渡
```

### 7.2 缓动函数
```scss
--ease-in: cubic-bezier(0.4, 0, 1, 1);
--ease-out: cubic-bezier(0, 0, 0.2, 1);
--ease-in-out: cubic-bezier(0.4, 0, 0.2, 1);
```

## 8. 响应式断点

### 8.1 断点定义
```scss
// 移动端优先
--breakpoint-xs: 480px;          // 超小屏幕
--breakpoint-sm: 768px;          // 小屏幕
--breakpoint-md: 1024px;         // 中等屏幕
--breakpoint-lg: 1200px;         // 大屏幕
--breakpoint-xl: 1440px;         // 超大屏幕
```

## 9. 图标规范

### 9.1 图标尺寸
```scss
--icon-size-xs: 12px;            // 超小图标
--icon-size-sm: 16px;            // 小图标
--icon-size-md: 20px;            // 中图标
--icon-size-lg: 24px;            // 大图标
--icon-size-xl: 32px;            // 超大图标
```

### 9.2 图标风格
- **线性图标**：简洁的线条风格
- **填充图标**：重要操作和状态
- **双色图标**：特殊功能和标识

## 10. 组件状态

### 10.1 按钮状态
```scss
// 默认状态
--button-bg-default: var(--primary-color);
--button-text-default: var(--text-primary);

// 悬停状态
--button-bg-hover: var(--primary-dark);
--button-shadow-hover: var(--shadow-medium);

// 激活状态
--button-bg-active: var(--primary-dark);
--button-transform-active: scale(0.98);

// 禁用状态
--button-bg-disabled: var(--bg-tertiary);
--button-text-disabled: var(--text-disabled);
```

### 10.2 输入框状态
```scss
// 默认状态
--input-border-default: var(--border-color);
--input-bg-default: var(--bg-secondary);

// 聚焦状态
--input-border-focus: var(--primary-color);
--input-shadow-focus: 0 0 0 2px rgba(255, 107, 53, 0.2);

// 错误状态
--input-border-error: var(--danger-color);
--input-shadow-error: 0 0 0 2px rgba(244, 67, 54, 0.2);
```

## 11. 身份标识设计

### 11.1 身份徽章
```scss
// 大师身份
--master-badge-bg: var(--master-color);
--master-badge-text: var(--text-primary);
--master-badge-border: 2px solid var(--master-color);

// 学徒身份
--apprentice-badge-bg: var(--apprentice-color);
--apprentice-badge-text: var(--text-primary);
--apprentice-badge-border: 2px solid var(--apprentice-color);

// 双重身份
--dual-badge-bg: linear-gradient(45deg, var(--master-color), var(--apprentice-color));
--dual-badge-text: var(--text-primary);
--dual-badge-border: 2px solid transparent;
```

### 11.2 身份切换器
```scss
--identity-switcher-bg: var(--bg-secondary);
--identity-switcher-border: 1px solid var(--border-color);
--identity-switcher-radius: var(--border-radius-medium);
--identity-switcher-shadow: var(--shadow-light);
```

## 12. 内容层级

### 12.1 信息层级
- **主要信息**：标题、重要数据、核心功能
- **次要信息**：描述、辅助数据、可选功能
- **辅助信息**：提示、说明、帮助文本

### 12.2 视觉层级
- **最高层级**：模态框、下拉菜单、通知
- **高层级**：卡片、按钮、输入框
- **中层级**：背景、分割线、图标
- **低层级**：阴影、边框、装饰元素

## 13. 无障碍设计

### 13.1 对比度
- 文字与背景对比度 ≥ 4.5:1
- 大号文字对比度 ≥ 3:1
- 图标与背景对比度 ≥ 3:1

### 13.2 交互设计
- 点击区域 ≥ 44px × 44px
- 支持键盘导航
- 提供焦点指示器
- 支持屏幕阅读器

## 14. 设计检查清单

### 14.1 视觉检查
- [ ] 色彩使用符合品牌规范
- [ ] 字体大小和字重合适
- [ ] 间距和留白合理
- [ ] 圆角和阴影一致
- [ ] 图标风格统一

### 14.2 交互检查
- [ ] 按钮状态反馈清晰
- [ ] 输入框验证及时
- [ ] 加载状态明确
- [ ] 错误提示友好
- [ ] 成功反馈及时

### 14.3 响应式检查
- [ ] 移动端布局合理
- [ ] 平板端适配良好
- [ ] 桌面端体验优化
- [ ] 横屏模式支持
- [ ] 不同设备兼容

---

**文档版本**：v1.0.0  
**创建日期**：2024年12月  
**设计负责人**：Sally (UX Expert) 