# Photon GDBC Starter

一個基於 GORM 的資料庫連接管理模組，提供簡單易用的資料庫抽象層，支援 MySQL 和 PostgreSQL。

- 🎯 **GORM 整合**：基於強大的 GORM ORM 框架

## 安裝

```bash
go get github.com/Phofuture/photon-gdbc-starter
```

## 使用方式

### 基本配置

在您的應用程式配置檔案中添加資料庫配置：

```yaml
database:
  # 主資料庫配置
  host: localhost
  port: "3306"
  username: root
  password: password
  schema: mydb
  driver: mysql
  ifPrimary: true
  maxIdleConns: 10
  maxOpenConns: 100
  maxLifetimeSecond: 3600

  # 其他資料源（可選）
  dataSources:
    secondary:
      host: localhost
      port: "5432"
      username: postgres
      password: password
      schema: secondary_db
      driver: postgres
      maxIdleConns: 5
      maxOpenConns: 50
      maxLifetimeSecond: 1800
```

### 程式碼整合

```go
package main

import (
    _ "github.com/Phofuture/photon-gdbc-starter"
    "github.com/Phofuture/photon-core-starter/core"
)

func main() {
    // 啟動 Photon 應用程式
    core.Start()
}
```

## 支援的資料庫驅動

- **MySQL**: `mysql`
- **PostgreSQL**: `postgres`

## 配置參數說明

### ConnectData 結構

| 參數 | 型別 | 說明 |
|------|------|------|
| `host` | string | 資料庫主機地址 |
| `port` | string | 資料庫連接埠 |
| `username` | string | 資料庫使用者名稱 |
| `password` | string | 資料庫密碼 |
| `schema` | string | 資料庫名稱/模式 |
| `driver` | string | 資料庫驅動類型 |
| `ifPrimary` | bool | 是否為主資料庫 |
| `maxIdleConns` | int | 最大閒置連線數 |
| `maxOpenConns` | int | 最大開啟連線數 |
| `maxLifetimeSecond` | int | 連線最大生命週期（秒） |

## 相依套件

- [GORM](https://gorm.io/) - Go 的強大 ORM 函式庫
- [Photon Core Starter](https://github.com/Phofuture/photon-core-starter) - 核心框架
- MySQL Driver - `gorm.io/driver/mysql`
- PostgreSQL Driver - `gorm.io/driver/postgres`

