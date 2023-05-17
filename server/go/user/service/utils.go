package service

import (
	"github.com/actliboy/hoper/server/go/user/confdao"
	"github.com/hopeio/pandora/protobuf/errorcode"
	"github.com/hopeio/pandora/utils/verification"
)

func LuosimaoVerify(vCode string) error {
	if err := verification.LuosimaoVerify(confdao.Conf.Customize.LuosimaoVerifyURL,
		confdao.Conf.Customize.LuosimaoAPIKey, vCode); err != nil {
		return errorcode.InvalidArgument.Message(err.Error())
	}
	return nil
}
