# 🚀 Railway Deployment Guide

## Quick Start (5 Minutes to Deployment)

### Prerequisites
- GitHub account
- Railway account ([railway.app](https://railway.app))
- Git installed locally

---

## Step 1: Push to GitHub

```bash
cd hardcourt
git init
git add .
git commit -m "Production-ready tennis live scores app"
git remote add origin https://github.com/persimmonlabs/hardcourt
git branch -M main
git push -u origin main
```

---

## Step 2: Create Railway Project

1. **Login to Railway**
   - Go to [railway.app](https://railway.app)
   - Sign in with GitHub

2. **Create New Project**
   - Click "New Project"
   - Select "Deploy from GitHub repo"
   - Choose `persimmonlabs/hardcourt`

---

## Step 3: Add PostgreSQL Database

1. **Add Database Service**
   - In your Railway project, click "New"
   - Select "Database" → "PostgreSQL"
   - Railway will automatically create `DATABASE_URL` variable

2. **Database is Ready!**
   - No manual schema setup needed
   - Migrations run automatically on backend startup
   - Tables, indexes, and constraints created automatically

---

## Step 4: Configure Backend Service

Railway auto-detects the `railway.toml` configuration, but verify:

### Service: backend
- **Root Directory:** `backend`
- **Build:** Uses root `Dockerfile`
- **Health Check:** `/health`

### Environment Variables (Auto-Set by Railway):
- `PORT` - Automatically set by Railway
- `DATABASE_URL` - Automatically set when you add PostgreSQL
- `REDIS_URL` - (Optional) Add Redis if needed

**No manual configuration needed!** Railway sets everything automatically.

---

## Step 5: Configure Frontend Service

### Service: frontend
- **Root Directory:** `frontend`
- **Build:** Nixpacks (auto-detected Next.js)

### Required Environment Variables:

After backend is deployed, get its URL and set these in Railway:

```bash
# Replace with your actual backend Railway URL
NEXT_PUBLIC_API_URL=https://hardcourt-backend.up.railway.app
NEXT_PUBLIC_WS_URL=wss://hardcourt-backend.up.railway.app/ws
```

**How to set:**
1. Click on frontend service in Railway
2. Go to "Variables" tab
3. Add the two variables above
4. Redeploy frontend

---

## Step 6: Verify Deployment

### Backend Health Check
```bash
curl https://hardcourt-backend.up.railway.app/health
```

Expected response:
```json
{
  "status": "healthy",
  "database": "connected",
  "redis": "disconnected",
  "timestamp": "2025-11-23T..."
}
```

### API Endpoints
```bash
# Get all live matches
curl https://hardcourt-backend.up.railway.app/api/matches?status=Live

# Get specific match
curl https://hardcourt-backend.up.railway.app/api/matches/match_0
```

### Frontend
Visit your frontend Railway URL:
```
https://hardcourt-frontend.up.railway.app
```

You should see:
- 5 live tennis matches
- Real-time score updates every 2 seconds
- Match statistics and probabilities

---

## What Happens Automatically

✅ **Backend Startup:**
1. Connects to PostgreSQL
2. Runs database migrations (creates tables, indexes)
3. Seeds initial tournaments and players
4. Creates 5 live matches
5. Starts simulator (updates every 2 seconds)
6. Starts WebSocket server
7. Starts REST API server

✅ **Frontend Startup:**
1. Fetches initial match data via REST API
2. Connects to WebSocket for live updates
3. Displays matches with real-time updates

---

## Optional: Add Redis (For Production Scalability)

If you want to enable Redis pub/sub:

1. In Railway project, click "New"
2. Select "Database" → "Redis"
3. Railway auto-sets `REDIS_URL`
4. Redeploy backend

**Benefits:**
- Enables multi-instance scaling
- Pub/sub for distributed updates

---

## Post-Deployment: Security Hardening

### Update CORS (Recommended)

After deployment, update backend CORS to your specific frontend URL:

**File:** `backend/cmd/server/main.go` (line ~100)

```go
// Before (permissive)
AllowedOrigins: []string{"*"},

// After (secure)
AllowedOrigins: []string{
    "https://hardcourt-frontend.up.railway.app",
},
```

Commit and push to update.

---

## Troubleshooting

### Backend Won't Start
1. Check Railway logs: `Backend service → Deployments → Logs`
2. Verify `DATABASE_URL` is set (should be automatic)
3. Check health endpoint for error details

### Frontend Can't Connect
1. Verify environment variables are set correctly
2. Check WebSocket URL uses `wss://` not `ws://`
3. Verify backend is deployed and healthy

### Database Connection Fails
1. PostgreSQL should be automatically connected via `DATABASE_URL`
2. Check Railway PostgreSQL service is running
3. Verify backend logs for connection errors

### No Matches Appear
1. Check backend logs - matches are created on startup
2. Verify `/health` shows `database: connected`
3. Check frontend console for API errors
4. Test API endpoint directly: `/api/matches`

---

## Monitoring & Logs

### Backend Logs
```
Railway → Backend Service → Deployments → Logs
```

Look for:
- "Successfully connected to PostgreSQL database"
- "Migrations completed successfully"
- "Initialized 5 matches"
- "Server starting on port XXXX"

### Frontend Logs
```
Railway → Frontend Service → Deployments → Logs
```

### Health Monitoring
Set up Railway monitoring to ping `/health` endpoint every 5 minutes.

---

## Environment Variables Summary

### Backend (Auto-Set by Railway)
| Variable | Source | Required |
|----------|--------|----------|
| `PORT` | Railway | Auto |
| `DATABASE_URL` | PostgreSQL service | Auto |
| `REDIS_URL` | Redis service | Optional |

### Frontend (Manual Setup Required)
| Variable | Example | Required |
|----------|---------|----------|
| `NEXT_PUBLIC_API_URL` | `https://hardcourt-backend.up.railway.app` | Yes |
| `NEXT_PUBLIC_WS_URL` | `wss://hardcourt-backend.up.railway.app/ws` | Yes |

---

## Cost Estimate (Railway)

**Starter Plan (Free):**
- $5 free credit/month
- Enough for testing and low traffic

**Pro Plan ($20/month):**
- Unlimited projects
- Custom domains
- Better resource limits

**Estimated Usage:**
- Backend: ~$10-15/month
- Frontend: ~$5-10/month
- PostgreSQL: ~$5/month
- Total: ~$20-30/month for production

---

## Share with Friends!

Once deployed, share your frontend URL:
```
https://hardcourt-frontend.up.railway.app
```

Your friends will see:
- ✅ Live tennis matches
- ✅ Real-time score updates
- ✅ Match statistics
- ✅ Win probabilities
- ✅ Leverage index
- ✅ Fatigue indicators

---

## Next Steps

After successful deployment:
1. ✅ Test all features with friends
2. ✅ Monitor Railway logs for errors
3. ✅ Update CORS for security
4. ✅ (Optional) Add custom domain
5. ✅ (Optional) Add authentication
6. ✅ (Optional) Add user favorites

---

**Deployment Complete! 🎾🎉**

Your production-ready tennis live scores app is now running on Railway!
