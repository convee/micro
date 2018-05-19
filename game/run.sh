#!/bin/sh

case $1 in
	build)
		go build -o gameServer
		echo "构建成功..."
		sleep 1
	;;
	start)
		nohup ./gameServer 2>&1 >> gameServer.log 2>&1 /dev/null &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall gameServer
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall gameServer
		sleep 1
		nohup ./gameServer 2>&1 >> gameServer.log 2>&1 /dev/null &
		echo "服务已重启..."
		sleep 1
	;;
	*) 
		echo "$0 {build|start|stop|restart}"
		exit 4
	;;
esac

