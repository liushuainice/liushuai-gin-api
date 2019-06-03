package ws

import "github.com/gorilla/websocket"

// OnlineService  在线服务接口  websocket
type OnlineService interface {
	AddUserx(uid string, socket *websocket.Conn)    //新用户上线
	RemoveUserx(uid string)                         //用户下线
	IsUserOnlinex(uid string) bool                  //单个用户是否在线
	IsUsersOnlinex(uids []string) map[string]bool   //多个用户是否在线
	Pushx(uid string, event string, content string) //信息单发
	Broadcastx(event string, content string)        //信息群发
	GetAllUsersOnlinex() map[string]bool            //所有在线用户
}
