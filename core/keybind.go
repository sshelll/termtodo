package core

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

type KeyBinder struct {
	mapping map[tcell.Key]func(*tcell.EventKey)
}

func (kb *KeyBinder) Bind(fn func(*tcell.EventKey), keys ...tcell.Key) {

	if kb.mapping == nil {
		kb.mapping = make(map[tcell.Key]func(*tcell.EventKey))
	}

	if len(keys) == 0 {
		return
	}

	for _, k := range keys {
		if _, ok := kb.mapping[k]; ok {
			panic(fmt.Sprintf("key [%d] bind twice", k))
		}
	}

	for _, k := range keys {
		kb.mapping[k] = fn
	}

}

func (kb *KeyBinder) Find(key tcell.Key) func(*tcell.EventKey) {
	return kb.mapping[key]
}
