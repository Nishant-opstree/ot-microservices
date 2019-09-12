package webapp

import (
    log "github.com/sirupsen/logrus"
    "fmt"
	// "io"
	"time"
    "gopkg.in/ini.v1"
	"os"
	"github.com/gomodule/redigo/redis"	
)

var pool *redis.Pool

func initializeCache() *redis.Pool {
    var redisHost string
	var redisPort string
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
			return redis.Dial("tcp", redisHost + ":" + redisPort)
		},
	}
}

func redisIndex() {
	pool = initializeCache()
	conn :=  pool.Get()
    emp := Employee{}
	res := []Employee{}
	// id := r.FormValue("uid")
	keys_list, err := redis.Int(conn.Do("KEYS", "*"))
	if err != nil {
		log.Error(err)
	}

	for _, key := range keys_list {
		name, err := redis.String(conn.Do("HGET", key, "name"))
		if err != nil {
			log.Error(err)
		}
		email, err := redis.String(conn.Do("HGET", key, "email"))
		if err != nil {
			log.Error(err)
		}
		date, err := redis.String(conn.Do("HGET", key, "date"))
		if err != nil {
			log.Error(err)
		}
		city, err := redis.String(conn.Do("HGET", key, "city"))
		if err != nil {
			log.Error(err)
		}
        emp.Id = key
        emp.Name = name
        emp.Email = email
        emp.Date = date
		emp.City = city
		res = append(res, emp)
		// for key, value := range reply {
		// 	log.Info("Data found in redis key is " + key + " and value is " + value)
		// }
	}
	fmt.Println(res)
}
