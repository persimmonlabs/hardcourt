# Hardcourt

High-performance live tennis scoring application.

## Structure

- `/backend`: Go (Chi, Redis, DDD)
- `/frontend`: Next.js 14 (Tailwind, Framer Motion)

## Running Locally

1. **Backend**:
   ```bash
   cd backend
   go mod tidy
   go run cmd/server/main.go
   ```
   *Requires Redis running on localhost:6379*

2. **Frontend**:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

## Architecture

- **Math Engine**: Calculates Win Probability using Markov Chains and Leverage Index.
- **Simulator**: Generates live match events every 2 seconds.
- **WebSockets**: Pushes updates to the frontend in real-time.
