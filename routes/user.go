package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"playground/db"
	"playground/models"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		conn := db.Connect()

		var us []models.User
		conn.Preload("Roll").Find(&us)

		results := []gin.H{}

		for i := 0; i < len(us); i++ {
			results = append(results, gin.H{
				"id":        us[i].ID,
				"firstName": us[i].FirstName,
				"lastName":  us[i].LastName,
				"rollName":  us[i].Roll.RollName,
			})
		}

		c.JSON(http.StatusOK, results)
	})

	r.GET("/:id", func(c *gin.Context) {
		conn := db.Connect()

		id := c.Param("id")

		var u models.User
		if err := conn.Preload("Roll").Find(&u, "id=?", id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid id",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":        u.ID,
			"firstName": u.FirstName,
			"lastName":  u.LastName,
			"rollName":  u.Roll.RollName,
		})
	})

	r.POST("/", func(c *gin.Context) {
		conn := db.Connect()

		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err.Error())
		}

		var u models.User
		err = json.Unmarshal([]byte(value), &u)
		if err != nil {
			fmt.Println(err.Error())
		}
		conn.Create(&u)

		c.JSON(http.StatusOK, gin.H{
			"createdId": u.ID,
		})
	})

	r.PUT("/:id", func(c *gin.Context) {
		conn := db.Connect()
		id := c.Param("id")

		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err.Error())
		}

		m := make(map[string]interface{}, 10)
		err = json.Unmarshal([]byte(value), &m)
		if err != nil {
			fmt.Println(err.Error())
		}

		var u models.User
		conn.Find(&u, "id=?", id)

		if v := m["firstName"]; v != nil {
			if s, ok := v.(string); ok {
				u.FirstName = s
			}
		}

		if v := m["lastName"]; v != nil {
			if s, ok := v.(string); ok {
				u.LastName = s
			}
		}

		if v := m["rollId"]; v != nil {
			if i, ok := v.(float64); ok {
				u.RollID = int(i)
			}
		}

		conn.Save(&u)

		c.JSON(http.StatusOK, gin.H{
			"createdId": u.ID,
		})
	})
}
