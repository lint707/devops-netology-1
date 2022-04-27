package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
)

var uploadFormTmpl = []byte(`
<html>
	<body>
	<form action="/upload" method="post" enctype="multipart/form-data">
		Image: <input type="file" name="my_file">
		<input type="submit" value="Upload">
	</form>
	</body>
</html>
`)

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write(uploadFormTmpl)
}
func uploadPage(w http.ResponseWriter, r *http.Request) {
	// f, _ :=os.Create("result_file.txt")
	// io.Cipy(f, r.Body)
	// r.Body.Close()
	// f.Sync()
	// f.Close()
	// fmt.Fprintln(w, "you enter: ")
	// retutn

	r.ParseMultipartForm(5 * 1024 * 1024)
	file, header, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "handler.Filename $v\n", header.Filename)
	fmt.Fprintf(w, "handler.Header $v\n", header.Header)

	hasher := md5.New()
	io.Copy(hasher, file)
	fmt.Fprintf(w, "md5 %x\n", hasher.Sum(nil))
}

type Params struct {
	ID   int
	User string
}

//func main(){} - НЕ ВЕСЬ КОД
