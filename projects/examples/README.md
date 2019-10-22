Install with cmd.exe
```
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/lovemew67/go-misc/master/projects/examples/example.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
```

Install with PowerShell.exe
```
Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/lovemew67/go-misc/master/projects/examples/example.ps1'))
```

Install with bash
```
curl -sL https://raw.githubusercontent.com/lovemew67/go-misc/master/projects/examples/example.sh | sudo bash -
source <(curl -sL https://raw.githubusercontent.com/lovemew67/go-misc/master/projects/examples/example.sh)
bash <(curl -sL https://raw.githubusercontent.com/lovemew67/go-misc/master/projects/examples/example.sh)
```