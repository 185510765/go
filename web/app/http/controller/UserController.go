package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	common "web/app"

	_ "github.com/go-sql-driver/mysql"
)

func Init() {}

// 首页
func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	// 注意:如果没有调用 ParseForm 方法，下面无法获取表单的数据
	// fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
}

// 登录
func Login(w http.ResponseWriter, r *http.Request) {
	// 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("resources/views/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")
		code, msg := loginValidate(username, password)
		if code != 200 {
			fmt.Println(msg)
		}

		// // 打印信息 ******************************************
		// fmt.Println(r.Form["username"])
		// fmt.Println(r.Form["password"])
		// fmt.Println(r.Form)
		// fmt.Println(r.Form.Get("username"))
		// fmt.Println(r.Form.Get("password"))

		// // 转义
		// fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 输出到服务器端
		// fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		// template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端

		// // request.Form 是一个 url.Values 类型，里面存储的是对应的类似 key=value 的信息
		// v := url.Values{}
		// v.Set("name", "zhangsan")
		// v.Set("friend", "lisi")
		// v.Set("friend", "wangwu")
		// v.Set("friend", "zhaoliu")

		// fmt.Println(v.Get("name"))
		// fmt.Println(v.Get("frien"))
		// fmt.Println(v["friend"])
	}
}

// 登录验证
func loginValidate(username string, password string) (int, string) {
	if common.StrLen(username) < 5 {
		return 500, "用户名不能小于5位数"
	}
	if common.StrLen(password) < 6 {
		return 500, "密码不能小于6位数"
	}

	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	common.CheckErr(err)
	stmt, err := db.Prepare("SELECT * FROM user WHERE username = ?")
	rows, err := stmt.Exec("张三")
	common.CheckErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	return 200, "success"
}
