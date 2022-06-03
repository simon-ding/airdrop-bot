# airdrop-bot


## Doing
* 捋顺整个流程
* 完善每个步骤的代码表示

## Todo

* ServerChan服务器通知
* 利用aws服务器更换ip
* gas费实时数据
* 步骤中加入随机时间

## Thoughts

* Xvfb可以实现虚拟的显示服务，用来驱动chrome不使用headless模式

## 技术选型

 * 浏览器通讯：chromedp
 * 配置管理： viper
 * 日志服务：uber/zap
 * 数据库：sqlite + gorm
 * 错误处理：pkg/errors