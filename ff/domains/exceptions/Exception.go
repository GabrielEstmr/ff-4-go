package ff_domains_exceptions

type LibException interface {
	GetCode() int
	GetMessages() []string
	Error() string
}
