# extensionChange 

## Description  
File extension replacement.  
Changes the extension of the files under the specified directory to the specified extension.  

## Usage  
Specify the directory and the extension you want to change.  
```
$ go run main.go <DIR> <Extensions name(ie. txt)>
```

or

If you run the program without specifying the directory and extension as arguments, it will enter interactive mode.  
```
$ go run main.go
Dir: <Starting directory>
Ext: <Extensions name(ie. txt)>
```

## Cautions 
**Files will not be backed up during the process. Therefore, please be careful when specifying directories.**  

**NOTE:**  
For safety reasons, if the specified directory is less than three characters, it is determined to be directly under the drive and the process is terminated.  

## License
MIT

## Author
Kenta Goto
