# Configuration Setup

## Important: Protecting Your Credentials

Your config files contain sensitive credentials that should **NOT** be committed to version control.

## Setup Instructions

### 1. Add Config Files to `.gitignore`

Add these lines to your `.gitignore` file:

```gitignore
# Config files with real credentials
internal/config/config.dev.yaml
internal/config/config.prod.yaml
internal/config/config.staging.yaml
```

### 2. Copy Example Files

For local development:
```bash
cp internal/config/config.dev.yaml.example internal/config/config.dev.yaml
cp internal/config/config.prod.yaml.example internal/config/config.prod.yaml
```

### 3. Update With Your Credentials

Edit the copied files with your actual credentials:

**`config.dev.yaml`:**
- Replace `your_auth0_client_id` with your Auth0 Client ID
- Replace `your_auth0_client_secret` with your Auth0 Client Secret
- Replace `your-tenant.us.auth0.com` with your Auth0 domain
- Update other service URLs as needed

**`config.prod.yaml`:**
- Use production credentials
- Update domain to your production domain
- Ensure `isProd: true` for production security settings

## Files Structure

```
internal/config/
├── config.dev.yaml.example     ✅ Commit this (no real credentials)
├── config.prod.yaml.example    ✅ Commit this (no real credentials)
├── config.dev.yaml             ❌ DO NOT COMMIT (real credentials)
├── config.prod.yaml            ❌ DO NOT COMMIT (real credentials)
└── config.go                   ✅ Commit this
```

## Environment Variables

Also set these environment variables for database access:

```bash
export DBUSER=your_db_username
export DBPASS=your_db_password
export ENVIRONMENT=dev  # or 'prod'
```

Or create a `.env` file (already in `.gitignore`):

```env
ENVIRONMENT=dev
DBUSER=your_db_username
DBPASS=your_db_password
```

