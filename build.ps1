$rndName = -join ((65..90) + (97..122) | Get-Random -Count 16 | ForEach-Object {[char]$_})

go build -o bin/$rndName.exe -ldflags='-X main.BUILD_TIMESTAMP="'$(Get-Date -f yyyyMMddHHmmss)'"' ./cmd/gosource/main.go 

Write-Host "Output: bin/$($rndName).exe"
