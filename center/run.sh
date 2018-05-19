#!/bin/sh

case $1 in
	build)
		go build -o centerServer
		echo "构建成功..."
		sleep 1
	;;
	start)
		nohup ./centerServer 2>&1 >> centerServer.log 2>&1 /dev/null &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall centerServer
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall centerServer
		sleep 1
		nohup ./centerServer 2>&1 >> centerServer.log 2>&1 /dev/null &
		echo "服务已重启..."
		sleep 1
	;;
	*) 
		echo "$0 {build|start|stop|restart}"
		exit 4
	;;
esac

