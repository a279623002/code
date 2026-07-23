# 认证授权面试笔记

> 项目背景：边端智能助手和机器人管控平台都需要用户登录、接口鉴权。JWT 和 OAuth2 是最常见的两种认证授权方案。

---

## 一、JWT（JSON Web Token）

### 1. 什么是 JWT？

**一句话**：JWT 是一种紧凑、自包含的方式，用于在各方之间安全地传输信息。

### 2. JWT 结构

```
header.payload.signature

例如：
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

| 部分 | 说明 |
|---|---|
| **Header** | 算法和类型，Base64Url 编码 |
| **Payload** | 声明（claims），如用户 ID、过期时间 |
| **Signature** | 用密钥对前两部分签名，防止篡改 |

### 3. Header 示例

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

### 4. Payload 常见声明

| 声明 | 含义 |
|---|---|
| `iss` | 签发者 |
| `sub` | 主题（用户 ID） |
| `aud` | 接收方 |
| `exp` | 过期时间 |
| `nbf` | 生效时间 |
| `iat` | 签发时间 |
| `jti` | 唯一标识 |

### 5. 签名过程

```
HMACSHA256(
  base64Url(header) + "." + base64Url(payload),
  secret
)
```

### 6. JWT 认证流程

```
客户端 ──登录（用户名/密码）──► 服务端
                                │
                                ↓ 验证成功，生成 JWT
客户端 ◄────────JWT──────────── 服务端

客户端 ──请求 API（Header: Authorization: Bearer <JWT>）──► 服务端
                                │
                                ↓ 验签、解析、判断过期
客户端 ◄────────响应─────────── 服务端
```

### 7. Go 生成和验证 JWT

```go
package main

import (
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var secret = []byte("my-secret-key")

// 生成 JWT
func generateToken(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 2).Unix(),
    })
    return token.SignedString(secret)
}

// 验证 JWT
func parseToken(tokenStr string) (*jwt.Token, error) {
    return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secret, nil
    })
}

func main() {
    token, _ := generateToken("1001")
    fmt.Println("JWT:", token)

    parsed, err := parseToken(token)
    if err != nil {
        fmt.Println("验证失败:", err)
        return
    }
    if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
        fmt.Println("user_id:", claims["user_id"])
    }
}
```

### 8. JWT 优点与缺点

| 优点 | 缺点 |
|---|---|
| 无状态，服务端不用存 session | 令牌一旦签发，无法主动失效（除非加黑名单） |
| 跨服务共享方便 | payload 可解码，不能放敏感信息 |
| 性能好 | 令牌长度比 session id 大 |

### 9. JWT 常见问题

**Q：JWT 被盗怎么办？**
- 设置较短过期时间
- 使用 HTTPS
- 敏感操作二次验证
- 配合 Refresh Token

**Q：怎么让 JWT 失效？**
- 黑名单（Redis 存失效 token）
- 缩短有效期
- 修改签 secret（全部失效）

---

## 二、OAuth2

### 1. 什么是 OAuth2？

**一句话**：OAuth2 是一种授权框架，允许第三方应用在不获取用户密码的情况下，访问用户在另一服务上的资源。

### 2. 角色

| 角色 | 说明 |
|---|---|
| **Resource Owner** | 资源所有者，即用户 |
| **Client** | 第三方应用 |
| **Authorization Server** | 授权服务器，发放 Token |
| **Resource Server** | 资源服务器，保存用户资源 |

### 3. 四种授权模式

| 模式 | 适用场景 |
|---|---|
| **授权码模式（Authorization Code）** | 最常用、最安全，Web 应用 |
| **简化模式（Implicit）** | 单页应用，已不推荐 |
| **密码模式（Password）** | 受信任应用，如自家 App |
| **客户端模式（Client Credentials）** | 服务间调用 |

### 4. 授权码模式流程

```
用户 ──浏览器──► 第三方应用（Client）
                    │
                    ↓ 重定向到授权服务器
用户 ──浏览器──► 授权服务器（Authorization Server）
                    │
                    ↓ 用户登录并同意授权
用户 ◄──授权码──── 授权服务器
                    │
                    ↓ 携带授权码请求 Token
