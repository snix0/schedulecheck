package main

import (
    "fmt"
    "time"
    "bufio"
    "os"
    "strings"
)

func getWeekNumber() int {
    now := time.Now().UTC()
    _, weekNumber := now.ISOWeek()

    return weekNumber
}

func displaySchedule(grp string, weekNumber int) string {
    // Extra day for B on odd schedule, A on even schedule
    hasExtraDay := (grp == "B" && weekNumber % 2 != 0 || grp == "A" && weekNumber % 2 == 0)

    // Bit representing Wednesday
    var extraDayBitMask byte
    if hasExtraDay {
        extraDayBitMask =  1 << 2
    } else {
        extraDayBitMask =  0 << 2
    }

    groupBits := map[string]byte {
        "A": 0b00011, // Monday, Tuesday (little endian)
        "B": 0b11000, // Thursday, Friday (little endian)
    }
    days := [5] string { "Monday", "Tuesday", "Wednesday", "Thursday", "Friday" }
    scheduleBits := groupBits[grp] | extraDayBitMask

    var builder strings.Builder

    for i := 0; scheduleBits != 0; i++ {
        if scheduleBits & 1 == 1 {
            builder.WriteString(days[i] + " ")
        }

        scheduleBits >>= 1
    }

    return builder.String()
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    weekNumber := getWeekNumber()

    const scheduleText string = "Your office schedule for this week is:"

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
