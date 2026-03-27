# GreenTraceChain

[English](./README.md) | [中文](./README_CN.md)

![Vue.js](https://img.shields.io/badge/Vue.js-3.x-4FC08D?logo=vue.js&logoColor=white)
![Go](https://img.shields.io/badge/Go-1.19+-00ADD8?logo=go&logoColor=white)
![ChainMaker](https://img.shields.io/badge/区块链-长安链-blue)
![License](https://img.shields.io/badge/许可证-MIT-green)

基于**长安链（ChainMaker）**的碳排放溯源与交易平台。系统连接四类参与角色——平台管理员、企业监管方、数据审核员、第三方监管机构——共同完成碳核算数据的提交、审核、碳信用交易与链上溯源全流程。

该项目获得2024年全国大学生服务外包创新创业大赛全国一等奖。

---

## 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                          前端                               │
│                                                             │
│  ┌───────────────┐  ┌───────────────┐  ┌────────────────┐  │
│  │ admin-portal  │  │supervisor-    │  │ auditor-portal │  │
│  │  平台管理员   │  │portal         │  │   数据审核员   │  │
│  └──────┬────────┘  │  企业监管方   │  └───────┬────────┘  │
│         │           └──────┬────────┘          │           │
│         │                  │         ┌─────────┴────────┐  │
│         │                  │         │ regulator-portal │  │
│         │                  │         │  第三方监管机构  │  │
│         │                  │         └─────────┬────────┘  │
└─────────┼──────────────────┼───────────────────┼───────────┘
          │                  │                   │
          ▼                  ▼                   ▼
┌─────────────────────────────────────────────────────────────┐
│                          后端                               │
│                                                             │
│  ┌──────────────────────┐   ┌───────────────────────────┐  │
│  │  chainmaker-service  │   │      management-api       │  │
│  │  区块链 SDK 交互服务  │   │    业务逻辑与管理后端     │  │
│  │       :8087          │   │          :9999            │  │
│  └──────────┬───────────┘   └──────────────┬────────────┘  │
│             │                              │               │
│             │               ┌──────────────┴────────────┐  │
│             │               │    blockchain-explorer    │  │
│             │               │        区块链浏览器        │  │
│             │               └──────────────┬────────────┘  │
└─────────────┼──────────────────────────────┼───────────────┘
              │                              │
              ▼                              ▼
┌─────────────────────────┐   ┌─────────────────────────────┐
│     长安链区块链节点     │   │        MySQL 数据库          │
└─────────────────────────┘   └─────────────────────────────┘
```

---

## 目录结构

```
greenTraceChain/
├── frontend/
│   ├── admin-portal/       # 平台管理员 — 区块链及合约管理
│   ├── supervisor-portal/  # 企业监管方 — 数据总览、碳信用、溯源
│   ├── auditor-portal/     # 数据审核员 — 碳核算数据提交与审核
│   └── regulator-portal/   # 第三方监管机构 — 监管合规门户
└── backend/
    ├── chainmaker-service/ # 长安链 SDK 交互 API 服务
    ├── management-api/     # 核心管理后端
    └── blockchain-explorer/# 区块链浏览器
```

---

## 用户角色

| 角色 | 前端模块 | 主要职责 |
|------|----------|----------|
| 平台管理员 | `admin-portal` | 管理区块链合约、节点、用户审批、链上交易数据 |
| 企业监管方 | `supervisor-portal` | 数据总览、用户管理、数据审核、碳信用管理、溯源中心 |
| 数据审核员 | `auditor-portal` | 提交碳核算数据、碳信用交易市场、知识教学、溯源查询 |
| 第三方监管机构 | `regulator-portal` | 监管数据分析、碳排放报告查看、碳信用追踪 |

---

## 快速开始

### 环境要求

- Node.js >= 16
- Go >= 1.19
- MySQL 5.7+
- 已运行的长安链节点（参考 [长安链文档](https://docs.chainmaker.org.cn)）

### 1. 启动后端服务

```bash
# chainmaker-service（区块链 SDK API）
cd backend/chainmaker-service
go mod download
go run src/main.go
# 监听端口 :8087

# management-api（管理后端）
cd backend/management-api
go mod download
go run src/main.go --config configs
# 监听端口 :9999

# 或使用 Docker Compose 启动 management-api
cd backend/management-api
cp .env.example .env   # 填写数据库密码
docker-compose up -d
```

### 2. 启动前端

```bash
cd frontend/<模块名>
npm install
npm run dev
```

> **注意：** 所有前端模块默认连接 `http://localhost:8087` 作为 API 地址。如果 `chainmaker-service` 部署在其他主机或端口，请在对应模块的源码中修改该地址。

---

## 技术栈

| 层级 | 技术 |
|------|------|
| 区块链 | 长安链（ChainMaker）v2.0+ |
| 前端 | Vue 3、Vite、Vue Router、Element Plus、Axios |
| 后端 | Go、Gin、GORM |
| 数据库 | MySQL 5.7 |
| 容器化 | Docker、Docker Compose |

---

## 许可证

[MIT](./LICENSE) — 本项目为 2024 年中国服务外包创新创业大赛（全国总决赛）参赛作品。
