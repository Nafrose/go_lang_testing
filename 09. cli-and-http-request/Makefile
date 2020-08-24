run: 
	# Should run with given arguments

build-main:
	# make build main
	go build main.go

run-main:
	go run main.go -run create-person-cli
	go run main.go -run create-person-json
	go run main.go -get people
	go run main.go -get people/5
	go run main.go -get "people:5,people:1"
	go run main.go -get people/firstname/alim
	go run main.go -get people-plain/firstname
	go run main.go -get people-plain/firstname-and-dateofbirth
	go run main.go -get people-plain/id-and-firstname
	go run main.go -throw "message"

run-main:
	go run main.go

browser-endpoints:
	localhost:8080/people
	localhost:8080/people/5
	localhost:8080/firstName/alim

run-Postman && post:
	localhost:8080/people/create

