package service

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author" gorm:"-"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count" gorm:"-"`
	CommentCount  int64  `json:"comment_count" gorm:"-"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`

	Userid     int64 `json:"-"`
	CreateTime int64 `json:"-"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user" gorm:"-"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`

	Userid int64 `json:"-"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`

	TotalFavorited string `json:"total_favorited,omitempty" gorm:"-"`
	WorkCount      int64  `json:"work_count,omitempty" gorm:"-"`
	FavoriteCount  int64  `json:"favorite_count,omitempty" gorm:"-"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
