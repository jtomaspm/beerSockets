package server

import (
	"beerSockets/web/controller"
	"log"
	"net"
)

type ServerConfig struct {
	Host string
	Port string
}

type Server struct {
	config ServerConfig
	ln     net.Listener
	quitch chan struct{}
	msgch  chan []byte
	router *Router
}

func New(config ServerConfig, controllerConfig controller.Config) *Server {
	r := NewRouter()
	cw := controller.NewControllerWrapper(controllerConfig)
	r.AddRoute("SAVE", cw.SAVE)
	r.AddRoute("GET", cw.GET)
	r.AddRoute("GETALL", cw.GETALL)

	return &Server{
		config: config,
		quitch: make(chan struct{}),
		msgch:  make(chan []byte, 100),
		router: r,
	}
}

func (self *Server) acceptLoop() {
	for {
		con, err := self.ln.Accept()
		if err != nil {
			log.Println("Accept Error:", err)
			continue
		}
		log.Println("New Connection:", con.RemoteAddr())
		go self.readLoop(con)
	}
}

func (self *Server) readLoop(con net.Conn) {
	buffer := make([]byte, 2048)
	defer con.Close()
	for {
		n, err := con.Read(buffer)
		if err != nil {
			log.Println("Read Error:", err)
			continue
		}
		self.msgch <- buffer[:n]
	}
}

func (self *Server) handleMessages() {
	for msg := range self.msgch {
		log.Println("Received Message:", string(msg))
		rmsg := NewRouterMessage(msg)
		res := self.router.HandleRoute(rmsg)
		log.Println("Sending Response:", string(res))
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
