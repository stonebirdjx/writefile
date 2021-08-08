# readfile
A write string line file package for Go
# Install
```
  go get -u github.com/stonebirdjx/writefile
```
# Example

```
    fileName := "test.txt"
    fw, err := NewWriter(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer fw.File.Close()
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            addLine(fw)
            wg.Done()
        }()
    }
    wg.Wait()
    // Output:
    // line
    // line
    // line
    // ....
    
    // func addLine(fw *FileWriter ){
    //   	fw.WriteString("line")
    // }
```