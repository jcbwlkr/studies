package main

import (
	"bytes"
	"fmt"
	"log"
)

const output = `Status: 202 Accepted
X-Powered-By: PHP/7.2.0
Content-type: text/html; charset=UTF-8

<!DOCTYPE html>
<html lang="en">
	<head>
		<title>Hello, World!</title>
	</head>
	<body>
		<h1>Hello, World!</h1>
		<p>PHP works!</p>
	</body>
</html>`

func main() {

	src := []byte(output)

	pos := bytes.Index(src, []byte("\n\n"))
	if pos == -1 {
		pos = bytes.Index(src, []byte("\r\n\r\n"))
	}
	if pos == -1 {
		log.Fatal("Um I don't know")
	}

	fmt.Println("Headers:")
	fmt.Println(string(src[:pos]))

	fmt.Println("Body:")
	fmt.Println(string(src[pos+2:]))

	//var res http.ResponseWriter = httptest.NewRecorder()

	//reg := regexp.MustCompile(`Status: (\d{3})`)

	//r := strings.NewReader(output)
	//s := bufio.NewScanner(r)

	//var status int

	//for s.Scan() {
	//line := s.Text()

	//// Status should be the first line
	//if status == 0 {
	//m := reg.FindStringSubmatch(line)
	//var err error
	//status, err = strconv.Atoi(m[1])
	//if err != nil {
	//log.Fatal("could not parse status", err)
	//}

	//continue
	//}

	//// Capture headers until we see a blank line
	//if line == "" {
	//break
	//}

	//sep := strings.Index(line, ": ")
	//res.Header().Add(line[:sep], line[sep+2:])
	//continue
	//}

	//if err := s.Err(); err != nil {
	//log.Fatalln("scan error", err)
	//}

	//// Done finding headers. Go ahead and write the header block
	//res.WriteHeader(status)

	//// Continue scanning to write out the body
	//for s.Scan() {
	//// Copy body to output buffer
	//if _, err := res.Write(s.Bytes()); err != nil {
	//log.Fatalln("could not copy body line", err)
	//}
	//if _, err := res.Write([]byte("\n")); err != nil {
	//log.Fatalln("could not write newline", err)
	//}
	//}

	//if err := s.Err(); err != nil {
	//log.Fatalln("scan error", err)
	//}

	//rec := res.(*httptest.ResponseRecorder)
	//rec.Flush()
	//fmt.Println("What was sent to response:")
	//fmt.Println("Status:", rec.Code)
	//fmt.Println(rec.Header())
	//fmt.Println("")
	//fmt.Println(rec.Body.String())
}
