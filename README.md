# Activity 活动框架

## 简介
一个基于领域驱动设计(DDD)的活动框架，支持多种活动玩法的快速开发和部署。

## 核心特性
- 领域驱动设计：清晰的领域边界和职责划分
- 配置驱动：通过 JSON 配置文件定义活动
- 插件化设计：支持新玩法的快速接入
- 解耦设计：不依赖特定的UI框架或后端服务
- 状态管理：统一的活动状态管理机制

## 架构设计

### 领域驱动设计架构
1. 领域层（Domain Layer）
   - 活动（Activity）领域
   - 玩法（Game）领域
   - 奖品（Prize）领域
   - 参与（Participation）领域

2. 应用层（Application Layer）
   - 应用服务
   - DTO对象
   - 领域对象转换

3. 基础设施层（Infrastructure Layer）
   - 持久化实现
   - 配置管理
   - 公共组件

4. 接口层（Interface Layer）
   - HTTP接口
   - RPC接口
   - 中间件

### 核心组件
1. 活动（Activity）
   - 活动生命周期管理
   - 多玩法组合支持
   - 统一的状态管理

2. 玩法（Game）
   - 可扩展的玩法接口
   - 独立的业务逻辑
   - 配置化支持

3. 奖品（Prize）
   - 统一的奖品模型
   - 灵活的奖品发放机制

### 目录结构
```
.
├── api/                    # 接口层
│   ├── handler/           # HTTP处理器
│   ├── middleware/        # 中间件
│   ├── service/           # 应用服务
│   └── error.go           # 错误定义
├── constant/              # 常量定义
│   └── constant.go        # 错误码和消息
├── models/                # 业务模型
│   ├── activity_config.go # 活动配置
│   └── game_config.go     # 玩法配置
├── storage/               # 基础设施层
│   └── mysql/
│       ├── entity/        # 数据实体
│       ├── repository/    # 仓储实现
│       └── migrations/    # 数据库迁移
├── config/                # 配置管理
├── main.go               # 程序入口
└── README.md             # 项目文档
```

### 配置说明
活动配置示例：
```json
{
  "category": "community",
  "version": "1.0",
  "name": "社区活动",
  "start_at": 1234567890,
  "end_at": 1234567890,
  "games": [
    {
      "type": "post",
      "name": "发帖赢奖品",
      "config": {
        // 具体玩法配置
      }
    }
  ]
}
```

## 开发指南

### 新增活动类型
1. 在 `models` 中定义活动配置结构
2. 在 `storage/mysql/entity` 中定义数据实体
3. 在 `storage/mysql/repository` 中实现仓储接口
4. 在 `api/service` 中实现应用服务

### 新增玩法
1. 在 `models` 中定义玩法配置结构
2. 实现玩法特定的业务逻辑
3. 定义玩法特定的配置结构

### 错误处理
- 统一错误码定义在 `constant/constant.go`
- 错误处理工具在 `api/error.go`
- 应用层错误处理在 `api/service`

