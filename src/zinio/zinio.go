package main



import ("net/http"
		"flag"
        "io/ioutil"
        "fmt"
        "strings"
        "os"
        "bytes"
        )

func main() {
	page := flag.String("page", "page1", "introduce la página")
	id := flag.String("id", "416405263", "introduce wl id")
	flag.Parse()
    // Declare http client
    client := &http.Client{}

    // Declare post data
    PostData := strings.NewReader("useId=5&age=12")


    uri := bytes.NewBufferString("http://imgs.zinio.com/repository/219299425/")
	uri.WriteString(*id)
	uri.WriteString("/SVG/")
	uri.WriteString(*page)
	uri.WriteString(".svg")
	fmt.Println(uri.String())

    // Declare HTTP Method and Url
    req, err := http.NewRequest("GET", uri.String(), PostData)

    // Set cookie
    req.Header.Set("Cookie", "DYN_USER_ID=8479190473; count=x")
    resp, err := client.Do(req)
    // Read response
    data, err := ioutil.ReadAll(resp.Body)

    // error handle
    if err != nil {
        fmt.Printf("error = %s \n", err);
    }

    // Print response
    fmt.Printf("Response = %s", string(data));
    d1 := []byte(string(data))
    err2 := ioutil.WriteFile(*page+"_"+*id+".html", d1, 0644)
    f, err2 := os.Create("dat2")
    if err2!=nil{
    	fmt.Printf("error = %s \n", err2);
    }
    defer f.Close()
}