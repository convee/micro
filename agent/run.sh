#!/bin/sh

case $1 in
	build)
		go build -o agentServer
		echo "构建成功..."
		sleep 1
	;;
	start)
		nohup ./agentServer 2>&1 >> agentServer.log 2>&1 /dev/null &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall agentServer
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall agentServer
		sleep 1
		nohup ./agentServer 2>&1 >> agentServer.log 2>&1 /dev/null &
		echo "服务已重启..."
		sleep 1
	;;
	*) 
		echo "$0 {build|start|stop|restart}"
		exit 4
	;;
esac

