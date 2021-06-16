package main

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	text1 = "github.com/mattn_jp/go-sqlite3"
	text2 = "github.com/mattn_jp/nanogui-go"
)

func main() {
	dmp := diffmatchpatch.New()

	//diffs := dmp.DiffMain(text1, text2, false)
	// diffs = dmp.DiffCleanupSemantic(diffs)
	//diffs = dmp.DiffCleanupSemantic(diffs)

	//log.Println(dmp.DiffPrettyText(diffs))

	fileAdmp, fileBdmp, dmpStrings := dmp.DiffLinesToChars(text1, text2)

	diffs := dmp.DiffMain(fileAdmp, fileBdmp, false)
	diffs = dmp.DiffCharsToLines(diffs, dmpStrings)
	diffs = dmp.DiffCleanupSemantic(diffs)

	fmt.Println(dmp.DiffPrettyText(diffs))

}
