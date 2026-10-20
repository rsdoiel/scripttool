#!/bin/bash
# generated with CMTools 0.0.1 7534e43, then fixed by hand (see below)
#

#
# Publish script for scripttool for GitHub pages. It expect the gh-pages
# branch to already exist.
#
# Every step is checked. The script used to rely on `git pull` having a default
# way to reconcile diverged branches, and carried on after a failed pull to push
# nothing and report "Everything up-to-date". Now each pull says --no-rebase
# --no-edit, a commit is made only when there is something to commit, the first
# failure stops the script and names the step, and the push is confirmed by
# comparing gh-pages with origin/gh-pages. A merge conflict leaves you on
# gh-pages in the middle of the merge, to resolve by hand.
#

fail() {
	echo "publish stopped: $1" >&2
	echo "You are on branch $(git branch --show-current)." >&2
	exit 1
}

# commit_all commits tracked changes, if there are any.
commit_all() {
	if ! git diff --quiet HEAD --; then
		git commit -am "$1" || fail "git commit"
	fi
}

WORKING_BRANCH=$(git branch --show-current)
if [ "${WORKING_BRANCH}" = "gh-pages" ]; then
	commit_all "publishing to gh-pages branch"
	git push origin gh-pages || fail "git push origin gh-pages"
else
	echo "You're in ${WORKING_BRANCH} branch"
	echo "You need to pull in changes to the gh-pages branch to publish"
	# shellcheck disable=SC2162
	read -p "process Y/n " YES_NO
	if [ "${YES_NO}" = "Y" ] || [ "${YES_NO}" = "y" ]; then
		echo "Committing and pushing to ${WORKING_BRANCH}"
		commit_all "commiting to ${WORKING_BRANCH}"
		git push origin "${WORKING_BRANCH}" || fail "git push origin ${WORKING_BRANCH}"
		echo "Changing branchs from ${WORKING_BRANCH} to gh-pages"
		git checkout gh-pages || fail "git checkout gh-pages"
		echo "Merging changes from origin gh-pages"
		git pull --no-rebase --no-edit origin gh-pages || fail "git pull origin gh-pages"
		echo "Merging changes from ${WORKING_BRANCH} into gh-pages"
		git pull --no-rebase --no-edit origin "${WORKING_BRANCH}" || fail "git pull origin ${WORKING_BRANCH} (resolve the conflict, commit, then git push origin gh-pages)"
		echo "Pushing changes up and publishing"
		git push origin gh-pages || fail "git push origin gh-pages"
		git fetch origin gh-pages
		if [ "$(git rev-parse gh-pages)" != "$(git rev-parse origin/gh-pages)" ]; then
			fail "gh-pages and origin/gh-pages differ after the push"
		fi
		echo "Published gh-pages at $(git rev-parse --short gh-pages)"
		echo "Changing back to your working branch ${WORKING_BRANCH}"
		git checkout "${WORKING_BRANCH}" || fail "git checkout ${WORKING_BRANCH}"
	fi
fi
