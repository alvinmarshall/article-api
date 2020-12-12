package repository

type BaseRepository interface {
	FindOne(filter interface{}) (interface{}, error)
	Find(filter ...interface{}) (interface{}, error)
	Remove(data interface{}) error
	Save(data interface{}) error
	Update(data interface{}) error
}
