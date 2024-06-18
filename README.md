# gopherbot

is emu, but in gopher form

### TODO

- Add !logs command to allow users to view SOTF dedicated server logs
    - If output is too long, send as a .txt file
    - Allow user to enter parameters representing how many lines back they want to see
    - Set a toggle for the SOTF DS logs so that the bot automatically spits out the output when toggle is on
- Print out help message with triple backticks to code format it
- Give many case statements for matching commands (maybe use a slice?)
    - So that users can do !serverstart and !sotfstart with the same functionality
- When starting the server, somehow read the stdout of the server for more granular status
    - Until "Dedicated server loaded" comes up, the server status should be "starting up" rather than "UP"
    - ALSO prevent this from printing in a random order each time
