package middleware

var MicroService = make(map[string]interface{})

func RegisterMiddleware(services []interface{}) {
	MicroService["userService"] = services[0]
	MicroService["taskService"] = services[1]
}
