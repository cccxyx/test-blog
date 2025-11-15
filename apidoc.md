## API DOC

[APIPOST 在线文档](https://docs.apipost.net/docs/54c01e5e0888000?locale=zh-cn)

### 统一响应格式

所有API返回格式：

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

### 错误码说明

- `200`: 请求成功
- `400`: 请求参数错误
- `401`: 未授权（token无效或未提供）
- `403`: 禁止访问（权限不足）
- `404`: 资源不存在
- `500`: 服务器内部错误

### 用户相关接口

#### 1. 用户注册

**接口**: `POST /api/register`

**请求体**:
```json
{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user_id": 1,
    "username": "testuser",
    "email": "test@example.com"
  }
}
```

#### 2. 用户登录

**接口**: `POST /api/login`

**请求体**:
```json
{
  "username": "testuser",
  "password": "123456"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2025-11-07 17:00:00"
    }
  }
}
```

#### 3. 获取当前用户信息

**接口**: `GET /api/user/info`

**请求头**: `Authorization: Bearer {token}`

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2025-11-07 17:00:00"
  }
}
```

### 文章相关接口

#### 4. 获取所有文章

**接口**: `GET /api/articles`

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "title": "文章标题",
      "content": "文章内容",
      "user_id": 1,
      "author": {
        "id": 1,
        "username": "testuser",
        "email": "test@example.com"
      },
      "created_at": "2025-11-07 17:00:00",
      "updated_at": "2025-11-07 17:00:00"
    }
  ]
}
```

#### 5. 获取指定用户的文章

**接口**: `GET /api/articles/user/:user_id`

**路径参数**: `user_id` - 用户ID

**响应**: 同获取所有文章

#### 6. 根据ID获取文章详情

**接口**: `GET /api/articles/:id`

**路径参数**: `id` - 文章ID

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "title": "文章标题",
    "content": "文章内容",
    "user_id": 1,
    "author": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com"
    },
    "created_at": "2025-11-07 17:00:00",
    "updated_at": "2025-11-07 17:00:00"
  }
}
```

**错误响应示例**:
```json
{
  "code": 404,
  "message": "文章不存在",
  "data": null
}
```

#### 7. 创建文章

**接口**: `POST /api/articles`

**请求头**: `Authorization: Bearer {token}`

**请求体**:
```json
{
  "title": "我的第一篇文章",
  "content": "这是文章的内容..."
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "创建成功",
  "data": {
    "id": 1,
    "title": "我的第一篇文章",
    "content": "这是文章的内容...",
    "user_id": 1,
    "author": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com"
    },
    "created_at": "2025-11-07 17:00:00",
    "updated_at": "2025-11-07 17:00:00"
  }
}
```

#### 8. 更新文章

**接口**: `PUT /api/articles/:id`

**请求头**: `Authorization: Bearer {token}`

**路径参数**: `id` - 文章ID

**请求体**:
```json
{
  "title": "更新后的标题",
  "content": "更新后的内容..."
}
```

**响应**: 同创建文章

#### 9. 删除文章

**接口**: `DELETE /api/articles/:id`

**请求头**: `Authorization: Bearer {token}`

**路径参数**: `id` - 文章ID

**响应示例**:
```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```
