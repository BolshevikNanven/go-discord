# API 文档

## 1. 用户认证相关

### 登录
- **URL**: `/v1/auth/login`
- **方法**: `POST`
- **请求参数**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "登录成功",
      "data": {
        "accessToken": "string",
        "refreshToken": "string"
      }
    }
    ```

### 注册
- **URL**: `/v1/auth/register`
- **方法**: `POST`
- **请求参数**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "注册成功",
      "data": {
        "id": 123,
        "username": "string"
      }
    }
    ```

### 刷新令牌
- **URL**: `/v1/auth/refresh`
- **方法**: `POST`
- **请求参数**:
    ```json
    {
      "refreshToken": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "刷新令牌成功",
      "data": {
        "accessToken": "string",
        "refreshToken": "string"
      }
    }
    ```

## 2. 空间相关

### 创建空间
- **URL**: `/v1/space`
- **方法**: `POST`
- **请求参数**:
    ```json
    {
      "name": "string",
      "avatar": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "创建成功",
      "data": {
        "id": 123,
        "name": "string",
        "avatar": "string",
        "owner": 123
      }
    }
    ```

### 加入空间
- **URL**: `/v1/space/join/:space_id`
- **方法**: `POST`
- **请求参数**: 
    - URL 参数: `space_id` (int64)
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "加入成功",
      "data": {
        // 返回的空间信息
      }
    }
    ```

### 离开空间
- **URL**: `/v1/space/leave/:space_id`
- **方法**: `POST`
- **请求参数**: 
    - URL 参数: `space_id` (int64)
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "离开成功",
      "data": {
        // 返回的空间信息
      }
    }
    ```

### 获取空间列表
- **URL**: `/v1/space`
- **方法**: `GET`
- **请求参数**: 无
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "获取成功",
      "data": [
        {
          "id": 123,
          "name": "string",
          "avatar": "string",
          "owner": 123
        }
      ]
    }
    ```

### 更新空间
- **URL**: `/v1/space/:space_id`
- **方法**: `PUT`
- **请求参数**:
    - URL 参数: `space_id` (int64)
    - 请求体:
    ```json
    {
      "name": "string",
      "avatar": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "更新成功",
      "data": {
        // 返回的空间信息
      }
    }
    ```

### 删除空间
- **URL**: `/v1/space/:space_id`
- **方法**: `DELETE`
- **请求参数**: 
    - URL 参数: `space_id` (int64)
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "删除成功",
      "data": {
        // 返回的空间信息
      }
    }
    ```

## 3. 频道相关

### 创建频道
- **URL**: `/v1/channel`
- **方法**: `POST`
- **请求参数**:
    ```json
    {
      "spaceId": 123,
      "name": "string",
      "type": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "创建成功",
      "data": {
        // 返回的频道信息
      }
    }
    ```

### 加入频道
- **URL**: `/v1/channel/join/:channel_id`
- **方法**: `POST`
- **请求参数**: 
    - URL 参数: `channel_id` (int64)
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "加入成功",
      "data": {
        // 返回的频道信息
      }
    }
    ```

### 离开频道
- **URL**: `/v1/channel/leave/:channel_id`
- **方法**: `POST`
- **请求参数**: 
    - URL 参数: `channel_id` (int64)
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "离开成功",
      "data": {
        // 返回的频道信息
      }
    }
    ```

### 获取频道列表
- **URL**: `/v1/channel`
- **方法**: `GET`
- **请求参数**: 无
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "获取成功",
      "data": [
        {
          "id": 123,
          "name": "string",
          "owner": 123
        }
      ]
    }
    ```

### 更新频道
- **URL**: `/v1/channel/:channel_id`
- **方法**: `PUT`
- **请求参数**:
    - URL 参数: `channel_id` (int64)
    - 请求体:
    ```json
    {
      "name": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "更新成功",
      "data": {
        // 返回的频道信息
      }
    }
    ```

### 删除频道
- **URL**: `/v1/channel/:channel_id`
- **方法**: `DELETE`
- **请求参数**: 
    - URL 参数: `channel_id` (int64)
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "删除成功",
      "data": {
        // 返回的频道信息
      }
    }
    ```

