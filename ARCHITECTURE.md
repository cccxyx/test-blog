# 博客后端架构设计文档

## 技术栈

- **Web框架**: Gin
- **数据库**: SQLite3
- **ORM**: GORM
- **配置管理**: Viper
- **身份认证**: JWT
- **密码加密**: bcrypt
- **日志**: Gin自带日志 + 自定义中间件

## 项目目录结构

```
test-blog/
├── main.go                 # 程序入口
├── go.mod                  # Go模块依赖
├── go.sum                  # 依赖校验
├── config.yaml            # 配置文件
├── config.example.yaml    # 配置示例
├── blog.db                # SQLite数据库文件（运行时生成）
├── README.md              # 项目说明
├── ARCHITECTURE.md        # 架构设计文档
│
├── config/                # 配置管理
│   └── config.go         # 配置加载和定义
│
├── models/               # 数据模型
│   ├── user.go          # 用户模型
│   └── article.go       # 文章模型
│
├── database/            # 数据库
│   └── database.go     # 数据库连接和初始化
│
├── middleware/          # 中间件
│   ├── cors.go         # CORS中间件
│   ├── logger.go       # 日志中间件
│   └── auth.go         # JWT认证中间件
│
├── utils/              # 工具函数
│   ├── jwt.go         # JWT工具
│   ├── password.go    # 密码加密工具
│   └── response.go    # 统一响应格式
│
├── services/           # 业务逻辑层
│   ├── user.go        # 用户业务逻辑
│   └── article.go     # 文章业务逻辑
│
├── handlers/           # 控制器层
│   ├── user.go        # 用户相关接口
│   └── article.go     # 文章相关接口
│
└── router/            # 路由配置
    └── router.go     # 路由设置
```

## 数据模型设计

### User（用户表）

| 字段 | 类型 | 说明 | 约束 |
|------|------|------|------|
| id | uint | 主键 | 自增 |
| username | string | 用户名 | 唯一，非空 |
| password | string | 密码（加密） | 非空 |
| email | string | 邮箱 | 唯一，非空 |
| created_at | time.Time | 创建时间 | 自动 |
| updated_at | time.Time | 更新时间 | 自动 |

### Article（文章表）

| 字段 | 类型 | 说明 | 约束 |
|------|------|------|------|
| id | uint | 主键 | 自增 |
| title | string | 标题 | 非空 |
| content | text | 内容 | 非空 |
| user_id | uint | 作者ID | 外键 |
| created_at | time.Time | 创建时间 | 自动 |
| updated_at | time.Time | 更新时间 | 自动 |

**关系**: User 1:N Article（一个用户可以有多篇文章）

## 系统架构流程

### 请求处理流程

```
客户端请求
    ↓
CORS中间件 (允许跨域)
    ↓
日志中间件 (记录请求日志)
    ↓
路由匹配
    ↓
JWT认证中间件 (需要认证的路由)
    ↓
Handler处理器 (参数验证)
    ↓
Service业务层 (业务逻辑)
    ↓
Model数据层 (数据库操作)
    ↓
统一响应格式
    ↓
返回客户端
```

### 用户认证流程

```
1. 注册流程:
   用户提交注册信息 → 验证用户名/邮箱唯一性 → bcrypt加密密码 → 保存到数据库

2. 登录流程:
   用户提交登录信息 → 查询用户 → bcrypt验证密码 → 生成JWT token → 返回token

3. 认证流程:
   客户端携带token → JWT中间件验证 → 解析用户信息 → 传递给Handler
```

## 配置项说明

配置文件使用YAML格式，包含以下配置项：

```yaml
server:
  port: 8080              # 服务端口
  mode: debug            # 运行模式：debug/release

database:
  path: "./blog.db"      # SQLite数据库文件路径

jwt:
  secret: "your-secret-key"  # JWT密钥
  expire: 168            # token过期时间（小时）

cors:
  allow_origins: ["*"]   # 允许的来源
```

## 安全考虑

1. **密码安全**: 使用bcrypt加密存储密码，不存储明文
2. **JWT安全**: 使用强密钥，设置合理的过期时间
3. **权限控制**: 用户只能修改/删除自己的文章
4. **输入验证**: 对所有用户输入进行验证和清理
5. **CORS配置**: 根据实际需求配置允许的来源

## 日志格式

日志中间件记录以下信息：
- 请求时间
- 请求方法
- 请求路径
- 客户端IP
- 响应状态码
- 处理耗时

示例：
```
[2025-11-07 17:00:00] INFO | POST /api/login | 200 | 15ms | 127.0.0.1
```

## 开发步骤

按照以下顺序实现：

1. 项目初始化和依赖安装
2. 配置管理模块
3. 数据库连接和模型定义
4. 工具函数（JWT、密码加密、响应格式）
5. 中间件（CORS、日志、认证）
6. 业务逻辑层
7. 控制器层
8. 路由配置
9. 主程序入口
10. 测试和文档