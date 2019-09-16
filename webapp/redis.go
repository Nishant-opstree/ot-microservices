package webapp

import (
    log "github.com/sirupsen/logrus"
    // "fmt"
	// "strconv"
	"time"
    "gopkg.in/ini.v1"
	"os"
	// "json"
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

func signUp(w http.ResponseWriter, r *http.Request) {
	pool = initializeCache()
	conn := pool.Get()
	if r.Method == "POST" {
		nId := r.FormValue("id")
		name := r.FormValue("name")
		password := r.FormValue("password")

		_, err := conn.Do("HMSET", nId, "name", name, "password", password)

		if err != nil {
			log.Error(err)
		}
	}
}

// func Signin(w http.ResponseWriter, r *http.Request) {
// 	var creds Credentials
// 	// Get the JSON body and decode into credentials
// 	err := json.NewDecoder(r.Body).Decode(&creds)
// 	if err != nil {
// 		// If the structure of the body is wrong, return an HTTP error
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Get the expected password from our in memory map
// 	expectedPassword, ok := users[creds.Username]

// 	// If a password exists for the given user
// 	// AND, if it is the same as the password we received, the we can move ahead
// 	// if NOT, then we return an "Unauthorized" status
// 	if !ok || expectedPassword != creds.Password {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	// Create a new random session token
// 	sessionToken := uuid.NewV4().String()
// 	// Set the token in the cache, along with the user whom it represents
// 	// The token has an expiry time of 120 seconds
// 	_, err = cache.Do("SETEX", sessionToken, "120", creds.Username)
// 	if err != nil {
// 		// If there is an error in setting the cache, return an internal server error
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// Finally, we set the client cookie for "session_token" as the session token we just generated
// 	// we also set an expiry time of 120 seconds, the same as the cache
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "session_token",
// 		Value:   sessionToken,
// 		Expires: time.Now().Add(120 * time.Second),
// 	})
// }

// func redisIndex(w http.ResponseWriter, r *http.Request) {
// 	pool = initializeCache()
// 	conn :=  pool.Get()
//     emp := Employee{}
// 	res := []Employee{}
// 	keys_list, err := redis.Strings(conn.Do("KEYS", "*"))
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	for _, key := range keys_list {
// 		name, err := redis.String(conn.Do("HGET", key, "name"))
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		email, err := redis.String(conn.Do("HGET", key, "email"))
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		date, err := redis.String(conn.Do("HGET", key, "date"))
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		city, err := redis.String(conn.Do("HGET", key, "city"))
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		id, err := strconv.Atoi(key)
// 		if err != nil {
// 			log.Error(err)
// 		}
//         emp.Id = id
//         emp.Name = name
//         emp.Email = email
//         emp.Date = date
// 		emp.City = city
// 		res = append(res, emp)
// 	}
// 	tmpl.ExecuteTemplate(w, "Index", res)
// }

// func redisUserShow(w http.ResponseWriter, r *http.Request) {
// 	pool = initializeCache()
// 	conn :=  pool.Get()
// 	nId := r.FormValue("id")
// 	emp := Employee{}
// 	name, err := redis.String(conn.Do("HGET", nId, "name"))
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	email, err := redis.String(conn.Do("HGET", nId, "email"))
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	date, err := redis.String(conn.Do("HGET", nId, "date"))
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	city, err := redis.String(conn.Do("HGET", nId, "city"))
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	id, err := strconv.Atoi(nId)
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	emp.Id = id
// 	emp.Name = name
// 	emp.Email = email
// 	emp.Date = date
// 	emp.City = city

// 	tmpl.ExecuteTemplate(w, "Show", emp)
// }

// func redisEditUser(w http.ResponseWriter, r *http.Request) {
// 	pool = initializeCache()
// 	conn :=  pool.Get()
// 	nId := r.FormValue("id")
// 	emp := Employee{}
// 	name, err := redis.String(conn.Do("HGET", nId, "name"))
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	email, err := redis.String(conn.Do("HGET", nId, "email"))
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	date, err := redis.String(conn.Do("HGET", nId, "date"))
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	city, err := redis.String(conn.Do("HGET", nId, "city"))
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	id, err := strconv.Atoi(nId)
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	emp.Id = id
// 	emp.Name = name
// 	emp.Email = email
// 	emp.Date = date
// 	emp.City = city

// 	tmpl.ExecuteTemplate(w, "Edit", emp)
// }

// func redisInsertUser(w http.ResponseWriter, r *http.Request) {
// 	pool = initializeCache()
// 	conn :=  pool.Get()
// 	if r.Method == "POST" {
// 		nId := r.FormValue("id")
// 		name := r.FormValue("name")
// 		city := r.FormValue("city")
// 		email := r.FormValue("email")
// 		date := r.FormValue("date")

// 		fmt.Println(nId)
// 		insForm, err := conn.Do("HMSET", nId, "name", name, "email", email, "date", date, "city", city)
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		fmt.Print(insForm)
// 	}
// 	http.Redirect(w, r, "/", 301)
// }

// func redisUpdateUser(w http.ResponseWriter, r *http.Request) {
// 	pool = initializeCache()
// 	conn :=  pool.Get()
// 	if r.Method == "POST" {
// 		nId := r.FormValue("id")
// 		name := r.FormValue("name")
// 		city := r.FormValue("city")
// 		email := r.FormValue("email")
// 		date := r.FormValue("date")
// 		insForm, err := conn.Do("HMSET", nId, "name", name, "email", email, "date", date, "city", city)
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		fmt.Print(insForm)
// 	}
// 	http.Redirect(w, r, "/", 301)
// }

// func redisDeleteUser(w http.ResponseWriter, r *http.Request) {
// 	pool = initializeCache()
// 	conn :=  pool.Get()
// 	nId := r.FormValue("id")
// 	insForm, err := conn.Do("DEL", nId)
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	fmt.Print(insForm)
// 	fmt.Println(insForm)
// 	http.Redirect(w, r, "/", 301)
// }
