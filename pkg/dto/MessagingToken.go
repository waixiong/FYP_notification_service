package dto

type UserToken struct {
	UserId string            `json:"userId" bson:"userId"`
	Tokens []*MessagingToken `json:"tokens" bson:"tokens"`
}

type MessagingToken struct {
	Id       string `json:"id" bson:"id"`
	Type     string `json:"type" bson:"type"`
	Token    string `json:"token" bson:"token"`
	Platform string `json:"platform" bson:"platform"`
	Domain   string `json:"domain" bson:"domain"`
}
