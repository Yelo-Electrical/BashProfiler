#
# BERGE
# Directories
# Directory
# Explorer
# Functions
# General
# Git
# My SQL
# Starting Services
#Deleted
alias ab="git push -u origin $1"
alias b="cd .."
alias be="yelo && command cd BE"
alias bp="repo && command cd BashProfiler"
alias bpc="rbp && bp && cd bashprofilefiles && cp ~/.bash_profile ."
alias bs="cd bashscripts"
alias c="clear"
alias cm="check master"
alias dev="cd /c/Users/Dell\ XPS/Dev"
alias dum="yelo && command cd DummyMS"
alias e="explorer ."
alias edit="subl ~/.bash_profile"
alias fresh="source ~/.bash_profile"
alias ga="git add . && clear && git status"
alias gb="clear && git branch"
alias gens="todir be && bash genSchema.sh"
alias gop="cd ~/go"
alias gp="git push origin master"
alias gr="git rebase $1"
alias gs="clear && git status"
alias gsa="git stash apply"
alias gst="git stash"
alias home="command cd ~"
alias la="ls -al"
alias ll="ls -l"
alias malta="repo && command cd Malta2/be"
alias mas="git checkout master"
alias ns="netstat -ano"
alias rbp="bp && rm -rf bashprofilefiles/.bash* && cp backup/.bash_profile_deleted bashprofilefiles/ && cp backup/.bash_profile bashprofilefiles/ && cp backup/.bash_profile_repo bashprofilefiles/"
alias repo="command cd /c/Users/Dell\ XPS/Dev/Repositories"
alias res="bp && cp -r ~/BergeSafetyVault/.bash_profile ~/. && cp ~/.bash_profile bashprofilefiles/"
alias sand="todir be && bash runSandbox.sh"
alias sh="bash"
alias sql="mysql -u root -p"
alias vbp="bp && cat bashprofilefiles/.bash_profile"
alias vbpd="bp && cat bashprofilefiles/.bash_profile_deleted"
alias vbpr="bp && cat bashprofilefiles/.bash_profile_repo"
alias yelo="cd ~/Dev/Repositories/YeloElectrical"
berge() {
	bp
	cd scripts
	bash run.sh
}
bn() {
	git checkout -b $1
}
check() {
	git checkout $1
}
export PATH=$PATH:"C:/Program Files/MySQL/MySQL Workbench 8.0 CE/"
export PATH=$PATH:C:/Dev/protoc-3.11.4-win64/bin
f() {
	grep -R "$1" *
}
ff () {
	grep -R "$1" $2
}
gbd() {
	git branch -D $1
}
gc() {
	clear
	git add .
	git commit -m "$1"
}
genp() {
	todir $1
	bash genProto.sh
}
gf() {
	git remote add $1
	git push origin master
}
ms() {
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
runc () {
	todir $1
	bash runClient.sh
}
runfe () {
	todir $1
	bash runFE.sh
}
runs () {
	todir $1
	bash runServer.sh
}
todir () {
	cmd=$1
	case $cmd in
	'be')
	be
	;;
	'dum')
	dum
	;;
	'malta')
	malta
	;;
	'bp')
	bp
	;;
	esac
	cd bashscripts
}
v() {
	vim $1
}

#Deleted
