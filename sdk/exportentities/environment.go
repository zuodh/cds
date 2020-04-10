package exportentities

import (
	"github.com/ovh/cds/sdk"
)

// Environment is a struct to export sdk.Environment
type Environment struct {
	Name   string                   `json:"name" yaml:"name" jsonschema_description:"The name of the environment."`
	Values map[string]VariableValue `json:"values,omitempty" yaml:"values,omitempty"`
	Keys   map[string]KeyValue      `json:"keys,omitempty" yaml:"keys,omitempty"`
}

//NewEnvironment returns an Environment from an sdk.Environment pointer
func NewEnvironment(e sdk.Environment, keys []EncryptedKey) Environment {
	env := Environment{
		Name:   e.Name,
		Values: make(map[string]VariableValue, len(e.Variables)),
	}
	for _, v := range e.Variables {
		env.Values[v.Name] = VariableValue{
			Type:  string(v.Type),
			Value: v.Value,
		}
	}
	env.Keys = make(map[string]KeyValue, len(keys))
	for _, k := range keys {
		env.Keys[k.Name] = KeyValue{
			Type:  k.Type,
			Value: k.Content,
		}
	}
	return env
}

//Environment returns a sdk.Environment entity
func (e *Environment) Environment() (env *sdk.Environment) {
	env = new(sdk.Environment)
	env.Name = e.Name
	env.Variables = make([]sdk.Variable, len(e.Values))
	var i int
	for k, v := range e.Values {
		if v.Type == "" {
			v.Type = sdk.StringVariable
		}
		env.Variables[i] = sdk.Variable{
			Name:  k,
			Type:  v.Type,
			Value: v.Value,
		}
		i++
	}

	return
}
