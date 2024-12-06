# xcli

Command line for X (Twitter).

## Installation

```bash
go install github.com/DanielLiu1123/xcli/cmd/xcli@latest
```

## Usage

1. Go [Developer Portal](https://developer.x.com/en/portal/dashboard) to get your secrets.
2. Use environment variables to set secrets or use flags to pass them.

```bash
# Create a new tweet
xcli tweet create --text "Hello, world!" --api-key <API_KEY> --api-secret <API_SECRET> --access-token <ACCESS_TOKEN> --access-secret <ACCESS_SECRET>

# Delete tweets
xcli tweet delete <tweet_id> <tweet_id>

# Lookup tweets
xcli tweet lookup <tweet_id> <tweet_id>
```

## License

The MIT License.
