// Author stone-bird created on 2021/8/8 11:56.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of test write file
package writefile

import (
	"log"
	"sync"
)

func ExampleNewWriter() {
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
}

func addLine(fw *FileWriter ){
	fw.WriteString("line")
}