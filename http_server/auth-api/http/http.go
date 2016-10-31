package http

import (
	"github.com/golang/glog"
	"github.com/oikomi/FishChatServer2/http_server/auth-api/conf"
	// "github.com/oikomi/FishChatServer2/common/ecode"
	"github.com/oikomi/FishChatServer2/common/net/xhttp"
	"github.com/oikomi/FishChatServer2/common/net/xhttp/router"
	wctx "github.com/oikomi/FishChatServer2/common/net/xweb/context"
	// "go-common/business/service/identify"
	"github.com/oikomi/FishChatServer2/common/ecode"
	"github.com/oikomi/FishChatServer2/http_server/auth-api/service"
	"net/http"
)

var (
	authSvc *service.Service
)

func Init(c *conf.Config) {
	// idfSvc = identify.New(c.Identify)
	// hisSvc = service.New(c)
	// init external router
	extM := http.NewServeMux()
	extR := router.New(extM)
	// extR.Identify = idfSvc
	// extR.Statsd = stats
	// extR.ELK = elk
	outerRouter(extR)
	// init Outer serve
	if err := xhttp.Serve(extM, c.MultiHTTP.Outer); err != nil {
		glog.Errorf("xhttp.Serve error(%v)", err)
		panic(err)
	}
	// init Inner router
	intM := http.NewServeMux()
	intR := router.New(intM)
	// intR.Identify = idfSvc
	// intR.Statsd = stats
	// intR.ELK = elk
	innerRouter(intR)
	// init Inner serve
	if err := xhttp.Serve(intM, c.MultiHTTP.Inner); err != nil {
		glog.Errorf("xhttp.Serve error(%v)", err)
		panic(err)
	}
	// init local router
	localM := http.NewServeMux()
	localR := router.New(localM)
	// localR.Identify = idfSvc
	// localR.Statsd = stats
	// localR.ELK = elk
	localRouter(localR)
	// init local server
	if err := xhttp.Serve(localM, c.MultiHTTP.Local); err != nil {
		glog.Errorf("xhttp.Serve error(%v)", err)
		panic(err)
	}
}

// outerRouter init outer router api path.
func outerRouter(r *router.Router) {
	// r.Degrade("/platform/history/degrade")
	// r.Get("/history/health/check", xhttp.SlbChecker(conf.Conf.CheckFile))
	// init api
	r.Group("/x/auth", func(cr *router.Router) {
		cr.GuestPost("/login", login)
	})
	return
}

// innerRouter init inner router api path.
func innerRouter(r *router.Router) {
	r.Group("/x/v2/history", func(cr *router.Router) {
		cr.VerifyPost("/add", addHistory)
		// cr.VerifyGet("/getAids", getAids)
	})
}

// innerRouter init local router api path.
func localRouter(r *router.Router) {
	// r.Get("/history/version", version)
	r.Get("/history/monitor/ping", ping)
}

// // getAids get a user history archive ids
// func getAids(c wctx.Context) {
// 	res := c.Result()
// 	mid, ok := c.Get("mid")
// 	if !ok {
// 		glog.Error("no mid")
// 		res["code"] = ecode.RequestErr
// 		return
// 	}
// 	res["data"], res["code"] = hisSvc.GetAids(c, mid.(int64))
// }

type aa struct {
	A string
}

// history get a user history, for mobile app service
func login(c wctx.Context) {
	res := c.Result()
	// pnStr := c.Request().Form.Get("pn")
	// psStr := c.Request().Form.Get("ps")
	// mid, ok := c.Get("mid")
	// if !ok {
	// 	log.Error("no mid")
	// 	res["code"] = ecode.RequestErr
	// 	return
	// }
	// // deal
	// pn, err := strconv.Atoi(pnStr)
	// if err != nil || pn < 1 {
	// 	pn = 1
	// }
	// ps, err := strconv.Atoi(psStr)
	// if err != nil || ps > conf.Conf.Max || ps <= 0 {
	// 	ps = conf.Conf.Max
	// }
	// res["data"], res["code"] = hisSvc.Get(c, mid.(int64), pn, ps, c.RemoteIP())
	var tmp []*aa
	tmp1 := &aa{
		A: "hello",
	}
	tmp = append(tmp, tmp1)
	res["data"], res["code"] = tmp, ecode.OK
}

// clearHistory user history
func clearHistory(c wctx.Context) {
	// res := c.Result()
	// mid, ok := c.Get("mid")
	// if !ok {
	// 	log.Error("no mid")
	// 	res["code"] = ecode.RequestErr
	// 	return
	// }
	// res["code"] = hisSvc.Clear(c, mid.(int64))
}

// deleteHistory delete aid from history
func deleteHistory(c wctx.Context) {
	// res := c.Result()
	// mid, ok := c.Get("mid")
	// if !ok {
	// 	log.Error("no mid")
	// 	res["code"] = ecode.RequestErr
	// 	return
	// }
	// aidStr := c.Request().Form.Get("aid")
	// aid, err := strconv.ParseInt(aidStr, 10, 64)
	// if err != nil {
	// 	log.Error("strconv.ParseInt(%s) error(%v)", aidStr, err)
	// 	res["code"] = ecode.RequestErr
	// 	return
	// }
	// res["code"] = hisSvc.Delete(c, mid.(int64), aid)
}

// addHistory add history into user redis set.
func addHistory(c wctx.Context) {
	// res := c.Result()
	// mid, ok := c.Get("mid")
	// if !ok {
	// 	log.Error("no mid")
	// 	res["code"] = ecode.RequestErr
	// 	return
	// }
	// aidStr := c.Request().Form.Get("aid")
	// aid, err := strconv.ParseInt(aidStr, 10, 64)
	// if err != nil {
	// 	log.Error("strconv.ParseInt(%s) error(%v)", aidStr, err)
	// 	res["code"] = ecode.RequestErr
	// 	return
	// }
	// res["code"] = hisSvc.Add(c, mid.(int64), aid, c.RemoteIP())
}

// ping check server ok.
func ping(c wctx.Context) {
	// res := c.Result()
	// if err := hisSvc.Ping(c); err != nil {
	// 	res["code"] = err
	// 	log.Error("history service ping error(%v)", err)
	// 	http.Error(c.Response(), "", http.StatusServiceUnavailable)
	// 	c.Done()
	// }
}

// version check server ver.
// func version(c wctx.Context) {
// 	res := c.Result()
// 	res["data"] = map[string]interface{}{
// 		"version": conf.Conf..Version,
// 	}
// }