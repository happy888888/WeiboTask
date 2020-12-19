#!/bin/sh
name="WeiboTask"

mkdir -p ./ipk/opt/bin

cat>./ipk/postinst<<EOF1
#!/bin/sh
[ ! -d "/etc/WeiboTask" ] && mkdir /etc/WeiboTask
if [ ! -f "/etc/WeiboTask/config.json" ]; then
  cat>/etc/WeiboTask/config.json<<EOF
`cat config.json`
EOF
fi
[ ! -d "/etc/init.d" ] && mkdir /etc/init.d
cat>/etc/init.d/wbt<<EOF
#!/bin/sh
START=50
start() {
    echo "begin start"
	pid=\\\`ps | grep WeiboTask | grep -v 'grep' | awk '{print \\\$1}' | head -n 1\\\`
	if [ -n "\\\$pid" ]; then
      echo "Already started!"
	else
	  /opt/bin/WeiboTask -D
    fi
}
stop() {
    echo "begin stop"
	pid=\\\`ps | grep WeiboTask | grep -v 'grep' | awk '{print \\\$1}' | head -n 1\\\`
    if [ -n "\\\$pid" ]; then
	  kill -9 \\\$pid
      echo "stopped"
    else
      echo "Error! not started!" 1>&2
    fi
}
case "\\\$1" in
    start)
        start
        exit 0
    ;;
    stop)
        stop
        exit 0
    ;;
    reload|restart|force-reload)
        stop
        start
        exit 0
    ;;
    **)
        echo "Usage: \\\$0 {start|stop|reload}" 1>&2
        exit 1
    ;;
esac
EOF
chmod 755 /etc/init.d/wbt
echo "安装成功，请在下方粘贴cookie(ALC)并回车，然后程序会立即启动"
echo "若直接回车，程序不会启动，请在/etc/WeiboTask/config.json中配置ALC后"
echo "使用命令/etc/init.d/wbt restart启动"
read -p "粘贴cookie(ALC):" ALC
if [ -n "\$ALC" ]; then
  nn=\`grep -n '"name": "ALC"' /etc/WeiboTask/config.json | head -1 | cut -d ':' -f 1\`
  let nn+=1
  sed -i "\${nn}c \"value\": \"\${ALC}\"," /etc/WeiboTask/config.json
  /opt/bin/WeiboTask -D
fi
EOF1

cat>./ipk/prerm<<EOF1
#!/bin/sh
[ -d "/etc/WeiboTask" ] && rm -rf /etc/WeiboTask
[ -e "/etc/init.d/wbt" ] && rm -rf /etc/init.d/wbt
EOF1

chmod 755 ./ipk/postinst
chmod 755 ./ipk/prerm

echo "2.0" >./ipk/debian-binary


/bin/cp -f ./bin/linux_mips/$name ./ipk/opt/bin/

echo "Package: ${name}" >./ipk/control
echo "Version: ${version}" >>./ipk/control
echo "Section: lang" >>./ipk/control
echo "Maintainer: 星辰 <stars888888@outlook.com>" >>./ipk/control
echo "Architecture: all" >>./ipk/control
echo "Installed-Size: `stat -c "%s" ./ipk/opt/bin/$name`" >>./ipk/control
echo "Description:  微博签到任务" >>./ipk/control

tar -zcvf ./ipk/data.tar.gz --transform s=/ipk== ./ipk/opt
tar -zcvf ./ipk/control.tar.gz --transform s=/ipk== ./ipk/control ./ipk/postinst ./ipk/prerm
tar -zcvf ./ipk/${name}_${version}_mips.ipk --transform s=/ipk== ./ipk/data.tar.gz ./ipk/control.tar.gz ./ipk/debian-binary


rm -f ./ipk/data.tar.gz ./ipk/control.tar.gz
/bin/cp -f ./bin/linux_mipsle/$name ./ipk/opt/bin/

echo "Package: ${name}" >./ipk/control
echo "Version: ${version}" >>./ipk/control
echo "Section: lang" >>./ipk/control
echo "Maintainer: 星辰 <stars888888@outlook.com>" >>./ipk/control
echo "Architecture: all" >>./ipk/control
echo "Installed-Size: `stat -c "%s" ./ipk/opt/bin/$name`" >>./ipk/control
echo "Description:  微博签到任务" >>./ipk/control

tar -zcvf ./ipk/data.tar.gz --transform s=/ipk== ./ipk/opt
tar -zcvf ./ipk/control.tar.gz --transform s=/ipk== ./ipk/control ./ipk/postinst ./ipk/prerm
tar -zcvf ./ipk/${name}_${version}_mipsle.ipk --transform s=/ipk== ./ipk/data.tar.gz ./ipk/control.tar.gz ./ipk/debian-binary

rm -rf ./ipk/data.tar.gz ./ipk/control.tar.gz ./ipk/control ./ipk/postinst ./ipk/prerm ./ipk/opt ./ipk/debian-binary