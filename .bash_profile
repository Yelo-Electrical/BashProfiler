export PATH=$PATH:C:/Dev/protoc-3.11.4-win64/bin
export PATH=$PATH:"C:/Program Files/MySQL/MySQL Workbench 8.0 CE/"
alias ab="git push -u origin $1"
alias b="cd .."
alias bs="cd bashscripts"
alias c="clear"
alias cm="check master"
alias dev="cd /c/Dev"
alias e="explorer ."
alias edit="subl ~/.bash_profile"
alias fresh="source ~/.bash_profile"
alias ga="git add . && clear && git status"
alias gb="clear && git branch"
alias gop="cd ~/go"
alias gp="git push origin master"
alias gr="git rebase $1"
alias gs="clear && git status"
alias gst="git stash"
alias gsa="git stash apply"
alias home="command cd ~"
alias la="ls -al"
alias ll="ls -l"
alias mas="git checkout master"
alias ns="netstat -ano"
alias repo="cd /c/Users/Dell\ XPS/Dev/Repositories"
alias sh="bash"
alias sql="mysql -u root -p"
alias yelo="cd /c/Users/Dell\ XPS/Dev/Repositories/YeloElectrical"

#Starting Services
alias be="yelo && cd BE"
alias dum="yelo && cd DummyMS"
alias malta="repo && cd Malta2/be"
alias sand="todir be && bash runSandbox.sh"

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
	esac

	cd bashscripts
}

genp() {
	todir $1
	bash genProto.sh
}

runs () {
	todir $1
	bash runServer.sh
}

runc () {
	todir $1
	bash runClient.sh
}

runfe () {
	todir $1
	bash runFE.sh
}

t () {
	todir $1
	cd ../pkg/service/v1
	go test
}

check() {
	git checkout $1
}

bn() {
	git checkout -b $1
}



f() {
	grep -R "$1" *
}

ff () {
	grep -R "$1" $2
}

#A hard version of fresh where we push the file to git repo
refresh() {
	cd ~
	cp  .bash_profile bash_profile/
	cd bash_profile
	git add .
	git commit -m "changes"
	git push -u origin master
	yelo
}

gc() {
	clear
	git add .
	git commit -m "$1"
}

gbd() {
	git branch -D $1
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

v() {
	vim $1
}



