# Development instructions

>**Recommended development OS version: Ubuntu 22.04.3 LTS**
**For other OS, such as MacOS, some details may be different, please use Google/Chatgpt to solve it**

![Backend Development Pattern](img/design_pic2.jpg)

[Go_Backend_Devlopment_Note](Go_Backend_Devlopment_Note.md) has my entire project learning process as a reference

_Continuously updated..._

## 1. Git

### 1.1 Install Git

- Enter the following code in the terminal

    ```bash
    sudo apt install git
    ```

- Set global name and mailbox

    ```bash
    git config --global user.name "Your Name"
    git config --global user.email "your_email@example.com"
    ```

    _Same as you registered on the Github website_

- Generate SSH public key

    _This step is not necessary, only for subsequent interaction with remote code libraries using the SSH protocol. You can also use the HTTP protocol, then you can skip this step, but it is recommended to use the SSH protocol because it has higher security and write permissions_

    ```bash
    ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
    ```

    >Enter file in which to save the key (/home/your_user/.ssh/id_rsa)
        _Prompt for the save path and file name of the key file (by default, it will be saved in the~/. ssh directory)_
    Enter passphrase (empty for no passphrase)
        _Prompt for the password used to encrypt the private key (passphrase). You can choose not to set the password and simply **press Enter** to leave it blank_
    Enter same passphrase again
        _Prompt to confirm password. If a password is set, enter the same password again to confirm_

    View Public Key

    ```bash
    cd ~/.ssh # 进入ssh文件夹内
    cat id_rsa.pub # 查看公钥
    ```

    _Alternatively, use other file viewing tools such as VSCode_

    Copy Public Key(starts with 'ssh rsa', followed by a large string of characters)

    Enter the Github website, Settings ->SSH and GPG keys ->New SSH key

    ![new ssh](img/1697423021183.png)

    Above is the title name, which can be set by yourself. Below is the place to paste your public key. Finally, clicking on Add SSH key.

### 1.2 Clone code

```bash
git clone git@github.com:ZhangZhengruiNUS/Red_Pocket_Game_Server.git
```

_If using the HTTP protocol, please modify the above command_

## 2. Docker & Postgresql

### 2.1 Install Docker

_If it's MacOS, download Docker Desktop directly from the official website_

For Ubuntu

```bash
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo systemctl start docker
sudo systemctl enable docker
```

### 2.2 Docker image

Search for postgres in <https://dbdiagram.io> to view the latest image version
![hub.docker.com](img/1697287476841.png)

Use the following commands in the terminal

```bash
docker pull postgres:16-alpine
```

_choose the Alpine version here, because it has a smaller volume_

### 2.3 Run postgres docker container

#### 2.3.1 Install make

Use the following commands in the terminal

```bash
sudo apt update
sudo apt install make
```

### 2.3.2 Use make to run docker

In the terminal, navigate to the project root directory and use the following commands

```bash
make postgres
```

This will trigger the instruction corresponding to "postgres" in the Makefile

![Makefile-postgres](img/1697431748554.jpg)

>--name database name
-e environment variable
-p exposed host and container ports
-d image name corresponding to the container to be launched

### 2.4 Create database

In the terminal, navigate to the project root directory and use the following commands

```bash
make createdb
```

This will trigger the instruction corresponding to "createdb" in the Makefile:
>docker exec -it postgres16 createdb --username=root --owner=root red_pocket_game

### 2.5 Create tables of database

In the terminal, navigate to the project root directory and use the following commands

```bash
make migrateup
```

This will trigger the instruction corresponding to "createdb" in the Makefile:
>migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable" -verbose up

This command will execute the SQL scripts in the latest **0XXXXX_init_schema.up.sql** file under the Red_Pocket_Game_Server/db/migration

![init_schema.up.sql](img/1697432658910.png)

## 3. Development Instruction

### 3.1 Preparation

#### 3.1.1 IDE

Recommend to use VSCode, other IDE is also acceptable
Recommended installation extensions: Go

#### 3.1.2 Component Installation

- Install Go

```bash
sudo snap install go --channel=1.18/stable
```

- Install godef (optional)

Code Navigation Tools

```bash
go get -u github.com/rogpeppe/godef
```

- Install gopls (optional)

Intelligent perception and code navigation

```bash
go get -u golang.org/x/tools/gopls
```

- Install sqlc

```bash
sudo snap install sqlc --channel=stable/1.22.0
```

### 3.2 Common rules

- Use CamelCase to name variable, like this:

>userName, inventoryId

- Write more comments to show indicate the work of the regions, like this:

![comments example](img/1697648283943.jpg)

### 3.3 Use Git to start team development

>_You can also use the git shortcut button in the IDE, as long as you ensure that your actions are correct_

#### 3.3.1 Use branch to handle major changes

- When you want to start your current work for the first time, please create a new branch, and it is recommended to summarize and name it with the approximate work you will complete this time.

```bash
git checkout -b my-work
```

- If you are unsure of the current branch, please use the  following command to confirm

```bash
git branch
```

_All branches will be listed and marked with * next to the current branch_

- If the current branch is not the feature/my feature branch you created you are preparing to write code, please use the following command to switch to that branch

