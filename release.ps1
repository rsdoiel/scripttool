<#
.SYNOPSIS
Release script for scripttool on GitHub using gh CLI.

.DESCRIPTION
The release is made from the commit HEAD (workspace DR-0012). The script commits
and pushes nothing: it stops unless tracked files are committed and HEAD is the
tip of its branch on origin, and it creates the draft with --target HEAD.

An interrupted run is resumed by running the script again (workspace DR-0013).
If the draft for the tag exists, targets HEAD and is still a draft, nothing is
created again: each file in dist is skipped when the draft already holds it
(same SHA-256), uploaded when missing, and uploaded again with --clobber when the
draft holds a different or unfinished copy. A published release, or a draft for
another commit, stops the script.

.PARAMETER DryRun
Do everything up to the prompt (notes, checksums), say what would be created or
uploaded and from which commit, and change nothing.
#>
param(
    [switch]$DryRun
)

function Invoke-Release {
    # Determine repository and group IDs
    $repoId = Split-Path -Leaf -Path (Get-Location)
    $groupId = (git config --get remote.origin.url) -replace '.*github\.com[:/]([^/]+)/.*', '$1'
    $repoUrl = "https://github.com/${groupId}/${repoId}"
    Write-Host "REPO_URL -> ${repoUrl}"

    # Generate release tag and notes
    $releaseTag = "v$(jq -r .version codemeta.json)"
    if ($releaseTag -notmatch '^v[0-9a-zA-Z._-]+$') {
        Write-Error "error: version contains unexpected characters: ${releaseTag}"
        return 1
    }
    Write-Host "tag: ${releaseTag}, notes:"
    jq -r .releaseNotes codemeta.json | Out-File -FilePath release_notes.tmp -Encoding utf8
    Get-Content release_notes.tmp | Out-Host

    # Generate checksums for distribution zip files
    $checksumFile = "${repoId}-${releaseTag}-checksums.txt"
    $hashes = Get-ChildItem -Path dist -Filter *.zip | ForEach-Object {
        $hash = (Get-FileHash -Path $_.FullName -Algorithm SHA256).Hash.ToLower()
        "$hash  $($_.Name)"
    }
    $hashes | Out-File -FilePath "dist/${checksumFile}" -Encoding utf8
    Write-Host "Checksums written to dist/${checksumFile}"

    # The release is built from one known commit: tracked files are committed
    # and HEAD is the tip of its branch on origin (DR-0012). Done with git, since
    # Windows may not have make.
    $failed = $false
    if (git status --porcelain --untracked-files=no) {
        Write-Error "check-committed: tracked files have uncommitted changes (commit and push them)"
        $failed = $true
    }
    $branch = git branch --show-current
    $releaseCommit = git rev-parse HEAD
    $remoteTip = ((git ls-remote origin "refs/heads/${branch}" 2>$null) -split '\s+' | Select-Object -First 1)
    if (-not $remoteTip) {
        Write-Error "check-committed: cannot find branch ${branch} on origin (push it)"
        $failed = $true
    } elseif ($remoteTip -ne $releaseCommit) {
        Write-Error "check-committed: HEAD is not the tip of origin/${branch} (push, or pull first)"
        $failed = $true
    }
    if ($failed) { return 1 }
    Write-Host "check-committed: $($releaseCommit.Substring(0, 7)) is committed and pushed"

    # A release that already exists is resumed only if it is a draft for this
    # commit (DR-0013); anything else stops the run, so a second release is never
    # created and a published one is never touched.
    $resume = $false
    $view = $null
    $viewJson = gh release view "${releaseTag}" --json isDraft,targetCommitish,assets 2>$null
    if ($LASTEXITCODE -eq 0) {
        $view = ($viewJson -join "`n") | ConvertFrom-Json
        if (-not $view.isDraft) {
            Write-Error "error: ${releaseTag} is not a draft on GitHub; it is published and is not touched"
            return 1
        }
        if ($view.targetCommitish -ne $releaseCommit) {
            Write-Error "error: the draft for ${releaseTag} targets $($view.targetCommitish), not HEAD ${releaseCommit}; delete the draft first"
            return 1
        }
        $resume = $true
    }

    # What to do with each file: upload, replace or skip. The checksums file first, then the zips.
    $files = @(Get-Item -Path "dist/${checksumFile}") + @(Get-ChildItem -Path dist -Filter *.zip)
    $plan = @()
    foreach ($file in $files) {
        $action = "upload"
        if ($resume) {
            $asset = @($view.assets | Where-Object { $_.name -eq $file.Name }) | Select-Object -First 1
            if ($asset) {
                $local = (Get-FileHash -Path $file.FullName -Algorithm SHA256).Hash.ToLower()
                if ($asset.state -ne "uploaded") {
                    $action = "replace"
                } elseif ($asset.digest) {
                    $action = if ($asset.digest -eq "sha256:${local}") { "skip" } else { "replace" }
                } else {
                    $action = if ($asset.size -eq $file.Length) { "skip" } else { "replace" }
                }
            }
        }
        $plan += [pscustomobject]@{ Action = $action; File = $file }
    }
    $todo = @($plan | Where-Object { $_.Action -ne "skip" }).Count
    if ($resume) {
        Write-Host "Resume: the draft ${releaseTag} exists and targets ${releaseCommit}; ${todo} of $($plan.Count) files to upload"
    } else {
        Write-Host "New release: draft ${releaseTag} at ${releaseCommit}; $($plan.Count) files to upload"
    }
    foreach ($item in $plan) { Write-Host "  $($item.Action): dist/$($item.File.Name)" }

    if ($DryRun) {
        Write-Host "dry run: nothing changed"
        return 0
    }
    if ($resume -and $todo -eq 0) {
        Write-Host "Nothing to upload: the draft holds every file. Finalize it at ${repoUrl}/releases"
        return 0
    }
    $prompt = if ($resume) { "Resume the draft release on GitHub with gh? (y/N)" } else { "Create the draft release on GitHub with gh? (y/N)" }
    $yesNo = Read-Host -Prompt $prompt
    if ($yesNo -ne "y") { return 0 }

    if (-not $resume) {
        Write-Host "Pushing release ${releaseTag} to GitHub"
        gh release create "${releaseTag}" `
            --draft `
            --target "${releaseCommit}" `
            --notes-file release_notes.tmp `
            --generate-notes | Out-Host
        if ($LASTEXITCODE -ne 0) { return 1 }
    }
    Write-Host "Uploading distribution files and checksums"
    foreach ($item in $plan) {
        $name = $item.File.Name
        if ($item.Action -eq "skip") {
            Write-Host "  Skipped (already uploaded): dist/${name}"
            continue
        }
        Write-Host "  Starting upload: dist/${name}"
        if ($item.Action -eq "replace") {
            gh release upload "${releaseTag}" $item.File.FullName --clobber | Out-Host
        } else {
            gh release upload "${releaseTag}" $item.File.FullName | Out-Host
        }
        if ($LASTEXITCODE -ne 0) {
            Write-Error "error: upload of dist/${name} failed; run release.ps1 again to resume"
            return 1
        }
        Write-Host "  Completed upload: dist/${name}"
    }

    Write-Host @"

Now go to repo release and finalize draft.

    ${repoUrl}/releases

"@
    return 0
}

# The notes file is removed however the script ends.
try {
    # Anything a command leaves on the output stream joins the return value, so take the last item.
    $code = @(Invoke-Release)[-1]
} finally {
    Remove-Item -ErrorAction SilentlyContinue release_notes.tmp
}
exit $code
