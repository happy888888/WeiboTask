<div align="center"> 
<h1 align="center">
WeiboTask
</h1>

[![](https://img.shields.io/badge/author-%E6%98%9F%E8%BE%B0-red "作者")](https://github.com/happy888888/ )

</div>

### 💥主要功能
* [x] 微博超话签到

</br>

### 🚀使用方式

本项目***不会使用任何账号密码***，仅需要[新浪网](https://www.sina.com.cn/) 的一个名为`ALC`的cookie(获取方式见下面说明)并存放到config.json文件中 <br>
config.json文件会缓存您的cookie以***保存***和刷新您的登录状态，使登录状态一直有效而***不必重复登录***，从而方便放在路由器等设备上持续运行 <br>

#### 一、Windows本地运行

* 1.1 获取cookie
        电脑浏览器登录[新浪网](https://www.sina.com.cn/) <br>
        进入[某404 NOT FOUND网页](https://login.sina.com.cn/sso/test) <br>
		通过下图所示获取一个名为`ALC`的cookie的值并保存 <br>
		`在上述网址按F12打开开发者工具--》application--》cookies` <br>
		![获取cookie ALC](https://user-images.githubusercontent.com/67217225/102229329-9f5e5a00-3f26-11eb-929d-174539c489c3.png)

* 1.2 下载本项目
        进入本项目[Release](https://github.com/happy888888/WeiboTask/releases) ，下载windows版本压缩包  <br>
		解压后获得`WeiboTask.exe`, `config.json`两个文件(这两个文件放在同一文件夹) <br>
		
* 1.3 启动
        进入本项目[Release](https://github.com/happy888888/WeiboTask/releases) ，下载windows版本压缩包  <br>
		用记事本打开`config.json`文件，把***步骤1.1***中获得的cookie值填写到`"name": "ALC"`下面的`"value": ""`字段 <br>
		![image](https://user-images.githubusercontent.com/67217225/102366467-a69f6980-3ff3-11eb-84f7-5933da15f9a8.png) <br>
		保存文件后直接双击`WeiboTask.exe`文件启动

#### 二、openwrt等路由器运行（推荐）

* 2.1 获取cookie
        同上面***步骤1.1***

* 2.2 安装
        使用`xshell`等工具登录路由器，执行下面的命令安装  <br>
		`wget -O /tmp/WeiboTask.ipk "https://github.com/happy888888/WeiboTask/releases/download/1.0.0/WeiboTask_1.0.0_`uname -m`_24kc.ipk" && opkg install /tmp/WeiboTask.ipk` <br>
		使用`xftp`(`WinScp`)等软件修改路由器目录`/etc/WeiboTask/config.json`文件，把***步骤1.1***中获得的cookie值填写到`"name": "ALC"`下面的`"value": ""`字段，内容参照上面***步骤1.3***中所示图片 <br>
		执行命令启动 <br>
		`WeiboTask -c /etc/WeiboTask/config.json`
		
* 3.3 设置定时启动
        使用`xshell`等工具登录路由器，执行下面的命令配置crontab  <br>
		`0 0 * * * /usr/bin/WeiboTask -c /etc/WeiboTask/config.json >>/etc/crontabs/root && /etc/init.d/cron reload` <br>
		然后每天凌晨0点会自动运行  <br>
		`当然路由器网页中有类似"定时任务"的功能可以把"0 0 * * * /usr/bin/WeiboTask -c /etc/WeiboTask/config.json"追加到后面一行保存，也有相同的效果`

#### 三、github Actions运行

***强烈不推荐，异地ip登录会异常***

* 3.1 获取cookie
        同上面***步骤1.1***

* 3.2 配置secrets
        本项目仅需要配置名为`ALC`(必须，值为上面获取的cookie)和`SCKEY`(非必须，用于消息推送)  <br>
		![image](https://user-images.githubusercontent.com/67217225/102372598-511a8b00-3ffa-11eb-81c2-216463f60a9a.png)
		
* 3.3 启动
        ![image](https://user-images.githubusercontent.com/67217225/102372899-a0f95200-3ffa-11eb-920b-4eec5d328037.png)