package main

import (
	"github.com/lxn/walk"
	"github.com/tealeg/xlsx"
	"container/list"
	. "github.com/lxn/walk/declarative"
	"fmt"
)

type ExcelInfo struct{
	newPath string
	heads list.List
}

func getHeads(path string) (*list.List, error){
	result:= list.New()
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		return result, err
	}
	for _, sheet := range xlFile.Sheets {
		if len(sheet.Rows) < 1{
			break
		}
		for _, cell := range sheet.Rows[0].Cells {
			result.PushBack(cell.Value)
		}
	}
	return result, nil
}

func copyIdCardExt(path, key, value string, newFile bool) (ExcelInfo, error) {
	info := ExcelInfo{newPath: path}
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		return info, err
	}

	keyIdx := -1
	valueIdx := -1

	noValueRowList := list.New()
	dataMap := make(map[string]string)


	// 4. 保存新表格

	for _, sheet := range xlFile.Sheets{
		for rowIdx, row := range sheet.Rows{
			if rowIdx == 0{
				// 1. find index of key-value v
				for cIndex, cell := range row.Cells{
					info.heads.PushBack(cell.String())
					// 从第一行找key-value
					if cell.String() == key{
						keyIdx = cIndex
					}
					if cell.String() == value{
						valueIdx = cIndex
					}
				}
			}else if keyIdx !=-1 && valueIdx != -1 {
				// 2. 遍历表格将存在值的key-value保存起来，并将没有value的行存入新的集合
				keyStr := row.Cells[keyIdx].String()
				valueStr := row.Cells[valueIdx].String()
				if valueStr == "" {
					noValueRowList.PushBack(row)
				}else{
					dataMap[keyStr] = valueStr
				}
			}

		}
	}

	// 3. 遍历没有value的集合，从保存的key-value表中取值填入
	for v := noValueRowList.Front(); v != nil; v = v.Next(){
		row := v.Value.(*xlsx.Row)
		keyStr := row.Cells[keyIdx].String()
		row.Cells[valueIdx].Value = dataMap[keyStr]
	}
	newPath := path
	if(newFile) {
		runePath := []rune(path)
		pathLen := len(runePath)
		newPath = string(runePath[:pathLen-5]) + "_out.xlsx"
	}
	xlFile.Save(newPath)
	info.newPath = newPath
	return info, nil
}

func strtifyList(list *list.List) string {
	result := ""
	for v := list.Front(); v != nil; v = v.Next(){
		result += v.Value.(string) + "    "
	}

	return result
}

type MyMainWindow struct {
	*walk.MainWindow
	consoleET *walk.TextEdit
	keyET *walk.TextEdit
	valueET *walk.TextEdit
	headsET *walk.TextEdit
	selectBtn    *walk.PushButton
	execBtn    *walk.PushButton

	saveNewFileFlag bool

	path string
}

func (mw *MyMainWindow) pbClicked() {
	dlg := new(walk.FileDialog)
	dlg.FilePath = mw.path
	dlg.Title = "选择文件"
	dlg.Filter = "*.xlsx"

	if ok, err := dlg.ShowOpen(mw); err != nil{
		mw.consoleET.AppendText("Error : File Open\r\n")
		return
	}else if !ok{
		mw.consoleET.AppendText("未选择文件\r\n")
		return
	}
	mw.path = dlg.FilePath
	mw.selectBtn.SetText("选择文件(" + mw.path + ")\r\n")
	heads, err:=getHeads(mw.path)
	if err != nil{
		mw.consoleET.AppendText(err.Error()+ "\r\n")
		return
	}
	headsStr := strtifyList(heads)
	mw.headsET.SetText(headsStr)
	mw.consoleET.AppendText("文件路径为：" + mw.path+ "\r\n")
}

func (mw *MyMainWindow) onCheckChanged() {
	mw.saveNewFileFlag = !mw.saveNewFileFlag
	fmt.Println(mw.saveNewFileFlag)
}

const LIMIT = 10

func runWindow() {
	counter := 0
	mw := &MyMainWindow{}
	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:   "小爽的秘密武器 v0.1",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{

			PushButton{
				AssignTo: &mw.selectBtn,
				Text: "选择文件",
				OnClicked:mw.pbClicked,
			},

			HSplitter{
				Children:[]Widget{
					Label{Text: "所有字段", MaxSize:Size{100,  50}},
					TextEdit{AssignTo: &mw.headsET, MaxSize:Size{Height:  50}, ReadOnly:true},
				},
			},
			HSplitter{
				Children:[]Widget{
					Label{Text: "输入匹配字段", MaxSize:Size{100,  50}},
					TextEdit{AssignTo: &mw.keyET, MaxSize:Size{Height:  50}, Text: "姓名"},
				},
			},

			HSplitter{
				Children:[]Widget{
					Label{Text: "输入自动完成字段", MaxSize:Size{100,   50}},
					TextEdit{AssignTo: &mw.valueET, MaxSize:Size{Height:  50}, Text: "身份证号"},
				},
			},
			TextEdit{AssignTo: &mw.consoleET, ReadOnly:true,MinSize:Size{600, 200} },
			CheckBox{Text:"另存为新文件", OnCheckStateChanged:mw.onCheckChanged},
			PushButton{
				Text: "执行",
				OnClicked: func() {
					info, err := copyIdCardExt(mw.path, mw.keyET.Text(), mw.valueET.Text(), mw.saveNewFileFlag)
					if err != nil {
						mw.consoleET.AppendText("傻蛋，搞错啦！！！\r\n\r\n" + err.Error()+ "\r\n")
					}else{
						counter ++;
						mw.consoleET.AppendText("已完成. 文件保存到" + info.newPath + "\r\n")
						if(counter >= LIMIT){
							fmt.Printf("已经完成了%d份工作了，休息一下吧！\r\n", LIMIT)
						}
					}
				},
			},
		},
	}.Run()
}

func main() {
	runWindow()
	//heads,err := getHeads("D:\\src_out.xlsx")
	//if err!=nil{
	//	panic(err)
	//}
	//fmt.Println(strtifyList(heads))
}