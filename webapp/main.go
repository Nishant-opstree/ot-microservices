package webapp

import (
    "net/http"
    log "github.com/sirupsen/logrus"
    "os"
    dbcheck "github.com/dimiro1/health/db"
    "github.com/dimiro1/health"
    "github.com/dimiro1/health/redis"
    "gopkg.in/ini.v1"
)

func Run() {

	propertyfile := "/etc/conf.d/ot-go-webapp/application.ini"
    if fileExists(propertyfile) {
        loggingInit()
        vaules, err := ini.Load(propertyfile)
        if err != nil {
            log.Error("No property file found in " + propertyfile)
        }
        redisHost = vaules.Section("redis").Key("REDIS_HOST").String()
        redisPort = vaules.Section("redis").Key("REDIS_PORT").String()
        log.Info("Reading properties file " + propertyfile)
    } else {
        redisHost = os.Getenv("REDIS_HOST")
        redisPort = os.Getenv("REDIS_PORT")
        loggingInit()
        log.Info("No property file found, using environment variables")
    }

    generateLogging()
    createDatabaseTable()
    db := dbConn()
    mysql := dbcheck.NewMySQLChecker(db)
    handler := health.NewHandler()
    handler.AddChecker("MySQL", mysql)
    handler.AddChecker("Redis", redis.NewChecker("tcp", redisHost + ":" + redisPort))
    http.Handle("/health", handler)
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}
