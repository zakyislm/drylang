$ErrorActionPreference = "Stop"

$repo = "zakyislm/drylang"
$binName = "y.exe"
$installDir = Join-Path $env:USERPROFILE ".drylang\bin"

Write-Host "Installing dryLang..." -ForegroundColor Cyan

# Create install directory
if (-not (Test-Path $installDir)) {
    New-Item -Path $installDir -ItemType Directory -Force | Out-Null
}

$binPath = Join-Path $installDir $binName

# Download latest release
$apiUrl = "https://api.github.com/repos/$repo/releases/latest"
try {
    Write-Host "Fetching latest release..."
    $release = Invoke-RestMethod -Uri $apiUrl -ErrorAction Stop
    $asset = $release.assets | Where-Object { $_.name -match "windows" }
    
    if (-not $asset) {
        Write-Host "No Windows binary found in latest release. Falling back to local build (if available)." -ForegroundColor Yellow
        # Fallback mechanism if no release exists yet
    } else {
        $downloadUrl = $asset.browser_download_url
        Write-Host "Downloading $downloadUrl ..."
        Invoke-WebRequest -Uri $downloadUrl -OutFile $binPath -UseBasicParsing
    }
} catch {
    Write-Host "Warning: Could not fetch from GitHub. If you are developing locally, run 'go build -o $binPath' manually." -ForegroundColor Yellow
}

# Update Path
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($userPath -notmatch [regex]::Escape($installDir)) {
    Write-Host "Adding $installDir to User PATH..."
    $newPath = $userPath + (if ($userPath.EndsWith(";")) { "" } else { ";" }) + $installDir
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
} else {
    Write-Host "Path already configured."
}

Write-Host "`nInstallation Complete!" -ForegroundColor Green
Write-Host "dryLang is installed at: $binPath"
Write-Host "Please restart your terminal or run: `$env:Path += `";$installDir`""
Write-Host "Then type 'y' to verify."
