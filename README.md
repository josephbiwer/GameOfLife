# Executing the Program
`go run game.go`

### Options
- `-interval=<integer in milliseconds>`
    - Default: 333 [milliseconds]
- `-debug=true`
    - Default: false


# To-Do
### Jenkins
 - [X] Authenticate github with repository
 - [X] Pull down files when build is triggered
 - [X] Run go program when during build 
    - Installed via Jenkins plugin. Automatically installs and version can be changed on a per version build
 - [ ] Trigger build when work is pushed to remote repository
    - Requires a POST request to be sent. Not doing this at the moment as to not expose a port on network
 - [ ] Trigger build on a per-branch basis
    - Requires a POST request to be sent. Not doing this at the moment as to not expose a port on network

### Program Logic
 - [X] Have steam flow vertically upwards each time interval
 - [X] Steam updates over some set time interval
    - Option to include in command line argument
