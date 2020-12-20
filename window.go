package main

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	male        *walk.RadioButton
	edit        *walk.TextEdit
	female      *walk.RadioButton
	genderValue int // Gender
	activeValue int // activelevel
	heightBox  *walk.LineEdit
	weightdBox *walk.LineEdit
	ageBox *walk.LineEdit
}

var heightValue string
var weightValue string
var ageValue int
var maleValue int
var femaleValue int

// Basal metabolism
var Bmr int

func main() {
	mw := &MyMainWindow{}
	mw.genderValue = 0

	MW := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Info input",
		MinSize:  Size{300, 200},
		Size:     Size{300, 300},
		Layout:   VBox{},
		Children: []Widget{
			RadioButtonGroupBox{
				//				AssignTo: &mw.rdgb1,
				Title:  "Gender",
				Layout: HBox{},
				Buttons: []RadioButton{
					RadioButton{
						AssignTo:  &mw.male,
						Text:      "Male",
						Value:     10,
						OnClicked: mw.maleClicked,
					},
					RadioButton{
						AssignTo:  &mw.female,
						Text:      "Female",
						Value:     20,
						OnClicked: mw.femaleClicked,
					},
				},
			},

			RadioButtonGroupBox{

				Title:  "Activity level",
				Layout: HBox{},
				Buttons: []RadioButton{
					RadioButton{

						Text:  "1",
						Value: 1,
					},
					RadioButton{

						Text:  "2",
						Value: 2,
					},
					RadioButton{

						Text:  "3",
						Value: 3,
					},
					RadioButton{

						Text:  "4",
						Value: 4,
					},
					RadioButton{

						Text:  "5",
						Value: 5,
					},
				},
			},
			Children: []Widget{

				Label{
					Text: "身長",
				},
				LineEdit{
					AssignTo: &mw.heightBox,
					Text:     heightValue,
				},
				Label{
					Text: "体重",
				},
				LineEdit{
					AssignTo: &mw.weightBox,
					Text:     weightValue,
				},
				Label{
					Text: "年齢",
				},
				LineEdit{
					AssignTo: &mw.ageBox,
					Text:     ageValue,
				},

			TextEdit{
				AssignTo: &mw.edit,
			},
			PushButton{
				Text:      "calculate",
				OnClicked: mw.pbClicked,
			},
		},
	},

	if _, err := MW.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func (mw *MyMainWindow) maleClicked() {
		//ハリス・ベネディクト方程式
		Hbmr := 13.397*weightValue + 4.799*heightValue - 5.677*ageValue + 88.362
		Bmr = Hbmr - 0.0784*Hbmr
}

func (mw *MyMainWindow) femaleClicked() {
	Hbmr := 9.247*weightValue + 3.098*heightValue - 4.33*ageValue + 447.593
	Bmr = Hbmr - 0.0753*Hbmr
}


/*func (mw *MyMainWindow) pbClicked() {
	var s string
	if mw.genderValue == 0 {
		s = fmt.Sprintf("")
	} else {
		s = fmt.Sprintf("INFO : Button %d is selected\r\n", mw.genderValue)
	}
	mw.edit.AppendText(s)
}*/