## 4. 消息相关

### 发送消息
- **URL**: `/v1/message`
- **方法**: `POST`
- **请求参数**:
    ```json
    {
      "space_id": 123,
      "to": 123,
      "type": "string",
      "content": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "消息发送成功",
      "data": {
        "message_id": 123
      }
    }
    ```

### 确认消息
- **URL**: `/v1/message/ack`
- **方法**: `POST`
- **请求参数**:
    ```json
    {
      "space_id": 123,
      "message_ids": [123, 456]
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "消息确认成功",
      "data": {
        "success": true
      }
    }
    ```

### 获取历史消息
- **URL**: `/v1/message`
- **方法**: `GET`
- **请求参数**:
    ```json
    {
      "space_id": 123,
      "channel_id": 456,
      "from": 123,
      "cursor": 0,
      "limit": 10
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "历史消息获取成功",
      "data": {
        "messages": [
          {
            "id": 123,
            "space_id": 123,
            "from": 123,
            "to": 456,
            "type": "string",
            "content": "string",
            "created_at": 1234567890
          }
        ],
        "cursor": 0,
        "limit": 10
      }
    }
    ```

### 获取收件箱消息
- **URL**: `/v1/message/inbox`
- **方法**: `GET`
- **请求参数**:
    ```json
    {
      "space_id": 123,
      "limit": 10
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "收件箱消息获取成功",
      "data": {
        "messages": [
          {
            "id": 123,
            "space_id": 123,
            "from": 123,
            "to": 456,
            "type": "string",
            "content": "string",
            "created_at": 1234567890
          }
        ]
      }
    }
    ```

### 发送频道消息
- **URL**: `/v1/message/channel/:channel_id`
- **方法**: `POST`
- **请求参数**:
    - URL 参数: `channel_id` (int64)
    - 请求体:
    ```json
    {
      "space_id": 123,
      "to": 123,
      "type": "string",
      "content": "string"
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "消息发送成功",
      "data": {
        "message_id": 123
      }
    }
    ```

### 确认频道消息
- **URL**: `/v1/message/channel/:channel_id/ack`
- **方法**: `POST`
- **请求参数**:
    - URL 参数: `channel_id` (int64)
    - 请求体:
    ```json
    {
      "space_id": 123,
      "message_ids": [123, 456]
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "消息确认成功",
      "data": {
        "success": true
      }
    }
    ```

### 获取频道历史消息
- **URL**: `/v1/message/channel/:channel_id`
- **方法**: `GET`
- **请求参数**:
    - URL 参数: `channel_id` (int64)
    - 请求体:
    ```json
    {
      "space_id": 123,
      "from": 123,
      "cursor": 0,
      "limit": 10
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "历史消息获取成功",
      "data": {
        "messages": [
          {
            "id": 123,
            "space_id": 123,
            "from": 123,
            "to": 456,
            "type": "string",
            "content": "string",
            "created_at": 1234567890
          }
        ],
        "cursor": 0,
        "limit": 10
      }
    }
    ```

### 获取频道收件箱消息
- **URL**: `/v1/message/channel/:channel_id/inbox`
- **方法**: `GET`
- **请求参数**:
    - URL 参数: `channel_id` (int64)
    - 请求体:
    ```json
    {
      "space_id": 123,
      "limit": 10
    }
    ```
- **响应结果**:
    ```json
    {
      "code": 0,
      "message": "收件箱消息获取成功",
      "data": {
        "messages": [
          {
            "id": 123,
            "space_id": 123,
            "from": 123,
            "to": 456,
            "type": "string",
            "content": "string",
            "created_at": 1234567890
          }
        ]
      }
    }
    ```
