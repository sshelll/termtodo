package todolist

import (
	"fmt"
	"github.com/SCU-SJL/termtodo/util"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"sort"
)

var inst *todoList

type todoList struct {
	filePath  string      `yaml:"-"`
	items     []*todoItem `yaml:"items"`
	doneItems []*todoItem `yaml:"-"`
	todoItems []*todoItem `yaml:"-"`
}

type todoItem struct {
	Key    string `yaml:"key"`
	Desc   string `yaml:"desc"`
	IsDone bool   `yaml:"is_done"`
}

func Add(key, desc string) {

	if util.IsBlank(key) {
		key = "default"
	}

	inst.items = append(inst.items, &todoItem{
		Key:    key,
		Desc:   desc,
		IsDone: false,
	})

	regroup()

}

func DelByIdx(idx int) {

	if idx < 0 || idx > len(inst.items)-1 {
		return
	}

	inst.items = append(inst.items[:idx], inst.items[idx+1:]...)

	regroup()

}

func Count() (doneCnt, todoCnt int) {
	return len(inst.doneItems), len(inst.todoItems)
}

func Reload() error {

	data, err := ioutil.ReadFile(inst.filePath)
	if os.IsNotExist(err) {
		util.WithFatalf(Save, "init file")
		return nil
	}

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &inst.items); err != nil {
		return err
	}

	regroup()

	return nil

}

func Lines() []string {

	lines := make([]string, 0, 2+len(inst.items))

	lines = append(lines, "[TODO]:")

	for _, item := range inst.items {
		ln := fmt.Sprintf("[🕑] %s", item.Desc)
		if item.IsDone {
			ln = fmt.Sprintf("[👌] %s", item.Desc)
		}
		lines = append(lines, ln)
	}

	return lines

}

func TodoLines() []string {

	lines := make([]string, 0, 2+len(inst.items))

	lines = append(lines, "[TODO]:")

	for _, item := range inst.todoItems {
		ln := fmt.Sprintf("[🕑] %s", item.Desc)
		if item.IsDone {
			ln = fmt.Sprintf("[👌] %s", item.Desc)
		}
		lines = append(lines, ln)
	}

	return lines

}

func Save() error {

	f, err := os.Create(inst.filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	d, err := yaml.Marshal(inst.items)
	if err != nil {
		return err
	}

	_, err = f.Write(d)
	return err
}

func SwitchStatusByIdx(idx int) {

	if idx < 0 || idx > len(inst.items)-1 {
		return
	}

	inst.items[idx].IsDone = !inst.items[idx].IsDone

	regroup()

}

func regroup() {

	sort.SliceStable(inst.items, func(i, j int) bool {
		return !inst.items[i].IsDone && inst.items[j].IsDone
	})

	inst.doneItems = make([]*todoItem, 0, len(inst.items))
	inst.todoItems = make([]*todoItem, 0, len(inst.items))

	for _, item := range inst.items {
		if item.IsDone {
			inst.doneItems = append(inst.doneItems, item)
		} else {
			inst.todoItems = append(inst.todoItems, item)
		}
	}

}

func init() {

	inst = new(todoList)

	util.WithFatalf(func() error {
		p, err := util.JoinHomePath(`/local/lib/todo.yml`)
		inst.filePath = p
		return err
	}, "get file path")

	util.WithFatalf(func() error {
		_, err := os.Stat(inst.filePath)
		return err
	}, "file stat")

	util.WithFatalf(Reload, "reload")

}
