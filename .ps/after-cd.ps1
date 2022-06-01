Write-Host
Write-Host "Recent Branches:"
Write-Host
git branch --list --sort=-committerdate --format='%(committerdate:short) %(refname:short)' | select -first 20 | Write-Host
