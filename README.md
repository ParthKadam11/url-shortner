# URL Shortener

My first Go program — a simple URL shortener server.

Built with Go's standard library only (no frameworks), using Go 1.22+ enhanced routing patterns.

## API

### Shorten a URL

```
POST /short
Content-Type: application/json

{"url": "https://example.com/very/long/url"}
```

Response:

```json
{"short_url": "0916a431"}
```

### Redirect to original URL

```
GET /redirect/{short_url}
```

Visiting `/redirect/0916a431` in a browser returns `302 Found` and redirects to the original URL.

### Root

```
GET /
```

Response: `Hello / Route`

## Run

```bash
go run main.go
```

Server starts on `http://localhost:3000`.

## How it works

- Short codes are the first 8 characters of the MD5 hash of the original URL.
- URLs are stored in an in-memory map (`urlDB`).
- No persistence — data is lost on restart.

## Go concepts in this project

| Concept | Where |
|---|---|
| Structs & JSON tags | `URL` struct |
| Maps | `urlDB` |
| Pointers | `json.Decode(&data)`, `json.Encode(w)` |
| Error handling | `getUrl` returns `(URL, error)` |
| HTTP server | `net/http` with pattern-based routing |
| Visibility | Uppercase → exported, lowercase → private |
| `:=` short declaration | Used throughout |

## TODO

- [ ] Persist to a database
- [ ] Custom short URLs
- [ ] Click tracking
- [ ] Unit tests
