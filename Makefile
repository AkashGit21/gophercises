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
	cd quiz-game && $(GORUN) main.go -limit 10