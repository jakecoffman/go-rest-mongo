package framework

type Repository interface {
	New() Resource

	List(query map[string]interface{}, limit int, sort ...string) (interface{}, error)
	Get(id string) (interface{}, error)
	Insert(data interface{}) (interface{}, error)
	Update(id string, data interface{}) (interface{}, error)
	Delete(id string) error
}
