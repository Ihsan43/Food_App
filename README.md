# Food Delivery App

## 📌 Introduction
Food Delivery App is a modern web application that allows users to order food from various suppliers. This project is built with:
- **Backend:** Golang with PostgreSQL
- **Frontend:** Vue.js (Vite + Vue 3)

## 🚀 Features
### **User Features:**
- User registration & authentication (JWT-based)

### **Admin Features:**
- Manage suppliers, categories, and products
- View and process orders
- Manage users and deliveries

## 🛠 Tech Stack
### **Backend (Golang + PostgreSQL):**
- PostgreSQL
- JWT Authentication
- Docker

### **Frontend (Vue.js):**
- Vue 3 + Vite
- Vue Router
- Pinia (State Management)
- Axios (API requests)
- TailwindCSS

## ⚙️ Installation
### **Prerequisites:**
- Go 1.20+
- Node.js 16+
- PostgreSQL
- Docker (Optional)

### **Backend Setup:**
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/food-delivery-app.git
   cd food-delivery-app/backend
   ```
2. Configure environment variables (`.env`):
   ```env
# Server Configuration
PORT=8080

# Authentication
ACCESS_SECRET=your_access_secret_key
ACCESS_LIFETIME_MINUTES=60
REFRESH_SECRET=your_refresh_secret_key
REFRESH_LIFETIME_MINUTES=1440

# Database Configuration
DB_USER=postgres
DB_HOST=localhost
DB_PASSWORD=yourpassword
DB_PORT=5432
DB_NAME=food_delivery

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

### **Frontend Setup:**
1. Navigate to frontend directory:
   ```bash
   cd ../frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run the development server:
   ```bash
   npm run dev
   ``
---
✨ Happy Coding! 🍽️

