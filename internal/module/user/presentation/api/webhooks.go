package api

import (
	"github.com/gin-gonic/gin"
)

func (s *apiServer) RegisterWebhooks() {
	s.gin.POST("/webhooks/fb/new-account", s.SetupFirebaseWebhook)
}

func (s *apiServer) SetupFirebaseWebhook(c *gin.Context) {
	//account := new(entity.User)
	//err := c.BindJSON(account)
	//
	//if err != nil {
	//	s.lg.WithError(err).Error("could not bind json")
	//	c.Status(http.StatusBadRequest)
	//
	//	return
	//}
	//
	//err = s.services.Auth.ProcessAccountChange(context.Background(), account)
	//
	//if err != nil {
	//	s.lg.WithError(err).Error("could not process account change")
	//	c.Status(http.StatusInternalServerError)
	//
	//	return
	//}
	//
	//c.Status(http.StatusOK)
}
