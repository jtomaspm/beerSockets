package server

import (
	"beerSockets/web/controller"
	"log"
	"net"
)

type ServerMessage struct {
	con     *net.Conn
	payload []byte
}

type ServerConfig struct {
	Host string
	Port string
}

type Server struct {
	config      ServerConfig
	ln          net.Listener
	quitch      chan struct{}
	msgch       chan ServerMessage
	router      *Router
	connections map[*net.Conn]bool
}

func New(config ServerConfig, controllerConfig controller.Config) *Server {
	r := NewRouter()
	cw := controller.NewControllerWrapper(controllerConfig)
	r.AddRoute("SAVE", cw.SAVE)
	r.AddRoute("GET", cw.GET)
	r.AddRoute("GETALL", cw.GETALL)

	return &Server{
		config:      config,
		quitch:      make(chan struct{}),
		msgch:       make(chan ServerMessage, 100),
		router:      r,
		connections: make(map[*net.Conn]bool),
	}
}

func (self *Server) acceptLoop() {
	for {
		con, err := self.ln.Accept()
		if err != nil {
			log.Println("Accept Error:", err)
			continue
		}
		self.connections[&con] = true
		log.Println("New Connection:", con.RemoteAddr())
		go self.readLoop(con)
	}
}

func (self *Server) readLoop(con net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := con.Read(buffer)
		if err != nil {
			break
		}
		self.msgch <- ServerMessage{
			con:     &con,
			payload: buffer[:n],
		}
	}
}

func (self *Server) handleMessages() {
	for msg := range self.msgch {
		log.Println("Received Message:", string(msg.payload))
		rmsg := NewRouterMessage(msg.payload)
		res := self.router.HandleRoute(rmsg)
		log.Println("Sending Response:", string(res))
		(*msg.con).Write(res)
		self.connections[msg.con] = false
		(*msg.con).Close()
		delete(self.connections, msg.con)
	}
}

func (self *Server) Run() error {
	ln, err := net.Listen("tcp", self.config.Host+":"+self.config.Port)
	if err != nil {
		return err
	}
	defer ln.Close()
	self.ln = ln

	go self.acceptLoop()
	log.Println("Server is running on", self.config.Host+":"+self.config.Port)
	go self.handleMessages()

	<-self.quitch
	close(self.msgch)
	return nil
}
