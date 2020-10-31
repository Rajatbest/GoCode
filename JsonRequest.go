import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "os"
)

// .....

type Student struct {
    Name    string `json:"name"`
    Address string `json:"address"`
}

// .....

body := &Student{
    Name:    "abc",
    Address: "xyz",
}

payloadBuf := new(bytes.Buffer)
json.NewEncoder(payloadBuf).Encode(body)
req, _ := http.NewRequest("POST", url, payloadBuf)

client := &http.Client{}
res, e := client.Do(req)
if e != nil {
    return e
}

defer res.Body.Close()

fmt.Println("response Status:", res.Status)
// Print the body to the stdout
io.Copy(os.Stdout, res.Body)
