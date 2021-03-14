package tagsalot

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// PrincipalEntry in an IAM Policy
type PrincipalEntry struct {
	AWS []string
}

// StatementEntry in an IAM Policy
type StatementEntry struct {
	Sid       string
	Effect    string
	Principal PrincipalEntry
	Action    []string
}

// PolicyDocument in an IAM Policy
type PolicyDocument struct {
	Version   string
	Statement []StatementEntry
}

func CreateECRRepository(repoName string) {
	svc := ecr.New(session.New())

	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(repoName),
		EncryptionConfiguration: &ecr.EncryptionConfiguration{
			EncryptionType: aws.String("AES256"),
		},
		ImageScanningConfiguration: &ecr.ImageScanningConfiguration{
			ScanOnPush: aws.Bool(true),
		},
		ImageTagMutability: aws.String("MUTABLE"),
		Tags: []*ecr.Tag{
			{
				Key:   aws.String("CostTech"),
				Value: aws.String("ecr"),
			},
			{
				Key:   aws.String("CostProduct"),
				Value: aws.String("mixed"),
			},
		},
	}
	repo, err := svc.CreateRepository(input)
	if err != nil {
		if ecrErr, ok := err.(awserr.Error); ok && ecrErr.Code() != "RepositoryAlreadyExistsException" {
			fmt.Println(err)
		}
	} else {
		fmt.Println(repo)
	}

	policy, err := createPolicyStatement()
	if err != nil {
		fmt.Println(err)
	}

	svc.SetRepositoryPolicy(&ecr.SetRepositoryPolicyInput{
		RepositoryName: aws.String(repoName),
		PolicyText:     aws.String(string(policy)),
	})

}

func createSession() *session.Session {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return sess
}

func createPolicyStatement() ([]byte, error) {
	policy := PolicyDocument{
		Version: "2008-10-17",
		Statement: []StatementEntry{
			StatementEntry{
				Sid:    "Cross Account Access",
				Effect: "Allow",
				Principal: PrincipalEntry{
					// TODO: make this from a config map
					AWS: []string{
						"arn:aws:iam::187042759598:root",
						"arn:aws:iam::012364521670:root",
						"arn:aws:iam::640532232197:root",
						"arn:aws:iam::963188529772:root",
						"arn:aws:iam::548768115844:root",
						"arn:aws:iam::032693105358:root",
					},
				},
				Action: []string{
					"ecr:GetDownloadUrlForLayer",
					"ecr:BatchGetImage",
					"ecr:BatchCheckLayerAvailability",
					"ecr:DescribeImages",
				},
			},
		},
	}

	b, err := json.Marshal(&policy)
	if err != nil {
		return nil, err
	}

	return b, nil
}
