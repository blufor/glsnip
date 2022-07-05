package sources

type Source interface {
	Content() (string, error)
	Path() string
	Size() int64
	Close() error
}
