package operation

import "time"

type PublicAttr struct {
	Id          int64
	GmtModified time.Time `gorm:"autoUpdateTime"`
}

type TVideoInfo struct {
	AuthorId      int64
	PublishTime   time.Time
	CoverUrl      string
	PlayUrl       string
	FavoriteCount int64
	CommentCount  int64
	Title         string
	PublicAttr
}

type TUserInfo struct {
	Username      string
	Password      string
	FollowCount   int64
	FollowerCount int64
	PublicAttr
}

type TFollowList struct {
	UserId     int64
	FollowerId int64
	Status     bool
	PublicAttr
}

type TLikedVideo struct {
	UserId  int64
	VideoId int64
	PublicAttr
}

type TComment struct {
	UserId     int64
	VideoId    int64
	Content    string
	GmtCreated time.Time
	PublicAttr
}
