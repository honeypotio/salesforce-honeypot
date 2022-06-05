package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/services/oauth2/token", postToken)
	router.GET("/services/data/v53.0/query", getQuery)

	fmt.Println("Running on port 4000")
	router.Run(":4000")
}

func postToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"access_token": "fake_access_token"})
}

func getQuery(c *gin.Context) {
	params := c.Request.URL.Query()
	query := params["q"][0]

	if strings.Contains(query, "opportunity") {
		getOpportunities(c)
	} else if strings.Contains(query, "Account") {
		getAccount(c)
	} else if strings.Contains(query, "User") {
		getOwner(c)
	}
}

func getOpportunities(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"records": []interface{}{
			gin.H{
				"Start_of_Billed_Membership_Date__c": "2022-01-01",
				"Contract_End_Date__c":               "2042-12-31",
				"Number_of_Hires__c":                 99,
				"Active_opportunity__c":              true,
				"Automatically_Renews__c":            true,
				"Last_Day_to_Cancel__c":              "2042-12-01",
				"Bear_Type__c":                       "Big Bear",
				"End_of_Trial_Date__c":               "2022-02-01",
				"Trial_Period_Included__c":           true,
				"Membership_Duration_Months__c":      240,
				"Billing_Cycle__c":                   "Monthly",
				"Amount":                             5000,
				"Total_Value__c":                     100000,
			},
		},
	})
}

func getAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"records": []interface{}{
			gin.H{
				"Account_ID_18__c": 123,
				"OwnerId":          234,
			},
		},
	})
}

func getOwner(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"records": []interface{}{
			gin.H{
				"Name":           "John Doe",
				"MediumPhotoUrl": "https://www.fakepersongenerator.com/Face/female/female20171026166672180.jpg",
				"Username":       "jane.doe@honeypot.io",
				"Phone":          "+49123123123",
				"MobilePhone":    "+49234234234",
			},
		},
	})
}
