package s3

import (
	"log"
	"fmt"
	"os"
	"nequi.com/poc-reading-files/internal/domain"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	//"errors"
	"bufio"
)

type S3Repository struct {
	d *s3manager.Downloader
}

type LineRecord struct {
	Cuenta string
	CifId string
	TipoCRedito string
}

func NewS3Repository() *S3Repository {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
    downloader := s3manager.NewDownloader(sess)
	return &S3Repository{
		d : downloader,
	}
}

//requiere la variable de entorno 
//export NEQUI_BUCKET_NAME="nequi-s3-select-tmp"
//export NEQUI_FILE_KEY="resource/FINACLE_NEQUICARTERA_20200508_VENCIDOS.csv"
func (r *S3Repository) GetTransactions() (transaction.Transaction, error) {
	transaction, error := transaction.NewTransaction(500,1, 2, 3, "debentura" )
	file, err := os.Create("myname")
    if err != nil {
        fmt.Println("Unable to open file %q, %v", "myname", err)
    }
    defer file.Close()
	numBytes, err := r.d.Download(file,
        &s3.GetObjectInput{

            Bucket: aws.String(os.Getenv("NEQUI_BUCKET_NAME")),
            Key:    aws.String(os.Getenv("NEQUI_FILE_KEY")),
        })
    if err != nil {
        fmt.Println("pailas", err)
    }
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	file2, err := os.Open("myname")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file2)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	var numeros1 []LineRecord;
	for _, eachline := range txtlines {
		fmt.Println(eachline)		
		fmt.Println("--------------------------------")
	}
	fmt.Println(numeros1);
	e := os.Remove("myname")
    if e != nil {
        log.Fatal(e)
    }
	return transaction, error
}