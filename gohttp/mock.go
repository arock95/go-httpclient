package gohttp

var (
	mocks map[string]*Mock
)

type Mock struct {
	Method             string
	Url                string
	RequestBody        string
	
	ResponseBody       string
	Error              error
	ResponseStatusCode int
}

func AddMock(mock Mock) {
	key := mock.Method + mock.Url + mock.RequestBody
	mocks[key] = &mock
}
