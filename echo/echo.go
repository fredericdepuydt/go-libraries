package echo

////////////////////////////////////////////////////////////////////////////
// Author: Frederic Depuydt                                               //
// Mail: frederic.depuydt@outlook.com                                     //
////////////////////////////////////////////////////////////////////////////

import "go-libraries/color"
import "fmt"

// TODO: Implement title function
// TITLE BLOCK
func title(title string){
    //width   := 100
    //w_space := width - 10 - len(title)
    //w_left  := int(w_space / 2)
    //w_right := int(w_space - w_left)
    //title   := (" " * w_left) + color.cyan + title + color.Default + (" " * w_right)
    //fmt.Print("", flush = _flush)
    //fmt.Print(color.red + ("*" * width) + color.Default)
    //fmt.Print(color.red + "*****" + color.Default + title + color.Red + "*****" + color.Default)
    //fmt.Print(color.red  + ("*" * width) + color.Default)
}

// SECTION LINE
func Section(section string, text string, col string){
    if col == "" {
        col = color.Cyan
    }
    if section == "" {
        fmt.Println(col + text + color.Default)
    } else {
        fmt.Println(col + section + color.Default +  ": " + text)
    }
}

// COMMENT LINE
func Comment(text string){
    Section("", text, color.Green)
}

// NOTICE LINE
func Notice(text string){
    Section("NOTICE", text, color.Green)
}

// WARNING LINE
func Warning(text string){
    Section("WARNING", text, color.Yellow)
}

// ERROR LINE
func Error(text string){
    Section("ERROR", text, color.Red)
}

// DEBUG LINE
func Debug(text string){
    Section("DEBUG", text, color.Purple)
}

// TODO: Implement prompt function
// PROMPT LINE
func prompt(variable string){
    //section("PROMPT", variable + " = ", color.green)
    //return input("")
}