第三方应用 ◄──Access Token── 授权服务器
                    │
                    ↓ 用 Access Token 访问资源
第三方应用 ◄──用户资源──── 资源服务器
```

### 5. 授权码模式详解

```
Step 1: 申请授权码
GET /authorize?response_type=code
            &client_id=CLIENT_ID
            &redirect_uri=CALLBACK_URL
            &scope=read
            &state=xyz

Step 2: 用户同意后，重定向到回调地址
GET /callback?code=AUTH_CODE&state=xyz

Step 3: 用授权码换 Token
POST /token
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=AUTH_CODE
&redirect_uri=CALLBACK_URL
&client_id=CLIENT_ID
&client_secret=CLIENT_SECRET

Step 4: 返回 Token
{
  "access_token": "ACCESS_TOKEN",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "REFRESH_TOKEN"
}
```

### 6. Access Token 与 Refresh Token

| Token | 作用 | 有效期 |
|---|---|---|
| **Access Token** | 访问资源 | 短，如 1 小时 |
| **Refresh Token** | 换取新的 Access Token | 长，如 7 天 |

```
Access Token 过期后：
    Client ──Refresh Token──► Authorization Server
    Client ◄──新 Access Token── Server
```

### 7. JWT vs OAuth2 的关系

| 对比 | JWT | OAuth2 |
|---|---|---|
| 本质 | Token 格式 | 授权框架 |
| 关系 | OAuth2 返回的 Access Token 可以是 JWT | OAuth2 不限定 Token 格式 |
| 用途 | 认证、信息传递 | 授权第三方访问 |

---

## 三、SSO 单点登录

### 1. 概念

**一句话**：用户登录一次，即可访问多个互相信任的系统。

### 2. 基于 Cookie 的 SSO

```
系统 A ──登录──► 认证中心
                  │
                  ↓ 设置 Cookie（domain=.example.com）
系统 A ◄──登录成功── 认证中心

用户访问系统 B：
系统 B ──检查 Cookie──► 认证中心
                  │
                  ↓ Cookie 有效
系统 B ◄──已登录── 认证中心
```

### 3. 基于 JWT 的 SSO

```
用户 ──登录──► 认证中心
                  │
                  ↓ 签发 JWT
用户 ◄──JWT─── 认证中心

访问系统 A：Header 带 JWT
访问系统 B：Header 带 JWT
多个系统共享同一套 JWT 验证逻辑
```

---

## 四、面试高频问题

### Q1：JWT 和 Session 的区别？

| 特性 | JWT | Session |
|---|---|---|
| 存储位置 | 客户端 | 服务端 |
| 状态 | 无状态 | 有状态 |
| 扩展性 | 好 | 需共享 session |
| 安全性 | payload 可解码 | 不暴露数据 |
| 失效 | 难主动失效 | 服务端直接删 |

### Q2：OAuth2 授权码模式为什么安全？

**答**：
- 授权码通过前端浏览器传递
- Access Token 在服务端之间交换，不暴露给浏览器
- 需要 client_secret 才能换 Token
- 可以校验 redirect_uri 防止伪造

### Q3：Access Token 和 Refresh Token 为什么要分开？

**答**：
- Access Token 短有效期，降低被盗风险
- Refresh Token 长期有效，但只用于换 Access Token
- 即使 Access Token 泄露，有效期短影响有限

### Q4：JWT 放在哪里？Cookie 还是 Header？

**答**：
- 移动端/前后端分离：放 Header（`Authorization: Bearer <token>`）
- Web 页面：可放 Cookie（httpOnly + secure），防 XSS

### Q5：怎么防止 Token 被伪造？

**答**：
- 服务端用密钥签名
- 验证签名是否匹配
- 使用 HTTPS 防止传输中被截取
- 设置合理的过期时间

---

## 五、一句话总结

- **JWT**：自包含的 Token，适合无状态认证
- **OAuth2**：授权框架，让第三方安全访问用户资源
- **授权码模式**：最安全，Web 应用首选
- **Refresh Token**：长有效期，用于刷新短效的 Access Token
- **JWT 缺点**：无法主动失效、payload 可解码，需配合 HTTPS 和短有效期使用
