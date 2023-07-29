package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg"
	"github.com/sunflower10086/TikTok/http/internal/user"
)

var handler = &Handler{}

// 自注册到Ioc层
func init() {
	pkg.RegistryGin(handler)
}

//func NewHostHTTPHandler() *Handler {
//	return &Handler{}
//}

type Handler struct {
	svc user.Service
}

func (h *Handler) Name() string {
	return pkg.UserAppName
}

func (h *Handler) Config() error {
	hostService := pkg.GetImpl(pkg.UserAppName)
	if v, ok := hostService.(user.Service); ok {
		h.svc = v
		return nil
	}

	return fmt.Errorf("%s does not implement the %s Service interface", pkg.UserAppName, pkg.UserAppName)
}

func (h *Handler) RouteRegistry(g *gin.Engine) {
}
