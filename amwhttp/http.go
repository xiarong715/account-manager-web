package amwhttp

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type resData struct {
	Status int8              `json:"status"`
	Msg    string            `json:"msg"`
	Data   map[string]string `json:"data"`
}

func initKey(c *gin.Context) {
	// get data
	pubkey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFMnRzbE91Z2hmQXNXK28rVDQvcDkrVExMdWxMRwpCSjZIY1J4L01tSEt0MkRIZDg3WGpBYmJBdzdXamRTdWZVRjAxMjhlMVZTL01QT3NFRUNNZDhvNlJRPT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	prikey := "LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUxxRExYdVBpSHA3TU1QTU5PNG5jTnR3Z1FnWUNyRVE4bmJhdm9FWm1qczVvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFMnRzbE91Z2hmQXNXK28rVDQvcDkrVExMdWxMR0JKNkhjUngvTW1IS3QyREhkODdYakFiYgpBdzdXamRTdWZVRjAxMjhlMVZTL01QT3NFRUNNZDhvNlJRPT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo="
	res := resData{Status: 0, Msg: "good", Data: map[string]string{"pubkey": pubkey, "prikey": prikey}}
	c.JSON(http.StatusOK, res)
}

func setKey(c *gin.Context) {
	// parsing data
	type keyStruct struct {
		Pubkey string `form:"pubkey" json:"pubkey"`
		Prikey string `form:"prikey" json:"prikey"`
	}
	var ks keyStruct
	err := c.ShouldBindBodyWith(&ks, binding.JSON)
	if err != nil {
		log.Println("set key", "should bind: ", err)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "parsing data err",
			"data":   map[string]string{}, // TODO
		})
		return
	}

	log.Println("public key: ", ks.Pubkey)
	log.Println("private key: ", ks.Prikey)

	// save data

	// reply data
	res := resData{Status: 0, Msg: "success", Data: map[string]string{"pubkey": ks.Pubkey, "prikey": ks.Prikey}}
	c.JSON(http.StatusOK, res)
}

func initAPPToken(c *gin.Context) {
	type item struct {
		Id       int64  `json:"id"`
		APPName  string `json:"appname"`
		APPToken string `json:"apptoken"`
	}
	type data struct {
		Items []item `json:"items"`
		Total int64  `json:"total"`
	}
	type resData struct {
		Status int8   `json:"status"`
		Msg    string `json:"msg"`
		Data   data   `json:"data"`
	}
	datas := data{
		Items: append(make([]item, 0, 3),
			item{Id: 1, APPName: "gogs-user", APPToken: "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiZ29ncy11c2VyIn0.3Q_Al6-gPNdsHPVRLdjWtIgYsXCZC_nCAC2DgruQNVXOkQKxvAQSaIpjynWw0wRoaZBV-oxdJ5ptdd8J8siGHg"},
			item{Id: 2, APPName: "gogs-code", APPToken: "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiZ29ncy1jb2RlIn0.bZUu2VQaaklv9oDHXXN-xUqI6TI2MZ5eHALgR263HvjlwIZYeYivq9sO1c-35dryaUNCLJkZ-jkzbkKW31Ohnw"},
			item{Id: 3, APPName: "gogs-tpm", APPToken: "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiZ29ncy10cG0ifQ.lbynmqWVo_GgPmNWTynnUqakeE_JTke3MERt_oDBlbhx5xnWEgcyUjYChnJzWTsBn4D-EN3_NJ5hsHIrthV6SA"},
		),
		Total: 3,
	}

	res := resData{Status: 0, Msg: "good", Data: datas}
	c.JSON(http.StatusOK, res)
}

func addAPPToken(c *gin.Context) {
	type tokenData struct {
		APPName string `json:"appname"`
	}
	var td tokenData
	err := c.ShouldBindBodyWith(&td, binding.JSON)
	if err != nil {
		log.Println("add app token: ", err)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "parsing data err",
			"data":   map[string]string{},
		})
		return
	}

	log.Println("appname: ", td.APPName)

	// save data
	// xxx
	// auto gen token
	payload, err := json.Marshal(struct {
		U string `json:"u"`
	}{
		U: td.APPName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 1,
			"msg":    "marshal err",
			"data":   map[string]string{},
		})
		return
	}
	log.Println(payload)
	token := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiZ29ncy1sc3kifQ.c_SzngQJNQVJB5cFQNO6Q8z6VVHFdDecQMYM17K7jhrPjDU7KtKfY0OIuDrYuVptxKUOFbCWdfJjujkWQEj_Aw"

	// reply data
	type item struct {
		Id       int64  `json:"id"`
		APPName  string `json:"appname"`
		APPToken string `json:"apptoken"`
	}
	type data struct {
		Items []item `json:"items"`
		Total int64  `json:"total"`
	}
	type resData struct {
		Status int8   `json:"status"`
		Msg    string `json:"msg"`
		Data   data   `json:"data"`
	}

	datas := data{
		Items: append(make([]item, 0),
			item{APPName: td.APPName, APPToken: token},
		),
		Total: 1,
	}

	res := resData{Status: 0, Msg: "success", Data: datas}
	c.JSON(http.StatusOK, res)
}

