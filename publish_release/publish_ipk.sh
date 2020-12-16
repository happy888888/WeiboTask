#!/bin/sh
name="WeiboTask"

mkdir -p ./ipk/usr/bin

cat>./ipk/postinst<<EOF1
#!/bin/sh
[ ! -d "/etc/WeiboTask" ] && mkdir /etc/WeiboTask
if [ ! -f "/etc/WeiboTask/config.json" ]; then
  cat>/etc/WeiboTask/config.json<<EOF
`cat config.json`
EOF
fi
EOF1

cat>./ipk/prerm<<EOF1
#!/bin/sh
[ -d "/etc/WeiboTask" ] && rm -rf /etc/WeiboTask
EOF
fi
EOF1

echo "2.0" >./ipk/debian-binary


/bin/cp -f ./bin/linux_mips/$name ./ipk/usr/bin/

echo "Package: ${name}" >./ipk/control
echo "Version: ${version}" >>./ipk/control
echo "Section: lang" >>./ipk/control
echo "Maintainer: 星辰 <stars888888@outlook.com>" >>./ipk/control
echo "Architecture: mips_24kc" >>./ipk/control
echo "Installed-Size: `stat -c "%s" ./ipk/usr/bin/$name`" >>./ipk/control
echo "Description:  微博签到任务" >>./ipk/control

tar -zcvf ./ipk/data.tar.gz --transform s=/ipk== ./ipk/usr
tar -zcvf ./ipk/control.tar.gz --transform s=/ipk== ./ipk/control ./ipk/postinst ./ipk/prerm
tar -zcvf ./ipk/${name}_${version}_mips_24kc.ipk --transform s=/ipk== ./ipk/data.tar.gz ./ipk/control.tar.gz ./ipk/debian-binary


rm -f ./ipk/data.tar.gz ./ipk/control.tar.gz
/bin/cp -f ./bin/linux_mipsle/$name ./ipk/usr/bin/

echo "Package: ${name}" >./ipk/control
echo "Version: ${version}" >>./ipk/control
echo "Section: lang" >>./ipk/control
echo "Maintainer: 星辰 <stars888888@outlook.com>" >>./ipk/control
echo "Architecture: mipsle_24kc" >>./ipk/control
echo "Installed-Size: `stat -c "%s" ./ipk/usr/bin/$name`" >>./ipk/control
echo "Description:  微博签到任务" >>./ipk/control

tar -zcvf ./ipk/data.tar.gz --transform s=/ipk== ./ipk/usr
tar -zcvf ./ipk/control.tar.gz --transform s=/ipk== ./ipk/control ./ipk/postinst ./ipk/prerm
tar -zcvf ./ipk/${name}_${version}_mipsle_24kc.ipk --transform s=/ipk== ./ipk/data.tar.gz ./ipk/control.tar.gz ./ipk/debian-binary

rm -rf ./ipk/data.tar.gz ./ipk/control.tar.gz ./ipk/control ./ipk/postinst ./ipk/prerm ./ipk/usr ./ipk/debian-binary