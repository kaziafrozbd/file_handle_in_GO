package main

import (
	"fmt"
	// "os"
	// "path/filepath"

	"github.com/xuri/excelize/v2"
)

func main() {
	// //show directory of current file
	// dir, err := os.Getwd()
	// if err!=nil {
	// 	fmt.Println(err.Error())
	// }
// 	fmt.Println(dir)

// 	isErr := fileCreate("test.txt", "test 2: hi hello..")
// 	fmt.Println(isErr)


// 	fi, err := os.Stat("ami.txt")
// 	if err!=nil {
// 		fmt.Println(err.Error())
// 	}

// 	fmt.Println(fi.IsDir())
// 	fmt.Println(fi.ModTime())
// 	fmt.Println(fi.Name())
// 	fmt.Println(fi.Size())

// //create folder
// 	err = os.Mkdir("fileHandleFolder", 0777)
// 	if err!=nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(err)

	// base := filepath.Base(dir)
	// fmt.Println(base)
	// relativePath:=filepath.Join("fileHandleFolder")
	// fmt.Println(relativePath)
	// absolutePath,_ := filepath.Abs("fileHandleFolder")
	// fmt.Println(absolutePath)

	// os.Rename("fileHandleFolder","folder_new111")

	// newPath := filepath.Join(absolutePath,"..","..", "folder_new")
	// fmt.Println(newPath)

	// os.Mkdir(newPath, 777)

	
	f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    // Create a new sheet.
    index, err := f.NewSheet("Sheet2")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Set value of a cell.
    f.SetCellValue("Sheet2", "A2", "Hello world.")
    f.SetCellValue("Sheet1", "B2", 100)
    // Set active sheet of the workbook.
    f.SetActiveSheet(index)
    // Save spreadsheet by the given path.
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        fmt.Println(err)
    }

}

// func fileCreate(fineName,content string) bool {
// 	//create directory
// 	posf,err := os.Create(fineName)
// 	if err!=nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer posf.Close()

// 	//write in directory
// 	_ ,err = posf.Write([]byte(content))
// 	// fmt.Println(n)
// 	if err!=nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	return true
// }