package main

import (
	"fmt"
	"time"
)

func main() {
	expiration := time.Now()
	fmt.Println(expiration)
	expiration = expiration.AddDate(1, 0, 0)
	fmt.Println(expiration)
	// cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	// http.SetCookie(w, &cookie)
}
