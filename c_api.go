package configurator

// Configurator
// API
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

//"errors"
//"flag"
//"fmt"
//"os"
//"reflect"
//"strconv"
//"strings"

import "github.com/BurntSushi/toml"

const (
	// DELIMITER_COMMAND - is used to bind the section name and key name (command line)
	DELIMITER_COMMAND string = "/"
	// DELIMITER_PARAM - is used to separate key and value (command line)
	DELIMITER_PARAM string = "="
	// ARGUMENT_NAME_CONF_FILE - an argument path name of the configuration file on the command line
	ARGUMENT_NAME_CONF_FILE = "confile"
)

// Configurator structure
//Configuring using the configuration file, environment variables, and command-line variables.
//The default configuration is loaded from the specified file ( `config.toml`).
//The configuration file can be changed from the command line like this:
//
//	`yourservice -confile config.toml`
//
//If the operating system environment variables are set up, they have a higher priority than variables from the configuration file. Command line //parameters are most important priority. To change a parameter in the command line you need to specify its name in the form of a section name and //parameter name (with capital letters!). Here is an example to change the port and the name:
//
//	'yourservice -Main/Port 85 -Main/Name Happy`
//
type Configurator struct {
	config
}

// NewConfigurator - create a new Configurator-struct
func NewConfigurator(path string) (*Configurator, error) {
	t := &Configurator{}
	argCommLine, err := t.parseCommandLine()
	if err != nil {
		return nil, err
	}
	if s, ok := argCommLine[ARGUMENT_NAME_CONF_FILE]; ok {
		path = s
	}
	if _, err := toml.DecodeFile(path, t); err != nil {
		return nil, err
	}
	if err := t.env(argCommLine); err != nil {
		return nil, err
	}
	return t, nil
}
