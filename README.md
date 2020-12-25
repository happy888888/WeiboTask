<div align="center"> 
<h1 align="center">
🔔WeiboTask⏰
</h1>

[![](https://img.shields.io/badge/author-%E6%98%9F%E8%BE%B0-red "作者")](https://github.com/happy888888/ )
![](https://img.shields.io/github/downloads/happy888888/WeiboTask/total?style=flat-square "下载量")
<br>
![](https://img.shields.io/docker/pulls/happy888888/weibotask?color=purple "docker拉取数")
![](https://img.shields.io/docker/image-size/happy888888/weibotask "docker镜像大小")

</div>

### 💥主要功能
* [x] 💰微博超话签到(pc web)⭐
* [x] 💯超话积分获取(超话app)🌀
* [x] 🚅超话评论帖子并删除评论(超话app)⚽
* [x] ⛄超话转发帖子并删除转发(超话app)⚾

<img src="https://user-images.githubusercontent.com/67217225/102878310-0037e600-4483-11eb-85a8-ee0ab4b496c1.png" width="500" height="600" title="任务列表" style="display:block;" />

* [x] Ⓜ签到(微博app)☀
* [x] ☁评论帖子并删除(微博app,需要S参数)☎
* [x] ☕点赞帖子并删除(微博app,需要S参数)☔
* [x] ♠转发帖子并删除(微博app,需要S参数)♣
* [x] ♦关注用户并取关(微博app,需要S参数)♨
* [x] ⭐刷关注微博(微博app,需要S参数)⭕

<img src="https://user-images.githubusercontent.com/67217225/103014306-8898b280-4579-11eb-8935-04602a8a7e9d.png" width="500" height="900" title="任务列表" style="display:block;" />

🔊📒📓📔📕📗📘📙📚📖

</br>

### 🚀使用方式

本项目***不会使用任何账号密码***❌，仅需要[新浪网👼](https://www.sina.com.cn/) 的一个名为`ALC`✌的cookie(获取方式见下面说明)并存放到config.json文件中(微博app任务需要S参数,只能靠手机抓包得到) <br>
config.json文件会缓存您的cookie以***保存***和刷新您的***登录状态***，使登录状态一直有效而***不必重复登录***，从而方便放在路由器(opwnwrt,padavan...)，nas(群辉,威联通...)，电视盒子(armbian等linux系统)等设备上持续运行 <br>

#### ➰一、Windows本地运行

* 1.1 获取cookie🆗
        电脑浏览器登录[新浪网👼](https://www.sina.com.cn/) <br>
        进入[某404 NOT FOUND网页💀](https://login.sina.com.cn/sso/test) <br>
		通过下图所示获取一个名为`ALC`的cookie的值并保存 <br>
		`在上述网址按F12打开开发者工具--》application--》cookies` <br>
		![获取cookie ALC](https://user-images.githubusercontent.com/67217225/102229329-9f5e5a00-3f26-11eb-929d-174539c489c3.png)

* 1.2 下载本项目⛳
        进入本项目[Release](https://github.com/happy888888/WeiboTask/releases) ，下载windows版本压缩包  <br>
		解压后获得`WeiboTask.exe`, `config.json`两个文件(这两个文件放在同一文件夹) <br>
		
* 1.3 启动⚡
        进入本项目[Release](https://github.com/happy888888/WeiboTask/releases) ，下载windows版本压缩包  <br>
		用记事本打开`config.json`文件，把***步骤1.1***中获得的cookie值填写到`"name": "ALC"`下面的`"value": ""`字段 <br>
		![image](https://user-images.githubusercontent.com/67217225/103126469-6f068080-46c9-11eb-890e-a6f4cb6088bb.png) <br>
		保存文件后直接双击`WeiboTask.exe`文件启动💎

#### 🚩✅二、openwrt等路由器运行（推荐）

* 2.1 获取cookie🆗
        同上面***步骤1.1☝***

* 2.2 安装⛳
        使用`xshell`等工具登录路由器，执行下面的命令安装  <br>
		```wget -O /tmp/WeiboTask.ipk "https://github.com/happy888888/WeiboTask/releases/download/1.0.1/WeiboTask_1.0.1_`uname -m`.ipk" && opkg install /tmp/WeiboTask.ipk``` <br>
		在安装时会提示填入***步骤2.1***获取的cookie(ALC)，然后程序会自动运行，默认在每天⏰`00:00`⏰自动签到

* 2.3 其他⛅
        启动程序的命令为`/etc/init.d/wbt start`
        关闭程序的命令为`/etc/init.d/wbt stop`
		重启程序的命令为`/etc/init.d/wbt restart`
		安装完成如果输入了cookie(ALC)程序会自动启动不需要使用命令再启动一次
		配置文件存放在`/etc/WeiboTask/config.json`文件中，包括保存的cookie和server酱推送的SCKEY以及每天的签到时间
		卸载程序直接运行命令`opkg remove WeiboTask`💔

#### 🚢三、docker运行

[docker hub地址](https://registry.hub.docker.com/repository/docker/happy888888/weibotask) 

|  平台   | tag标签  |
|  ----  | ----  |
| windows/linux(x64)  | latest |
| linux(arm32)  | arm-latest |
| linux(arm64)  | arm64-latest |

* 3.1 获取cookie🆗
        同上面***步骤1.1☝***

* 3.2 填写配置文件⛳
        下载config.json文件并填写ALC(cookie)那一项，存放到`/etc/WeiboTask/config.json`(其他目录也行)
		
* 3.3 启动⚡
        运行命令`docker run -d -v /etc/WeiboTask:/etc/WeiboTask happy888888/weibotask:latest`

#### 🚧四、github Actions运行

***强烈不推荐❌，异地ip登录会异常🚫***

* 4.1 获取cookie🆗
        同上面***步骤1.1☝***

* 4.2 配置secrets⛳
        本项目仅需要配置名为`ALC`(必须，值为上面获取的cookie)和`SCKEY`(非必须，用于消息推送)  <br>
		![image](https://user-images.githubusercontent.com/67217225/102372598-511a8b00-3ffa-11eb-81c2-216463f60a9a.png)
		
* 4.3 启动⚡
        ![image](https://user-images.githubusercontent.com/67217225/102372899-a0f95200-3ffa-11eb-920b-4eec5d328037.png)