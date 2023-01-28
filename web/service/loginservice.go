package service

import (
	"errors"
	"taskmanager/web/constant"
	"taskmanager/web/dao"
	"taskmanager/web/entity"
	"taskmanager/web/entity/cmd"
	"taskmanager/web/entity/dto"
	"taskmanager/web/security"
	"taskmanager/web/token"
	"time"
)

type LoginService interface {
	Login(cmd *cmd.LoginCmd) entity.Response
}

type loginserviceimpl struct {
	accountdao      dao.AccountDao
	passwordencoder security.PasswordEncoder
	tokenservice    token.AuthTokenService
}

func NewLoginService() (LoginService, error) {
	accountdao, err := dao.NewAccountDao()
	if err != nil {
		return nil, err
	}

	return &loginserviceimpl{
		accountdao:      accountdao,
		passwordencoder: security.NewBcryptPasswordEncoder(),
		tokenservice:    token.NewAuthTokenService(),
	}, nil
}

func (impl loginserviceimpl) Login(cmd *cmd.LoginCmd) entity.Response {
	// 参数校验
	if cmd.User == "" || cmd.Pass == "" {
		return entity.Fail(constant.ParamErr)
	}

	// 根据用户名从数据库查询用户
	user, err := impl.accountdao.SelectByUsername(cmd.User)
	if err != nil {
		return entity.Fail(constant.Err)
	}

	// 密码校验
	matches := impl.passwordencoder.Matches(cmd.User, user.Password)
	if matches {
		token, err := impl.createtoken(cmd.User)
		if err != nil {
			return entity.Fail(constant.TokenGenerateFailed)
		}

		logindto := dto.LoginDto{
			Id:    user.Id,
			User:  cmd.User,
			Token: token,
		}

		return entity.OKWithData(logindto)
	}

	return entity.Fail(constant.PasswordMismatch)
}

func (impl loginserviceimpl) createtoken(user string) (string, error) {
	var payload map[string]string = map[string]string{}
	payload["user"] = user
	// 颁发Token
	tokenpalyload := token.TokenPayload{
		Expire:  time.Now().Add(7 * 24 * time.Hour),
		Payload: payload,
	}
	token := impl.tokenservice.CreateToken(tokenpalyload)
	if token != "" {
		return token, nil
	}

	return "", errors.New("create access token error")
}
