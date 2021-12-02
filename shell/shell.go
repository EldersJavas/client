package shell

import (
	"github.com/ratel-online/client/ctx"
	"github.com/ratel-online/client/model"
	"github.com/ratel-online/core/log"
	"time"
)

type shell struct {
	ctx  *ctx.Context
	addr string
}

func New(addr string) *shell {
	return &shell{
		addr: addr,
	}
}

func (s *shell) Start() error {
	//fmt.Printf("usr: ")
	//username, err := util.Readline()
	//if err != nil {
	//    panic(err)
	//}
	//fmt.Printf("pwd: ")
	//password, err := gopass.GetPasswd()
	//if err != nil {
	//    panic(err)
	//}
	//resp, err := api.Login(string(username), string(password))
	//if err != nil {
	//    log.Error(err)
	//    return err
	//}
	s.ctx = ctx.New(model.LoginRespData{
		ID:       time.Now().UnixNano(),
		Name:     "Nico",
		Score:    100,
		Username: "Nico",
		Token:    "aeiou",
	})
	err := s.ctx.Connect("tcp", s.addr)
	if err != nil {
		log.Error(err)
		return err
	}
	err = s.ctx.Auth()
	if err != nil {
		log.Error(err)
		return err
	}
	return s.ctx.Listener()
}
