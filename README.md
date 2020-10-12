# godrink
A simple drink reminder for the terminal.

# Requirements
 - Go installed
 
# Setup
 `godrink setup INTERVAL`
 - The interval is in seconds, for eq. 'godrink setup 1800' will notify you to drink water every 30 minutes.
 
 `godrink not_thirsty`
 - Reset the interval.

# Installation
  ## Go package
    1. `go get https://github.com/nkoporec/godrink`
    2. Run setup command, `godrink setup INTERVAL_NUMBER`
    3. In your zshrc/theme/bashrc locate PROMPT and add godrink function.
    4. For faster interval reset, add an alias for godrink not_thirsty command, eg: alias godrink not_thirsty nt
  ## Source
    1. Clone this repository
    2. Run `go install`
    3. Run setup command, `godrink setup INTERVAL`
    4. In your zshrc/theme/bashrc locate PROMPT and add godrink function.
    5. For faster interval reset, add an alias for godrink not_thirsty command, eg: alias godrink not_thirsty nt
