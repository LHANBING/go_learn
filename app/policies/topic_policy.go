// Package policies 用户授权
package policies

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/models/topic"
	"go_learn/pkg/auth"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
