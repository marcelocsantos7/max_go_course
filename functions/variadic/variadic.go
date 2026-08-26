package variadic

type Server struct {
    Host string
    Port int
}

type Option func(*Server)

func WithPort(port int) Option {
    return func(s *Server) {
        s.Port = port
    }
}

func NewServer(host string, opts ...Option) *Server {
    // 1. Set default values
    srv := &Server{
        Host: host,
        Port: 8080, // Default Port
    }

    // 2. Override defaults with provided options
    for _, opt := range opts {
        opt(srv)
    }

    return srv
}

