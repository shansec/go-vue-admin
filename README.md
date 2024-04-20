# go-vue-admin

## 📦 本地开发

### 环境要求

go 1.18

node版本:  v16.15.0

pnpm版本: 8.7.1

### 后端环境搭建

#### 1、创建后端代码目录

```bash
mkdir go-vue-admin
cd go-vue-admin
```

#### 2、获取后端代码

> 推荐前后端项目代码放在同一文件夹下；

```bash
# 获取后端代码
git clone https://github.com/shansec/go-vue-admin.git
```

#### 3、后端服务启动说明

```bash
# 进入 go-vue-admin 后端项目
cd ./go-vue-admin

# 更新整理依赖
go mod tidy

# 修改配置 
# 文件路径  ./config.yml
vi ./config.yml

# 编译项目
go build main.go

# 注意: config.mysql 下对应的数据库配置信息
```

### 前端启动说明

#### 1、创建前端代码目录

```bash
mkdir go-vue
cd go-vue
```

#### 2、获取后端代码

> 推荐前后端项目代码放在同一文件夹下；

```bash
# 获取前端代码
git clone https://github.com/shansec/go-vue.git
```

#### 3、安装依赖

```bash
pnpm install
```

#### 4、运行服务

```bash
# 启动服务
pnpm dev
```

### 放在最后的提醒
**为了确保生成代码功能的正常使用，需要把前后端项目放在同一个文件夹下**

