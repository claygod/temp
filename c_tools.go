package configurator

// Microservice
// Config
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

//Config structure
//When you change the structure of the `Config`, make
//sure the same changes need to be made to `config.toml`
type config struct {
	Main    confMain
	Session confSession
}

// confMain - basic configuration
type confMain struct {
	Port string
	Name string
}

// confSession parameters
type confSession struct {
	Secure string
	Name   string
}
