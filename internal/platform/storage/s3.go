package s3

import (
	//"log"
	"fmt"
	"os"
	"nequi.com/poc-reading-files/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	//"errors"
)

type S3Repository struct {
	d *s3manager.Downloader
}

func NewS3Repository() *S3Repository {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")},
    )
    downloader := s3manager.NewDownloader(sess)
	return &S3Repository{
		d : downloader,
	}
}

//requiere la variable de entorno export NEQUI_CREDITS_TABLE_NAME=credit-customer-product-qa
func (r *S3Repository) GetTransactions() (transaction.Transaction, error) {
	transaction, error := transaction.NewTransaction(500,1, 2, 3, "debentura" )
	
	file, err := os.Create("myname")
    if err != nil {
        fmt.Println("Unable to open file %q, %v", "myname", err)
    }

    defer file.Close()
	
	numBytes, err := r.d.Download(file,
        &s3.GetObjectInput{
            Bucket: aws.String("nequi-s3-select-tmp"),
            Key:    aws.String("resource/FINACLE_NEQUICARTERA_20200508_VENCIDOS.csv"),
        })
    if err != nil {
        fmt.Println("pailas", err)
    }

    fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	
	return transaction, error
}