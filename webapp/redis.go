package webapp

import (
    log "github.com/sirupsen/logrus"
    "fmt"
	"strconv"
	"time"
    "gopkg.in/ini.v1"
	"os"
    "net/http"
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

func redisIndex(w http.ResponseWriter, r *http.Request) {
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
	tmpl.ExecuteTemplate(w, "Index", res)
}

func redisUserShow(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn :=  pool.Get()
	nId := r.URL.Query().Get("id")
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

	tmpl.ExecuteTemplate(w, "Show", emp)
}

func redisEditUser(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn :=  pool.Get()
	nId := r.URL.Query().Get("id")
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

	tmpl.ExecuteTemplate(w, "Edit", emp)
}

func redisInsertUser(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn :=  pool.Get()
	if r.Method == "POST" {
		nId := r.URL.Query().Get("id")
		name := r.FormValue("name")
		city := r.FormValue("city")
		email := r.FormValue("email")
		date := r.FormValue("date")

		fmt.Println(nId)
		insForm, err := conn.Do("HMSET", nId, "name", name, "email", email, "date", date, "city", city)
		if err != nil {
			log.Error(err)
		}
		fmt.Print(insForm)
	}
	http.Redirect(w, r, "/", 301)
}

func redisUpdateUser(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn :=  pool.Get()
	if r.Method == "POST" {
		nId := r.URL.Query().Get("id")
		name := r.FormValue("name")
		city := r.FormValue("city")
		email := r.FormValue("email")
		date := r.FormValue("date")
		insForm, err := conn.Do("HMSET", nId, "name", name, "email", email, "date", date, "city", city)
		if err != nil {
			log.Error(err)
		}
		fmt.Print(insForm)
	}
	http.Redirect(w, r, "/", 301)
}

func redisDeleteUser(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn :=  pool.Get()
	nId := r.URL.Query().Get("id")
	insForm, err := conn.Do("DEL", nId)
	if err != nil {
		log.Error(err)
	}
	fmt.Print(insForm)
	fmt.Println(insForm)
	http.Redirect(w, r, "/", 301)
}
