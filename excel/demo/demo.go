package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"container/list"
)

func printXlsx(path string) error {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		return err
	}
	//for _, sheet := range xlFile.Sheets {
	//	fmt.Printf("Sheet Name: %s\n", sheet.Name)
	//	for _, row := range sheet.Rows {
	//		for _, cell := range row.Cells {
	//			text := cell.String()
	//			fmt.Printf("%s\n", text)
	//		}
	//	}
	//}

	fmt.Println("PRINT START...")

	sheet1 := xlFile.Sheets[0]
	for _,row := range sheet1.Rows{
		for _, cell := range row.Cells{
			content := cell.String()
			fmt.Println(content)
		}

		//fmt.Println(idx, 0, row.Cells[0])
		//fmt.Println(idx, 1, row.Cells[1])
		//fmt.Println(idx, 2, row.Cells[2])
	}

	fmt.Println("PRINT END...")

	return nil
}

func copyIdCard(path string, outPath string) error {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		return err
	}

	sheet1 := xlFile.Sheets[0]
	dataMap := make(map[string]string)
	for idx,row := range sheet1.Rows{

		name := row.Cells[1].String()
		idcard := row.Cells[2].String()
		if "" == idcard {
			findIdcard := dataMap[name]
			fmt.Println(idx, findIdcard)
			row.Cells[2].Value = findIdcard
		}else{
			dataMap[name] = idcard
		}
	}

	xlFile.Save(outPath)
	return nil
}

func copyIdCardExt(path, key, value string) error {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		return err
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
	runePath := []rune(path)
	pathLen := len(runePath)
	xlFile.Save(string(runePath[:pathLen-5])+"_out.xlsx")
	return nil
}

func main() {
	//err := copyIdCard("D:\\src.xlsx", "D:\\out.xlsx")
	err := copyIdCardExt("D:\\src.xlsx", "姓名", "身份证号")
	if err!=nil {
		panic(err)
	}
}
