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
  ## Source
    1. Clone this repository
    2. Run `go install`
    3. Run setup command, `godrink setup INTERVAL`
    4. In your zshrc/theme/bashrc locate PROMPT and add godrink function.
    5. For faster interval reset, add an alias for godrink not_thirsty command, eg: alias godrink not_thirsty nt
