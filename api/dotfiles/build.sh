echo "# generating cores from design..."
goagen bootstrap -d github.com/made2591/horizon/api/design > /dev/null

echo "# building application..."
go build -o horizon

echo "# generating updated docs..."
./generate-docs.sh > /dev/null

echo "# petting unicorns..."
echo "# ..."
sleep 1
echo "# done!!!"
