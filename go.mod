module example.com

go 1.17

require github.com/gorilla/mux v1.8.0

replace example.com/mocks => ../mocks

replace example.com/pkg/mocks => ../pkg/mocks
