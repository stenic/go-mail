package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"log"

	"github.com/NoUseFreak/go-vembed"
	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
)

var smtp string
var from string
var cc []string
var bcc []string
var subject string
var body string

func init() {
	rootCmd.PersistentFlags().StringVar(&smtp, "smtp", "smtp://127.0.0.1:25", "Define the smtp server")
	rootCmd.PersistentFlags().StringVarP(&subject, "subject", "s", "Test mail from go-mail", "Subject of the mail")
	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "go-mail@example.com", "Send the mail from this address")
	rootCmd.PersistentFlags().StringVar(&body, "body", "Hello from go-mail", "Define the message body")
	rootCmd.PersistentFlags().StringArrayVarP(&cc, "cc-addr", "c", []string{}, "Send a carbon copy to this address")
	rootCmd.PersistentFlags().StringArrayVarP(&bcc, "bcc-addr", "b", []string{}, "Send a blink carbon copy to this address")

	rootCmd.Version = fmt.Sprintf(
		"%s, build %s",
		vembed.Version.GetGitSummary(),
		vembed.Version.GetGitCommit(),
	)
}

func main() {
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "go-mail [-s subject] to-addr ...",
	Short: "go-mail",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		stdin, err := readStdin()
		if err != nil {
			log.Fatalf("Failed capturing input - %s", err.Error())
			return err
		}
		if stdin != "" {
			body = stdin
		}

		u, err := url.Parse(smtp)
		if err != nil {
			return err
		}
		var port int
		if port, err = strconv.Atoi(u.Port()); err != nil {
			return err
		}

		dailer := gomail.Dialer{Host: u.Hostname(), Port: port}
		m := gomail.NewMessage()

		m.SetHeader("To", args...)
		if len(cc) > 0 {
			m.SetHeader("Cc", cc...)
		}
		if len(bcc) > 0 {
			m.SetHeader("Bcc", bcc...)
		}
		m.SetHeader("From", from)
		m.SetHeader("Subject", subject)
		m.SetBody("text/plain", body)

		if err := dailer.DialAndSend(m); err != nil {
			log.Fatalf("Failed to send message - %s", err.Error())
			return err
		}
		log.Printf("Message sent to %s", strings.Join(args, ", "))
		return nil
	},
}

func readStdin() (string, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {

		return "", nil
	}
	var stdin []byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = append(stdin, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return string(stdin), nil
}
