package todolist

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/sshelll/termtodo/util"
	"gopkg.in/yaml.v3"
)

var inst *todoList

type todoList struct {
	Items      []*todoItem          `yaml:"Items"`
	CateKeys   []string             `yaml:"cate_keys"`
	filePath   string               `yaml:"-"`
	categories map[string]*category `yaml:"-"`
}

type category struct {
	key   string
	fold  bool
	items []*todoItem
}

type todoItem struct {
	Key    string `yaml:"key"`
	Desc   string `yaml:"desc"`
	IsDone bool   `yaml:"is_done"`
}

func newCategory(key string) *category {
	return &category{
		key:  key,
		fold: true,
	}
}

func (c *category) Key() string {
	return c.key
}

func (c *category) Fold() bool {
	return c.fold
}

func (c *category) Size() int {
	return len(c.items)
}

func (c *category) Append(item *todoItem) *category {
	c.items = append(c.items, item)
	return c
}

func AddNewItem(key, desc string) {

	if util.IsBlank(key) {
		key = "default"
		inst.CateKeys = append(inst.CateKeys, key)
	}

	inst.Items = append(inst.Items, &todoItem{
		Key:    key,
		Desc:   desc,
		IsDone: false,
	})

	regroup()

}

func CateKeys() []string {
	return inst.CateKeys
}

func GetCategory(key string) *category {
	return inst.categories[key]
}

func AddNewCategory(key string) {
	inst.CateKeys = append(inst.CateKeys, key)
	regroup()
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

	if err := yaml.Unmarshal(data, &inst); err != nil {
		return err
	}

	regroup()

	return nil

}

func DelWithCategoryKey(cateKey string, offset int) {

	var cate *category
	var idx int

	for i, k := range inst.CateKeys {
		if k == cateKey {
			cate = inst.categories[k]
			idx = i
		}
	}

	// cursor is at category -> del category
	if offset == 0 {
		delete(inst.categories, cateKey)

		inst.CateKeys = append(inst.CateKeys[:idx], inst.CateKeys[idx+1:]...)

		return
	}

	if cate == nil {
		return
	}

	// cursor is below category -> del item of category
	cate.items = append(cate.items[:offset-1], cate.items[offset:]...)

	// rewrite item list
	inst.Items = make([]*todoItem, 0, len(inst.Items)-1)
	for _, cate := range inst.categories {
		for _, item := range cate.items {
			inst.Items = append(inst.Items, item)
		}
	}

}

// Content convert todolist to content for show.
func Content() []string {

	lines := make([]string, 0, 2+len(inst.Items))

	lines = append(lines, "[TODO]:")

	for _, cateKey := range inst.CateKeys {
		if cate, ok := inst.categories[cateKey]; !ok {
			lines = append(lines, fmt.Sprintf("▾ %s/ (empty)", cateKey))
		} else if cate.fold {
			lines = append(lines, fmt.Sprintf("▸ %s/", cate.key))
		} else if !cate.fold {
			lines = append(lines, fmt.Sprintf("▾ %s/", cate.key))
			for _, item := range cate.items {
				ln := fmt.Sprintf("[🕑] %s", item.Desc)
				if item.IsDone {
					ln = fmt.Sprintf("[👌] %s", item.Desc)
				}
				lines = append(lines, ln)
			}
		}
	}

	return lines
}

func DoingContent() []string {
	return groupedContent(false)
}

func DoneContent() []string {
	return groupedContent(true)
}

func groupedContent(isDone bool) []string {

	lines := make([]string, 0, len(inst.Items))
	if isDone {
		lines = append(lines, "[What's done!✅]")
	} else {
		lines = append(lines, "[What's doing!❎]")
	}

	if len(inst.Items) == 0 {
		lines[0] = fmt.Sprintf("%s/(empty)", lines[0])
		return lines
	}

	tmpl := util.If(isDone, "[👌] %s - (%s)", "[🕑] %s - (%s)").(string)

	for _, item := range inst.Items {
		if isDone == item.IsDone {
			lines = append(lines, fmt.Sprintf(tmpl, item.Desc, item.Key))
		}
	}

	return lines

}

func Save() error {

	if err := util.CreateFile(filepath.Dir(inst.filePath)); err != nil {
		return err
	}

	f, err := os.Create(inst.filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	d, err := yaml.Marshal(inst)
	if err != nil {
		return err
	}

	_, err = f.Write(d)
	return err

}

func SwitchFoldStatus(cateKey string) {
	for _, cate := range inst.categories {
		if cate.Key() == cateKey {
			cate.fold = !cate.fold
		}
	}
}

func RemarkItemStatus(cateKey string, offset int) {
	for _, cate := range inst.categories {
		if cate.Key() == cateKey && offset-1 < len(cate.items) {
			cate.items[offset-1].IsDone = !cate.items[offset-1].IsDone
		}
	}
}

func regroup() {

	// sort by 'is_done' and ascii
	sort.SliceStable(inst.Items, func(i, j int) bool {
		if !inst.Items[i].IsDone && inst.Items[j].IsDone {
			return true
		}
		return inst.Items[i].Key < inst.Items[i].Key
	})

	oldCate := inst.categories

	// 5 categories should be enough :)
	inst.categories = make(map[string]*category, 5)
	cateMap := make(map[string]*category, 5)

	for _, item := range inst.Items {
		// update categories
		if cateMap[item.Key] == nil {
			cateMap[item.Key] = newCategory(item.Key).Append(item)
			if c, ok := oldCate[item.Key]; ok {
				cateMap[item.Key].fold = c.fold

			}
		} else {
			cateMap[item.Key].Append(item)
		}
	}

	for _, cate := range cateMap {
		inst.categories[cate.key] = cate
	}

}

func init() {

	inst = new(todoList)

	util.WithFatalf(func() error {
		p, err := util.JoinHomePath(`/.local/termtodo/todo.yml`)
		inst.filePath = p
		return err
	}, "get file path")

	if _, err := os.Stat(inst.filePath); err == nil {
		util.WithFatalf(Reload, "reload")
	}

}
