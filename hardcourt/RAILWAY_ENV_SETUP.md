# Railway Environment Variables Setup

## Backend Service (`accomplished-hope-production`)

**Auto-configured by Railway:**
- `PORT` - Automatically set by Railway
- `DATABASE_URL` - Automatically set when you add PostgreSQL database
- `REDIS_URL` - (Optional) Set when you add Redis database

**Optional Configuration:**
- `ENABLE_SIMULATOR` - Set to `"true"` to enable simulator when no live matches available (default: off)
- `SCRAPER_INTERVAL` - Scraper interval (default: `"1m"` for 1 minute)

## Frontend Service (`hardcourt-production`)

**REQUIRED - Must be manually set in Railway dashboard:**

1. Go to Railway Dashboard → Frontend Service → Variables tab

2. Add these environment variables:

```bash
NEXT_PUBLIC_API_URL=https://accomplished-hope-production.up.railway.app
NEXT_PUBLIC_WS_URL=wss://accomplished-hope-production.up.railway.app/ws
```

**Important Notes:**
- Use `https://` for API URL (Railway provides SSL)
- Use `wss://` for WebSocket URL (secure WebSocket)
- Path is `/ws` (not `/wss`)
- After adding variables, **redeploy frontend** for changes to take effect

## Verification

### Check Backend Health:
```bash
curl https://accomplished-hope-production.up.railway.app/health
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

### Check Frontend Connectivity:
1. Open browser console on `https://hardcourt-production.up.railway.app`
2. Should see successful API calls to backend
3. Should see WebSocket connected message

## Common Issues

### 502 Bad Gateway
- Backend is still deploying (wait 2-5 minutes)
- Database connection failed (check DATABASE_URL)
- Healthcheck timeout (backend takes >300s to start)

### CORS Errors
- Fixed in latest deployment (allows all Railway URLs)
- If issue persists, check frontend environment variables

### WebSocket Connection Failed
- Check `NEXT_PUBLIC_WS_URL` is set correctly
- Verify path is `/ws` not `/wss`
- Ensure backend is running (check /health endpoint)

### No Matches Displayed
- Backend may need seeding: Run `go run cmd/seed/main.go` locally or on Railway
- Set `ENABLE_SIMULATOR=true` to test with simulated matches
- Check browser console for API errors
