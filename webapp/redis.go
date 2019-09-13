package webapp

import (
    log "github.com/sirupsen/logrus"
    "fmt"
	// "io"
	"strconv"
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
	keys_list, err := redis.Strings(conn.Do("KEYS", "*"))
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
		id, err := strconv.Atoi(key)
		if err != nil {
			log.Error(err)
		}
        emp.Id = id
        emp.Name = name
        emp.Email = email
        emp.Date = date
		emp.City = city
		res = append(res, emp)
	}
	fmt.Println(res)
}

func redisUserShow() {
	pool = initializeCache()
	conn :=  pool.Get()
	// nId := r.URL.Query().Get("id")
	nId := "1"
	emp := Employee{}
	name, err := redis.String(conn.Do("HGET", nId, "name"))
	if err != nil {
		log.Error(err)
	}
	email, err := redis.String(conn.Do("HGET", nId, "email"))
	if err != nil {
		log.Error(err)
	}
	date, err := redis.String(conn.Do("HGET", nId, "date"))
	if err != nil {
		log.Error(err)
	}
	city, err := redis.String(conn.Do("HGET", nId, "city"))
	if err != nil {
		log.Error(err)
	}
	id, err := strconv.Atoi(nId)
	if err != nil {
		log.Error(err)
	}
	emp.Id = id
	emp.Name = name
	emp.Email = email
	emp.Date = date
	emp.City = city

	fmt.Println(emp)
}

func redisEditUser() {
	pool = initializeCache()
	conn :=  pool.Get()
	// nId := r.URL.Query().Get("id")
	nId := "1"
	emp := Employee{}
	name, err := redis.String(conn.Do("HGET", nId, "name"))
	if err != nil {
		log.Error(err)
	}
	email, err := redis.String(conn.Do("HGET", nId, "email"))
	if err != nil {
		log.Error(err)
	}
	date, err := redis.String(conn.Do("HGET", nId, "date"))
	if err != nil {
		log.Error(err)
	}
	city, err := redis.String(conn.Do("HGET", nId, "city"))
	if err != nil {
		log.Error(err)
	}
	id, err := strconv.Atoi(nId)
	if err != nil {
		log.Error(err)
	}
	emp.Id = id
	emp.Name = name
	emp.Email = email
	emp.Date = date
	emp.City = city

	fmt.Println(emp)
}

func redisInsertUser() {
	pool = initializeCache()
	conn :=  pool.Get()
	// nId := r.URL.Query().Get("id")
	// name := r.FormValue("name")
	// city := r.FormValue("city")
	// email := r.FormValue("email")
	// date := r.FormValue("date")
	nId := "3"
	name := "Abhishek Dubey"
	city := "Delhi"
	email := "abhishek.dubey@opstree.com"
	date := "2019-08-01"

	insForm, err := conn.Do("HMSET", nId, "name", name, "email", email, "date", date, "city", city)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(insForm)
}

func redisUpdateUser() {
	pool = initializeCache()
	conn :=  pool.Get()
	// nId := r.URL.Query().Get("id")
	// name := r.FormValue("name")
	// city := r.FormValue("city")
	// email := r.FormValue("email")
	// date := r.FormValue("date")
	nId := "3"
	name := "Anuj Lohani"
	city := "Delhi"
	email := "abhishek.dubey@opstree.com"
	date := "2019-08-01"
	insForm, err := conn.Do("HMSET", nId, "name", name, "email", email, "date", date, "city", city)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(insForm)
}

func redisDeleteUser() {
	pool = initializeCache()
	conn :=  pool.Get()
	nId := "3"
	insForm, err := conn.Do("DEL", nId)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(insForm)
}
