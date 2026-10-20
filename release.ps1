<#
generated with CMTools 0.0.1 7534e43

.SYNOPSIS
Release script for scripttool on GitHub using gh CLI.
#>

# Determine repository and group IDs
$repoId = Split-Path -Leaf -Path (Get-Location)
$groupId = (git config --get remote.origin.url) -replace '.*github\.com[:/]([^/]+)/.*', '$1'
$repoUrl = "https://github.com/${groupId}/${repoId}"
Write-Output "REPO_URL -> ${repoUrl}"

# Generate release tag and notes
$releaseTag = "v$(jq -r .version codemeta.json)"
if ($releaseTag -notmatch '^v[0-9a-zA-Z._-]+$') {
    Write-Error "error: version contains unexpected characters: ${releaseTag}"
    exit 1
}
Write-Output "tag: ${releaseTag}, notes:"
jq -r .releaseNotes codemeta.json | Out-File -FilePath release_notes.tmp -Encoding utf8
Get-Content release_notes.tmp

# Generate checksums for distribution zip files
$checksumFile = "${repoId}-${releaseTag}-checksums.txt"
$hashes = Get-ChildItem -Path dist -Filter *.zip | ForEach-Object {
    $hash = (Get-FileHash -Path $_.FullName -Algorithm SHA256).Hash.ToLower()
    "$hash  $($_.Name)"
}
$hashes | Out-File -FilePath "dist/${checksumFile}" -Encoding utf8
Write-Output "Checksums written to dist/${checksumFile}"

# Prompt user to push release to GitHub
$yesNo = Read-Host -Prompt "Push release to GitHub with gh? (y/N)"
if ($yesNo -eq "y") {
    Write-Output "Saving working state for ${releaseTag}"
    git commit -am "prep for ${releaseTag}"
    git push
    Write-Output "Pushing release ${releaseTag} to GitHub"
    gh release create "${releaseTag}" `
        --draft `
        --notes-file release_notes.tmp `
        --generate-notes
    Write-Output "Uploading distribution files and checksums"
    Write-Output "  Starting upload: dist/${checksumFile}"
    gh release upload "${releaseTag}" "dist/${checksumFile}"
    Write-Output "  Completed upload: dist/${checksumFile}"
    foreach ($file in (Get-ChildItem -Path dist -Filter *.zip)) {
        Write-Output "  Starting upload: $($file.Name)"
        gh release upload "${releaseTag}" $file.FullName
        Write-Output "  Completed upload: $($file.Name)"
    }

    Write-Output @"

Now go to repo release and finalize draft.

    ${repoUrl}/releases

"@

    Remove-Item release_notes.tmp
}
