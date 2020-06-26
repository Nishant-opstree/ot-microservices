package main

import (
	"ot-go-webapp/config"
	"ot-go-webapp/elastic"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"encoding/json"
	"net/http"
	"strings"
)

// EmployeeInfo struct will be the data structure for employee's information
type EmployeeInfo struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	JoiningDate   string  `json:"joining_date"`
	Addresss      string  `json:"address"`
	EmailID       string  `json:"email_id"`
	AnnualPackage float64 `json:"annual_package"`
	PhoneNumber   string  `json:"phone_number"`
}

func main() {
	logrus.Infof("Running employee-management in webserver mode")
	logrus.Infof("employee-management is listening on port: 8080")
	logrus.Infof("Endpoint is available now - http://0.0.0.0:8080/create")
	router := gin.Default()
	router.POST("/create", pushEmployeeData)
	router.GET("/search", fetchEmployeeData)
	router.GET("/search/all", fetchALLEmployeeData)
	router.Run(":8080")
}

func pushEmployeeData(c *gin.Context) {
	var request EmployeeInfo
	if err := c.BindJSON(&request); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		logrus.Errorf("Error parsing the request body in JSON: %v", err)
		return
	}

	info := EmployeeInfo{
		ID:            request.ID,
		Name:          request.Name,
		JoiningDate:   request.JoiningDate,
		Addresss:      request.Addresss,
		EmailID:       request.EmailID,
		AnnualPackage: request.AnnualPackage,
		PhoneNumber:   request.PhoneNumber,
	}
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	elastic.PostDataInSearch(conf, request.ID, info)
	logrus.Infof("Successfully pushed employee's data to elasticsearch")
}

func fetchEmployeeData(c *gin.Context) {
	searchQuery := c.Request.URL.Query()
	var searchValue string
	response := &EmployeeInfo{}

	for _, value := range searchQuery {
		searchValue = strings.Join(value, "")
	}
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	data := elastic.SearchDataInElastic(conf, searchValue)

	for _, parsedData := range data["hits"].(map[string]interface{})["hits"].([]interface{}) {
		empData, err := json.Marshal(parsedData.(map[string]interface{})["_source"])
		if err != nil {
			logrus.Errorf("Unable to marshal response JSON: %v", err)
		}
		json.Unmarshal(empData, &response)
	}
	c.JSON(http.StatusOK, response)
}

func fetchALLEmployeeData(c *gin.Context) {
	conf, err := config.ParseFile("/go/src/ot-go-webapp/config.yaml")
	if err != nil {
		logrus.Errorf("Unable to parse configuration file for management: %v", err)
	}
	data := elastic.SearchALLDataInElastic(conf)

	var employeeInfo []EmployeeInfo
	for _, parsedData := range data["hits"].(map[string]interface{})["hits"].([]interface{}) {
		response := &EmployeeInfo{}
		empData, err := json.Marshal(parsedData.(map[string]interface{})["_source"])
		if err != nil {
			logrus.Errorf("Unable to marshal response JSON: %v", err)
		}
		json.Unmarshal(empData, &response)
		employeeInfo = append(employeeInfo, *response)
	}
	c.JSON(http.StatusOK, employeeInfo)
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}