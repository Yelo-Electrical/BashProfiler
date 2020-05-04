Berge is our Bash Profile Merger.
Add the project to your Repos.
Add this below code to your .bash_profile.


// Repo must take you to your home repo where Berge lives
alias bp="repo && command cd BashProfiler"

berge() {
	bp 
	command cd bashprofilefiles
	cp -r ~/.bash_profile .
	command cd ../bashscripts
	bash run.sh
	mv ../bashprofilefiles/.bash_profile ~/
	git pull origin master
	git add .
	git commit -m "Berge!"
	git push origin master
}