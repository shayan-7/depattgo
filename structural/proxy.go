package structural

type Server interface {
	HandleRequest(url, method string) (status int, message string)
}

type WebServer struct {
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (ws *WebServer) HandleRequest(url, method string) (int, string) {
	if url == "/status" && method == "GET" {
		return 200, "OK"
	} else if url == "/users" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Found"

}

type Nginx struct {
	ws                *WebServer
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginx() *Nginx {
	return &Nginx{
		ws:                NewWebServer(),
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	if allowed := n.CheckRateLimit(url); !allowed {
		return 403, "Not Allowed"
	}
	return n.ws.HandleRequest(url, method)
}

func (n *Nginx) CheckRateLimit(url string) bool {
	rate := n.rateLimiter[url]
	if rate >= n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] += 1
	return true
}
