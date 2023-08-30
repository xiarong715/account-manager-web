package amwhttp

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var routes = map[string]func(c *gin.Context){
	// "init_key":             initKey,
	// "set_key":              setKey,
	"init_apptoken":                   initAPPToken,
	"add_apptoken":                    addAPPToken,
	"modify_apptoken":                 modifyAPPToken,
	"del_apptoken":                    delAPPToken,
	"get_appname":                     getAppName,
	"gen_token":                       genToken,
	"init_user_permission":            initUserPermission,
	"init_user_list":                  initUserList,
	"init_not_configured_permissions": initNotConfiguredPermissions,
	"add_user_permission":             addUserPermission,
}

type reqData struct {
	Action string `json:"action"`
}

func ApiFunc(c *gin.Context) {
	var rd reqData
	err := c.ShouldBindBodyWith(&rd, binding.JSON)
	if err != nil || rd.Action == "" {
		log.Println("should bind: ", err)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "parsing data err or action is empty",
			"data":   map[string]string{},
		})
		return
	}
	log.Println("action: ", rd.Action)

	routes[rd.Action](c)
}
