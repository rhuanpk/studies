package resource

import (
	"fmt"
	"log"
	"net/http"

	"dev/internal/core/model"
	db "dev/internal/infra/database"

	"github.com/gin-gonic/gin"
)

// Resource handler to resource.
func Resource(ctx *gin.Context) {
	var pg model.Pagination
	if err := ctx.ShouldBindQuery(&pg); err != nil {
		log.Fatalln("error in bind pagination:", err)
	}
	query := `select id from tb `
	if pg.IsUsing() {
		query += fmt.Sprintf(`limit %d offset %d `, pg.PageSize, pg.Offset())
	}
	result, err := db.DB.Query(query)
	defer result.Close()
	if err != nil {
		log.Fatalln("error in select query:", err)
	}
	var rows []db.TB
	for result.Next() {
		var row db.TB
		if err := result.Scan(&row.ID); err != nil {
			log.Fatalln("error in scan results:", err)
		}
		rows = append(rows, row)
	}
	// pg.SetInfos(db.DB, `select count(*) total from tb`, true)
	pg.SetInfos(db.DB, query, false)
	ctx.JSON(http.StatusOK, gin.H{"result": rows, "pagination": pg})
}
