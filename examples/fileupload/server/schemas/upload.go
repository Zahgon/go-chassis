package schemas

import (
	"net/http"

	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

// RestFulUpload is a struct used to implement restful upload
type RestFulUpload struct {
}

// UploadFile is a method used to reply user hello in json format
func (r *RestFulUpload) UploadFile(b *rf.Context) {
	_ = "STUB: not implemented"
	// USAGE: curl -X POST http://127.0.0.1:8083/uploadfile -H 'content-type: application/octet-stream' --data-binary '@input.txt'
	return
}

func uploadFile(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"

	//POST takes the uploaded file(s) and saves it to disk.
	return nil
}

// UploadForm is a method used to reply user hello in json format
func (r *RestFulUpload) UploadForm(b *rf.Context) {
	_ = "STUB: not implemented"
	// USAGE: curl -X POST http://127.0.0.1:8083/uploadform -H 'content-type: multipart/form-data' -F uploadfile=@input.txt
	return
}

func uploadForm(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"

	//POST takes the uploaded file(s) and saves it to disk.
	return nil
}

//parse the multipart form in the request

//get a ref to the parsed multipart form

//get the *fileheaders

//for each fileheader, get a handle to the actual file

//create destination file making sure the path is writeable.

//copy the uploaded file to the destination file

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulUpload) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }
