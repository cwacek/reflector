package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3View() *S3View {
	view := &S3View{
		awsSession: session.Must(session.NewSession()),
	}
	view.svc = s3.New(view.awsSession)
	return view
}

type S3View struct {
	awsSession *session.Session
	svc        *s3.S3
}

func (s *S3View) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
