package anysrv

const (
	GET = iota
	POST
	PUT
	HEAD
	DELETE
	PATCH
	OPTIONS
)

func getMethodCode(method string) int {
	switch method {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "HEAD":
		return HEAD
	case "DELETE":
		return DELETE
	case "PATCH":
		return PATCH
	case "OPTIONS":
		return OPTIONS
	default:
		return GET
	}
}
