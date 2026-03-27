# chainmaker-service

A Go service that wraps the **ChainMaker SDK** and exposes a unified HTTP API for all GreenTraceChain frontends. All on-chain read/write operations — including contract invocation, transaction queries, company registration, carbon accounting, and traceability — are handled through this service.

## Architecture

```
Frontend Portals
      │
      │  POST /chainmaker?cmb=<CommandName>
      ▼
chainmaker-service  (:8087)
      │
      │  ChainMaker Go SDK v3
      ▼
ChainMaker Node
      │
      ▼
   MySQL (off-chain data cache)
```

## API

All requests go through a single endpoint:

```
POST /chainmaker?cmb=<CommandName>
Content-Type: application/json
```

The `cmb` query parameter selects the command to execute. Examples:

| Command | Description |
|---------|-------------|
| `CompanyLogin` | Authenticate an enterprise account |
| `GetCompanyByCompanyId` | Fetch company info by ID |
| `GetCompanyList` | List all registered companies |
| `APDistributeCoin` | Admin distributes carbon coins to a company |
| `APRegisiterForDAPAndTPRP` | Admin approves company registration |
| `APReviewFailed` | Admin rejects a company registration |

## Prerequisites

- Go 1.19+
- A running ChainMaker node with a deployed carbon contract
- MySQL 5.7+ (connection configured in `src/dao/`)
- ChainMaker SDK config and crypto certificates in `config/`

## Configuration

The SDK connection is configured via `config/sdk_config.yml`. Crypto certificates for the ChainMaker organization are located under `config/crypto-config/`.

Update `sdk_config.yml` to point to your ChainMaker node's address and chain ID before running.

## Setup

```bash
# Install dependencies
go mod download

# Run the service
go run src/main.go
# Runs on :8087
```

## Tech Stack

- **Go** 1.19
- **Gin** — HTTP framework
- **ChainMaker Go SDK** v3
- **GORM** — MySQL ORM
