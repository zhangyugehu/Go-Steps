package mock

type Retriever struct {
	Constants string
}

func (r Retriever) Get(url string) string {
	return r.Constants
}

