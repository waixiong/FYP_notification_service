package model

import (

	//pb "./proto"

	// "github.com/syndtr/goleveldb/leveldb"

	"net/http"

	firebase "firebase.google.com/go"
	"getitqec.com/server/mailnotification/pkg/commons"
	"getitqec.com/server/mailnotification/pkg/dao"
)

type MailModel struct {
	// UserDAO dao.IUserDAO
	Client *http.Client
}

type NotificationModel struct {
	// Client *http.Client
	TokenDAO dao.IMessagingTokenDAO
	App      *firebase.App
}

// InitModel ...
func InitMailModel(m commons.MongoDB, client *http.Client) MailModelI {
	// dao := &dao.UserDAO{}
	// _dao := dao.InitUserDAO(m)
	return &MailModel{Client: client}
}

func InitNotificationModel(m commons.MongoDB, app *firebase.App) NotificationModelI {
	// dao := &dao.UserDAO{}
	_notificationDao := dao.InitMessagingTokenDAO(m)
	return &NotificationModel{TokenDAO: _notificationDao, App: app}
}
