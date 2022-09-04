package publish

import "github.com/gin-gonic/gin"

type ActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func Action(c *gin.Context) {

}
