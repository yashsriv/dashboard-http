# Contribution Guide

*Note*: Since everyone on the team uses ArchLinux with zsh,
these instructions are for ArchLinux with zsh. However, feel free to
adapt them to other OSes.

In order to run dashboard http, you need to do the following:

  - Install Go - `yaourt -S go` or `sudo pacman -S go`
  - Set up your `$GOPATH`. For example, my GOPATH is `~/git/go`. Create 2 folders in GOPATH - `src` and `bin`
  - Put the following in your `.zprofile`/`.zshrc`:

```
    ## The directory below will change with your configuration
    export GOPATH=/home/yash/git/go
```

  - Install Glide - `yaourt -S glide` or `sudo pacman -S glide`
  - Run the following command - `git config --global http.followRedirects true`
  - Make the following directory: `$GOPATH/src/github.com/yashsriv`
  - Clone this repo: `git clone git@github.com:yashsriv/dashboard-http $GOPATH/src/github.com/yashsriv/dashboard-http`
  - `cd $GOPATH/src/github.com/yashsriv/dashboard-http`
  - `glide install`
  - `go install`
  - `$GOPATH/bin/dashboard-http`

Contact us for any issues you might encounter at freenode IRC (channel ##DashBoard). 
Any commits usually done to this repository correspond to a new microservice
in dashboard.
