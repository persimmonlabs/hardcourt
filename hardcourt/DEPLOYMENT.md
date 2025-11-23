# Railway Deployment Guide for Hardcourt

## Pre-Deployment Checklist

### ✅ Completed
- [x] Railway configuration (`railway.toml`)
- [x] Multi-stage Dockerfile for backend
- [x] CORS middleware properly configured
- [x] Health check endpoint (`/health`)
- [x] Environment variable templates

### ⚠️ Critical Missing Features
- [ ] **PostgreSQL Integration** - Database connection code is not implemented
- [ ] **Database Migrations** - No migration runner in the code
- [ ] **Environment-specific WebSocket URLs** - Needs configuration for production

---

## Railway Deployment Steps

### 1. Prerequisites
- GitHub account with your repository pushed
- Railway account (sign up at [railway.app](https://railway.app))

### 2. Create Railway Project

1. **Connect GitHub Repository**
   ```
   Railway Dashboard → New Project → Deploy from GitHub → Select 'hardcourt' repo
   ```

2. **Add PostgreSQL Database**
   ```
   Railway Project → New → Database → PostgreSQL
   ```
   - Railway will automatically provide `DATABASE_URL` environment variable

3. **Add Redis**
   ```
   Railway Project → New → Database → Redis
   ```
   - Railway will automatically provide `REDIS_URL` environment variable

### 3. Configure Backend Service

Railway should auto-detect your `railway.toml` configuration, but verify:

**Service: backend**
- Root Directory: `backend`
- Dockerfile Path: `../Dockerfile`
- Environment Variables:
  ```
  PORT=8080 (Railway sets this automatically)
  DATABASE_URL=<automatically set by Railway PostgreSQL>
  REDIS_URL=<automatically set by Railway Redis>
  ```

**Health Check:**
- Path: `/health`
- Timeout: 100s

### 4. Configure Frontend Service

**Service: frontend**
- Root Directory: `frontend`
- Builder: Nixpacks (Railway auto-detects Next.js)
- Build Command: `npm run build`
- Start Command: `npm start`

**Critical Environment Variable:**
After backend is deployed, get its Railway URL (e.g., `backend.railway.app`) and set:
```
NEXT_PUBLIC_WS_URL=wss://your-backend-url.railway.app/ws
```

⚠️ **Note:** Use `wss://` (secure WebSocket) for production, not `ws://`

### 5. Database Setup

**IMPORTANT:** You need to run your schema migrations manually after first deployment:

1. Connect to Railway PostgreSQL using CLI or GUI:
   ```bash
   # Install Railway CLI
   npm i -g @railway/cli
   
   # Login and link project
   railway login
   railway link
   
   # Connect to database
   railway connect postgres
   ```

2. Run your `schema.sql`:
   ```sql
   \i backend/schema.sql
   ```

---

## Production URLs

After deployment, you'll have two services:

- **Frontend:** `https://hardcourt-frontend.up.railway.app`
- **Backend:** `https://hardcourt-backend.up.railway.app`

Make sure to update `NEXT_PUBLIC_WS_URL` in frontend environment variables to point to your backend WebSocket endpoint.

---

## Known Issues & Missing Features

### 🚨 Critical Issues

1. **No Database Connection Logic**
   - The `main.go` file doesn't connect to PostgreSQL
   - You have `schema.sql` but no code to execute queries
   - **Action Required:** Implement database connection and query logic

2. **No Migration Runner**
   - Need to manually run `schema.sql` via Railway CLI
   - Consider adding automatic migrations on startup

3. **CORS AllowedOrigins Set to Wildcard**
   - Currently set to `*` for flexibility
   - **Recommended:** After deployment, update to specific Railway frontend URL:
     ```go
     AllowedOrigins: []string{"https://your-frontend.railway.app"},
     ```

### ⚠️ Recommendations

4. **Add Database Health Check**
   - Current `/health` endpoint doesn't verify database connectivity
   
5. **Environment Validation**
   - Add startup checks to ensure all required env vars are set

6. **Logging**
   - Consider structured logging (e.g., `zap` or `logrus`)
   - Add request ID tracking for debugging

7. **Graceful Shutdown**
   - Implement signal handling to close connections cleanly

---

## Local Development

### Backend
```bash
cd backend
cp ../.env.example .env
# Edit .env with local database credentials
go run cmd/server/main.go
```

### Frontend
```bash
cd frontend
cp .env.example .env.local
# Edit .env.local to point to local backend
npm install
npm run dev
```

---

## Monitoring

Railway provides:
- **Logs:** Real-time logs for each service
- **Metrics:** CPU, Memory, Network usage
- **Deployments:** Rollback capability

Access via: `Railway Dashboard → Your Project → Service → Logs/Metrics`

---

## Support

If you encounter issues:
1. Check Railway logs for error messages
2. Verify environment variables are set correctly
3. Ensure database migrations have run
4. Test WebSocket connection using browser dev tools

---

**Last Updated:** 2025-11-22
