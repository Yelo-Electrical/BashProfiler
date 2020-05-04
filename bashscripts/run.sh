echo Pulling bash profiles...
# Copy .bash_profile to working directory
# cp ~/.bash_profile ../bashprofilefiles

# Build and run Go project to manipulate bash_profile
mkdir -p ../build
cd ../main
go build .
mv main.exe ../build
cd ../build
./main.exe

# Move .bash_profile back to working directory
``