package boot

import (
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gsession"
)

// Session 全局配置
func init() {
	s := g.Server()

	s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute,
		"SessionStorage": gsession.NewStorageMemory(),
	})
}
