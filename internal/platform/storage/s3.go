package s3

import (
	//"log"
	"fmt"
	//"os"
	"nequi.com/poc-reading-files/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	//"errors"
)

type S3Repository struct {
	client *s3.S3
}

func NewS3Repository() *S3Repository {
	svc := s3.New(session.New(),&aws.Config{Region: aws.String("us-east-1")})
	return &S3Repository{
		client : svc,
	}
}

//requiere la variable de entorno export NEQUI_CREDITS_TABLE_NAME=credit-customer-product-qa
func (r *S3Repository) GetTransactions() (transaction.Transaction, error) {
	transaction, error := transaction.NewTransaction(500,1, 2, 3, "debentura" )
	
	obj, err := r.client.GetObject(&s3.GetObjectInput{
        Bucket: aws.String("nequi-s3-select-tmp"),
        Key:    aws.String("resource/FINACLE_NEQUICARTERA_20200508_VENCIDOS.csv"),
    })
	fmt.Println(obj)
    if err != nil {
        fmt.Println("pailas", err)
    }	
	return transaction, error
}