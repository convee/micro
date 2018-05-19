package main

import (
	"flag"
	"github.com/micro/go-grpc"
	"micro/agent/pb"
	"net"
	"golang.org/x/net/context"
	"github.com/satori/uuid"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
)


type sid string

type Session struct {
	logger *log.Logger
	id sid
	conn *net.TCPConn
	rpcStream pb.GameService_StreamClient
	dieChan chan sid
}

type AgentService struct {
	logger *log.Logger
	rpcClient pb.GameServiceClient
	sessions map[sid]*Session
	sessionDieChan chan sid
}

func newAgentService(rpcClient pb.GameServiceClient, logger *log.Logger) AgentService {
	return AgentService{
		rpcClient: rpcClient,
		sessions: make(map[sid]*Session),
		logger:logger,
	}
}

func newSession(id sid, conn *net.TCPConn, rpcStream pb.GameService_StreamClient, dieChan chan sid, logger *log.Logger) *Session {
	return &Session{
		id:id,
		conn:conn,
		rpcStream:rpcStream,
		dieChan:dieChan,
		logger:logger,
	}
}


func (a AgentService) closeSession(id sid)  {
	a.logger.Infof("close a session, id:%s", id)
	a.sessions[id].rpcStream.Close()
	a.sessions[id].conn.Close()
	delete(a.sessions, id)
}


func (a AgentService) tcpServer(agentAddr string)  {
	a.logger.Info("start a tcp server, address：", agentAddr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", agentAddr)
	if err != nil {
		a.logger.Errorf("create tcp error: %v", err)
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		a.logger.Errorf("listen tcp error: %v", err)
		return
	}
	a.logger.Infof("listening on:%v", listener.Addr())
	rpcStream, err := a.rpcClient.Stream(context.Background())

	if err != nil {
		a.logger.Errorf("rpc stream error: %v", err)
		return
	}
	// loop accepting
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			a.logger.Warnf("tcp accept failed: %v", err)
			continue
		}
		id := sid(uuid.NewV4().String())
		a.logger.Infof("start a session, id:%s", id)
		a.sessions[id] = newSession(id, conn, rpcStream, a.sessionDieChan, a.logger)
		go a.sessions[id].forwardToClient()
		go a.sessions[id].forwardToServer()
		go a.closeSession(id)
	}
}

func (s *Session) forwardToClient()  {
	s.logger.Info("a session forward to client")
	for {
		frame, err := s.rpcStream.Recv()
		if err != nil {
			s.logger.Errorf("rpc recv error: %v", err)
			s.dieChan <- s.id
			return
		}
		if _, err := s.conn.Write(frame.Payload); err != nil {
			s.logger.Errorf("tpc write error: %v", err)
			s.dieChan <- s.id
			return
		}
		select {
		case <-s.dieChan:
			return
		default:
		}
	}
}

func (s *Session) forwardToServer()  {
	s.logger.Info("a session forward to server")
	for {
		bytes, err := ioutil.ReadAll(s.conn)
		if err != nil {
			s.logger.Errorf("tcp read error: %v", err)
			s.dieChan <- s.id
			return
		}
		if err := s.rpcStream.Send(&pb.Frame{Payload:bytes}); err != nil {
			s.logger.Errorf("rpc stream send error: %v", err)
			s.dieChan <- s.id
			return
		}
		select {
		case <-s.dieChan:
			return
		default:
		}
	}
}

func main() {
	var (
		//定义ServerAgent服务端口
		agentAddr = flag.String("agent.addr", "127.0.0.1:10001", "agent server tcp address")
	)
	flag.Parse()

	//日志初始化
	var logger = log.New()
	logger.WithFields(log.Fields{
		"log_name": "agent",
	})
	logger.Info("start a grpc client")
	gSrv := grpc.NewService()
	gSrv.Init()
	client := pb.NewGameServiceClient("go.micro.srv.game", gSrv.Client())
	logger.Info("start a agent service")
	agentService := newAgentService(client, logger)

	//启动goroutine建立tcp连接并通过grpc转发给gameServer
	agentService.tcpServer(*agentAddr)

}
