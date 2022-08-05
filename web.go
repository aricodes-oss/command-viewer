package main

import (
	"html/template"
	"net/http"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gin-gonic/gin"
)

func ListenAndServe(validBots *mapset.Set[string]) {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"StringsJoin": strings.Join,
	})
	r.LoadHTMLGlob("templates/*")

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		if !(*validBots).Contains(name) {
			c.String(http.StatusNotFound, "%v is not a valid bot name", name)
			return
		}

		commands, err := LoadDb(name)
		if err != nil {
			c.String(http.StatusBadRequest, "Error loading commands for %v", name)
		}

		c.HTML(http.StatusOK, "list.tmpl", Bot{BotName: name, Commands: commands})
	})

	r.Run(":3000")
}
