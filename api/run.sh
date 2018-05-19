#!/bin/sh

case $1 in
	build)
		go build -o apiServer
		echo "构建成功..."
		sleep 1
	;;
	start)
		nohup ./apiServer 2>&1 >> apiServer.log 2>&1 /dev/null &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall apiServer
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall apiServer
		sleep 1
		nohup ./apiServer 2>&1 >> apiServer.log 2>&1 /dev/null &
		echo "服务已重启..."
		sleep 1
	;;
	*) 
		echo "$0 {build|start|stop|restart}"
		exit 4
	;;
esac

