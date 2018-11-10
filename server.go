package main
import (
	"net/http"
)

type User struct{
	username string
	id string
	password string
}

var userlist []User

func FindUser(key string) (pwd string, is_present bool){
	for _,usr := range userlist{
		if usr.username == key{
			return usr.password, true
		}
	}
	return "", false
}

func SigninHandler(w http.ResponseWriter ,r *http.Request){
	err :=r.ParseForm()
	if err!= nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad request!"))
		return
	}
	// now r.Form has valid data
	name := r.FormValue("username")
	mailid := r.FormValue("mailid")
	pwd :=r.FormValue("password")
	re_pwd := r.FormValue("passwordconfirm")
	if(len(name)==0 || len(pwd)==0 || len(re_pwd)==0){
		w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("400 Bad request!"))
        return
	}
	/// all values intact. check pwd and confirm match
	if(pwd != re_pwd){
	    w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("400 mismatch password and conf fields!"))
        return
	}
	userlist = append(userlist, User{name,mailid,pwd})
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("201 . created successfully"))
}	

func LoginHandler(w http.ResponseWriter ,r *http.Request){
    err :=r.ParseForm()
    if err!= nil{
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("400 Bad request!"))
        return
    }
    name := r.FormValue("username")
    pwd :=r.FormValue("password")
	if(len(name)==0 || len(pwd)==0){
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("400 Bad request!"))
        return
    }
	saved_pwd,present := FindUser(name) 
	if present == false{
        w.WriteHeader(http.StatusForbidden)
        w.Write([]byte("404 no user found!"))
        return
	}
	if saved_pwd != pwd{
        w.WriteHeader(http.StatusForbidden)
        w.Write([]byte("404 wrong username or password!"))
        return
	}
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("200 . Login success! Cheers!!"))
}

func main(){
	routemux := http.NewServeMux()
	routemux.HandleFunc("/login",LoginHandler)
	routemux.HandleFunc("/signin",SigninHandler)
	http.ListenAndServe(":8088", routemux)
}
