package model

type User struct {
	Id       int      `json:"-"`
	UserId   int      `json:"id" gorm:"uniqueIndex:idx_id_name,priority:1"`
	Name     string   `json:"name" gorm:"uniqueIndex:idx_id_name,priority:2"`
	Intro    string   `json:"intro"`
	Created  string   `json:"created" gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00';index"`
	State    int      `json:"state" gorm:"int2;default:0"`
	IconUrl  string   `json:"iconUrl" gorm:"size:255;default:''"`
	CoverUrl string   `json:"coverUrl" gorm:"size:255;default:''"`
	Badges   []*Badge `json:"badges" gorm:"-"`
}

func (receiver *User) TableName() string {
	return UserTableName
}

type Badge struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id" gorm:"index"`
	BadgeId int    `json:"badge_id" gorm:"index"`
	Created string `json:"created" gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00';index"`
	Title   string `json:"title" gorm:"size:255;default:''"`
	IconUrl string `json:"iconUrl" gorm:"size:255;default:''"`
}

func (receiver *Badge) TableName() string {
	return BadgeTableName
}

type Diary struct {
	Id              int       `json:"id"`
	UserId          int       `json:"user_id" gorm:"index"`
	NoteBookId      int       `json:"notebook_id" gorm:"index"`
	NoteBookSubject string    `json:"notebook_subject" gorm:"index"`
	Content         string    `json:"content" gorm:"type:text"`
	Created         string    `json:"created" gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00';index"`
	Updated         string    `json:"updated" gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00'"`
	Type            int       `json:"type" gorm:"int2;default:0"`
	CommentCount    int       `json:"comment_count" gorm:"default:0"`
	PhotoUrl        string    `json:"photoUrl" gorm:"size:255;default:''"`
	PhotoThumbUrl   string    `json:"photoThumbUrl" gorm:"-"`
	LikeCount       int       `json:"like_count" gorm:"default:0"`
	Liked           bool      `json:"-" gorm:"-"`
	User            *User     `json:"user,omitempty" gorm:"-"`
	NoteBook        *NoteBook `json:"notebook,omitempty" gorm:"-"`
}

func (receiver *Diary) TableName() string {
	return DiaryTableName
}

type IndexDiary struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id" gorm:"index"`
	NoteBookId      int    `json:"notebook_id" gorm:"index"`
	NoteBookSubject string `json:"notebook_subject" gorm:"index"`
	Content         string `json:"content" gorm:"type:text"`
	Created         string `json:"created" gorm:"timestamptz(6);default:'0001-01-01 00:00:00';index"`
}

func (diary *Diary) IndexDiary() *IndexDiary {

	return &IndexDiary{
		Id:              diary.Id,
		UserId:          diary.UserId,
		NoteBookId:      diary.NoteBookId,
		NoteBookSubject: diary.NoteBookSubject,
		Content:         diary.Content,
		Created:         diary.Created,
	}
}

type TinyDiary struct {
	Id       int
	UserId   int
	PhotoUrl string
	Created  string
}

type NoteBook struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id" gorm:"index"`
	Subject     string `json:"subject" gorm:"size:255;index"`
	Description string `json:"description" gorm:"index"`
	Created     string `json:"created" gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00';index"`
	Updated     string `json:"updated" gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00'"`
	Expired     string `json:"expired" gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00'"`
	Privacy     int    `json:"privacy" gorm:"int2;default:0"`
	Cover       int    `json:"cover" gorm:"int2;default:0"`
	CoverUrl    string `json:"coverUrl" gorm:"size:255;default:''"`
	IsPublic    bool   `json:"isPublic" gorm:"-"`
}

func (receiver *NoteBook) TableName() string {
	return NoteBookTableName
}

type Comment struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id" gorm:"index"`
	RecipientId int    `json:"recipient_id" gorm:"index"`
	DairyId     int    `json:"dairy_id" gorm:"index"`
	Content     string `json:"content" gorm:"type:text"`
	Created     string `json:"created" gorm:"type:timestamptz(6);default:'0001-01-01 00:00:00';index"`
	User        *User  `json:"User" gorm:"-"`
	Recipient   *User  `json:"recipient" gorm:"-"`
}

func (receiver *Comment) TableName() string {
	return CommentTableName
}

type CoverType int

const (
	BookCoverType CoverType = iota
	UserCoverType
)

func (c CoverType) String() string {
	if c == BookCoverType {
		return "/book_cover/"
	}
	return "/user_icon/"
}

type Face struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id" gorm:"index"`
	DairyId int    `json:"dairy_id" gorm:"index"`
	Created string `json:"created" gorm:"type:timestamptz(6);default:now();index"`
}

func (receiver *Face) TableName() string {
	return FaceTableName
}
