# Function to log command into database
log_command() {
    local cmd="${1[0, -2]}"
    go run ~/dev/midori/cmd/journal/journal.go add "$cmd" $PWD "command"

}

add-zsh-hook zshaddhistory log_command

