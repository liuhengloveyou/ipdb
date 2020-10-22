package api

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/liuhengloveyou/ipdb/common"
	"github.com/liuhengloveyou/ipdb/db"

	gocommon "github.com/liuhengloveyou/go-common"
	"go.uber.org/zap"
)

func InitHttpApi(addr string) error {
	http.Handle("/", &IPDB{})

	s := &http.Server{
		Addr:           addr,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}

type IPDB struct{}

func (p *IPDB) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer common.Logger.Sync()

	switch r.Method {
	case "GET":
		p.get(w, r)
	default:
		gocommon.HttpErr(w, http.StatusMethodNotAllowed, 0, "")
		return
	}
}

func (p *IPDB) get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ip := strings.TrimSpace(r.FormValue("ip"))
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	common.Logger.Debug("get ipip ", zap.String("ip", ip))

	ipRst, err := db.FindIP(ip)
	if err != nil {
		common.Logger.Sugar().Warnf("find ERR: %v\n", err.Error())
	}
	if ipRst == nil {
		ipRst = &db.IpRecord{}
	}
	ipRst.IP = ip
	common.Logger.Sugar().Infof("ipip: %v %v %v\n", ip, ipRst, err)

	// 以UA决定返回json或是页面
	ua := strings.ToLower(r.UserAgent())
	if strings.Index(ua, "curl") >= 0 {
		gocommon.HttpErr(w, 200, 0, ipRst) // JSON
		return
	}

	tmpl, err := template.New("index.html").ParseFiles("tmpl/index.html")
	if err != nil {
		common.Logger.Sugar().Errorf("template ERR: %v\n", err.Error())
		gocommon.HttpErr(w, 200, -1, "服务错误")
		return
	}

	if err = tmpl.Execute(w, ipRst); err != nil {
		common.Logger.Sugar().Errorf("template ERR: %v\n", err.Error())
	}

	return
}
