package ff_mongo_exceptions

type LibException interface {
	GetCode() int
	GetMessages() []string
	Error() string
}
