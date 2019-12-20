#!/bin/bash
set -e
set -x

# Enable bash completion
echo "source /etc/bash_completion" >> "/root/.bashrc"

# Add autocomplete to kubectl
echo "alias k=kubectl" >> "/root/.bashrc"
echo "source <(kubectl completion bash)" >> "/root/.bashrc"
echo "source <(kubectl completion bash | sed 's/kubectl/k/g')" >> "/root/.bashrc"
# Add kubectx 
git clone https://github.com/ahmetb/kubectx.git /root/.kubectx 
COMPDIR=$(pkg-config --variable=completionsdir bash-completion) 
ln -sf /root/.kubectx/completion/kubens.bash $COMPDIR/kubens 
ln -sf /root/.kubectx/completion/kubectx.bash $COMPDIR/kubectx 

# Bash history search 
cat >> "/root/.bashrc" <<BINDINGS
bind '"\e[A": history-search-backward'
bind '"\e[A": history-search-backward'
bind '"\e[B": history-search-forward'
bind '"\eOA": history-search-backward'
bind '"\eOB": history-search-forward'
BINDINGS

# Save bash history each time a command is used ... /root/commandhistory is mounted to a volume so it
# survives between container restarts
echo "export PROMPT_COMMAND='history -a'" >> "/root/.bashrc"
echo "export HISTFILE=/root/commandhistory/.bash_history" >> "/root/.bashrc"
mkdir -p /root/commandhistory
touch /root/commandhistory/.bash_history

# Git command prompt
git clone https://github.com/magicmonty/bash-git-prompt.git ~/.bash-git-prompt --depth=1 
echo "if [ -f \"$HOME/.bash-git-prompt/gitprompt.sh\" ]; then GIT_PROMPT_ONLY_IN_REPO=1 && source $HOME/.bash-git-prompt/gitprompt.sh; fi" >> "/root/.bashrc"


