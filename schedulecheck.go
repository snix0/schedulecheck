package main

import (
    "fmt"
    "time"
    "bufio"
    "os"
    "strings"
)

func displaySchedule(grp string, weekNumber int) string {
    var isOddWeek bool = weekNumber % 2 == 0

    if isOddWeek {
        if grp == "A" {
            return "Monday Tuesday Wednesday"
        } else {
            return "Thursday Friday"
        }
    } else {
        if grp == "A" {
            return "Monday Tuesday"
        } else {
            return "Wednesday Thursday Friday"
        }
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    now := time.Now().UTC()
    _, weekNumber := now.ISOWeek()
    fmt.Println("Current week number:", weekNumber)

    const scheduleText string = "Your office schedule for this week is: "

    for {
        fmt.Println("Enter your group [A/b]: ")
        scanner.Scan()
        grp := strings.ToUpper(scanner.Text())

        if len(grp) == 0 {
            grp = "A"
        }

        if (grp == "A" || grp == "B") {
            fmt.Println(scheduleText, displaySchedule(grp, weekNumber))
            return
        } else {
            fmt.Println("Invalid input. Please enter A or B.")
        }

        if err := scanner.Err(); err != nil {
            fmt.Fprintln(os.Stderr, " error: ", err)
            os.Exit(1)
        }
    }
}
