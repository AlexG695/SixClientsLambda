git add .
git commit -m "upload lambda"
git push

go build -tags lambda.norpc -o bootstrap main.go
rm main.zip
zip main.zip main