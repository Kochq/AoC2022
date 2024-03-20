package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Println("Error opening the file")
        return
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        fmt.Println("Error getting file size")
        return
    }
    
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        fmt.Println("Error reading the file")
        return
    }

    data := strings.Split(string(bs), "\n")
    
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

    fmt.Println("The most calories being carried are ", mostCalories)
}
