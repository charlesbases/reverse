# reverse
```
数据库表结构生成结构体。支持 Mysql、Postgres
```

### 参数说明

```
-f string
  config file (default "reverse.toml")
```

- ##### reverse.toml

  [example](./example)

  ```
  - type: 数据库类型。支持 mysql、postgres
  - dir: models 文件生成目录
  - dsn: connect dsn
  ```

### 使用说明

```shell
# 使用默认当前目录下 reverse.toml
reverse

# 指定 toml
reverse -f other.toml
```

### 安装

```shell
go get github.com/charlesbases/reverse
```

