# Go WSL Development

![Backend Development Pattern](img/design_pic-1.jpg)

## 1. Windows开发环境配置

### 1.1 安装WSL

_便于后面部署到linux系统上_

打开Windows PowerShell（管理员），输入安装命令

```bat
wsl --install
```

_若执行遇到报错：WslRegisterDistribution failed with error: 0x8007273f_
解决方案：

- 桌面新建一个记事本文件，将它的后缀改成cmd或bat，复制下面的代码

```bat
pushd "%~dp0"

dir /b %SystemRoot%\servicing\Packages\*Hyper-V*.mum >hyper-v.txt

for /f %%i in ('findstr /i . hyper-v.txt 2^>nul') do dism /online /norestart /add-package:"%SystemRoot%\servicing\Packages\%%i"

del hyper-v.txt

Dism /online /enable-feature /featurename:Microsoft-Hyper-V-All /LimitAccess /ALL
```

- 右键选择：以管理员身份运行。代码执行结束后，重启电脑。

- 打开Windows PowerShell（管理员），输入命令

```bat
wsl --set-default-version 1
```

- 去最近安装的程序中，打开刚才安装好的Ubuntu程序，输入账号密码，成功进入系统

- 依次输入命令

```bash
sudo apt update
sudo apt dist-upgrade
```

- 退出

- 打开Windows PowerShell（管理员），输入命令

```bat
wsl --set-default-version 2
```

_如果切换版本报错，尝试在Windows PowerShell（管理员）输入命令netsh winsock reset，重启电脑之后再尝试_

- 重启电脑

### 1.2 安装make

进入Ubuntu，输入命令

```bash
sudo apt update
sudo apt install make
```

### 1.3 安装go

```bash
sudo snap install go --classic
```

查看go版本

```bash
go version
```

当前安装版本为：go1.18.1 linux/amd64

### 1.4 安装sqlc

```bash
sudo snap install sqlc
```

当前安装版本为：sqlc 1.22.0 from Kyle Conroy (kjconroy)

### 1.5 下载安装VSCODE

在Windows系统中，下载安装
安装插件：WSL，GO，Remote Development

### 1.6 下载安装Docker Desktop

打开Docker Desktop，在设置->General中，勾选“Add the * docker ...”
_这将建立容器和主机之间的通信_

![Docker Desktop](img/image.png)

重新打开Ubuntu终端，检查安装是否正确

![Check](img/image-1.png)

## 2. 开始Go开发

- 新建项目文件夹project_name

- 进入文件夹，终端输入创建命令

```bash
go mod init project_name
```
  
- 创建main.go文件，编写代码

- 运行go主程序，终端输入创建命令

```bash
go run main.go 
```

## 3. 数据库

使用PostgreSQL数据库
_相较于MySQL，支持更多数据类型、更强大的扩展性和并发性_

使用 <https://dbdiagram.io> 作为数据库辅助设计工具

使用DBeaver作为数据库管理工具

### 3.1 数据库安装

_在WSL创建PostgreSQL的Docker Container_

#### 3.1.1 下载docker image

在 <https://hub.docker.com> 中搜索postgres，查看最新的image版本

![hub.docker.com](img/1697287476841.png)

返回WSL中（可以通过终端或者VSCode连接），输入命令

```bash
docker pull postgres:16-alpine
```

_这里选择alpine版本，是因为其体积较小，也可下载其它版本_

#### 3.1.2 启动docker容器

在WSL中，输入命令

```bash
docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin123 -p 5432:5432 -d postgres:16-alpine
```

>--name 数据库名称
-e 环境变量
-p 暴露的主机端口和容器端口
-d 要启动的容器对应的image

### 3.2 数据库设计

dbdiagram.io

![database dsgin](img/image-4.png)

Export to PostgreSQL

### 3.3 Database Migration

用于快速迁移和创建数据库、数据表

#### 3.3.1 安装Migration

在WSL中，切换到root权限（sudo -i），依次输入命令
_参考<https://github.com/golang-migrate/migrate/tree/master/cmd/migrate>_

```bash
curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate
```

#### 3.3.2 创建Migration文件

- 在项目根目录下，创建文件夹db，在db下创建文件夹migration

输入命令

```bash
migrate create -ext sql -dir db/migration -seq init_schema
```

这将新建1个init_schema.up.sql文件和1个init_schema.down.sql文件

其大概意义参照下图
![Migration up and down](img\1697353218235.jpg)

- 将之前dbdiagram导出的SQL语句复制进入init_schema.up.sql文件中

- 在init_schema.down.sql文件中，要恢复up脚本的更改，故输入以下脚本

```sql
DROP TABLE IF EXISTS inventories;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS accounts;
```

_删除表时的顺序要注意外键_

#### 3.3.3 创建Makefile文件

在项目根目录新建Makefile文件，输入如下代码

```makefile
postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin123 -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root red_pocket_game

dropdb:
	docker exec -it postgres16 dropdb --username=root --owner=root red_pocket_game

migrateup:
        migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable" -verbose up

migratedown:
        migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable" -verbose down 

.PHONY: postgres createdb dropdb migrateup migratedown
```