```bash
git checkout my-work
```

- When completing this stage of code writing

Use the following command to get latest code on main branch **(Important)**
_Please do this frequently to ensure your code has the latest version of the main branch_

```bash
git fetch origin
git merge origin/main
```

If there's no more modification, use the following command to add your changes to the local staging area

```bash
git add .
```

_This will add all modified files to the staging area. If you only want to add specific files, you can Replace "." with the file name._

- Submit your changes using the following command

```bash
git commit -m "A brief description of the work"
```

- Push to remote repository

```bash
git push origin my-work
```

- Pull request to merge branch

>a. Enter the GitHub repository page, open the Pull Request page, Click the "New Pull Request" button
b. Select the branch to merge
c. Compare changes
d. Fill in the modified information
e. Click the "Create Pull Request" button
f. Approve and merge the pull request on your own

#### 3.3.2 Directly handle small changes in the main branch

>_Be cautious when making modifications on the main branch. Ensure that no other person is making changes that conflict with your changes at the same time._

- First, make sure you are on the main branch

```bash
git checkout main
```

- Get the latest code **(Important)**

```bash
git pull origin main
```

- When completing modification, use the following command to add your changes to the local staging area

```bash
git add .
```

- Submit your changes using the following command

```bash
git commit -m "A brief description of the work"
```

- Push to remote repository

```bash
git push origin main
```

### 3.4 Server Development

#### 3.4.1 Run project

navigate to the project root directory and use the following commands

- Install the external libraries required for the project

```bash
go mod download
```

- Run project

```bash
make server
```

This will trigger the instruction corresponding to "createdb" in the Makefile:
>go run main.go

- Verify

Download the Postman tool for service verification

![verify example](img/5249398125634204d8c40607c285935.png)

#### 3.4.2 Devlopment

##### 3.4.2.1 Router

When you want to add a new route, modify "NewServer" function in "handlers/server.go"

![server.go example](img/1697479547897.png)

##### 3.4.2.2 Service

When you want to add a service in a new field, add a file under "handlers"

![add new service](img/1697479982877.png)

##### 3.4.2.3 Sql

(1) Non-Transactions

>Only in the following two cases, you can use this method:
>a. Change(Create/Update/Delete) only one table with only one record
>b. Only read tables

- Write the sqlc scripts in the "db/query/table_name.sql"(follow the syntax of sqlc), like this:
_(ref: <https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html>)_

![sqlc scripts example](img/1697650119551.jpg)

>-- name: GoFunctionName -> The Go function name you want generate automatically
>:one -> Return only one record
>:many -> Return some records
>:exec -> Return zero record

- in the terminal, navigate to the project root directory and use the following commands

```bash
make sqlc
```

This will generate the affected files under "./db/sqlc", then you can find the sqlc sctipts corresponding Go functions there and use them in the handler services

(2) Transactions

- Change(Create/Update/Delete) only one table with only one record

>In the following two cases, but not only the two cases, you can use this method:
>a. When you want to change(Create/Update/Delete) more than one table
>b. When you want to change(Create/Update/Delete) one table with more than one record

### 3.4 Database

>In this way, every change to the database structure is recorded in detail and can be easily applied or rolled back.

When you want to change the structure of the tables

- In the terminal, navigate to the project root directory and use the following commands

```bash
make migratecreate name=migration_name
```

_describe the changes you want to make in the tables in the "name" parameter_
This will trigger the instruction corresponding to "migratecreate" in the Makefile:
>migrate create -ext sql -dir db/migrations -seq $(name)

This command will generate **0XXXXX_migration_name.up.sql** & **0XXXXX_migration_name.down.sql** file under the Red_Pocket_Game_Server/db/migration (the serail number will automatically increase)

- Write the difference script in the 0XXXXX_migration_name.up.sql file and write rollback scripts in the 0XXXXX_migration_name.up.sql
_if the changes are quite significant, you can try to use some database compare tools to automatically generate the scripts_

>**Important: Do consider the impact of these changes on existing data**

- In the terminal, navigate to the project root directory and use the following commands to apply the changes

```bash
make migrateup
```

If you want to rollback the changes, use the following commands

```bash
make migratedown
```

After you run the command, if there's any mistake in the scripts, the migration will stop and when you run the migraion command again, it will probably show the error message:
>error: Dirty database version 1. Fix and force version.

This is because before the migration runs, each database will be set with a dirty flag, if the migration fails, the dirty state will persist. This will prevent future migration operations.
You can update the change the version of the database into the lastest version forcefully to solve this, use the following commands

```bash
make migrateforce version=latest version number
```

_put the latest version number in the "version" parameter_
This will trigger the instruction corresponding to "migrateforce" in the Makefile:
>migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable" force $(version)

Then you remove the dirty flag, you can continue the migration operations.

- check if your "db/query/table.sql" need to change

If you make some changes in the files
In the terminal, navigate to the project root directory and use the following commands to apply the changes

```bash
make sqlc
```

This will regenerate the affected files under "./db/sqlc"

![account.sql.go](img/1697375254302.png)

### 3.4 Test

## 4. CICD
