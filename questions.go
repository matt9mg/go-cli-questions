package questions

import (
	"bufio"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
)

type ConfigFunc func(*Config)

type Config struct {
	template Writer
}

type Writer interface {
	Write(data []byte) error
}

type Stdout struct {
	writer io.Writer
}

func NewStdout() *Stdout {
	return &Stdout{
		writer: io.Writer(os.Stdout),
	}
}

func (s *Stdout) Write(data []byte) error {
	_, err := s.writer.Write(data)
	s.writer.Write([]byte("\n"))

	return err
}

type Question struct {
	config *Config
	reader *bufio.Reader
}

// NewQuestion returns a Question Client based on the config functional options. Provide
// additional config functional options to further configure the behavior of the Question client,
// such as changing the render.
func NewQuestion(configs ...ConfigFunc) *Question {
	c := setDefaults()

	for _, fn := range configs {
		fn(c)
	}

	return &Question{
		reader: bufio.NewReader(os.Stdin),
		config: c,
	}
}

func setDefaults() *Config {
	return &Config{
		template: NewStdout(),
	}
}

// WithCustomTemplate takes a Writer interface that allows you to customise the look and feel of the questions asked.
func WithCustomTemplate(w Writer) ConfigFunc {
	return func(c *Config) {
		c.template = w
	}
}

// Ask takes a question and returns the answer or an error
func (q *Question) Ask(question string) (string, error) {
	for {
		if err := q.config.template.Write([]byte(question)); err != nil {
			return "", err
		}

		return q.reader.ReadString('\n')
	}
}

// AskSecurely asks your question and returns the answer or an error
// The terminal input is hidden when entering the answer
func (q *Question) AskSecurely(question string) (string, error) {
	for {
		if err := q.config.template.Write([]byte(question)); err != nil {
			return "", err
		}

		pw, err := terminal.ReadPassword(0)

		if err != nil {
			return "", err
		}

		return string(pw), nil
	}
}
