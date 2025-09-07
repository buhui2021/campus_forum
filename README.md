# 校园论坛系统

一个基于 Go 和 Gin 框架构建的校园论坛系统，支持用户发帖、评论、点赞和管理员审核功能。

## 功能特性

- 用户注册和登录
- 发布和管理帖子
- 评论功能
- 点赞功能
- 管理员审核和置顶帖子
- JWT 认证
- SQLite 数据库

## 安装和运行

1. 克隆项目
2. 安装依赖：`go mod download`
3. 设置环境变量（复制 `text.env.example` 为 `text.env` 并修改）
4. 运行项目：`go run main.go`
5. 访问 `http://localhost:8080`

## API 文档

详见 [API.md](API.md) 文件

## 项目结构
