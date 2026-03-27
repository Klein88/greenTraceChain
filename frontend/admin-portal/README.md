# admin-portal

The **platform administrator** frontend for GreenTraceChain. Admins manage the blockchain infrastructure, review enterprise user registrations, deploy and invoke smart contracts, and monitor on-chain transaction and carbon accounting data.

## Features

- **User Management** — review enterprise registration requests, approve or reject accounts, distribute carbon coins
- **Contract Management** — deploy, update, and invoke ChainMaker smart contracts
- **Chain Management** — monitor blockchain nodes and chain status
- **Transaction Data** — browse all on-chain transactions
- **Carbon Accounting Data** — view enterprise carbon submissions with multi-level traceability (TraceSource 1–4)

## Tech Stack

![Vue.js](https://img.shields.io/badge/Vue.js-3.x-4FC08D?logo=vue.js&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-4.x-646CFF?logo=vite&logoColor=white)

- Vue 3 + Vue Router
- Vite
- Element Plus
- Axios

## Setup

```bash
npm install
npm run dev      # development server → http://localhost:5173
npm run build    # production build → dist/
```

## Backend Dependencies

| Service | Default Address |
|---------|----------------|
| chainmaker-service | `http://localhost:8087` |
| management-api | `http://localhost:9999` |

> **Important:** API calls are currently hardcoded to `http://47.97.176.174:8087`.
> Before running locally, replace all occurrences of that address with your local `chainmaker-service` address.
