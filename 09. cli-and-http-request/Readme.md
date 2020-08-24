### Use golang httprouter library
Create HTTP endpoints which can serve REST Content:
Create data-struct as person.sample.json
create 5 person using the same struct. Use cli command to take input from user and then add it to the JSON list (Hint use slice peopleSlice := make([]PersonType, length: 0, capacity: 10) and then peopleSlice = append(peopleSlice, newPerson) eg. Slice Example)

### CLI options:
  cli.go -run create-person-cli: Will prompt to the user asking for first name, lastname and so on from the json and then app will take that input and create a struct and push it to json list and the finally writes all json back to file people.list.json day folder\data\people.list.json
  cli.go -run create-person-json: take a json input from cli and does the same as above. [Use mutex while writing to file.]
  cli.go -get people: will result all the person in the json file in json format in the console as json.
  cli.go -get people/5: will return a single person by search through id:5. 
  cli.go -get "people:5,people:1": will return a single person by search through id:5 and id:1.
  cli.go -get people/firstname/alim: will return a single person by search through firstName:alim do a contains search.
  cli.go -get people-plain/firstname: will all the first name in the system with new line separator. Note: We don't need to create dynamic search for other fields such as lastname (not JSON).
  cli.go -get people-plain/firstname-and-dateofbirth: will all the first name and date of birth in the system with 2 new line separator (not JSON).
  cli.go -get people-plain/id-and-firstname: will all the id and first in the system with 2 new line separator.
  cli.go -throw "message": will throw an error with given message.

### (Person, create a person Struct, use the CLI to delegate) with end points as follows (Hint for http endpoint create you can help from aukgit - Go modules training)
  localhost:8080/people Http-GET : will return list of all the person in the system as JSON array
  localhost:8080/people/create Http-POST : will create and store person to the json file.
  localhost:8080/people/5 Http-GET : will get a person as JSON with id : 5.
  localhost:8080/firstName/alim Http-GET : will get a person as JSON with first name contains alim.