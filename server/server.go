package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/walkergriggs/hellosb/api"
)

type Options struct {
	Host string
	Port int
	API  *api.API
}

type Server struct {
	Host         string
	Port         int
	KeepAlive    time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	listener     net.Listener

	api     *api.API
	handler http.Handler

	shutdown     chan struct{}
	shuttingDown int32
	interrupted  bool
	interrupt    chan os.Signal
}

func New(opts Options) *Server {
	return &Server{
		Host:         opts.Host,
		Port:         opts.Port,
		KeepAlive:    3 * time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,

		api:       opts.API,
		shutdown:  make(chan struct{}),
		interrupt: make(chan os.Signal, 1),
	}
}

func (s *Server) Listen() error {
	if s.listener != nil {
		return nil
	}

	listener, err := net.Listen("tcp", net.JoinHostPort(s.Host, strconv.Itoa(s.Port)))
	if err != nil {
		return err
	}

	s.listener = listener
	return nil
}

func (s *Server) Serve() (err error) {
	if s.listener == nil {
		if err = s.Listen(); err != nil {
			return err
		}
	}

	if s.handler == nil {
		s.handler = s.api.Configure()
	}

	wg := new(sync.WaitGroup)
	once := new(sync.Once)
	signalNotify(s.interrupt)
	go handleInterrupt(once, s)

	server := new(http.Server)
	server.ReadTimeout = s.ReadTimeout
	server.WriteTimeout = s.WriteTimeout
	server.SetKeepAlivesEnabled(int64(s.KeepAlive) > 0)
	server.Handler = s.handler

	wg.Add(1)
	log.Printf("Serving hellosb at http://%s\n", s.listener.Addr())

	go func(l net.Listener) {
		defer wg.Done()

		if err := server.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Fatalf("%v", err)
		}

		log.Printf("Stopped serving hellosb at http://%s\n", l.Addr())
	}(s.listener)

	wg.Add(1)
	go s.handleShutdown(wg, server)

	wg.Wait()
	return nil
}

func (s *Server) Shutdown() error {
	if atomic.CompareAndSwapInt32(&s.shuttingDown, 0, 1) {
		close(s.shutdown)
	}
	return nil
}

func (s *Server) handleShutdown(wg *sync.WaitGroup, server *http.Server) {
	defer wg.Done()
	<-s.shutdown

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	shutdownChan := make(chan bool)

	go func() {
		var success bool
		defer func() {
			shutdownChan <- success
		}()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v\n", err)
		} else {
			success = true
		}
	}()
}

func handleInterrupt(once *sync.Once, s *Server) {
	once.Do(func() {
		for range s.interrupt {
			if s.interrupted {
				continue
			}

			s.interrupted = true
			log.Println("Shutting down... ")
			if err := s.Shutdown(); err != nil {
				log.Printf("HTTP server Shutdown: %v\n", err)
			}
		}
	})
}

func signalNotify(interrupt chan<- os.Signal) {
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
}
