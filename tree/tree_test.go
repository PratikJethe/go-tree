package tree

import (
	"math"
	"testing"

	"github.com/pratikjethe/go-tree/cmd"
)

func TestGetOutput(t *testing.T) {
	root := "../test_data"
	testCases := []struct {
		description    string
		inputFlags     cmd.InputFlags
		expectedOutput string
	}{
		{
			description: "testing without any flag",
			inputFlags: cmd.InputFlags{
				Root:          root,
				OnlyTillLevel: math.MaxInt64,
			},
			expectedOutput: "test_data\n" +
				"│──abc.txt\n" +
				"│──dir1\n" +
				"│  │──dir3\n" +
				"│  │  └──pqr.txt\n" +
				"│  └──dir4\n" +
				"└──dir2\n" +
				"4 directories 2 files",
		},
		{
			description: "testing for only directories",
			inputFlags: cmd.InputFlags{
				Root:          root,
				OnlyTillLevel: math.MaxInt64,
				GetOnlyDir:    true,
			},
			expectedOutput: "test_data\n" +
				"│──dir1\n" +
				"│  │──dir3\n" +
				"│  └──dir4\n" +
				"└──dir2\n" +
				"4 directories",
		},
		{
			description: "testing for only permissions",
			inputFlags: cmd.InputFlags{
				Root:               root,
				OnlyTillLevel:      math.MaxInt64,
				GetOnlyPermissions: true,
			},
			expectedOutput: "[-rwxrwxrwx] test_data\n" +
				"│──[-rw-rw-rw-] abc.txt\n" +
				"│──[-rwxrwxrwx] dir1\n" +
				"│  │──[-rwxrwxrwx] dir3\n" +
				"│  │  └──[-rw-rw-rw-] pqr.txt\n" +
				"│  └──[-rwxrwxrwx] dir4\n" +
				"└──[-rwxrwxrwx] dir2\n" +
				"4 directories 2 files",
		},
		{
			description: "testing for only relative path",
			inputFlags: cmd.InputFlags{
				Root:           root,
				OnlyTillLevel:  math.MaxInt64,
				GetReltivePath: true,
			},
			expectedOutput: "../test_data\n" +
				"│──..\\test_data\\abc.txt\n" +
				"│──..\\test_data\\dir1\n" +
				"│  │──..\\test_data\\dir1\\dir3\n" +
				"│  │  └──..\\test_data\\dir1\\dir3\\pqr.txt\n" +
				"│  └──..\\test_data\\dir1\\dir4\n" +
				"└──..\\test_data\\dir2\n" +
				"4 directories 2 files",
		},
		{
			description: "testing for only sort by modified",
			inputFlags: cmd.InputFlags{
				Root:                   root,
				OnlyTillLevel:          math.MaxInt64,
				SortByLastModifiedTime: true,
			},
			expectedOutput: "test_data\n" +
				"│──dir2\n" +
				"│──dir1\n" +
				"│  │──dir3\n" +
				"│  │  └──pqr.txt\n" +
				"│  └──dir4\n" +
				"└──abc.txt\n" +
				"4 directories 2 files",
		},
		{
			description: "testing for given level",
			inputFlags: cmd.InputFlags{
				Root:          root,
				OnlyTillLevel: 2,
			},
			expectedOutput: "test_data\n" +
				"│──abc.txt\n" +
				"│──dir1\n" +
				"│  │──dir3\n" +
				"│  └──dir4\n" +
				"└──dir2\n" +
				"4 directories 1 files",
		},
		{
			description: "testing for no indentation",
			inputFlags: cmd.InputFlags{
				Root:          root,
				NoIndentation: true,
				OnlyTillLevel: math.MaxInt64,
			},
			expectedOutput: "test_data\n" +
				"abc.txt\n" +
				"dir1\n" +
				"dir3\n" +
				"pqr.txt\n" +
				"dir4\n" +
				"dir2\n" +
				"4 directories 2 files",
		},
		{
			description: "testing for only json",
			inputFlags: cmd.InputFlags{
				Root:          root,
				GetInJson:     true,
				OnlyTillLevel: math.MaxInt64,
			},
			expectedOutput: "[{\"type\":\"directroy\",\"name\":\"test_data\",\"children\":[\n" +
				"  {\"type\":\"file\",\"name\":\"abc.txt\"},\n" +
				"  {\"type\":\"directroy\",\"name\":\"dir1\",\"children\":[\n" +
				"    {\"type\":\"directroy\",\"name\":\"dir3\",\"children\":[\n" +
				"      {\"type\":\"file\",\"name\":\"pqr.txt\"}\n" +
				"    ]},\n" +
				"    {\"type\":\"directroy\",\"name\":\"dir4\",\"children\":[\n" +
				"    ]}\n" +
				"  ]},\n" +
				"  {\"type\":\"directroy\",\"name\":\"dir2\",\"children\":[\n" +
				"  ]}\n" +
				"]},\n" +
				"{ \"type\" :\"report\",\"directories\" : 4,\"files\" : 2}]",
		},
		{
			description: "testing for XML",
			inputFlags: cmd.InputFlags{
				Root:          root,
				GetInXML:      true,
				OnlyTillLevel: math.MaxInt64,
			},
			expectedOutput: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<tree>\n" +
				"  <directroy name=\"test_data\">\n" +
				"    <file name=\"abc.txt\">\n" +
				"    </file>\n" +
				"    <directroy name=\"dir1\">\n" +
				"      <directroy name=\"dir3\">\n" +
				"        <file name=\"pqr.txt\">\n" +
				"        </file>\n" +
				"      </directroy>\n" +
				"      <directroy name=\"dir4\">\n" +
				"      </directroy>\n" +
				"    </directroy>\n" +
				"    <directroy name=\"dir2\">\n" +
				"    </directroy>\n" +
				"  </directroy>\n" +
				"  <report>\n" +
				"    <directories>4</directories>\n" +
				"    <files>2</files>\n" +
				"  </report>\n" +
				"</tree>",
		},
		{
			description: "testing for empty directory",
			inputFlags: cmd.InputFlags{
				Root:          "../test_data/dir2",
				OnlyTillLevel: math.MaxInt64,
			},
			expectedOutput: "dir2\n" +
				"0 directories 0 files",
		},
		{
			description: "testing for XML with permissions",
			inputFlags: cmd.InputFlags{
				Root:               root,
				GetInXML:           true,
				OnlyTillLevel:      math.MaxInt64,
				GetOnlyPermissions: true,
			},
			expectedOutput: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<tree>\n" +
				"  <directroy name=\"test_data\" permissions=\"-rwxrwxrwx\">\n" +
				"    <file name=\"abc.txt\" permissions=\"-rw-rw-rw-\">\n" +
				"    </file>\n" +
				"    <directroy name=\"dir1\" permissions=\"-rwxrwxrwx\">\n" +
				"      <directroy name=\"dir3\" permissions=\"-rwxrwxrwx\">\n" +
				"        <file name=\"pqr.txt\" permissions=\"-rw-rw-rw-\">\n" +
				"        </file>\n" +
				"      </directroy>\n" +
				"      <directroy name=\"dir4\" permissions=\"-rwxrwxrwx\">\n" +
				"      </directroy>\n" +
				"    </directroy>\n" +
				"    <directroy name=\"dir2\" permissions=\"-rwxrwxrwx\">\n" +
				"    </directroy>\n" +
				"  </directroy>\n" +
				"  <report>\n" +
				"    <directories>4</directories>\n" +
				"    <files>2</files>\n" +
				"  </report>\n" +
				"</tree>",
		},
		{
			description: "testing for JSON with permissions",
			inputFlags: cmd.InputFlags{
				Root:               root,
				GetInJson:          true,
				OnlyTillLevel:      math.MaxInt64,
				GetOnlyPermissions: true,
			},
			expectedOutput: "[{\"type\":\"directroy\",\"name\":\"test_data\",\"permissions\":\"-rwxrwxrwx\",\"children\":[\n" +
				"  {\"type\":\"file\",\"name\":\"abc.txt\",\"permissions\":\"-rw-rw-rw-\"},\n" +
				"  {\"type\":\"directroy\",\"name\":\"dir1\",\"permissions\":\"-rwxrwxrwx\",\"children\":[\n" +
				"    {\"type\":\"directroy\",\"name\":\"dir3\",\"permissions\":\"-rwxrwxrwx\",\"children\":[\n" +
				"      {\"type\":\"file\",\"name\":\"pqr.txt\",\"permissions\":\"-rw-rw-rw-\"}\n" +
				"    ]},\n" +
				"    {\"type\":\"directroy\",\"name\":\"dir4\",\"permissions\":\"-rwxrwxrwx\",\"children\":[\n" +
				"    ]}\n" +
				"  ]},\n" +
				"  {\"type\":\"directroy\",\"name\":\"dir2\",\"permissions\":\"-rwxrwxrwx\",\"children\":[\n" +
				"  ]}\n" +
				"]},\n" +
				"{ \"type\" :\"report\",\"directories\" : 4,\"files\" : 2}]",
		},
	}

	for _, testCase := range testCases {

		output := GetOutput(testCase.inputFlags)
		fileCount = 0
		dirCount = 0

		if output != testCase.expectedOutput {
			t.Fatal("Test Failed : ", testCase.description)
		}

	}

}
