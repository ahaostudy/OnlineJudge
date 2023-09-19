// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {
// 	url := "https://openai.ahao.ink/v1/chat/completions"
// 	headers := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Authorization": "Bearer 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",
// 	}
// 	data := map[string]interface{}{
// 		"model": "gpt-3.5-turbo",
// 		"messages": []map[string]string{
// 			{"role": "user", "content": "hello"},
// 		},
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		panic(err)
// 		// 处理错误
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		panic(err)
// 		// 处理错误
// 	}

// 	for key, value := range headers {
// 		req.Header.Set(key, value)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 		// 处理错误
// 	}

// 	defer resp.Body.Close()

// 	// 处理响应
// 	bytes, err := io.ReadAll(resp.Body)

// 	fmt.Printf("string(bytes): %v\n", string(bytes))

// }
package main

import (
	"fmt"
	"net/url"
)

func main() {
    baseURL := "https://example.com/"
    endpoint := "/v1/chat/completions"

	base, err := url.Parse(baseURL)
    if err != nil {
        panic(err)
    }

    resolvedURL := base.ResolveReference(&url.URL{Path: endpoint})
    fmt.Println(resolvedURL.String())
}