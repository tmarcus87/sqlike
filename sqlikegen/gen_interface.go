package main

type Generator interface {
	Generate(string, *Schema) error
}
