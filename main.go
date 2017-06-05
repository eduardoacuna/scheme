package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	header := "welcome to the scheme interactive interpreter"
	footer := "farewell schemer"
	prompt := "› "
	iport := bufio.NewReader(os.Stdin)
	oport := bufio.NewWriter(os.Stdout)
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-signals
		fmt.Println()
		fmt.Println(footer)
		os.Exit(0)
	}()

	fmt.Println(header)

	for {
		fmt.Print(prompt)

		inputData, err := read(iport)
		if err != nil {
			panic(fmt.Errorf("REPL: %v", err))
		}

		outputData, err := eval(inputData)
		if err != nil {
			panic(fmt.Errorf("REPL: %v", err))
		}

		err = print(outputData, oport)
		if err != nil {
			panic(fmt.Errorf("REPL: %v", err))
		}

		_, err = oport.WriteRune('\n')
		if err != nil {
			panic(fmt.Errorf("REPL: Encountered error while printing an empty line... %v", err))
		}
	}
}

func read(iport *bufio.Reader) (string, error) {
	data := make([]rune, 0, 512)

	for {
		r, _, err := iport.ReadRune()
		if err != nil {
			if err == io.EOF {
				return string(data), nil
			}
			return "", fmt.Errorf("Encountered error while reading... %v", err)
		}
		data = append(data, r)
	}

	return "", fmt.Errorf("Something really bad just happened while reading...")
}

func eval(data string) (string, error) {
	return data, nil
}

func print(data string, oport *bufio.Writer) error {
	for _, r := range data {
		_, err := oport.WriteRune(r)
		if err != nil {
			return fmt.Errorf("Encountered error while printing... %v", err)
		}
	}
	err := oport.Flush()
	if err != nil {
		return fmt.Errorf("Encountered error while flushing the writer... %v", err)
	}
	return nil
}
