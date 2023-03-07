# Collar

![](https://cdn.dvkunion.cn/collar/collar.png)

Collar 是长亭牧云主机助手(Collie) 的 CLI 工具，方便使用者在终端管理 Collie。

Collar 作为 Collie 牧羊犬的项圈，希望能让用户更加舒适的管理和使用 Collie。

会逐步从 Collie 核心功能(文件管理、在线终端)为主，抽空慢慢填齐其余的全部的功能(Docker、资源负载显示、登陆历史、进程清单等)

在这个基础上，可能还会做一些功能的优化，比如多主机的数据聚合、多主机同步执行命令(类ansible)等。

目前已支持的功能模块：

+ 主机列表
+ 进程列表
+ shell终端

## 开始使用

首先注册百川云平台，并开通 牧云主机助手 应用。

然后点击工作台-Token管理-生成Token，勾选所有牧云主机助手相关权限，生成您的Token信息。

![](https://cdn.dvkunion.cn/collar/c1d25c402e94487b8e0dcfd18d4c297a.png)

> 请注意，token 存在有效期。为了方便使用，您可以申请一个时间比较长的 token。  
> 您的 token 十分珍贵，可以获取您所有主机的权限！请妥善保管～

然后在github-release页面下载符合自己操作系统的二进制文件，放置在$PATH目录下

执行:  
`collar auth -t YOUT_TOKEN`

初始化身份认证成功，即可开始使用。

## 使用手册

### 主机列表

`collar hosts`

获取主机列表信息。

### 进程列表
`collar top [hostId/host_name/host_ip/host_inner_ip]`

获取主机进程信息， 每3s更新一次。

### 登陆主机 Terminal

`collar shell [hostId/host_name/host_ip/host_inner_ip]`

可以通过 主机ID/主机名/主机IP/主机内网IP 进行登录。  

**使用自动登陆模式**：

`collar shell -a [hostId/host_name/host_ip/host_inner_ip]`

这将会使用您配置的自动登录用户名进行登陆（暂不支持通过cli设置自动登录用户名)