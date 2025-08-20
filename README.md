# 抖音类短视频应用

一个基于Go和Vue 3开发的全栈短视频应用，类似于抖音的社交视频平台。

## 项目特点

- 🚀 高性能：后端采用Go语言和Fiber框架，前端使用Vue 3和Vite
- 📱 移动端优化：使用Vant UI组件库，专为移动端设计
- 🔐 安全认证：JWT Token认证机制
- 📦 现代化技术栈：Go + Fiber + MongoDB + Vue 3 + Vite + Pinia
- 🌐 响应式设计：适配不同屏幕尺寸

## 技术架构

### 后端技术栈
- **语言**: Go 1.21
- **框架**: Fiber v2
- **数据库**: MongoDB
- **认证**: JWT
- **工具**: 
  - github.com/gofiber/fiber/v2
  - go.mongodb.org/mongo-driver
  - github.com/golang-jwt/jwt/v5

### 前端技术栈
- **框架**: Vue 3
- **构建工具**: Vite
- **状态管理**: Pinia
- **路由**: Vue Router
- **UI组件库**: Vant
- **HTTP客户端**: Axios

## 项目结构

```
.
├── backend/          # 后端代码
│   ├── controllers/  # 控制器层
│   ├── models/       # 数据模型
│   ├── routes/       # 路由定义
│   ├── middleware/   # 中间件
│   ├── config/       # 配置文件
│   ├── utils/        # 工具函数
│   ├── uploads/      # 上传文件目录
│   ├── main.go       # 程序入口
│   └── go.mod        # Go模块定义
├── frontend/         # 前端代码
│   ├── src/          # 源代码目录
│   │   ├── assets/   # 静态资源
│   │   ├── components/ # 组件
│   │   ├── views/    # 页面视图
│   │   ├── router/   # 路由配置
│   │   ├── store/    # 状态管理
│   │   ├── utils/    # 工具函数
│   │   ├── App.vue   # 根组件
│   │   └── main.js   # 入口文件
│   ├── package.json  # 依赖配置
│   └── vite.config.js # 构建配置
└── README.md         # 项目说明文件
```

## 快速开始

### 后端服务

1. 进入后端目录：
   ```bash
   cd backend
   ```

2. 安装依赖：
   ```bash
   go mod tidy
   ```

3. 配置环境变量（创建 `.env` 文件）：
   ```bash
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET=your_secret_key
   ```

4. 运行服务：
   ```bash
   go run main.go
   ```

   服务将运行在 `http://localhost:8080`

### 前端服务

1. 进入前端目录：
   ```bash
   cd frontend
   ```

2. 安装依赖：
   ```bash
   yarn install
   ```

3. 运行开发服务器：
   ```bash
   yarn dev
   ```

   服务将运行在 `http://localhost:5173`

## API接口

### 用户相关
- `POST /api/users/register` - 用户注册
- `POST /api/users/login` - 用户登录
- `GET /api/users/profile` - 获取用户信息
- `PUT /api/users/profile` - 更新用户信息

### 视频相关
- `POST /api/videos/upload` - 上传视频
- `GET /api/videos/feed` - 获取视频流
- `GET /api/videos/:id` - 获取视频详情
- `POST /api/videos/:id/like` - 点赞视频
- `POST /api/videos/:id/comment` - 评论视频

## 部署

### 后端部署

```bash
go build -o douyin_app main.go
./douyin_app
```

### 前端部署

```bash
yarn build
```

构建产物位于 `dist/` 目录，可部署到Nginx、Apache等Web服务器。

## 开发规范

### Git提交规范
- `feat`: 新功能
- `fix`: 修复bug
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

### 目录命名规范
- 统一使用小写字母
- 多个单词使用下划线分隔

## 项目截图

（此处可添加项目运行截图）

## 贡献

欢迎提交Issue和Pull Request来改进本项目。

## 许可证

MIT