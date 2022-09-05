package request

import "mime/multipart"

type PublishAction struct {
	Title   string
	LoginID int64
	Data    *multipart.FileHeader
}

type PublishList struct {
	LoginID int64
	UserID  int64
}
