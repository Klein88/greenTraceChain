# auditor-portal

The **data auditor** frontend for GreenTraceChain. Data auditors are responsible for submitting carbon accounting data on behalf of enterprises, reviewing records, trading carbon credits, and accessing carbon knowledge resources.

## Features

- **Registration & Login** — enterprise account registration and authentication
- **Dashboard** — personal carbon data overview
- **Carbon Accounting** — submit carbon accounting data with industry-specific forms:
  - Chemical Industry: fossil fuels, greenhouse gases, HNO₃, HOOCCH₂, M₂CO₃, electrical processes
  - Electrical Industry: CO₂ emissions, 6FS, transmission and distribution losses
- **Carbon Credit Shop** — buy and sell carbon credits on the platform marketplace
- **Teaching Modules** — carbon accounting knowledge base and guided tutorials
- **Traceability** — look up carbon traceability records on-chain
- **Personal Account** — manage profile and submission history

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

> **Important:** API calls are currently hardcoded to `http://47.97.176.174:8087`.
> Before running locally, replace all occurrences of that address with your local `chainmaker-service` address.
