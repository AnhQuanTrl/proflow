package system

import (
	"github.com/AnhQuanTrl/proflow/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type System struct {
	cfg    config.AppConfig
	mux    *chi.Mux
	rpc    *grpc.Server
	logger zerolog.Logger
}

func New(cfg config.AppConfig) (*System, error) {
	s := &System{cfg: cfg}
	
}
