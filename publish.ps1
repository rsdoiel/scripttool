<#
generated with CMTools 0.0.10 b59480b

.SYNOPSIS
Publish script for GitHub pages. It expects the gh-pages branch to already exist.
#>

$workingBranch = git branch | Select-String -Pattern "\* " | ForEach-Object { $_ -replace '\* ', '' }
if ($workingBranch -eq "gh-pages") {
    git commit -am "publishing to gh-pages branch"
    git push origin gh-pages
} else {
    Write-Output "You're in $workingBranch branch"
    Write-Output "You need to pull in changes to the gh-pages branch to publish"
    $yesNo = Read-Host "Process Y/n"
    if ($yesNo -eq "Y" -or $yesNo -eq "y") {
        Write-Output "Committing and pushing to $workingBranch"
        git commit -am "committing to $workingBranch"
        git push origin $workingBranch
        Write-Output "Changing branches from $workingBranch to gh-pages"
        git checkout gh-pages
        Write-Output "Merging changes from origin gh-pages"
        git pull origin gh-pages
        git commit -am "merging origin gh-pages"
        Write-Output "Pulling changes from $workingBranch into gh-pages"
        git pull origin $workingBranch
        Write-Output "Merging changes from $workingBranch"
        git commit -am "merging $workingBranch with gh-pages"
        Write-Output "Pushing changes up and publishing"
        git push origin gh-pages
        Write-Output "Changing back to your working branch $workingBranch"
        git checkout $workingBranch
    }
}
