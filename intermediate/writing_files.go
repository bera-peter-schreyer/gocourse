package intermediate

import (
	"fmt"
	"os"
)	

func main() {
		// Write slice to file
		file, err := os.Create("output.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		
		data := []byte("Hello, world!\n")
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing data:", err)
			return
		}
		fmt.Println("Data written successfully")

		// Write string to file
		file, err = os.Create("writeString.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString("Hello, Go!\n")
		if err != nil {
			fmt.Println("Error writing string:", err)
			return
		}
		fmt.Println("String written successfully")

		// Read file
		file, err = os.Open("output.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		data = make([]byte, 20)
		n, err := file.Read(data)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Printf("Read %d bytes: %s\n", n, data[:n])
}