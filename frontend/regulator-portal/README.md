# regulator-portal

The **third-party regulatory authority** frontend for GreenTraceChain. Regulators have read-only access to carbon data across all enterprises for independent oversight, compliance checking, and regulatory reporting.

## Features

- **Dashboard** — platform-wide carbon emissions summary
- **Data Analysis** — cross-enterprise trend analysis and statistics
- **Data Examination** — inspect submitted carbon accounting records
- **Carbon Emission Report** — view detailed emission reports by enterprise
- **Carbon Credit** — monitor carbon credit distribution and usage across the platform

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