这样便可以使团队中的开发者，在WSL里快捷地启动docker容器，创建/删除数据库，执行数据库migrateup/down
![make example](img/1697356608549.png)

### 3.4 数据库管理工具

在Windows系统中，下载安装DBeaver（或其它数据库管理工具），打开并新建Postgres连接，输入数据库名称，用户名和密码

![DBeaver create connection](img/1697356951882.jpg)

测试连接无误后，点击完成，展开数据库表，可以看到上一步创建的表

![DB example](img/1697357042444.png)

### 3.5 生成数据库CRUD的GO代码

_使用SQLC工具将数据库表的CRUD操作转换为Go代码_

- 在WSL中，将路径切换到项目根目录，输入命令

```bash
sqlc init
```

- 这将新建1个sqlc.yaml文件，打开该文件，输入相关代码

_参考<https://docs.sqlc.dev/en/stable/reference/config.html#version-1>_

```yaml
version: "1"
packages:
  - name: "db"
    path: "./db/sqlc"
    queries: "./db/query/"
    schema: "./db/migration/"
    engine: "postgresql"
    emit_db_tags: false
    emit_prepared_queries: true
    emit_interface: false
    emit_exact_table_names: false
    emit_empty_slices: false
    emit_exported_queries: false
    emit_json_tags: true
    emit_result_struct_pointers: false
    emit_params_struct_pointers: false
    emit_methods_with_db_argument: false
    emit_pointers_for_null_types: false
    emit_enum_valid_method: false
    emit_all_enum_values: false
    json_tags_case_style: "camel"
    omit_unused_structs: false
    output_batch_file_name: "batch.go"
    output_db_file_name: "db.go"
    output_models_file_name: "models.go"
    output_querier_file_name: "querier.go"
    output_copyfrom_file_name: "copyfrom.go"
    query_parameter_limit: 1
```

- 在项目根目录新建Makefile文件，增加"sqlc: sqlc generate"

![Makefile add sqlc](img/1697372104205.png)

- 在db目录下，创建query文件夹，在query目录下，创建account.sql，items.sql，inventories.sql文件，并填入SQLC语句

这里后期可以根据需求增加更多的SQL语句，以供业务程序使用
_参考<https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html>_

![account.sql](img/1697376920509.png)

>--name表示自定义生成的Go方法名
:one表示返回一个对象
:many表示返回多个对象
:exec表示仅执行SQL语句，而不返回对象

- 在项目根目录下，输入命令

```bash
make sqlc
```

这将在"./db/sqlc"下自动生成5个文件

![account.sql.go](img/1697375254302.png)

>models.go 包含表的结构体
db.go 包含DBTX接口，让开发者可以自由地使用数据库进行查询
account.sql.go 在account.sql里对应自动生成的Go方法
inventories.sql.go 在inventories.sql里对应自动生成的Go方法
items.sql.go 在items.sql里对应自动生成的Go方法

## 4. Go服务器开发

### 4.1 handlers.go

在项目根目录下，新建handlers文件夹，里面存放各业务代码，如处理登录的login.go

![login.go](img/1697383836657.png)

### 4.2 main.go

配置路由规则，开始监听

![main.go](img/1697383884401.png)

### 4.3 API测试

使用Postman测试工具

![account.sql.go](img/1697384390928.png)

## x. Go语言特殊语法（持续更新）

### (1) := 和 =

```go
x := 10  // := 声明并初始化x为整数类型
        //     只能在函数内部使用
x = 10  // = 将10赋给x
        //   可以在函数内外使用，即可以用于package中全局变量的赋值
```

### (2) 首字母大小写

如果一个变量（或函数、结构体、接口等）的名称以大写字母开头，那么该变量就是导出的（Public），意味着它可以在该包以外的代码中被访问和使用（注意，在函数内部声明的变量（包括大写字母开头的）是局部变量，其作用域仅限于该函数内）

```go
package math

var Pi = 3.14159  //Pi 是一个导出的变量，因为它以大写字母开头
                  //可以在外部通过math.Pi访问（也可修改）

```

## 5. CICD

### 5.1 Dockerfile

### 5.2 AWS

#### 5.2.1 CICD by Github Action

##### 5.2.1.1 ECR (Elastic Container Registry)

_存储、管理和部署Docker容器镜像_

- 搜索ECR，并创建新的仓库

![search ECR](img/1697952517730.jpg)

- 输入仓库名，检查其它选项有无需要改变

![create ECR](img/1697952649477.png)

- 在项目根目录创建".github/workflows/deploy.yml"，并编写基础内容：

```yml
name: Deploy to production

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:

  build:
    name: Build image
    runs-on: ubuntu-latest

    steps:
```

- 在Github market里选择"Actions Type"，并搜索"aws ecr"，选择官方编写的"Amazon ECR "Login" Action for GitHub Actions"，并复制其"AWS credentials and region","login example usage"

![Github market action ECR](img/1697953167625.png)
_AWS credentials and region在下方_

将其粘贴到deploy.yml中

##### 5.2.1.2 Secrets

##### 5.2.1.3 yml file for Github Action

##### 5.2.1.4 K8s (Kubernetes)

#### 5.2.2 DB
