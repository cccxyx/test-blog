# 系统设计可视化

## 系统架构图

```mermaid
graph TB
    Client[前端客户端] --> |HTTP请求| Server[Gin服务器]
    
    Server --> CORSMiddleware[CORS中间件]
    CORSMiddleware --> LogMiddleware[日志中间件]
    LogMiddleware --> Router[路由器]
    
    Router --> |公开接口| PublicHandler[公开处理器]
    Router --> |需要认证| AuthMiddleware[JWT认证中间件]
    
    PublicHandler --> |注册/登录| UserHandler[用户处理器]
    AuthMiddleware --> |认证通过| ProtectedHandler[受保护处理器]
    
    ProtectedHandler --> UserInfoHandler[用户信息处理器]
    ProtectedHandler --> ArticleHandler[文章处理器]
    
    UserHandler --> UserService[用户服务层]
    UserInfoHandler --> UserService
    ArticleHandler --> ArticleService[文章服务层]
    
    UserService --> |密码加密| PasswordUtil[密码工具]
    UserService --> |生成Token| JWTUtil[JWT工具]
    UserService --> UserModel[用户模型]
    
    ArticleService --> ArticleModel[文章模型]
    
    UserModel --> |GORM| Database[(SQLite3数据库)]
    ArticleModel --> |GORM| Database
    
    Config[Viper配置] --> |加载配置| Server
    Config --> Database
    Config --> JWTUtil
    
    style Client fill:#e1f5ff
    style Server fill:#fff4e6
    style Database fill:#f3e5f5
    style Config fill:#e8f5e9
```

## 请求流程图

### 用户注册流程

```mermaid
sequenceDiagram
    participant C as 客户端
    participant M as 中间件层
    participant H as 处理器层
    participant S as 服务层
    participant D as 数据库
    
    C->>M: POST /api/register
    M->>M: CORS验证
    M->>M: 记录日志
    M->>H: 路由到UserHandler
    H->>H: 验证参数
    H->>S: 调用注册服务
    S->>D: 检查用户名/邮箱唯一性
    D-->>S: 返回查询结果
    S->>S: bcrypt加密密码
    S->>D: 保存用户信息
    D-->>S: 返回保存结果
    S-->>H: 返回用户数据
    H-->>C: 返回成功响应
```

### 用户登录流程

```mermaid
sequenceDiagram
    participant C as 客户端
    participant M as 中间件层
    participant H as 处理器层
    participant S as 服务层
    participant J as JWT工具
    participant D as 数据库
    
    C->>M: POST /api/login
    M->>M: CORS验证
    M->>M: 记录日志
    M->>H: 路由到UserHandler
    H->>H: 验证参数
    H->>S: 调用登录服务
    S->>D: 查询用户
    D-->>S: 返回用户信息
    S->>S: bcrypt验证密码
    S->>J: 生成JWT Token
    J-->>S: 返回Token
    S-->>H: 返回Token和用户信息
    H-->>C: 返回登录成功响应
```

### 创建文章流程

```mermaid
sequenceDiagram
    participant C as 客户端
    participant M as 中间件层
    participant A as JWT中间件
    participant H as 处理器层
    participant S as 服务层
    participant D as 数据库
    
    C->>M: POST /api/articles<br/>Authorization: Bearer token
    M->>M: CORS验证
    M->>M: 记录日志
    M->>A: JWT认证
    A->>A: 验证Token
    A->>A: 解析用户信息
    A->>H: 传递用户ID到Context
    H->>H: 验证参数
    H->>S: 调用创建服务
    S->>D: 保存文章
    D-->>S: 返回保存结果
    S-->>H: 返回文章数据
    H-->>C: 返回创建成功响应
```

## 数据库关系图

```mermaid
erDiagram
    User ||--o{ Article : creates
    User {
        uint id PK
        string username UK
        string password
        string email UK
        datetime created_at
        datetime updated_at
    }
    Article {
        uint id PK
        string title
        text content
        uint user_id FK
        datetime created_at
        datetime updated_at
    }
```

## 项目分层架构

```mermaid
graph LR
    subgraph 表现层
        Handler[处理器层<br/>handlers/]
    end
    
    subgraph 业务层
        Service[服务层<br/>services/]
    end
    
    subgraph 数据层
        Model[模型层<br/>models/]
        DB[(数据库)]
    end
    
    subgraph 基础设施
        Middleware[中间件<br/>middleware/]
        Utils[工具函数<br/>utils/]
        Config[配置管理<br/>config/]
    end
    
    Handler --> Service
    Service --> Model
    Model --> DB
    Handler --> Utils
    Service --> Utils
    Middleware --> Handler
    Config --> Service
    Config --> Model
```

## API路由结构

```mermaid
graph TD
    Root["/api"] --> Public[公开路由]
    Root --> Protected[受保护路由]
    
    Public --> Register["POST /register<br/>用户注册"]
    Public --> Login["POST /login<br/>用户登录"]
    
    Protected --> UserInfo["GET /user/info<br/>获取用户信息"]
    Protected --> ArticleGroup[文章相关]
    
    ArticleGroup --> GetAllArticles["GET /articles<br/>获取所有文章"]
    ArticleGroup --> GetUserArticles["GET /articles/user/:id<br/>获取用户文章"]
    ArticleGroup --> CreateArticle["POST /articles<br/>创建文章"]
    ArticleGroup --> UpdateArticle["PUT /articles/:id<br/>更新文章"]
    ArticleGroup --> DeleteArticle["DELETE /articles/:id<br/>删除文章"]
    
    style Protected fill:#ffe0b2
    style Public fill:#c8e6c9
```

## 技术栈选择理由

| 技术 | 选择理由 |
|------|---------|
| Gin | 高性能、简洁的API、丰富的中间件生态 |
| GORM | 功能强大的ORM、支持多种数据库、易于使用 |
| SQLite3 | 轻量级、无需配置、适合小型项目 |
| JWT | 无状态认证、易于扩展、前后端分离友好 |
| bcrypt | 专为密码设计、慢速哈希、防止暴力破解 |
| Viper | 强大的配置管理、支持多种格式、易于使用 |

## 关键设计决策

### 1. 分层架构
- **Handler层**: 负责HTTP请求处理和参数验证
- **Service层**: 包含业务逻辑，可复用
- **Model层**: 数据模型和数据库操作

### 2. 中间件设计
- **CORS**: 允许前端跨域访问
- **Logger**: 统一日志格式，便于调试
- **Auth**: JWT认证，保护需要登录的接口

### 3. 统一响应格式
所有API返回一致的JSON格式，便于前端处理

### 4. 安全性考虑
- 密码使用bcrypt加密
- JWT token有效期限制
- 文章操作权限控制

### 5. 代码组织
- 按功能模块划分目录
- 每个模块职责单一
- 便于测试和维护