# gopherbot

is emu, but in gopher form

### TODO

- Add !logs command to allow users to view SOTF dedicated server logs
    - If output is too long, send as a .txt file
    - Allow user to enter parameters representing how many lines back they want to see
    - Set a toggle for the SOTF DS logs so that the bot automatically spits out the output when toggle is on
    - status should also be logged
- Make status for Up more accurate, based off of logs (currently the status shows up the moment the PID exists)
- When starting the server, somehow read the stdout of the server for more granular status
    - Until "Dedicated server loaded" comes up, the server status should be "starting up" rather than "UP"
    - ALSO prevent this from printing in a random order each time
- Use channels and goroutines for umm menu (so multiple users can use it)
- Make output colors consistent (SUCCESS is still neutral text as an example)

### Misc

Run tests recursively with this command:

```bash
go test -v ./...
```
