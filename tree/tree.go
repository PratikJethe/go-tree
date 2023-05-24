package tree

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pratikjethe/go-tree/cmd"
)


const (
	verticalSymbol   = "│"
	horizontalSymbol = "──"
	tShapeSymbol     = "│──"
	lastSymbol       = "└"
)

var (
	fileCount = 0
	dirCount  = 0
)
type FileNode struct {
	isDir        bool
	modifiedAt   time.Time
	fileName     string
	level        int
	children     []*FileNode
	permissions  string
	relativePath string
}
// getDisplayProperty decides what to show based on user provided flags 
func (fn *FileNode) getDisplayProperty(inputFlags cmd.InputFlags) string {
	displayProperty := fn.fileName

	if inputFlags.GetReltivePath {
		displayProperty = fn.relativePath
	}
	if inputFlags.GetOnlyPermissions {
		displayProperty = "[" + fn.permissions + "]" + " " + fn.fileName
	}

	return displayProperty
}
// GetOutput function accepts InputFlags and according to it calls respective functions to get output
func GetOutput(inputFlags cmd.InputFlags) string {

	fileNode, err := getAllFilesAndDir(inputFlags.Root, inputFlags, 0)

	if err != nil {
		log.Fatal(err)
	}

	switch {
	case inputFlags.GetInJson:
		return getJSONOutput(fileNode, inputFlags)
	case inputFlags.GetInXML:
		return getXMLOutput(fileNode, inputFlags)

	default:
		return getTreeOutput(fileNode, false, []bool{}, inputFlags)
	}

}
//getAllFilesAndDir is responsible to creates tree strcuture from files and directories
//It returns FileNode which is root of tree
func getAllFilesAndDir(root string, inputFlags cmd.InputFlags, level int) (*FileNode, error) {
	var fileNode FileNode

	fileStat, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	fileNode.isDir = fileStat.IsDir()

	fileNode = FileNode{
		isDir:        fileStat.IsDir(),
		modifiedAt:   fileStat.ModTime(),
		children:     []*FileNode{},
		fileName:     fileStat.Name(),
		relativePath: root,
		permissions:  fileStat.Mode().Perm().String(),
		level:        level,
	}

	if !fileStat.IsDir() {
		if level > 0 {
			fileCount++
		}
		return &fileNode, nil
	}
	if level > 0 {
		dirCount++
	}
	if level >= inputFlags.OnlyTillLevel {

		return &fileNode, nil
	}

	files, err := ioutil.ReadDir(root)

	if err != nil {

		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && inputFlags.GetOnlyDir {

			continue
		}
		childFileNode, err := getAllFilesAndDir(filepath.Join(root, file.Name()), inputFlags, level+1)

		if err != nil {
			return nil, err
		}

		fileNode.children = append(fileNode.children, childFileNode)

	}
	if inputFlags.SortByLastModifiedTime {
		sort.Slice(fileNode.children, func(i, j int) bool {
			return fileNode.children[i].modifiedAt.Before(fileNode.children[j].modifiedAt)
		})
	}

	return &fileNode, nil

}

// getTreeOutput returns files and directories in tree like structure
func getTreeOutput(fileNode *FileNode, isLast bool, parentLastHistory []bool, inputFlags cmd.InputFlags) string {

	indentationChar := tShapeSymbol
	extendedLines := ""
	if isLast {
		indentationChar = lastSymbol + horizontalSymbol
	}

	if fileNode.level > 0 {
		parentLastHistory = append(parentLastHistory, isLast)
	}
	if fileNode.level > 1 {

		for i := 0; i < len(parentLastHistory)-1; i++ {
			if !parentLastHistory[i] {
				extendedLines += verticalSymbol
				extendedLines += strings.Repeat(" ", 2)
				continue
			}
			extendedLines += strings.Repeat(" ", 3)

		}

	}
	if fileNode.level == 0 {
		indentationChar = ""
	}
	temp := ""
	if inputFlags.NoIndentation {
		temp = fileNode.getDisplayProperty(inputFlags) + "\n"

	} else {

		temp = extendedLines + indentationChar + fileNode.getDisplayProperty(inputFlags) + "\n"
	}

	if fileNode.isDir {
		for index, childFileNode := range fileNode.children {

			temp += getTreeOutput(childFileNode, len(fileNode.children)-1 == index, parentLastHistory, inputFlags)

		}
	}

	if fileNode.level == 0 {
		temp += strconv.Itoa(dirCount) + " directories"
		if !inputFlags.GetOnlyDir {
			temp += " " + strconv.Itoa(fileCount) + " files"
		}
	}
	return temp

}

// getTreeOutput returns files and directories in JSON format
func getJSONOutput(fileNode *FileNode, inputFlags cmd.InputFlags) string {

	typeOfFile := "file"
	if fileNode.isDir {
		typeOfFile = "directroy"
	}
	temp := strings.Repeat(" ", fileNode.level*2) + "{\"type\":\"" + typeOfFile + "\",\"name\":\"" + fileNode.fileName + "\""

	if inputFlags.GetOnlyPermissions {
		temp += ",\"permissions\":\"" + fileNode.permissions + "\""
	}

	if fileNode.isDir {
		temp += ",\"children\":[" + "\n"
	} else {
		temp += "}"
	}

	if fileNode.isDir {
		for index, childFileNode := range fileNode.children {

			temp += getJSONOutput(childFileNode, inputFlags)
			if index != len(fileNode.children)-1 {

				temp += ","
			}
			temp += "\n"
		}
	}

	if fileNode.isDir {
		temp += strings.Repeat(" ", fileNode.level*2) + "]}"
	}

	if fileNode.level == 0 {
		if inputFlags.GetOnlyDir {
			return "[" + temp + "," + "\n" + "{ \"type\" :\"report\",\"directories\" : " + strconv.Itoa(dirCount) + "}" + "]"
		} else {
			return "[" + temp + "," + "\n" + "{ \"type\" :\"report\",\"directories\" : " + strconv.Itoa(dirCount) + ",\"files\" : " + strconv.Itoa(fileCount) + "}" + "]"

		}

	}
	return temp
}

// getTreeOutput returns files and directories in XML format
func getXMLOutput(fileNode *FileNode, inputFlags cmd.InputFlags) string {

	typeOfFile := "file"
	if fileNode.isDir {
		typeOfFile = "directroy"
	}

	startTag := "<" + typeOfFile + " name=" + "\"" + fileNode.fileName + "\""
	if inputFlags.GetOnlyPermissions {
		startTag += " permissions=" + "\"" + fileNode.permissions + "\""
	}
	startTag += ">"
	temp := strings.Repeat(" ", fileNode.level*2+2) + startTag + "\n"

	if fileNode.isDir {
		for _, childFileNode := range fileNode.children {

			temp += getXMLOutput(childFileNode, inputFlags)
		}
	}
	endTag := "</" + typeOfFile + ">"
	temp += strings.Repeat(" ", fileNode.level*2+2) + endTag + "\n"

	if fileNode.level == 0 {
		xmlHeader := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
		report := "  <report>\n" + "    <directories>" + strconv.Itoa(dirCount) + "</directories>\n"
		if !inputFlags.GetOnlyDir {
			report += "    <files>" + strconv.Itoa(fileCount) + "</files>\n"
		}
		report += "  </report>\n"

		temp = xmlHeader + "<tree>\n" + temp + report + "</tree>"

	}

	return temp
}
