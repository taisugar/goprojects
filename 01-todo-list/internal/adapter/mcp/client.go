package adapter_mcp

// import "fmt"

// type MCPClient struct {
// 	server *MCPServer
// }

// func NewMCPClient(s *MCPServer) *MCPClient {
// 	return &MCPClient{server: s}
// }

// func (c *MCPClient) CallTool(name string, input map[string]interface{}) {
// 	for _, t := range c.server.Tools {
// 		if t["name"] == name {
// 			handler := t["handler"].(func(map[string]interface{}) (interface{}, error))
// 			res, err := handler(input)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				return
// 			}
// 			fmt.Println("Result:", res)
// 		}
// 	}
// }
