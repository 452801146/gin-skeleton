# gin-skeleton
####一个gin项目骨架,包含以下组件,具体看go.mod:

- orm(mysql)
- redis 
- viper
- zap
- ratelimit


#### 目录结构

- `conf` 配置文件目录
- `constant` 静态常量配置目录
- `controllers` 控制器目录
- `g` 公共变量/方法目录
- `initialize` 服务器初始化目录
- `logs` 日志目录
- `models` 模型目录
- `pkg` 第三方扩展包目录
- `routers` 路由目录
- `service` service目录
- `task` 任务/定时任务目录
- `utils` 工具目录
- `ws` ws逻辑目录