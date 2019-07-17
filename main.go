package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Print("Usage: aws-assume-run [role_arn] [command] [args...]")
		os.Exit(1)
	}

	sess := session.Must(session.NewSession())
	creds := stscreds.NewCredentials(sess, os.Args[1])
	values, err := creds.Get()

	must(err)

	binary, err := exec.LookPath(os.Args[2])

	must(err)

	cmd := exec.Command(binary, os.Args[3:]...)
	roleCreds := fmt.Sprintf(
		"AWS_ACCESS_KEY_ID=%s,AWS_SECRET_ACCESS_KEY=%s,AWS_SESSION_TOKEN=%s",
		values.AccessKeyID,
		values.SecretAccessKey,
		values.SessionToken,
	)
	cmd.Env = append(os.Environ(), strings.Split(roleCreds, ",")...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	must(err)

}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
