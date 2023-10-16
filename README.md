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

#### 1.3 Use Git to start team development

>_You can also use the git shortcut button in the IDE, as long as you ensure that your actions are correct_

#### 1.3.1 Use branch to handle major changes

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

- When completing this stage of code writing, use the following command to add your changes to the local staging area

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

>Enter the GitHub repository page, open the Pull Request page, Click the "New Pull Request" button
Select the branch to merge
Compare changes
Fill in the modified information
Click the "Create Pull Request" button
Approve and merge the pull request on your own

#### 1.3.2 Directly handle small changes in the main branch

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

- Before submitting, make sure again that your branch is up to date **(Important)**

```bash
git checkout main
git pull origin main
```

- Submit your changes using the following command

```bash
git commit -m "A brief description of the work"
```

- Push to remote repository

```bash
git push origin main
```

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

## 3. Development instruction

### 3.1 Server Development

#### 3.1.1 Install Go

```bash
sudo snap install go --channel=1.18/stable
```

#### 3.1.2 Install sqlc

```bash
sudo snap install sqlc --channel=stable/1.22.0
```

#### 3.1.3 Start project

navigate to the project root directory and use the following commands

- Install the external libraries required for the project

```bash
go mod download
```

- Run project

```bash
go run main.go
```

- Verify

Use a browser or download the Postman tool for service verification

![verify example](img/5249398125634204d8c40607c285935.png)

### 3.2 Database

### 3.3 Test

## 4. CICD
