package out

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

const (
	bucketName = "artwork-lambda-dev"
	region     = "us-east-1"
)

var (
	unableToSave = "Unable to save response"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Request struct {
	User                      User   `json:"user"`
	ProofOfOnwership          []byte `json:"proofOfOnwership"`
	CertificateOfAuthenticity []byte `json:"certificateOfAuthenticity"`
	ApprisalInfo              []byte `json:"apprisalInfo"`
	ProvenanceOfArtwork       []byte `json:"provenanceOfArtwork"`
}

func Handler(ctx context.Context, request Request) (string, error) {
	// Create a new UUID for this user
	user := request.User
	proofOfOnwership := request.ProofOfOnwership
	certificateOfAuthenticity := request.CertificateOfAuthenticity
	apprisalInfo := request.ApprisalInfo
	provenanceOfArtwork := request.ProvenanceOfArtwork

	uuid := uuid.New()

	// Upload the PDF files to S3
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Printf("Error unable to create a session: %s", err.Error())
		return "", errors.New(unableToSave)
	}

	uploader := s3manager.NewUploader(sess)
	pdfFiles := map[string][]byte{
		"8POO-" + uuid.String() + ".pdf": proofOfOnwership,
		"8COA-" + uuid.String() + ".pdf": certificateOfAuthenticity,
		"8AIN-" + uuid.String() + ".pdf": apprisalInfo,
		"8POA-" + uuid.String() + ".pdf": provenanceOfArtwork,
	}

	prefix := user.Email + "/" + uuid.String()
	for key, value := range pdfFiles {
		pdfKey := prefix + "/" + key
		_, ierr := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(pdfKey),
			Body:   bytes.NewReader(value),
		})
		if ierr != nil {
			log.Printf("Error: uploading the PDF file to S3: %s", ierr.Error())
			return "", errors.New(unableToSave)
		}
	}

	// Create the JSON file with the UUID in the file name
	jsonFileName := fmt.Sprintf("%s.json", uuid.String())
	jsonKey := prefix + "/" + "8INF-" + jsonFileName

	// Marshal the user struct into JSON
	jsonBod, err := json.Marshal(user)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New(unableToSave)
	}
	_, ierr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(jsonKey),
		Body:   bytes.NewReader(jsonBod),
	})

	if ierr != nil {
		// If there is an error in the uploading a json file, delete files
		for key, _ := range pdfFiles {
			deleteParams := &s3.DeleteObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(prefix + "/" + key),
			}
			_, delErr := s3.New(sess).DeleteObject(deleteParams)
			if delErr != nil {
				log.Printf("Error: removing the uploaded files: %s", delErr.Error())
			}
		}

		log.Printf("Error: uploading the JSON file to S3: %s", ierr.Error())
		return "", errors.New(unableToSave)
	}

	// Return the UUID as a string
	return uuid.String(), nil
}

func main() {
	lambda.Start(Handler)
}
