Berge is our Bash Profile Merger.
Add the project to your Repos.
Add this below code to your .bash_profile.


// Repo must take you to your home repo where Berge lives
alias bp="repo && command cd BashProfiler"

berge() {
	bp
	cd scripts
	bash run.sh
}