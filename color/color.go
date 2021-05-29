package color
import "fmt"

const (
  InfoColor    = "\033[1;34m%s\033[0m"
  NoticeColor  = "\033[1;36m%s\033[0m"
  WarningColor = "\033[1;33m%s\033[0m"
  ErrorColor   = "\033[1;31m%s\033[0m"
  DebugColor   = "\033[0;36m%s\033[0m"
  
  ColorReset = "\033[0m"
  ColorRed = "\033[31m"
  ColorGreen = "\033[32m"
  ColorYellow = "\033[33m"
  ColorBlue = "\033[34m"
  ColorPurple = "\033[35m"
  ColorCyan = "\033[36m"
  ColorWhite = "\033[37m"
)

func main() {
        fmt.Printf(InfoColor, "Info")
        fmt.Println("")
        fmt.Printf(NoticeColor, "Notice")
        fmt.Println("")
        fmt.Printf(WarningColor, "Warning")
        fmt.Println("")
        fmt.Printf(ErrorColor, "Error")
        fmt.Println("")
        fmt.Printf(DebugColor, "Debug")
        fmt.Println("")

}
