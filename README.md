# sky-log

A real-time log viewer and monitoring tool with a web UI and webhook alerting. Built with [Sky](https://github.com/anzellai/sky).

sky-log tails log sources (files, commands, or stdin), parses structured JSON log entries, and streams them to a live web dashboard with filtering. Optionally, it fires webhooks when log messages match configurable patterns — useful for headless server monitoring without the web UI.

## Installation

```sh
curl -fsSL https://raw.githubusercontent.com/anzellai/sky-log/main/install.sh | sh
```

Install to a custom directory:

```sh
curl -fsSL https://raw.githubusercontent.com/anzellai/sky-log/main/install.sh | sh -s -- --dir ~/.local/bin
```

Install a specific version:

```sh
curl -fsSL https://raw.githubusercontent.com/anzellai/sky-log/main/install.sh | SKY_LOG_VERSION=0.1.0 sh
```

### Supported platforms

| OS | Architecture |
|----|-------------|
| Linux | x64, arm64 |
| macOS | x64, arm64 |
| Windows | x64, arm64 |

## Quick start

```sh
# Watch a log file
sky-log /var/log/app.log

# Watch multiple files
sky-log /var/log/app.log /var/log/worker.log

# Pipe from a command
tail -f /var/log/syslog | sky-log

# Use a config file for sources + webhooks
sky-log
```

Then open `http://localhost:4000` in your browser.

## Log format

sky-log expects JSON log entries (one per line):

```json
{"timestamp": 1711036801, "level": "info", "scope": "server", "message": "listening on :8080"}
{"timestamp": 1711036802, "level": "error", "scope": "db", "message": "connection refused"}
```

| Field | Type | Description |
|-------|------|-------------|
| `timestamp` | `int` | Unix timestamp (seconds or milliseconds) |
| `level` | `string` | `debug`, `info`, `warn`, or `error` |
| `scope` | `string` | Component or module name |
| `message` | `string` | Log message text |

Lines that fail to parse as JSON are displayed as `info`-level entries with scope `raw`.

## Configuration

sky-log reads a `sky-log.toml` file from the current directory. This file defines log sources and webhook rules.

```toml
# Global webhook — applies to all sources unless overridden
[webhook]
url = "https://hooks.slack.com/services/T00/B00/xxx"
filter = "panic|fatal|OOMKilled"

# Each [[source]] defines a log stream
[[source]]
name = "api"
command = "docker logs -f api-container"

[[source]]
name = "worker"
command = "journalctl -u worker -f -o cat"

[[source]]
name = "nginx"
command = "tail -f /var/log/nginx/error.log"
filter = "502|503|upstream timed out"
webhook_url = "https://hooks.slack.com/services/T00/B00/yyy"
```

### Source fields

| Field | Required | Description |
|-------|----------|-------------|
| `name` | yes | Label shown in the UI and included in webhook payloads |
| `command` | yes | Shell command to execute; stdout is captured as log lines |
| `filter` | no | Regex pattern — only matching entries trigger webhooks for this source. Falls back to the global `[webhook].filter` if omitted |
| `webhook_url` | no | Per-source webhook URL. Falls back to the global `[webhook].url` if omitted |

### Webhook fields

| Field | Description |
|-------|-------------|
| `url` | Default webhook endpoint for all sources |
| `filter` | Default regex pattern for all sources |

## Environment variables

The web UI port and other Sky.Live settings can be overridden via environment variables or a `.env` file:

| Variable | Default | Description |
|----------|---------|-------------|
| `SKY_LIVE_PORT` | `4000` | Port for the web UI |
| `SKY_LIVE_INPUT` | `debounce` | Input mode: `debounce` or `blur` |
| `SKY_LIVE_SESSION_STORE` | `memory` | Session store: `memory`, `sqlite`, `redis`, or `postgresql` |
| `SKY_LIVE_SESSION_PATH` | — | Path for sqlite session store |
| `SKY_LIVE_SESSION_URL` | — | URL for redis/postgresql session store |
| `SKY_LIVE_STATIC_DIR` | — | Directory for static file serving |
| `SKY_LIVE_TTL` | — | Session TTL |

Example:

```sh
SKY_LIVE_PORT=9090 sky-log /var/log/app.log
```

Or create a `.env` file in the working directory:

```
SKY_LIVE_PORT=9090
```

## Web UI

The web dashboard provides:

- **Real-time streaming** — log entries appear as they are written, via server-sent events
- **Level filter** — show only entries at or above a severity (DEBUG, INFO, WARN, ERROR)
- **Scope filter** — filter by component/module name
- **Regex search** — search messages with regular expressions
- **Source filter** — when watching multiple sources, filter to a single source
- **Auto-scroll** — follows new entries (toggle to pause and scroll freely)
- **Light/dark theme** — toggle with the theme button
- **Clear** — reset the log view

## Webhooks

Webhooks enable headless monitoring — no browser needed. When a log entry matches a source's filter pattern, sky-log sends a POST request with a JSON payload:

```json
{
  "source": "api",
  "level": "error",
  "scope": "db",
  "message": "connection refused",
  "timestamp": 1711036802
}
```

### How matching works

1. Each source checks its own `filter` regex. If the source has no `filter`, it falls back to the global `[webhook].filter`
2. Each source checks its own `webhook_url`. If the source has no `webhook_url`, it falls back to the global `[webhook].url`
3. A webhook fires only when **both** a filter and a URL are defined (either per-source or global)
4. Filters are case-insensitive regex patterns matched against the log `message`

### Use cases

**Server monitoring without a browser:**

```toml
[webhook]
url = "https://hooks.slack.com/services/T00/B00/xxx"
filter = "panic|fatal|OOM|connection refused"

[[source]]
name = "production"
command = "ssh prod 'journalctl -u myapp -f -o cat'"
```

sky-log runs on your server (or locally over SSH), tails logs, and sends Slack alerts when critical errors appear. The web UI is available but optional.

**Multi-service monitoring with per-source alerts:**

```toml
[webhook]
url = "https://hooks.slack.com/services/T00/B00/general"
filter = "panic|fatal"

[[source]]
name = "api"
command = "docker logs -f api"
filter = "5[0-9]{2}|timeout|circuit.breaker"
webhook_url = "https://hooks.slack.com/services/T00/B00/api-team"

[[source]]
name = "payments"
command = "docker logs -f payments"
filter = "failed|declined|insufficient"
webhook_url = "https://hooks.slack.com/services/T00/B00/payments-team"

[[source]]
name = "worker"
command = "docker logs -f worker"
```

- `api` errors matching `5xx|timeout|circuit.breaker` go to the API team's Slack channel
- `payments` errors go to the payments team's channel
- `worker` uses the global filter (`panic|fatal`) and global webhook URL
- All sources are visible in the web UI

**CI/CD log monitoring:**

```toml
[[source]]
name = "deploy"
command = "tail -f /var/log/deploy.log"
filter = "FAILED|rollback|exit code [1-9]"
webhook_url = "https://discord.com/api/webhooks/xxx/yyy"
```

**Pipe any command and get alerts:**

```toml
[[source]]
name = "k8s-events"
command = "kubectl get events -w -o json"
filter = "CrashLoopBackOff|OOMKilled|FailedScheduling"
webhook_url = "https://hooks.slack.com/services/T00/B00/xxx"
```

## Development

sky-log is written in [Sky](https://github.com/anzellai/sky), an Elm-inspired language that compiles to Go.

```sh
sky build src/Main.sky    # compile
sky run src/Main.sky      # build and run
sky dev src/Main.sky      # watch mode
```

## License

MIT
