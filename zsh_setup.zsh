# Function to log command into database
log_command() {
    local cmd="${1[0, -2]}"
    if [[ -n "$cmd" ]]; then
        journal add --cmd "$cmd" 
    fi
}

add-zsh-hook zshaddhistory log_command
