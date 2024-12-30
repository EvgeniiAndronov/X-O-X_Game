package models

// TODO: Model data to send

// TODO: Model data on server
//
//type General struct {
//	RoomId string `json:"room_id"`
//	Data OnServer
//}

type OnServer struct {
	UserXId string       `json:"user_x_id"`
	UserOId string       `json:"user_y_id"`
	Map     [3][3]string `json:"map"`
}

type SendData struct {
	UserId         string  `json:"user_id"`
	CoordinateMove float32 `json:"coordinate_move"`
	Token          string  `json:"token"`
}

type SaveUserData struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	Nickname     string `json:"nickname"`
	PasswordHash string `json:"password_hash"`
	Token        string `json:"token"`
}
