# ghostproxy

A stealthy, TLS-enabled reverse proxy in Go to capture credentials and cookies across any website. Inspired by Evilginx, designed for automation and stealth.

## Features

- Reverse proxy ANY site via subdomain
- Capture POSTed credentials
- Capture `Set-Cookie` values
- TLS support via Let's Encrypt
- Dashboard to view active logs
- Optional Telegram/Email alerts

## Usage

Point `*.proxy.yourdomain.com` to your VPS IP in DNS.

Run:

```bash
go run main.go
