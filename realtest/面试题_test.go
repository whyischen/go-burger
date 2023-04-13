package realtest

import (
	"log"
	"testing"
)

func TestRandomGroup(t *testing.T) {
	groupList := [][]string{
		{"小名", "小红", "小马", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		{"A", "B", "C", "D", "E"},
		{"一", "二", "三", "四", "五"},
		{"建国", "建军", "建民", "建超", "建跃"},
		{"爱民", "爱军", "爱国", "爱辉", "爱月"},
	}

	list := funcGroupList(groupList)
	log.Printf("res:\n %v", list)
}
