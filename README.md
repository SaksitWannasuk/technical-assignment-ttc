# technical-assignment-ttc
technical assignment for job interview ttc

## Tech Stack

### Frontend
- Angular
- Angular Material

### Backend
- Go
- Fiber

### Database
- SQLite

### ORM
- GORM

---

# Features

- Display person list in table
- Add person using modal
- View person details using modal
- Calculate age from birth date
- Store data in SQLite database

---

# Project Structure

```txt
technical-assignment-ttc/
├── backend
└── frontend
```

---

# Backend Setup

## 1. Go to backend

cd backend

## 2. Install dependencies

go mod tidy

## 3. Run backend

go run cmd/main.go

Backend running at:

http://localhost:3000

---

# Frontend Setup

## 1. Go to frontend

cd frontend

## 2. Install dependencies

npm install

## 3. Run frontend

ng serve

Frontend running at:

http://localhost:4200

---

# API Endpoints

## Get all people

GET /api/people

## Get person by id

```http
GET /api/person/:id
```

## Create person

```http
POST /api/person
```

Example request:

```json
{
  "firstName": "John",
  "lastName": "Doe",
  "birthDate": "2000-01-15",
  "address": "Bangkok"
}
```

---

# Notes

- Age is calculated from birth date during API response.
- SQLite database file will be created automatically.
- Angular Material is used for UI components.