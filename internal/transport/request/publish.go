package request

import "mime/multipart"

type PublishAction struct {
	Title  string
	UserID int64
	Data   *multipart.FileHeader
}

type PublishList struct {
	LoginID int64
	UserID  int64
}
