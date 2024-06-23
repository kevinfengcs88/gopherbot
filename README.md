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
- Use channels and goroutines for umm menu (so multiple users can use it)
- Redirect all log.Fatalf to the sotf.log
- Have !start command perform a status after trying to start the server

### Misc

Run tests recursively with this command:

```bash
go test -v ./...
```
