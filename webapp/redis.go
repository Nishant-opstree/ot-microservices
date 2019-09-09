package webapp

import (
    log "github.com/sirupsen/logrus"
    // "fmt"
	// "io"
	"time"
    "gopkg.in/ini.v1"
	"os"
	"github.com/gomodule/redigo/redis"	
)

func initializeCache() {
    var redisHost string
	var redisPort string
	var pool *redis.Pool
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

	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisHost + ":" + "6379")
			if err != nil {
				log.Error("ERROR: fail initializing the redis pool: %s", err.Error())
			}
			return conn, err
		},
	}
}
