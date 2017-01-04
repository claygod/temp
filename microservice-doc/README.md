# Microservice

[![API documentation](https://godoc.org/github.com/claygod/temp/microservice-doc?status.svg)](https://godoc.org/github.com/claygod/temp/microservice-doc)

The framework for the creation of microservices, written in Golang.

Architecture microservice includes a handle, a tuner (configuration), place a couple of demonstration middleware and storage for them. All works is very simple: in the application configuration is loaded in view of the configuration file and command line environment. Created with the middleware storage and corresponding queues formed at the desired Route. Then run the server and request the application fulfills the desired queue.
