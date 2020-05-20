echo BERGE!
git pull origin master

pwd
cd ../bashprofilefiles && touch .bash_profile_deleted
echo Backing up old bash profile to ~/BergeSafetyVault
mkdir -p ~/BergeSafetyVault
cp -r ~/.bash_profile ~/BergeSafetyVault
vaultRename=".bash_profile"+$(date '+%Y-%H-%M-%S');
mv ~/BergeSafetyVault/.bash_profile ~/BergeSafetyVault/$vaultRename

echo Copying uses base bash_profile to working directory BashProfiler
cd ../bashprofilefiles
cp -r ~/.bash_profile .

echo Build and run Go project to manipulate bash_profile
cd ..
mkdir -p build
cd pkg/main
go build .
mv main.exe ../../build
cd ../../build
./main.exe

echo Copying .bash_profile back to working directory
cp -r ../bashprofilefiles/.bash_profile ~/


if [ -z $1]
then
	echo "Not pushing to master"
else
	echo Pushing up to master
	cd ..
	git add .
	git commit -m "Berge!"
	git push origin master
	echo "Operation complete!"
fi


