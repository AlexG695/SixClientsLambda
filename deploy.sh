git add .
git commit -m "upload lambda"
git push

go build main.go
rm main.zip
zip main.zip main