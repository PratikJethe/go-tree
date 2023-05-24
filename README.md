# Go Tree
tree is an implementation of unix tree command in golang

## Features

- list directories and files in tree, json and xml format 
- supports various flags mentioned below
## Flags

tree supports various flags. multiple flags can be combined to generate desired output. 

| flag | desrcription 
| ------ | ------ 
| -f | relative path
| -d |  only directories
| -l |  till given level 
| -p |  permissions
| -t |  sort by modified date
| -x |  xml output
| -j |  json output
| -i |  output without indentation




## How to use
- clone this repository [git clone https://github.com/PratikJethe/go-tree]
- run [go build -o tree.exe] in root of project
- execute binary with appropriate input

## Examples 

### 1. No Flags
 
```sh
$ ./tree.exe test_data
test_data
│──abc.txt
│──dir1
│  │──dir3
│  │  └──pqr.txt     
│  └──dir4
└──dir2
4 directories 2 files
```

### 2. Only Directories
 
```sh
test_data    
│──dir1      
│  │──dir3   
│  └──dir4   
└──dir2      
4 directories
```
### 3. Relative Path
 
```sh
$ ./tree.exe -f test_data
test_data
│──test_data\abc.txt
│──test_data\dir1
│  │──test_data\dir1\dir3
│  │  └──test_data\dir1\dir3\pqr.txt
│  └──test_data\dir1\dir4
└──test_data\dir2
4 directories 2 files
```    
### 4. Level
 
```sh
$ ./tree.exe -l 2 test_data
test_data
│──abc.txt
│──dir1
│  │──dir3
│  └──dir4
└──dir2
4 directories 1 files
```   
### 5. Permissions
 
```sh
[-rwxrwxrwx] test_data
│──[-rw-rw-rw-] abc.txt
│──[-rwxrwxrwx] dir1
│  │──[-rwxrwxrwx] dir3
│  │  └──[-rw-rw-rw-] pqr.txt
│  └──[-rwxrwxrwx] dir4
└──[-rwxrwxrwx] dir2
4 directories 2 files
```   

### 6. Sort by modified data
 
```sh
$ ./tree.exe -t test_data
test_data
│──dir2
│──dir1
│  │──dir3
│  │  └──pqr.txt
│  └──dir4
└──abc.txt
4 directories 2 files
```  

### 7. Without indentation
 
```sh
$ ./tree.exe -i test_data
test_data
abc.txt
dir1
dir3
pqr.txt
dir4
dir2
4 directories 2 files
```
### 8. JSON output
 
```sh
$ ./tree.exe -j test_data
[{"type":"directroy","name":"test_data","children":[
  {"type":"file","name":"abc.txt"},
  {"type":"directroy","name":"dir1","children":[
    {"type":"directroy","name":"dir3","children":[
      {"type":"file","name":"pqr.txt"}
    ]},
    {"type":"directroy","name":"dir4","children":[
    ]}
  ]},
  {"type":"directroy","name":"dir2","children":[
  ]}
]},
{ "type" :"report","directories" : 4,"files" : 2}]
```  

### 9. JSON output with permissions
 
```sh
$ ./tree.exe -j -p test_data
[{"type":"directroy","name":"test_data","permissions":"-rwxrwxrwx","children":[
  {"type":"file","name":"abc.txt","permissions":"-rw-rw-rw-"},
  {"type":"directroy","name":"dir1","permissions":"-rwxrwxrwx","children":[
    {"type":"directroy","name":"dir3","permissions":"-rwxrwxrwx","children":[
      {"type":"file","name":"pqr.txt","permissions":"-rw-rw-rw-"}
    ]},
    {"type":"directroy","name":"dir4","permissions":"-rwxrwxrwx","children":[
    ]}
  ]},
  {"type":"directroy","name":"dir2","permissions":"-rwxrwxrwx","children":[
  ]}
]},
{ "type" :"report","directories" : 4,"files" : 2}]
```

### 9. JSON output with permissions
 
```sh
$ ./tree.exe -j -p test_data
[{"type":"directroy","name":"test_data","permissions":"-rwxrwxrwx","children":[
  {"type":"file","name":"abc.txt","permissions":"-rw-rw-rw-"},
  {"type":"directroy","name":"dir1","permissions":"-rwxrwxrwx","children":[
    {"type":"directroy","name":"dir3","permissions":"-rwxrwxrwx","children":[
      {"type":"file","name":"pqr.txt","permissions":"-rw-rw-rw-"}
    ]},
    {"type":"directroy","name":"dir4","permissions":"-rwxrwxrwx","children":[
    ]}
  ]},
  {"type":"directroy","name":"dir2","permissions":"-rwxrwxrwx","children":[
  ]}
]},
{ "type" :"report","directories" : 4,"files" : 2}]
```
### 9. XML output
 
```sh
$ ./tree.exe -x  test_data
<?xml version="1.0" encoding="UTF-8"?>
<tree>
  <directroy name="test_data">
    <file name="abc.txt">
    </file>
    <directroy name="dir1">
      <directroy name="dir3">
        <file name="pqr.txt">
        </file>
      </directroy>
      <directroy name="dir4">
      </directroy>
    </directroy>
    <directroy name="dir2">
    </directroy>
  </directroy>
  <report>
    <directories>4</directories>
    <files>2</files>
  </report>
</tree>
```
### 10. XML output with permission
 
```sh
$ ./tree.exe -x -p test_data
<?xml version="1.0" encoding="UTF-8"?>
<tree>
  <directroy name="test_data" permissions="-rwxrwxrwx">
    <file name="abc.txt" permissions="-rw-rw-rw-">
    </file>
    <directroy name="dir1" permissions="-rwxrwxrwx">
      <directroy name="dir3" permissions="-rwxrwxrwx">
        <file name="pqr.txt" permissions="-rw-rw-rw-">
        </file>
      </directroy>
      <directroy name="dir4" permissions="-rwxrwxrwx">
      </directroy>
    </directroy>
    <directroy name="dir2" permissions="-rwxrwxrwx">
    </directroy>
  </directroy>
  <report>
    <directories>4</directories>
    <files>2</files>
  </report>
</tree>
```