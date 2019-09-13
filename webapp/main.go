package webapp

import (
    "net/http"
    // dbcheck "github.com/dimiro1/health/db"
    "github.com/dimiro1/health"
)

func Run() {
    generateLogging()
    createDatabaseTable()
    // redisIndex()
    // redisUserShow()
    // redisEditUser()
    // redisInsertUser()
    // redisUpdateUser()
    // redisDeleteUser()
    db := dbConn()
    mysql := dbcheck.NewMySQLChecker(db)
    handler := health.NewHandler()
    handler.AddChecker("MySQL", mysql)
    http.Handle("/health", handler)
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    // handler.AddChecker("MySQL", mysql)
    // http.Handle("/health", handler)
    // http.HandleFunc("/", redisIndex)
    // http.HandleFunc("/show", redisUserShow)
    // http.HandleFunc("/new", New)
    // http.HandleFunc("/edit", redisEditUser)
    // http.HandleFunc("/insert", redisInsertUser)
    // http.HandleFunc("/update", redisUpdateUser)
    // http.HandleFunc("/delete", redisDeleteUser)
    http.ListenAndServe(":8080", nil)
}
