# Go 后端服务

该目录提供官网公开 API 和后台管理 API，使用 Go 标准库 HTTP 服务、MySQL 持久化和本地 `/oss` 文件上传目录。

## 环境变量

```bash
SERVER_ADDR=:8080
DB_HOST=180.102.24.183
DB_PORT=3306
DB_NAME=wsd_social
DB_USER=root
DB_PASSWORD='jny@2216446'
ADMIN_USER=admin
ADMIN_PASSWORD='change-me'
TOKEN_SECRET='please-change-this-secret'
CORS_ORIGINS=http://localhost:5173
UPLOAD_DIR=oss
MAX_UPLOAD_MB=5
```

上面的 MySQL 地址、用户名和密码仅用于临时测试库示例。不要在生产环境复用该密码。

## 初始化数据库

```bash
mysql -h 180.102.24.183 -u root -p wsd_social < schema.sql
```

本次实现没有在开发过程中连接真实测试库，也没有执行破坏性数据库操作。

## 启动

```bash
go mod download
go run .
```

## 公开接口

| Method | Path | 说明 |
| --- | --- | --- |
| GET | `/api/public/site` | 站点配置 |
| GET | `/api/public/banner` | 首页 Banner |
| GET | `/api/public/cases?page=1&page_size=9` | 已发布案例列表 |
| GET | `/api/public/cases/{id-or-slug}` | 已发布案例详情 |
| GET | `/api/public/news?page=1&page_size=9` | 已发布资讯列表 |
| GET | `/api/public/news/{id-or-slug}` | 已发布资讯详情 |
| POST | `/api/public/forms` | 表单提交 |

案例列表支持 `industry`、`platform`、`strategy`、`keyword`；资讯列表支持 `category`、`keyword`。

## 管理接口

先登录获取 Token：

```bash
curl -X POST http://localhost:8080/api/admin/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"admin","password":"change-me"}'
```

后续请求使用 `Authorization: Bearer <token>`。

| Method | Path | 说明 |
| --- | --- | --- |
| GET / POST | `/api/admin/cases` | 案例列表、新增 |
| GET / PUT / DELETE | `/api/admin/cases/{id}` | 案例详情、更新、删除 |
| GET / POST | `/api/admin/news` | 资讯列表、新增 |
| GET / PUT / DELETE | `/api/admin/news/{id}` | 资讯详情、更新、删除 |
| GET / PUT | `/api/admin/banner` | Banner 查看、更新 |
| GET / PUT | `/api/admin/settings` | 站点设置查看、更新 |
| GET / PUT | `/api/admin/seo` | SEO 查看、更新 |
| GET | `/api/admin/forms` | 表单列表 |
| PATCH | `/api/admin/forms/{id}/status` | 表单状态更新，支持 `new`、`processed` |
| GET | `/api/admin/forms/export` | CSV 导出 |
| POST | `/api/admin/upload` | 文件上传，字段名为 `file` |

上传文件会写入当前运行程序目录下的 `oss/YYYYMMDD/`，接口返回 `/oss/...` URL。

## 安全说明

- 管理接口使用 HMAC Token 做基础鉴权。
- 所有 MySQL 写入和查询均使用参数化 SQL。
- 文本字段会去除 `<script>` 和普通 HTML 标签；富文本字段会去除 `<script>`。
- CORS 通过 `CORS_ORIGINS` 控制。
- 管理接口使用 Bearer Token，不依赖 Cookie，降低 CSRF 风险；生产环境仍建议配合 HTTPS、严格 Origin 和反向代理安全头。

## 测试

```bash
go test ./...
```

测试使用内存 Store，不连接真实 MySQL。
