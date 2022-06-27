# gin-admin-gorm
基于 GIN + GORM  + WIRE 实现的RBAC权限管理脚手架，目标是提供一套轻量的中后台开发框架，方便、快速的完成业务需求的开发。

* 以动态路由的方式实现不同的角色加载不同的菜单

> 账号：admin  密码：123456

> 账号：test   密码：123456

## 特性
* 遵循 `RESTful API` 设计规范 & 基于接口的编程规范
* 基于 `GIN` 框架，提供了丰富的中间件支持（JWTAuth、CORS、RequestRateLimiter、Recover、GZIP）
* 基于[jwt](https://github.com/appleboy/gin-jwt)，对API接口进行权限控制
* 基于[go-playground/validator](https://github.com/go-playground/validator)开源库简化gin的请求校验
* 用Docker上云
* 在token过期后的一个小时内，用户再次操作会要求重新登陆
* 基于[swaggo](https://github.com/swaggo)为Go工程生成自动化接口文档
* 基于[wire](https://github.com/google/wire)依赖注入
* 基于[gorm](https://gorm.io/zh_CN/)全功能ORM
* 基于[air](https://github.com/cosmtrek/air)自动编译，重启程序
* 基于redis限制请求频率

### 项目结构

<pre><code>
├─api             # API 处理层
│  └─v1            
├─cmd
│  └─GameAdmin
│    └─main.go    # 程序入口
├─configs         # 配置文件
├─docs            # 文档
├─internal
│  ├─conf         # 配置文件映射
│  ├─contextx     # 统一上下文处理
│  ├─errors       # 错误处理
│  ├─ginx         # gin 扩展模块
│  ├─middleware   # gin 中间件模块
│  ├─models       # 数据访问层
│  ├─routers      # 路由层
│  ├─schema       # 统一入参、出参对象映射
│  └─service      # 业务逻辑层
└─pkg
    ├─fileStore   # 文件存储
    ├─gormx       # gorm 扩展模块
    ├─helper      # 工具类
    ├─logger      # 日志模块
    ├─rabbitMQ    # 消息中间件
    ├─structure
    ├─timer       # 定时任务
    ├─types
    └─validator   # 参数校验库
</code></pre>




### 下载依赖
<pre><code>depend.cmd</code></pre>

### 代码生成与运行
##### 生成
<pre><code>build.cmd</code></pre>

##### 运行
<pre><code>run.cmd 或go run ./cmd/GameAdmin/ web -c ./configs/config.toml</code></pre>

## 前端工程

基于 [vue](https://github.com/vuejs/vue) 和 [element-ui](https://github.com/ElemeFE/element)实现：[gin-admin-vue](https://github.com/dot123/gin-admin-vue)
