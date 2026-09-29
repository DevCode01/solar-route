# Solar Route

**Go** · **v0.1.0**

_Lightweight HTTP server with JSON endpoints and uptime tracking_

---

## Features

- Root, /health and /api/info JSON endpoints
- Uptime and timestamp reporting
- Port configurable via PORT environment variable

## Installation

```bash
git clone https://github.com/DevCode01/solar-route.git
cd solar-route
```

## Usage

```bash
PORT=8080 go run main.go
```

## How it works

Starts a small HTTP server exposing a health check and a JSON info endpoint. The port is read from the `PORT` environment variable (default 8080). `/api/info` reports service name, version, server time and uptime.

## Options

- `PORT` - environment variable, listening port (default 8080)
- `GET /` - friendly index message
- `GET /health` - JSON `{"status":"ok"}`
- `GET /api/info` - JSON service info (version, time, uptime)

## Example

```bash
$ PORT=8080 go run main.go
Server listening on :8080
```

## Testing

```bash
bash tests/run.sh
```

Runs the build/syntax check (interactive/random/networked output isn't asserted exactly).

## License

MIT
