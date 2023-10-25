# go-admin

## 📦 本地开发

### 环境要求

go 1.18

node版本:  v16.15.0

pnpm版本: 8.7.1

### 开发目录创建

```bash
# 创建开发目录
mkdir go-vue-admin
cd go-vue-admin
```

### 获取代码

> 推荐两个项目必须放在同一文件夹下；

```bash
# 获取后端代码
git clone https://github.com/shansec/go-vue-admin.git

# 获取前端代码
git clone https://github.com/shansec/go-vue.git

```

### 后端启动说明

#### 服务端启动说明

```bash
# 进入 go-vue-admin 后端项目
cd ./go-vue-admin

# 更新整理依赖
go mod tidy

# 编译项目
go build

# 修改配置 
# 文件路径  go-vue-admin/config.yml
vi ./config.yml

# 1. 配置文件中修改数据库信息 
# 注意: config.mysql 下对应的配置数据
```

### 前端启动说明

```bash
# 安装依赖
pnpm install

# 启动服务
pnpm dev
```
