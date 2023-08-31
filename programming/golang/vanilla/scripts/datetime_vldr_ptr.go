package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const layout = "2006-01-02 15:04:05"

// Object object for binding data payload.
type Object struct {
	DateTime *any `json:"date_time" binding:"required,datetime=2006-01-02 15:04:05"`
}

func main() {
	router := gin.New()

	router.POST("/", func(ctx *gin.Context) {
		var obj *Object

		if err := ctx.ShouldBindJSON(&obj); err != nil {
			log.Println("error on struct binding:", err)
			return
		}
		fmt.Printf("after binding:\n%#v\n%#v\n", *obj, *obj.DateTime)

		datetime, err := time.Parse(layout, any(*obj.DateTime).(string))
		if err != nil {
			log.Println("error on time parsing:", err)
			return
		}
		// before save in database:
		*obj.DateTime = datetime
		fmt.Printf("\nbefore save in database:\n%#v\n%#v\n", *obj, *obj.DateTime)

		// to return to client:
		*obj.DateTime = any(*obj.DateTime).(time.Time).Format(layout)
		fmt.Printf("\nto return to client:\n%#v\n%#v\n", *obj, *obj.DateTime)

		ctx.JSON(http.StatusOK, &obj)
	})

	log.Fatalln(router.Run(":1111"))
}
