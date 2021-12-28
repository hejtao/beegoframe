package index

import (
	"beegoframe/app/dao"
	"beegoframe/app/shared"
	"beegoframe/pkg/encoding/json"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"regexp"
	"time"
)

func initFilter() {
	web.InsertFilter("*", web.BeforeStatic, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    allowMethods,
		AllowHeaders:    allowHeaders,
		ExposeHeaders:   exposeHeaders,
	}))
	web.InsertFilter("api/login", web.BeforeExec, login)
	web.InsertFilter("api/pri/*", web.BeforeExec, checkToken)
	web.InsertFilter("*", web.BeforeExec, formatRequestBody)
}

const (
	authHeader = "Authorization"
)

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginAccount struct {
	Id        int    `json:"id"`
	IsDeleted bool   `json:"is_deleted"`
	Password  string `json:"password"`
}

var allowMethods = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"OPTIONS",
}

var allowHeaders = []string{
	"Origin",
	"Access-Control-Allow-Origin",
	"Content-Type",
	authHeader,
}

var exposeHeaders = []string{
	"Content-Length",
	"Access-Control-Allow-Origin",
}

var matchTime = regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}`)

var formatRequestBody = func(ctx *context.Context) {
	b := ctx.Input.RequestBody
	ctx.Input.RequestBody = matchTime.ReplaceAllFunc(b, func(match []byte) []byte {
		t, _ := time.ParseInLocation(shared.TimeLayout, string(match), time.Local)
		return []byte(t.UTC().Format(shared.TimeLayoutZone))
	})
}

var login = func(ctx *context.Context) {
	var (
		form    loginForm
		account loginAccount
	)
	if err := json.Unmarshal(ctx.Input.RequestBody, &form); err != nil {
		ctx.Output.SetStatus(shared.Hc400)
		return
	}
	cond := dao.NewCond().
		Or(dao.Account.Fields.Email.Exact(), form.Username).
		Or(dao.Account.Fields.Phone.Exact(), form.Username)
	if err := dao.Account.SetCond(cond).FirstRow(&account); err != nil {
		ctx.Output.SetStatus(shared.Hc404)
		return
	}
	if account.IsDeleted {
		ctx.Output.SetStatus(shared.Hc200)
		ctx.Output.JSON(shared.Resp{
			Message: shared.RmAccountIsDeleted,
		}, false, false)
		return
	}
	// todo
}

var checkToken = func(ctx *context.Context) {
	token := ctx.Input.Header(authHeader)
	claims, err := parseToken(token)
	if err != nil {
		ctx.Output.SetStatus(shared.Hc401)
		return
	}
	_ = claims
	// todo
}
