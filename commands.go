package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"github.com/cheggaaa/pb"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGet,
}

var commandGet = cli.Command{
	Name:  "get",
	Usage: "g3 get <bucketname> <target file or directory>",
	Description: `
'get' will help download from S3. For example, If you want to download the 'test.txt' files from the 'test' directory, you can be achieved with the following command:

 g3 get <bucketname> test/test.txt
`,
	Action: doGet,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getBucketNameAndPrefix() (string, string) {
	bucketName := os.Args[2]

	var prefix string
	if len(os.Args) > 3 {
		prefix = os.Args[3]
	}

	return bucketName, prefix
}

func doGet(c *cli.Context) {
	svc := getSvc()

	bucketName, prefix := getBucketNameAndPrefix()
	objectNameList := getObjectNameList(svc, bucketName, prefix)

	count := len(objectNameList)
	bar := pb.StartNew(count)
	for _, objectName := range objectNameList {
		makeDirectory(objectName)

		isMatch, err := regexp.MatchString("/$", objectName)
		if err != nil {
			assert(err)
		}
		if !isMatch {
			getObject(svc, &s3.GetObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(objectName),
			})
		}

		bar.Increment()
	}
	bar.FinishPrint("Finished🍻")
}

func getObject(svc *s3.S3, o *s3.GetObjectInput) {
	resp, err := svc.GetObject(o)
	if err != nil {
		assert(err)
	}
	// have to close, when reuse tcp connection.
	defer resp.Body.Close()

	writeFile(*o.Key, resp.Body)
}

func makeDirectory(path string) {
	regex_str := "([0-9A-Za-z_-]*/)*"
	re, err := regexp.Compile(regex_str)
	if err != nil {
		assert(err)
	}

	p := re.FindString(path)
	if p != "" {
		isExists, err := exists(p)
		if err != nil {
			assert(err)
		}

		if !isExists {
			if err := os.Mkdir(p, 0777); err != nil {
				assert(err)
			}
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func writeFile(filename string, body io.ReadCloser) {
	file, err := os.Create(filename)
	if err != nil {
		assert(err)
	}
	defer file.Close()

	io.Copy(file, body)
}

// Returns some or all (up to 1000) of the objects
func getObjectNameList(svc *s3.S3, bucketName string, prefix string) []string {
	resp, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		assert(err)
	}

	var objectNameList []string
	for _, c := range resp.Contents {
		fmt.Println(*c.Key)
		objectNameList = append(objectNameList, *c.Key)
	}

	return objectNameList
}

func getSvc() *s3.S3 {
	return s3.New(&aws.Config{Region: "ap-northeast-1"})
}
