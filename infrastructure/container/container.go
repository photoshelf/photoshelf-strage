package container

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"reflect"
)

type singletonContainer struct {
	values map[string]interface{}
}

var instance *singletonContainer

func init() {
	instance = &singletonContainer{
		make(map[string]interface{}),
	}
}

func Set(val interface{}) {
	var key string
	if reflect.TypeOf(val).Kind() == reflect.Ptr {
		key = reflect.Indirect(reflect.ValueOf(val)).Type().String()
	} else {
		key = reflect.ValueOf(val).Type().String()
	}

	log.Debug(fmt.Sprintf("added %s to components container.", key))
	instance.values[key] = val
}

func Get(ptr interface{}) {
	val := reflect.ValueOf(ptr)
	key := reflect.Indirect(val).Type().String()
	component := instance.values[key]
	if component == nil {
		if reflect.Indirect(val).Type().Kind() != reflect.Interface {
			log.Warn(fmt.Sprintf("component not found. such type of %s.", key))
			return
		}
		for _, component := range instance.values {
			value := reflect.ValueOf(component)
			elm := reflect.ValueOf(ptr).Elem()
			if value.Type().Implements(elm.Type()) {
				log.Debug(fmt.Sprintf("found component of %s .", key))
				elm.Set(value)
				return
			}
		}
		return
	}
	log.Debug(fmt.Sprintf("found component of %s .", key))

	elm := reflect.ValueOf(ptr).Elem()
	if reflect.TypeOf(component).Kind() == reflect.Ptr {
		elm.Set(reflect.Indirect(reflect.ValueOf(component)))
	} else {
		elm.Set(reflect.ValueOf(component))
	}
}
