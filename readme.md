基于 python3.8 的新 api server

## 开发环境

python 版本: 3.8

使用[poetry](https://github.com/python-poetry/poetry)进行依赖管理。

```shell
git clone https://github.com/bangumi/server bangumi-server
cd bangumi-server
```

进入虚拟环境

```shell
python -m venv .venv # MUST use python 3.8
source .venv/bin/activate # enable virtualenv
```

安装依赖

```shell
poetry install --remove-untracked
```

安装 git hook

```shell
pre-commit install
```

### 设置

可设置的环境变量

- `MYSQL_HOST` 默认 `127.0.0.1`
- `MYSQL_PORT` 默认 `3306`
- `MYSQL_DB` 默认 `bangumi`
- `MYSQL_USER` **无默认值**
- `MYSQL_PASS` **无默认值**

## 开发

### 项目结构

Web 框架 [fastapi](https://github.com/tiangolo/fastapi)

ORM 类定义在 [pol/db/tables.py](./pol/db/tables.py) 文件。

路由位于 [pol/api](./pol/api) 文件夹。

### 后端环境

https://github.com/bangumi/dev-env

启动开发服务器

```shell
uvicorn pol.server:app --reload --port 3000
```

## 测试

测试基于 pytest

### 运行测试(需要数据库)

```shell
pytest
```

### 编写测试

参照 [tests/app/test_base_router.py](./tests/app/test_base_router.py) 文件。在测试函数中添加`client`参数获取对应的 HTTP 测试客户端。`client` 是一个 `requests.Session` 的实例，可以使用 `requests` 的各种函数参数。

[详细文档](https://www.starlette.io/testclient/)

## 代码风格

以 LF 为换行符

### python

启用 [pre-commit](https://github.com/pre-commit/pre-commit)

```shell
pre-commit insall
```

pre-commit 会在当前仓库安装一个 git hook，在每次 commit 前自动运行。

也可以手动运行

```shell
pre-commit run #only check changed files
pre-commit run --all-files # check all files
```

lint: flake8

### 配置文件

非 python 文件(yaml, json, markdown 等)使用 [prettier](https://prettier.io/) 进行格式化。

## pol

pol 来源于我的旧项目名，没有特殊含义。

## License

BSD 3-Clause License

[LICENSE](https://github.com/bangumi/server/blob/master/LICENSE.md)
