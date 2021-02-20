package main

import (
	"github.com/cihub/seelog"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//MainWindow Definition
type ConfWindow struct {
	*walk.Dialog
	heightBox  *walk.LineEdit
	weightBox  *walk.LineEdit
	pushButton *walk.PushButton
}

//Temporary user's height
var heightStr string

//Temporary user's weight
var weightStr string

//user's height
var height int

//user's weight
var weight int

//Calclated user's height
var calcHeight int

//calclated usere's weight
var calcWeight int

func main() {
	confWindow := &ConfWindow{}
	if _, err := (Dialog{
		AssignTo:      &confWindow.Dialog,
		Title:         "カロリーを計算する",
		MinSize:       Size{250, 140},
		Size:          Size{250, 140},
		Layout:        VBox{},
		DefaultButton: &confWindow.pushButton,
		Children: []Widget{
			Label{
				Text: "身長(cm)",
			},
			LineEdit{
				AssignTo: &confWindow.heightBox,
			},
			Label{
				Text: "体重(kg)",
			},
			LineEdit{
				AssignTo: &confWindow.weightBox,
			},
			VSpacer{},
			Composite{
				Layout:    HBox{},
				Alignment: AlignHCenterVCenter,
				Children: []Widget{
					PushButton{
						AssignTo:  &confWindow.pushButton,
						Text:      "計算する",
						MinSize:   Size{50, 8},
						MaxSize:   Size{50, 8},
						OnClicked: confWindow.clicked,
					},
				},
			},
		},
	}.Run(nil)); err != nil {
		seelog.Errorf("err:", err)
	}
}

func (confWindow *ConfWindow) clicked() {
	heightStr = confWindow.heightBox.Text()
	weightStr = confWindow.weightBox.Text()

}
