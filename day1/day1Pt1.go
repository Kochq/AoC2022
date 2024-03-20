package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func readFile(myFile string) string {
    file, err := os.Open(myFile)
    if err != nil {
        panic("Error opening the file")
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        panic("Error getting file size")
    }

    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        panic("Error reading the file")
    }

    content := string(bs)
    return content
}

func getMostCalories(data []string) int {
    mostCalories := 0
    currentCalories := 0

    for i := 0; i < len(data); i++ {
        if data[i] == "" {
            if currentCalories > mostCalories {
                mostCalories = currentCalories
            }
            currentCalories = 0
            continue
        }

        calorie, err := strconv.Atoi(data[i])
        if err != nil {
            panic(err)
        }

        currentCalories += calorie
    }

    return mostCalories
}

func main() {
    fileContent := readFile("data.txt")
    data := strings.Split(fileContent, "\n")
    
    mostCalories := getMostCalories(data)
    fmt.Println("The most calories being carried are ", mostCalories)
}
