# test-blog

一个简洁的博客后端，用于前端同学对接练习

## 技术栈

- **Go**: 1.25.0
- **Web框架**: Gin
- **数据库**: SQLite3
- **ORM**: GORM
- **认证**: JWT (golang-jwt/jwt/v5)
- **密码加密**: bcrypt
- **配置管理**: Viper

## 项目结构

```
test-blog/
├── config/                # 配置管理
│   └── config.go
├── database/             # 数据库连接
│   └── database.go
├── handlers/             # 请求处理器
│   ├── user.go
│   └── article.go
├── middleware/           # 中间件
│   ├── auth.go          # JWT认证
│   ├── cors.go          # CORS
│   └── logger.go        # 日志
├── models/              # 数据模型
│   ├── user.go
│   └── article.go
├── router/              # 路由配置
│   └── router.go
├── services/            # 业务逻辑层
│   ├── user.go
│   └── article.go
├── utils/               # 工具函数
│   ├── jwt.go
│   ├── password.go
│   └── response.go
├── main.go              # 程序入口
├── config.yaml          # 配置文件
└── blog.db              # SQLite数据库（运行时生成）
```

## 快速开始

### 1. 克隆项目

```bash
cd test-blog
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置文件

复制`config.example.yaml`到`config.yaml`，修改一下配置：

```yaml
server:
  port: "8080"      # 服务端口
  mode: "debug"     # 运行模式

database:
  path: "./blog.db" # 数据库路径

jwt:
  secret: "test-blog-secret-key-2025"  # JWT密钥
  expire: 168       # Token过期时间（小时）

cors:
  allow_origins:
    - "*"           # 允许的来源
```

### 4. 启动服务

```bash
go run main.go
```

## API 文档

[API文档](./apidoc.md)

## 功能特性

### 用户部分

- 用户注册
- 用户登录
- 根据token获取用户信息

### 文章部分

- 获取所有文章列表（按时间倒序）
- 获取指定用户的文章列表
- 创建新文章
- 更新文章内容
- 删除文章

## 拓展开发说明

### 添加新的API端点

1. 在 `handlers/` 中创建处理函数
2. 在 `services/` 中实现业务逻辑
3. 在 `router/router.go` 中注册路由
4. 需要认证的接口放在 `auth` 路由组下


## 项目文档

- [`ARCHITECTURE.md`](ARCHITECTURE.md) - 详细的架构设计文档
- [`SYSTEM_DESIGN.md`](SYSTEM_DESIGN.md) - 系统设计可视化图表