# Activity 活动框架

## 简介
一个灵活可扩展的活动框架，支持多种活动玩法的快速开发和部署。

## 核心特性
- 配置驱动：通过 JSON 配置文件定义活动
- 插件化设计：支持新玩法的快速接入
- 解耦设计：不依赖特定的UI框架或后端服务
- 状态管理：统一的活动状态管理机制

## 架构设计

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
1. 实现 `ActivityInterface` 接口
2. 在 `NewActivityFromConfig` 中注册新活动类型
3. 创建对应的配置文件

### 新增玩法
1. 实现 `GameInterface` 接口
2. 在 `NewGameFromConfig` 中注册新玩法
3. 定义玩法特定的配置结构

## 存储结构

### 目录结构
```
storage/
  ├── mysql/
  │   ├── migrations/        # 数据库迁移文件
  │   │   └── 001_init_schema.sql
  │   ├── models/           # GORM模型
  │   │   └── activity.go
  │   ├── repository/       # 仓储层
  │   │   └── activity_repository.go
  │   └── db.go            # 数据库连接配置
```
