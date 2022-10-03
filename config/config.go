package config

import (
	"io/ioutil"
	"reflect"

	yaml "gopkg.in/yaml.v2"
)

// action that should be fired on a hotkey
type Action struct {
	Command string
	Target  string
}

// keyboard is a set of actions grouped by keys
type Keyboard map[int]Action

// file types
type Types struct {
	Supported   []string
	Unsupported []string
}

type ActionSet map[string]string
type Shortcuts struct {
	Clip ActionSet
	Kit  ActionSet
}

type Config struct {
	Types
	Shortcuts
}

func ReadConfig(filename string) (config Config) {
	yamlFile, err := ioutil.ReadFile(filename)
	check(err)
	err = yaml.Unmarshal(yamlFile, &config)
	check(err)
	return
}

func CollectShortCuts(config Config) (keyboard *Keyboard) {
	k := make(Keyboard)
	keyboard = &k
	gatherActions(&config.Shortcuts, keyboard)
	return keyboard
}

// this functions collects all actions to a single map
// to call them further when key is received
func gatherActions(s *Shortcuts, k *Keyboard) {
	fs := reflect.ValueOf(s).Elem()
	// iterate over struct fields to get all targets
	for i := 0; i < fs.NumField(); i++ {
		fieldName := fs.Type().Field(i).Name
		field := fs.Field(i)
		// collect all keys
		for action, key := range field.Interface().(ActionSet) {
			(*k)[getKey(key)] = Action{Command: action, Target: fieldName}
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getKey(s string) int {
	s1 := []rune(s)
	return int(s1[0])
}
