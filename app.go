package main

import (
	"fmt"
	"net/http"
	"strconv"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	// "log"
	"time"
	"github.com/rusinadaria/Logging/pkg/service"
)

// func appServerRun() {
// 	router := mux.NewRouter()
//     router.HandleFunc("/users/login/{guid}", handleGetGuidAndIP).Methods("Get")
//     router.HandleFunc("/users/refresh", refreshHandler).Methods("Get")
//     http.Handle("/", router)
// 	http.ListenAndServe(":80", nil)
// }

func handleGetGuidAndIP(w http.ResponseWriter, r *http.Request){
	// vars := mux.Vars(r)
	// guidParams := vars["guid"]
	// response := fmt.Sprintf("User %s", guidParams)
	// fmt.Fprint(w, response)
	guid := getGuid(r)
	fmt.Println(guid)
	ip := getIp(r)
	fmt.Println(ip)

	//вызвать генерацию токенов
	// tokenService.generateTokens(guid, ip)
	accessToken, refreshToken := service.GenerateTokens(ip)
	// _, refreshToken := generateTokens(ip)


	access_cookie := &http.Cookie{Name: "accessToken", Value: accessToken, Expires:  time.Now().Add(24 * time.Hour), HttpOnly: true}
    http.SetCookie(w, access_cookie)
	refresh_cookie := &http.Cookie{Name: "refreshToken", Value: refreshToken, Expires:  time.Now().Add(24 * time.Hour), HttpOnly: true}
    http.SetCookie(w, refresh_cookie)

	// store := sessions.NewCookieStore([]byte("secret-key"))
	// session, _ := store.Get(r, "session-name")

	// // Добавление значения в сессию
	// session.Values["accessToken"] = accessToken, refreshToken

	// // Сохранение изменений
	// err := session.Save(r, w)

	// service.saveToken(refreshToken, guid)
}

func getGuid(r *http.Request) int {
	// vars := mux.Vars(r)
	// guidParams := vars["guid"]
	guidParams := chi.URLParam(r, "guid")
	fmt.Printf("guidParams in getGuid: %s\n", guidParams)
	guid, err := strconv.Atoi(guidParams)
	if err != nil {
		fmt.Println("Ошибка с guid")
	}
	return guid
}

func getIp(r *http.Request) string {
	ip:= r.RemoteAddr
	// response := fmt.Sprintf("User ip    %s", ip)
	// fmt.Fprint(w, response)
	fmt.Printf("ip in getIp: %s\n", ip)
	return ip
}

// type MyCustomClaims struct {
//     jwt.StandardClaims
//     Username string `json:"username"`
//     Email    string `json:"email"`
// }

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "REFRESH ROUT")
	
	cookies := readCookie(r)
	fmt.Printf("Cookies: %s", cookies)

	
	// tokenString := "<YOUR TOKEN STRING>"    

	// claims := jwt.MapClaims{}
	// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("<YOUR VERIFICATION KEY>"), nil
	// })
	// // ... error handling
	
	// // do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

	// validTokens()

}


func readCookie(r *http.Request) []*http.Cookie{
	//получить куки
	cookies := r.Cookies()
	// fmt.Printf("Cookies: %s", cookies)
	return cookies
}