package domain

// DomainModel層(実クラス)
// DDDではEntityと呼ばれる
// 基本的にクラス(構造体を定義し、データと簡単な振る舞いのみを持たせる)

// モデルはドメインの知識を持ち、不変条件を満たすと良い
// (モデル設計ドメインの知識を持つということ http://little-hands.hatenablog.com/entry/2017/10/04/201201 )
type User struct {
	Id   uint `gorm:"primary_key"`
	Name string
	Age  int
}
