package boot

import (
	"strconv"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

func init() {
	// 模板绑定函数
	g.View().BindFunc("personDate", personDate)
}

// a. 留言时间的显示是根据留言时间与当前时间相比的差值不同而显示不同的结果，设N为留言时间和当前时间的差值
// b. N＜60分钟，显示XX分钟前
// c. 1小时≦N＜24小时，显示XX小时前
// d. 1天≦N＜30天，显示XX天前
// e. 30天≦N，显具体的留言时间 年-月-日 时：分：秒
func personDate(t *gtime.Time) string {
	// 时间间隔
	diff := gtime.Now().Sub(t)
	// tTimestamp := t.Timestamp()

	//  N＜60分钟，显示XX分钟前
	if diff < 60*time.Minute {
		return strconv.Itoa(int(diff.Minutes())) + " 分钟前"
	}

	// 1小时≦N＜24小时，显示XX小时前
	if diff >= 60*time.Minute && diff < time.Hour*24 {
		return strconv.Itoa(int(diff.Hours())) + " 小时前"
	}

	// 1天≦N＜30天，显示XX天前
	if diff >= 24*time.Hour && diff < time.Hour*24*30 {
		return strconv.Itoa(int(diff.Hours()/24)) + " 天前"
	}

	return t.Format("Y年m月d日 H时i分s秒")
}
