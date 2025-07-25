# Go Auth Backend (Gin + PostgreSQL + Redis)

A secure authentication backend built with Go, Gin, PostgreSQL, and Redis. Supports Signup, Login, Protected Routes, Logout (with Redis token blacklist).

---

## 🚀 Features

- ✅ Signup & Login with hashed passwords (bcrypt)
- ✅ JWT-based authentication
- ✅ Token revocation via Redis
- ✅ Auth middleware for protected routes
- ✅ PostgreSQL database with GORM ORM
- ✅ Redis with Docker (no local install required)

---

## 📦 Tech Stack

- **Go (Gin framework)**
- **PostgreSQL** (can use AWS RDS or local)
- **Redis** (via Docker)
- **JWT** for auth
- **bcrypt** for password hashing

---


---

## 🔧 Setup Instructions

### 1. Clone the repo

```bash
git clone https://github.com/yourname/go-auth-backend.git
cd go-auth-backend
PORT=8080

# PostgreSQL (can be local or RDS)
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=auth
DB_PORT=5432

# JWT
JWT_SECRET=your_jwt_secret

# Redis (running via Docker)
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
```
# for running redis 
```
docker run --name redis-server -p 6379:6379 -d redis

```

---



## 🙌 Final Words

If you found this project helpful, feel free to **give it a star ⭐**, **fork it**, and **follow me** for more full-stack and backend projects!

> Let's build secure, scalable, and developer-friendly APIs — the right way 💪

Stay connected:
- 📬 Drop feedback or ideas — contributions are always welcome!
- 🔁 Share this repo with others who want to learn real-world Go backend development.





---

## 🤝 Let’s Connect!

I’d love to connect and collaborate — feel free to reach out or follow me:

- 🔗 **LinkedIn**: [linkedin.com/in/nikhil6392](https://linkedin.com/in/nikhil6392)

> Let’s grow together as developers — share ideas, build cool stuff, and push each other forward 🚀

**Follow for more tech, code, and backend architecture goodness!**

**Thanks for checking this out — happy coding! 🚀**
**Developed by Nikhil Pathak**
