# 🚀 Urlshortener-redis
- URL Shortening service using Redis

## 🛠️ Tech Stack

### Backend

- **Fiber**: v2
- **Database**: Redis

---

### 📃API Reference

## GET `shorten`
```bash
http://localhost:3000/api/v1
```
> This endpoint makes an HTTP `POST` request to the specified URL. The request payload should be sent in raw format, with a key `"url"` containing the URL value.

### Request `raw`
```javascript
{
    "url":"https://www.youtube.com/watch?v=5La12L8g1Ys"
}
```

### Response
The response to this request is in JSON format with the following schema:

```JSON
{
    "url": "",
    "short": "",
    "expiry": 0,
    "rate_limit": 0,
    "rate_limit_reset": 0
}
```
> The response includes the keys `"url", "short", "expiry", "rate_limit", and "rate_limit_reset"`. The `"url"` and `"short"` values are `strings`, while `"expiry"`, `"rate_limit"`, and `"rate_limit_reset"` are `numeric` values.

## GET `originalUrl`
```bash
http://localhost:3000/d00603
```
### Response
The url will be ressolved by the backend and you will be redirected to the original url , within seconds.

---

## ⚙️ Getting Started
### Prerequisites
- Docker
- Golang
### Installation
```bash
git clone https://github.com/Adityasinghvats/url-redis.git
```
```bash
cd url-redis
```
```bash
docker-compose up -d
```
## 🤝 Contributing

> **Guidelines**: Outline how others can contribute to your project. Include steps for forking the repository, creating branches, submitting pull requests, and any coding standards or guidelines to follow.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---
