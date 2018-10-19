package http

import (
	"../g"
	"github.com/cihub/seelog"
	"net/http"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func buildChain(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	return m[0](buildChain(f, m[1:cap(m)]...))
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, err := r.Cookie("username")
		if err !=nil{
			seelog.Info(err)
			http.Error(w, "Not authorized", 401)
			return
		}
		password, err := r.Cookie("password")
		if err !=nil{
			seelog.Info(err)
			http.Error(w, "Not authorized", 401)
			return
		}
		seelog.Info(username, password,g.Cfg.UserName, g.Cfg.UserPassword,  g.Cfg.UserName == username.Value, g.Cfg.UserPassword == password.Value)
		if (g.Cfg.UserName == username.Value) && (g.Cfg.UserPassword == password.Value) {
			next.ServeHTTP(w, r)
		}else{
			http.Error(w, "Not authorized", 401)
			return
		}
	})
}
