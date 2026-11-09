# install.ps1 - aerc Windows installer
$ErrorActionPreference = 'Stop'

Write-Host "Installing aerc for Windows..." -ForegroundColor Cyan

# 1. Target directory: ~/.local/bin
$installDir = Join-Path $HOME ".local\bin"
if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir -Force | Out-Null
}

# 2. Download latest release
$apiUrl = "https://api.github.com/repos/hardikxk/aerc-win/releases/latest"
$release = Invoke-RestMethod -Uri $apiUrl -Headers @{ "User-Agent" = "aerc-installer" }
$asset = $release.assets | Where-Object { $_.name -eq "aerc.exe" } | Select-Object -First 1

if (-not $asset) {
    throw "Could not find aerc.exe in the latest release ($($release.tag_name))"
}

$destPath = Join-Path $installDir "aerc.exe"
Write-Host "Downloading aerc $($release.tag_name)..." -ForegroundColor Yellow
Invoke-WebRequest -Uri $asset.browser_download_url -OutFile $destPath -UseBasicParsing

# 3. Add to PATH permanently if not present
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($currentPath -split ';' -notcontains $installDir) {
    Write-Host "Adding $installDir to User PATH..." -ForegroundColor Green
    [Environment]::SetEnvironmentVariable("Path", "$currentPath;$installDir", "User")
    $env:Path = "$env:Path;$installDir"
}

Write-Host "`n✓ aerc installed successfully to $destPath" -ForegroundColor Green
Write-Host "Restart your terminal or run 'aerc' to get started!" -ForegroundColor Cyan