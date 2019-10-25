package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/kataras/iris"
	"github.com/liov/hoper/go/v2/gateway/internal/client"
	"github.com/liov/hoper/go/v2/gateway/internal/service"
	"github.com/liov/hoper/go/v2/protobuf/user"
	"github.com/liov/hoper/go/v2/utils/log"
)

type UserController struct{
	Controller
}
var userService = &service.UserService{}
func (u *UserController) Add() {
	u.api(
		path(""),
		method(http.MethodPost),
		describe("新增用户"),
		auth("jyb"),
		version(1),
		handle(
			func(ctx iris.Context) {
				var req user.SignupReq
				ctx.ReadJSON(&req)
				gctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				rep,err:=client.UserClient.Signup(gctx,&req)
				if err != nil {
					log.Errorf("could not greet: %v", err)
				}
				ctx.JSON(rep)
			}),
	)

}

func (u *UserController) Get() {
	u.api(
		path("/:id"),
		method(http.MethodGet),
		describe("get"),
		auth("jyb"),
		version(1),
		handle(
			func(ctx iris.Context) {
				ctx.Writef("返回")
			}),
	)
}
