# 南京灵动信息技术有限公司官方网站

南京灵动信息技术有限公司官方网站前台、管理端和 Go 后端服务整合项目。前端使用 Vue 3 + Vite，后端位于 `server/`，提供公开内容接口、管理端接口、MySQL 持久化和本地文件上传存储。网站域名按建设文档为 `www.ilingdong.cn`。

## 前端启动

安装依赖：

```bash
npm install
```

本地开发：

```bash
npm run dev
```

生产构建：

```bash
npm run build
```

如需让前端请求本地后端，可在启动前设置：

```bash
$env:VITE_PUBLIC_API_BASE="http://localhost:8080"
npm run dev
```

管理端路由挂载在 `/admin`，登录页为 `/admin/login`。

## 后端启动

进入后端目录：

```bash
cd server
go mod download
go run .
```

默认监听 `:8080`，可通过环境变量调整：

```bash
$env:SERVER_ADDR=":8080"
$env:CORS_ORIGINS="http://localhost:5173"
go run .
```

后端测试：

```bash
cd server
go test ./...
```

## MySQL 测试库配置

后端通过环境变量读取 MySQL 配置，默认库名为 `wsd_social`。本地或测试环境可按需设置：

```bash
$env:DB_HOST="127.0.0.1"
$env:DB_PORT="3306"
$env:DB_NAME="wsd_social"
$env:DB_USER="root"
$env:DB_PASSWORD="your-password"
```

初始化表结构：

```bash
cd server
mysql -h 127.0.0.1 -P 3306 -u root -p wsd_social < schema.sql
```

`server/README.md` 中记录了临时测试库示例配置。该配置仅用于联调，不应复用到生产环境。

## 本地 `/oss` 存储

后端上传接口会把文件写入运行目录下的 `oss/YYYYMMDD/`，并返回 `/oss/...` URL。默认上传目录可通过 `UPLOAD_DIR` 修改：

```bash
$env:UPLOAD_DIR="oss"
go run .
```

后端会以 `/oss/` 路径静态托管该目录。开发时如果从前端页面访问上传资源，请确保前端 API Base 指向同一个后端地址，或在反向代理中转发 `/oss/`。

## 常用入口

- 前台首页：`/`
- 服务列表：`/service`
- 合作咨询：`/cooperate`
- 案例列表：`/case`
- 资讯列表：`/message`
- 关于我们：`/about`
- 管理端：`/admin`
