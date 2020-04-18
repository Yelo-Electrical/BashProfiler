export PATH=$PATH:C:/Dev/protoc-3.11.4-win64/bin
export PATH=$PATH:"C:/go Program Files/MySQL/MySQL Workbench 8.0 CE/"

alias ab="git push -u origin $1"
alias build="go build hello.go && ./hello.exe"
alias c="clear"
alias cm="check master"
alias dev="cd /c/Dev"
alias e="explorer ."
alias edit="subl ~/.bash_profile"
alias fresh="source ~/.bash_profile"
alias ga="git add . && clear && git status"
alias gb="clear && git branch"
alias genproto="yelo && cd bashscripts && bash genProto.sh"
alias gopath="cd ~/go"
alias gp="genproto"
alias gr="git rebase $1"
alias gs="clear && git status"
alias gst="git stash"
alias gsta="git stash apply"
alias yelo="cd /c/Dev/MicroService"
alias la="ls -al"
alias ll="ls -l"
alias mas="git checkout master"
alias runClient="yelo && cd bashscripts && bash runClient.sh"
alias runServer="yelo && cd bashscripts && bash runServer.sh"
alias sql="mysql -u root -p"

check() {
	git checkout $1
}

checknew() {
	git checkout -b $1
}

dd () {
	cd $1
	la
}

//A hard version of fresh where we push the file to git repo
refresh() {
	cd ~
	cp  .bash_profile bash_profile/
	cd bash_profile
	git add .
	git commit -m "changes"
	git push -u origin master
}

gc() {
	clear
	git add .
	git commit -m "$1"
}

gbd() {
	git branch -D $1
}

mergesquash() {
    checknew "$1Merge"
    git merge --squash "$1"
    git add .
    git commit
    check master
    git merge "$1Merge"
    git branch -D "$1Merge"
}


o () {
	subl $1
}

v() {
	vim $1
}