# Makefile for Go project

build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run
