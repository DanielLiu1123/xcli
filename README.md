# xcli

Command line for X (Twitter).

## Installation

```bash
go install github.com/DanielLiu1123/xcli/cmd/xcli@latest
```

Build from source:

```bash
make
go build -o xcli cmd/xcli/main.go 
```

## Usage

1. Go [Developer Portal](https://developer.x.com/en/portal/dashboard) to get your secrets.
2. Use environment variables or use flags to pass your secrets.

```bash
# Use environment variables, has higher priority than flags
API_KEY=<API_KEY> API_SECRET=<API_SECRET> ACCESS_TOKEN=<ACCESS_TOKEN> ACCESS_SECRET=<ACCESS_SECRET> xcli user lookup --by-auth

# Use flags
xcli user lookup --by-auth --api-key <API_KEY> --api-secret <API_SECRET> --access-token <ACCESS_KEY> --access-secret <ACCESS_SECRET>
```

```bash
# Create a new tweet
xcli tweet create --text "Hello, world!"

# Delete tweets
xcli tweet delete <tweet_id> <tweet_id>

# Lookup tweets
xcli tweet lookup <tweet_id> <tweet_id>

# Lookup users by user ID
xcli user lookup --by-id <user_id>,<user_id>

# Lookup users by username
xcli user lookup --by-username <username>,<username>

# Lookup authenticated user
xcli user lookup --by-auth
```

## License

The MIT License.
