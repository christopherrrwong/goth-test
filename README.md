

### 1. Install Dependencies

```bash
go mod download
```


### 2. Environment Configuration

#### Step 1: Create `.env` file

Copy the example environment file:


Edit `.env` with your values:

```env
ENVIRONMENT=dev
DBUSER=your_mysql_username
DBPASS=your_mysql_password
DBHOST=your_mysql_host
DBNAME=your_mysql_database
```

environment will be dev if you don't set it.
please set it to ENVIRONMENT='prod' for production.

### YAML Configuration

#### Step 1: Copy Config Files

```bash
cp internal/config/config.dev.yaml.example internal/config/config.dev.yaml
cp internal/config/config.prod.yaml.example internal/config/config.prod.yaml
```

#### Step 2: Configure Development Settings

for development, edit `internal/config/config.dev.yaml`:
for production, edit `internal/config/config.prod.yaml`:


```yaml
server:
  port: 3000

session:
  maxAge: 500  # 30 days in seconds
  isProd: false
  httpOnly: true

auth0:
  clientId: "your_auth0_client_id"
  clientSecret: "your_auth0_client_secret"
  domain: "your-tenant.us.auth0.com"
  callbackUrl: "http://localhost:3000/sso-auth/auth0/callback"

saml:
  idpMetadataURL: "https://your-idp.com/metadata"
  rootURL: "http://localhost:3000"
  entityID: "http://localhost:3000"

cors:
  allowedOrigins:
    - "https://*"
    - "http://*"
  allowedMethods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
    - "OPTIONS"
    - "PATCH"
  allowedHeaders:
    - "Accept"
    - "Authorization"
    - "Content-Type"
  allowCredentials: true
  maxAge: 300
```

**Important**: Never commit these files with real credentials! and feel free to change the values to your own.



## API Endpoints

- `GET /sso-auth/{provider}?uuid={uuid}` - Start authentication
- `GET /sso-auth/{provider}/callback` - OAuth callback (handled by provider)


## Adding new providers

- add the provider to the `internal/auth/auth.go` file
- add the provider to the `internal/config/config.dev.yaml` file
- add the provider to the `internal/config/config.prod.yaml` file

as this app uses goth, you can add any provider that goth supports.
please check the goth documentation for more information as different providers will have different values and keys.
https://github.com/markbates/goth/blob/master/examples/main.go

currently, this app only supports auth0 for testing purposes. to add other providers your need to  append more provider to the `internal/auth/auth.go` file
and also add value for config file for example if you need to add google, you need to add the following to the `internal/config/config.dev.yaml` file:

google:
  GOOGLE_KEY: "your_google_client_key"
  GOOGLE_SECRET: "your_google_client_secret"
  callbackUrl: "http://localhost:3000/sso-auth/google/callback"

  

