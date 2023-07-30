package impl

import (
	"log"
	"os"

	"github.com/sunflower10086/TikTok/http/internal/pkg"
)

var impl = &UserServiceImpl{}

/*
	之前都是手动把实现的服务注册到IOC层的
	apps.HostService = impl.NewHostServiceImpl()
*/

type UserServiceImpl struct {
	l *log.Logger
}

// 通过匿名引入可以动态注册我们实现的服务
func init() {
	pkg.RegistryImpl(impl)
}

func (h *UserServiceImpl) Name() string {
	return pkg.UserAppName
}

func (h *UserServiceImpl) Config() {
	h.l = log.New(os.Stderr, "  [user] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewHostServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		l: log.New(os.Stderr, "  [user] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
