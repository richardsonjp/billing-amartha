---

# mini service wallet JULO

---

## Table of Contents

1. [Installation](#installation)
2. [Running the Project](#running-the-project)

---

## Getting Started

Follow these instructions to set up the project on your local machine for development and testing purposes.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/richardsonjp/billing-amartha
   ```

2. Navigate to the project directory:

   ```bash
   cd billing-amartha
   ```

3. Install dependencies:

    - For Go projects, run:

      ```bash
      go mod tidy
      ```

---

## Running the Project

1. For local development:

    from root project directory, run:
   ```bash
   _scripts/run.sh
   ```

3. Visit the application on `http://localhost:8000` (or the specified port). -> check your .env file

---

## Environment Variables

Before running the project, prepare the environment variables. Create a `.env` file in the root directory and add the following variables (example):

```env
# =======================
# SYSTEM CONFIG
# =======================
SYSTEM_SERVER: "127.0.0.1"
SYSTEM_ADDR: ":8000"
#debug,release,test
SYSTEM_MODE: debug
SYSTEM_TIME_ZONE: "Asia/Jakarta"

# =======================
# PATH
# =======================
# Locale file path location
APISERVER_LOCALE_PATH: "cmd/apiserver/app/locale/"

# =======================
# URL
# =======================

# =======================
# DATABASE CONFIG
# =======================
DB_USERNAME:user
DB_PASSWORD:passwordxxx
DB_NAME:database_name
DB_HOST:localhost
DB_PORT:5432
#true,false
DB_DEBUG:true
```

---