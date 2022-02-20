SHELL := /bin/bash

GOCMD=go
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install

.PHONY: quiz-game
quiz-game:
	echo "Running Quiz Game!" 
	cd quiz-game && $(GORUN) main.go -limit 20

.PHONY: url-shortener
url-shortener:
	echo "Running URL Shortener!"
	cd url-shortener && $(GORUN) main/main.go

.PHONY: cyoa
cyoa:
	echo "Running Choose Your Own Adventure Web Aplication!"
	cd choose-your-own-adventure && $(GORUN) main.go -port 3000 -file "gophers.json"