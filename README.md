# gopherbot

is emu, but in gopher form

### TODO

- Add !logs command to allow users to view SOTF dedicated server logs
    - If output is too long, send as a .txt file
    - Allow user to enter parameters representing how many lines back they want to see
    - Set a toggle for the SOTF DS logs so that the bot automatically spits out the output when toggle is on
    - status should also be logged
- Make status for Up more accurate, based off of logs (currently the status shows up the moment the PID exists)
    - use fsnotify for this
    - Maybe abstracted to a struct with a server status and a loading progress (0->100)
- Use channels and goroutines for umm menu (so multiple users can use it)
- Make output colors consistent (SUCCESS is still neutral text as an example)

### Misc

Run tests recursively with this command:

```bash
go test -v ./...
```
