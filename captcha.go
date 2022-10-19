package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// Response format
type CaptchaResponse struct {
	Id   string      `json:"id"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var store = base64Captcha.DefaultMemStore

// MakeGinResponse return gin json response object
func (r *CaptchaResponse) MakeGinResponse() gin.H {
	return gin.H{"code": r.Code, "data": r.Data, "msg": r.Msg, "id": r.Id}
}

func CaptchaHandle(c *gin.Context) {

	res := new(CaptchaResponse)

	driver := &base64Captcha.DriverString{
		Length:          4,
		Height:          30,
		Width:           100,
		ShowLineOptions: 2,
		NoiseCount:      0,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()

	if err != nil {
		res.Code = 404
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.MakeGinResponse())
		return
	}
	res.Code = 200
	res.Id = id
	res.Data = b64s
	res.Msg = "获取成功"
	c.JSON(http.StatusOK, res.MakeGinResponse())
}

func CaptchaVerify(id string, code string) bool {
	return store.Verify(id, code, false)
}
