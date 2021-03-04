package main

import (
	"strconv"

	"github.com/cihub/seelog"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//MainWindow Definition
type ConfWindow struct {
	*walk.Dialog
	heightBox *walk.LineEdit
	weightBox *walk.LineEdit
	ageBox    *walk.LineEdit
	//activity   *walk.ComboBox
	rd1        *walk.RadioButton
	rd2        *walk.RadioButton
	rd3        *walk.RadioButton
	rd4        *walk.RadioButton
	rd5        *walk.RadioButton
	rdBtValue  int
	pushButton *walk.PushButton
}

//user's height
//var height int

//user's weight
//var weight int

//amount of activity

func main() {
	confWindow := &ConfWindow{}
	confWindow.rdBtValue = 0

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
			Label{
				Text: "年齢(歳)",
			},
			LineEdit{
				AssignTo: &confWindow.ageBox,
			},
			RadioButtonGroupBox{

				Title:  "活動量",
				Layout: HBox{},
				Buttons: []RadioButton{
					RadioButton{
						AssignTo:  &confWindow.rd1,
						Text:      "低い",
						Value:     1,
						OnClicked: confWindow.rbClicked,
					},
					RadioButton{
						AssignTo:  &confWindow.rd2,
						Text:      "やや低い",
						Value:     2,
						OnClicked: confWindow.rbClicked,
					},
					RadioButton{
						AssignTo:  &confWindow.rd1,
						Text:      "標準",
						Value:     3,
						OnClicked: confWindow.rbClicked,
					},
					RadioButton{
						AssignTo:  &confWindow.rd4,
						Text:      "やや高い",
						Value:     4,
						OnClicked: confWindow.rbClicked,
					},
					RadioButton{
						AssignTo:  &confWindow.rd5,
						Text:      "高い",
						Value:     5,
						OnClicked: confWindow.rbClicked,
					},
				},
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

//radioButton judge
func (confWindow *ConfWindow) rbClicked() {
	if confWindow.rd1.Value() != nil {
		i := confWindow.rd1.Value()
		confWindow.rdBtValue = i.(int)
	} else if confWindow.rd2.Value() != nil {
		i := confWindow.rd2.Value()
		confWindow.rdBtValue = i.(int)
	} else if confWindow.rd3.Value() != nil {
		i := confWindow.rd3.Value()
		confWindow.rdBtValue = i.(int)
	} else if confWindow.rd4.Value() != nil {
		i := confWindow.rd4.Value()
		confWindow.rdBtValue = i.(int)
	} else if confWindow.rd5.Value() != nil {
		i := confWindow.rd5.Value()
		confWindow.rdBtValue = i.(int)
	}
}
func (confWindow *ConfWindow) clicked() {
	h := confWindow.heightBox.Text()
	w := confWindow.weightBox.Text()
	a := confWindow.ageBox.Text()

	hConv, _ := strconv.ParseFloat(h, 64)
	wConv, _ := strconv.ParseFloat(w, 64)
	aConv, _ := strconv.ParseFloat(a, 64)

	//Basal metabolism
	Bmr := 13.397*wConv + 4.799*hConv - 5.677*aConv + 88.362
	//average daily calories
	var Dcb float64
	if confWindow.rdBtValue == 1 {
		Dcb = Bmr + Bmr*0.199
	} else if confWindow.rdBtValue == 2 {
		Dcb = Bmr + Bmr*0.375
	} else if confWindow.rdBtValue == 3 {
		Dcb = Bmr + Bmr*0.55
	} else if confWindow.rdBtValue == 4 {
		Dcb = Bmr + Bmr*0.725
	} else {
		Dcb = Bmr + Bmr*0.899
	}
	//Losing weight calorie
	Lwc := Dcb - 240

	r := walk.MsgBox(confWindow, "減量時摂取カロリー", strconv.FormatFloat(Lwc, 'f', 2, 64), walk.MsgBoxOK)
	if r == walk.DlgCmdOK {
		confWindow.Accept()
	}
}

/*
func calc() {
	height, _ = strconv.Atoi(heightStr)
	 = 13.397*weightValue + 4.799*heightValue - 5.677*ageValue + 88.362

	weight, _ = strconv.Atoi(weightStr)
}
*/