func modifyAPPToken(c *gin.Context) {
}

func delAPPToken(c *gin.Context) {
	type appName struct {
		AppName string `json:"appname"`
	}
	var an appName
	err := c.ShouldBindBodyWith(&an, binding.JSON)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "parsing data err",
			"data":   map[string]string{},
		})
		return
	}

	log.Println("appname: ", an.AppName)
}

func initUserPermission(c *gin.Context) {
	type item struct {
		UserName   string `json:"username"`
		APPName    string `json:"appname"`
		Permission bool   `json:"switch"`
	}
	type data struct {
		Items []item `json:"items"`
		Total int64  `json:"total"`
	}
	type resData struct {
		Status int8   `json:"status"`
		Msg    string `json:"msg"`
		Data   data   `json:"data"`
	}

	type params struct {
		UserName string `json:"username"`
	}
	var param params
	err := c.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "initUserPermission should bind body with error",
			"data":   map[string]string{},
		})
		return
	}

	log.Println("username: ", param.UserName)

	// 测试数据
	datas := data{
		Items: append(make([]item, 0, 3), item{UserName: param.UserName, APPName: "gogs-user", Permission: true},
			item{UserName: param.UserName, APPName: "gogs-code", Permission: true},
			item{UserName: param.UserName, APPName: "gogs-tpm", Permission: false},
		),
		Total: 3,
	}

	res := resData{Status: 0, Msg: "good", Data: datas}
	c.JSON(http.StatusOK, res)
}

func getAppName(c *gin.Context) {

}

func genToken(c *gin.Context) {
	type payload struct {
		Payload string `json:"payload"`
	}
	var pl payload
	err := c.ShouldBindBodyWith(&pl, binding.JSON)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "parsing data err",
			"data":   map[string]string{},
		})
		return
	}

	log.Println("payload: ", pl.Payload)

	// gen token
	token := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiZ29ncy1jb2RlIn0.bZUu2VQaaklv9oDHXXN-xUqI6TI2MZ5eHALgR263HvjlwIZYeYivq9sO1c-35dryaUNCLJkZ-jkzbkKW31Ohnw"

	// reply data
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "success",
		"data":   map[string]string{"token": token},
	})
}

func initUserList(c *gin.Context) {
	type item struct {
		UserName    string `json:"username"`
		Permissions string `json:"permissions"`
	}
	type data struct {
		Items []item `json:"items"`
		Total int64  `json:"total"`
	}
	type resData struct {
		Status int8   `json:"status"`
		Msg    string `json:"msg"`
		Data   data   `json:"data"`
	}

	// 测试数据
	datas := data{
		Items: append(make([]item, 0, 3), item{UserName: "public", Permissions: strings.Join([]string{"gogs-user", "gogs-code", "gogs-tpm"}, ", ")},
			item{UserName: "code", Permissions: strings.Join([]string{"gogs-user", "gogs-code"}, ", ")},
			item{UserName: "user", Permissions: strings.Join([]string{"gogs-user"}, ", ")},
		),
		Total: 3,
	}

	res := resData{Status: 0, Msg: "good", Data: datas}
	c.JSON(http.StatusOK, res)
}

func initNotConfiguredPermissions(c *gin.Context) {
	type item struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}
	type data struct {
		Items []item `json:"items"`
		Total int64  `json:"total"`
	}
	type resData struct {
		Status int8   `json:"status"`
		Msg    string `json:"msg"`
		Data   data   `json:"data"`
	}

	type userStruct struct {
		UserName string `json:"username"`
	}
	var user userStruct
	err := c.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "parsing data err",
			"data":   map[string]string{},
		})
		return
	}

	log.Println("username: ", user.UserName)

	// 测试数据
	datas := data{
		Items: append(make([]item, 0, 3), item{Label: "gogs-tpm", Value: "gogs-tpm"},
			item{Label: "gogs-code", Value: "gogs-code"},
			item{Label: "gogs-lsy", Value: "gogs-lsy"},
		),
		Total: 3,
	}

	res := resData{Status: 0, Msg: "good", Data: datas}
	c.JSON(http.StatusOK, res)
}
func addUserPermission(c *gin.Context) {
	type item struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}
	type data struct {
		Items []item `json:"items"`
		Total int64  `json:"total"`
	}
	type resData struct {
		Status int8   `json:"status"`
		Msg    string `json:"msg"`
		Data   data   `json:"data"`
	}

	type userPermissionStruct struct {
		UserName string `json:"username"`
		AppNames string `json:"appnames"`
	}
	var user userPermissionStruct
	err := c.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "parsing data err",
			"data":   map[string]string{},
		})
		return
	}

	log.Println("username: ", user.UserName)
	log.Println("appnames: ", user.AppNames)

	// 测试数据
	datas := data{
		Items: append(make([]item, 0, 3), item{Label: "gogs-tpm", Value: "gogs-tpm"},
			item{Label: "gogs-code", Value: "gogs-code"},
			item{Label: "gogs-lsy", Value: "gogs-lsy"},
		),
		Total: 3,
	}

	res := resData{Status: 0, Msg: "good", Data: datas}
	c.JSON(http.StatusOK, res)
}
