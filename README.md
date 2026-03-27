# GreenTraceChain

[English](./README.md) | [中文](./README_CN.md)

![Vue.js](https://img.shields.io/badge/Vue.js-3.x-4FC08D?logo=vue.js&logoColor=white)
![Go](https://img.shields.io/badge/Go-1.19+-00ADD8?logo=go&logoColor=white)
![ChainMaker](https://img.shields.io/badge/Blockchain-ChainMaker-blue)
![License](https://img.shields.io/badge/License-MIT-green)

A carbon traceability and trading platform built on the **ChainMaker (长安链)** blockchain. The system connects four types of participants — platform admins, enterprise supervisors, data auditors, and third-party regulators — to manage carbon accounting, auditing, credit trading, and on-chain traceability.

---

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         Frontend                            │
│                                                             │
│  ┌───────────────┐  ┌───────────────┐  ┌────────────────┐  │
│  │ admin-portal  │  │supervisor-    │  │ auditor-portal │  │
│  │ Platform Admin│  │portal         │  │ Data Auditor   │  │
│  └──────┬────────┘  │ Enterprise    │  └───────┬────────┘  │
│         │           │ Supervisor    │          │           │
│         │           └──────┬────────┘          │           │
│         │                  │         ┌─────────┴────────┐  │
│         │                  │         │ regulator-portal │  │
│         │                  │         │ Third-party      │  │
│         │                  │         │ Regulator        │  │
│         │                  │         └─────────┬────────┘  │
└─────────┼──────────────────┼───────────────────┼───────────┘
          │                  │                   │
          ▼                  ▼                   ▼
┌─────────────────────────────────────────────────────────────┐
│                         Backend                             │
│                                                             │
│  ┌──────────────────────┐   ┌───────────────────────────┐  │
│  │  chainmaker-service  │   │      management-api       │  │
│  │  Blockchain SDK API  │   │   Business Logic & Admin  │  │
│  │       :8087          │   │          :9999            │  │
│  └──────────┬───────────┘   └──────────────┬────────────┘  │
│             │                              │               │
│             │               ┌──────────────┴────────────┐  │
│             │               │    blockchain-explorer    │  │
│             │               │      Block Browser        │  │
│             │               └──────────────┬────────────┘  │
└─────────────┼──────────────────────────────┼───────────────┘
              │                              │
              ▼                              ▼
┌─────────────────────────┐   ┌─────────────────────────────┐
│   ChainMaker Node(s)    │   │           MySQL             │
│   长安链区块链节点        │   │         Database            │
└─────────────────────────┘   └─────────────────────────────┘
```

---

## Project Structure

```
greenTraceChain/
├── frontend/
│   ├── admin-portal/       # Platform admin — blockchain & contract management
│   ├── supervisor-portal/  # Enterprise supervisor — overview, credits, traceability
│   ├── auditor-portal/     # Data auditor — carbon accounting review & submission
│   └── regulator-portal/   # Third-party regulator — regulatory oversight portal
└── backend/
    ├── chainmaker-service/ # ChainMaker blockchain interaction API
    ├── management-api/     # Core management backend
    └── blockchain-explorer/# Blockchain block explorer
```

---

## User Roles

| Role | Portal | Responsibilities |
|------|--------|-----------------|
| Platform Admin | `admin-portal` | Manage blockchain contracts, nodes, user approvals, transaction data |
| Enterprise Supervisor | `supervisor-portal` | Dashboard overview, user management, data examination, carbon credits, traceability |
| Data Auditor | `auditor-portal` | Submit & review carbon accounting data, carbon credit shop, traceability |
| Third-party Regulator | `regulator-portal` | Regulatory oversight, data analysis, carbon emission reports |

---

## Quick Start

### Prerequisites

- Node.js >= 16
- Go >= 1.19
- MySQL 5.7+
- A running ChainMaker node (see [ChainMaker docs](https://docs.chainmaker.org.cn))

### 1. Start the backend services

```bash
# chainmaker-service (blockchain SDK API)
cd backend/chainmaker-service
go mod download
go run src/main.go
# Runs on :8087

# management-api
cd backend/management-api
go mod download
go run src/main.go --config configs
# Runs on :9999

# Or use Docker Compose (management-api)
cd backend/management-api
docker-compose up -d
```

### 2. Start a frontend

```bash
cd frontend/<portal-name>
npm install
npm run dev
```

> **Note:** All frontend portals default to `http://localhost:8087` as the API endpoint. If your `chainmaker-service` runs on a different host or port, update the address in each portal's source files accordingly.

---

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Blockchain | ChainMaker (长安链) v2.0+ |
| Frontend | Vue 3, Vite, Vue Router, Element Plus, Axios |
| Backend | Go, Gin, GORM |
| Database | MySQL 5.7 |
| Container | Docker, Docker Compose |

---

## License

[MIT](./LICENSE) — developed for the 2024 China Service Outsourcing Innovation & Entrepreneurship Competition (National Finals).
