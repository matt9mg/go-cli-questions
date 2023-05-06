[![Go Reference](https://pkg.go.dev/badge/github.com/matt9mg/go-cli-questions.svg)](https://pkg.go.dev/github.com/matt9mg/go-cli-questions)

# CLI Questions
Simple CLI questions and answers

### Installation
```
go get github.com/matt9mg/go-cli-questions
```

### Examples
```go
q := questions.NewQuestion()
answer, err := q.Ask("hi there")

if err != nil {
    log.Fatalln(err)
}

log.Println(answer)

// asks a question but the terminal does not display the input
pw, err := q.AskSecurely("password")

if err != nil {
    log.Fatalln(err)
}

log.Println(pw)
```

Custom Template Handler
```go
type CustomTemplate struct {}

func (*CustomTemplate) Write(data []byte) error {
	t, err := template.ParseGlob("/*.tmpl")

	if err != nil {
		return err
	}

	return t.ExecuteTemplate(os.Stdout, "hello.tmpl", data)
}

func main() {
	q := questions.NewQuestion(
		questions.WithCustomTemplate(&CustomTemplate{}),
	)
	a, err := q.Ask("hi there")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(a)
}
```

### LICENSE

This project is licensed under the MIT License - see the LICENSE.md file for details

### Disclaimer

We take no legal responsibility for anything this code is used for.