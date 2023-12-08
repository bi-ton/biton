package webapi

import (
	"biton/core"

	"github.com/msw-x/moon/webs"
)

type Server struct {
	s *webs.Server
}

func New(opts Options, c *core.Core, version string) *Server {
	o := new(Server)
	o.s = webs.New().WithSecretDir(opts.CertDir)
	if opts.CertHost != "" {
		o.s.WithAutoSecret(opts.CertDir, opts.CertHost)
	}
	o.s.Run(opts.UiAddr, Routes(c, opts, version))
	return o
}

func (o *Server) Close() {
	o.s.Shutdown()
}
