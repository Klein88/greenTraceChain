# supervisor-portal

The **enterprise supervisor** frontend for GreenTraceChain. Supervisors have full visibility into the platform's carbon data, manage enterprise user accounts, audit submitted data, and oversee carbon credit allocation and traceability.

## Features

- **Dashboard** — aggregated carbon emissions overview and key metrics
- **Data Analysis** — charts and trend analysis on carbon data
- **User Management** — manage registered enterprise accounts
- **Data Examination** — review and approve carbon accounting submissions
- **Carbon Emission Report** — generate and view emission reports per enterprise
- **Carbon Credit** — monitor and manage carbon credit balances
- **Traceability Center** — trace carbon data back to its source on-chain

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
