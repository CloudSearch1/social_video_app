# 抖音后端服务 (Go语言实现)

## 项目结构

```
backend/
├── main.go          # 应用入口
├── go.mod           # 依赖管理
├── .env             # 环境配置
├── config/          # 配置文件
├── controllers/     # 控制器
├── middleware/      # 中间件
├── models/          # 数据模型
├── routes/          # 路由定义
├── utils/           # 工具函数
└── uploads/         # 上传文件存储目录
```

## 技术栈

- **语言**: Go 1.13+
- **Web框架**: Fiber v2.52.0
- **数据库**: MongoDB
- **认证**: JWT v5.2.0
- **环境管理**: godotenv v1.5.1

## 运行项目

1. 安装依赖:
   ```
   go mod tidy
   ```

2. 配置环境变量:
   编辑 `.env` 文件设置数据库连接和其他配置

3. 运行服务:
   ```
   go run main.go
   ```

服务将在 `http://localhost:8080` 启动

## API端点

### 用户相关
- `POST /api/users/register` - 用户注册
- `POST /api/users/login` - 用户登录
- `GET /api/users/profile` - 获取用户资料 (需要认证)
- `PUT /api/users/profile` - 更新用户资料 (需要认证)

### 视频相关
- `POST /api/videos/upload` - 上传视频 (需要认证)
- `GET /api/videos/list` - 获取视频列表
- `GET /api/videos/:id` - 获取视频详情
- `POST /api/videos/:id/like` - 点赞视频 (需要认证)
- `POST /api/videos/:id/comment` - 评论视频 (需要认证)

### 健康检查
- `GET /health` - 服务健康检查

## 开发指南

1. 添加新功能时，请遵循现有代码结构
2. 在 `routes/` 中定义路由
3. 在 `controllers/` 中实现业务逻辑
4. 在 `models/` 中定义数据模型
5. 在 `middleware/` 中实现中间件

## 数据库设计

### 用户模型 (User)
- ID: MongoDB ObjectId
- Username: 用户名
- Email: 邮箱
- Password: 密码 (加密存储)
- CreatedAt: 创建时间
- UpdatedAt: 更新时间

### 视频模型 (Video)
- ID: MongoDB ObjectId
- Title: 视频标题
- Description: 视频描述
- URL: 视频文件地址
- UserID: 上传用户ID
- Likes: 点赞用户ID列表
- Comments: 评论列表
- ViewCount: 观看次数
- CreatedAt: 创建时间
- UpdatedAt: 更新时间

### 评论模型 (Comment)
- ID: MongoDB ObjectId
- UserID: 评论用户ID
- Content: 评论内容
- CreatedAt: 创建时间

## 认证流程

1. 用户注册时，密码会被bcrypt加密存储
2. 用户登录时，服务器验证邮箱和密码，生成JWT token
3. 客户端在后续请求中通过Authorization头携带token
4. 服务器通过中间件验证token有效性

## 依赖管理

使用 Go Modules 进行依赖管理:

```bash
go mod tidy  # 整理依赖
go mod download  # 下载依赖
go mod vendor  # 创建 vendor 目录
```