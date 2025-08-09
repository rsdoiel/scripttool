<#
generated with CMTools 0.0.10 b59480b

.SYNOPSIS
Release script for scripttool on GitHub using gh CLI.
#>

# Determine repository and group IDs
$repoId = Split-Path -Leaf -Path (Get-Location)
$groupId = Split-Path -Leaf -Path (Split-Path -Parent -Path (Get-Location))
$repoUrl = "https://github.com/$groupId/$repoId"
Write-Output "REPO_URL -> $repoUrl"

# Generate a new draft release using jq and gh
$releaseTag = "v$(jq -r .version codemeta.json)"
$releaseNotes = jq -r .releaseNotes codemeta.json | ForEach-Object { $_ -replace "`n", " " -replace "`'", "'" }
Write-Output "tag: $releaseTag, notes:"
jq -r .releaseNotes codemeta.json | Out-File -FilePath release_notes.tmp -Encoding utf8
Get-Content release_notes.tmp

# Prompt user to push release to GitHub
$yesNo = Read-Host -Prompt "Push release to GitHub with gh? (y/N)"
if ($yesNo -eq "y") {
    # Assuming 'make save' is a placeholder for a command you have
    # Replace 'make save' with the appropriate PowerShell command or function
    Write-Output "Preparing for $releaseTag, $releaseNotes"
    # Create a draft release
    Write-Output "Pushing release $releaseTag to GitHub"
    gh release create "$releaseTag" `
        --draft `
        --notes-file release_notes.tmp `
        --generate-notes
    Write-Output "Uploading distribution files"
    gh release upload "$releaseTag" dist/*.zip

    @"
Now go to repo release and finalize draft.

    $repoUrl/releases

"@

    Remove-Item release_notes.tmp
}
