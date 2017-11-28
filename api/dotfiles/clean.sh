# my datetime format
dts() { date +%d-%m-%Y-%H-%M-%S; }

# save the date :D
mydate=$(dts);

# my personal create directory function
dmkdir() { mkdir bak/"$mydate"; }

echo "# creating unique bak dir..."
dmkdir

echo "# moving my core files to new unique bak dir..."
#TODO CHANGE
read -rsp $'# setup your clean script first: press escape to continue...\n' -d $'\e'
exit

# cp application.go bak/$mydate/application.go > /dev/null
# rm application.go

cp main.go bak/$mydate/main.go > /dev/null
rm main.go