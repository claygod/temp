package configurator

// Configurator
// Configurator
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	//"flag"
	//"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	//"github.com/BurntSushi/toml"
)

func (t *Configurator) parseCommandLine() (map[string]string, error) {
	argCommLine := make(map[string]string)
	ln := len(os.Args)
	//if ln%2 != 1 && ln != 1 {
	//	return nil, errors.New("Command Line Error (check the number of arguments)")
	//}
	for i := 1; i < ln-1; i++ {
		key := strings.TrimLeft(os.Args[i], "-")
		if len(key) == len(os.Args[i])-1 {
			i++
			argCommLine[key] = os.Args[i]
		} // else {
		//	return nil, errors.New("Command Line Error (no dashes)")
		//}
	}
	return argCommLine, nil
}

func (t *Configurator) env(argCommLine map[string]string) error {
	for _, str := range os.Environ() {
		parEnv := strings.Split(str, DELIMITER_PARAM)
		parKey := strings.Split(parEnv[0], DELIMITER_COMMAND)
		if len(parKey) == 2 {
			if err := t.reflecting(parKey[0], parKey[1], parEnv[1]); err != nil {
				return err
			}
		}
	}
	for k, str := range argCommLine {
		parKey := strings.Split(k, DELIMITER_COMMAND)
		if len(parKey) == 2 {
			if err := t.reflecting(parKey[0], parKey[1], str); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Configurator) reflecting(key1 string, key2 string, str string) error {
	t1 := reflect.TypeOf(*t)
	v1 := reflect.ValueOf(t)
	v1 = reflect.Indirect(v1)
	if m, ok := t1.FieldByName(key1); ok {
		v2 := v1.FieldByName(key1).Addr()
		v2 = reflect.Indirect(v2)
		if _, okk := m.Type.FieldByName(key2); okk {
			v3 := v2.FieldByName(key2).Addr()
			v3 = reflect.Indirect(v3)
			return t.switchType(v3, str)
		}
	}
	return nil
}

func (t *Configurator) switchType(v3 reflect.Value, str string) error {
	tp := v3.Type()
	switch tp {
	case reflect.TypeOf(""):
		v3.SetString(str)
	case reflect.TypeOf(1):
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		v3.SetInt(num)
	case reflect.TypeOf(0.1):
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		v3.SetFloat(num)
	default:
		return errors.New("Not supported by type!")
	}
	return nil
}
