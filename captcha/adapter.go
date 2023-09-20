package captcha

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response Gin response format
type Response struct {
	Id   string      `json:"id"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// MakeGinResponse return gin json response object
func (r *Response) MakeGinResponse() gin.H {
	return gin.H{"code": r.Code, "data": r.Data, "msg": r.Msg, "id": r.Id}
}

func (ca *Captcha) GinHandle(c *gin.Context) {

	res := new(Response)

	id, b64s, err := ca.Generate()

	if err != nil {
		res.Code = 404
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.MakeGinResponse())
		return
	}
	res.Code = 200
	res.Id = id
	res.Data = b64s
	res.Msg = "success"
	c.JSON(http.StatusOK, res.MakeGinResponse())
}
