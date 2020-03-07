package main

import (
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func NewWriter(dir, name string) Writer {
	if dir == "stdout" {
		return &StdoutWriter{name: name}
	}
	return &FileWriter{fp: filepath.Join(dir, name)}
}

type Writer interface {
	Write(format string, params ...interface{}) Writer
	Writeln(format string, params ...interface{}) Writer
	Ln() Writer
	Close() error
}

type StdoutWriter struct {
	name string
	buf  string
}

func (w *StdoutWriter) Write(format string, params ...interface{}) Writer {
	w.buf += fmt.Sprintf(format, params...)
	return w
}

func (w *StdoutWriter) Writeln(format string, params ...interface{}) Writer {
	w.Write(format+"\n", params...)
	return w
}

func (w *StdoutWriter) Ln() Writer {
	w.Write("\n")
	return w
}

func (w *StdoutWriter) Close() error {
	os.Stdout.WriteString("[" + w.name + "]\n")
	defer os.Stdout.WriteString("\n\n")
	return write(os.Stdout, w.buf)
}

type FileWriter struct {
	buf string
	fp  string
}

func (w *FileWriter) Write(format string, params ...interface{}) Writer {
	w.buf += fmt.Sprintf(format, params...)
	return w
}

func (w *FileWriter) Writeln(format string, params ...interface{}) Writer {
	w.Write(format+"\n", params...)
	return w
}

func (w *FileWriter) Ln() Writer {
	w.Write("\n")
	return w
}

func (w *FileWriter) Close() (err error) {
	dir := filepath.Dir(w.fp)
	if err = os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(w.fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer func() {
		err = f.Close()
	}()

	return write(f, w.buf)
}

func write(f *os.File, source string) error {
	formatted, err := format.Source([]byte(source))
	if err != nil {
		for i, l := range strings.Split(source, "\n") {
			log.Printf("%4d | %s\n", i, l)
		}
		log.Printf("%+v\n", err)
		return err
	}

	_, err = f.Write(formatted)
	return err
}
