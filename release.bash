#!/bin/bash
# generated with CMTools 0.0.1 7534e43

#
# Release script for scripttool on GitHub using gh cli.
#
# Usage: release.bash [--dry-run]
#   --dry-run  do everything up to the prompt (notes, checksums), say what
#              would be created or uploaded and from which commit, and change nothing.
#
# The release is made from the commit HEAD (workspace DR-0012). The script
# commits and pushes nothing: it stops unless the tree is committed and HEAD is
# pushed (make check-committed), and creates the draft with --target HEAD.
#
# An interrupted run is resumed by running the script again (workspace DR-0013).
# If the draft for the tag exists, targets HEAD and is still a draft, nothing is
# created again: each file in dist/ is skipped when the draft already holds it
# (same SHA-256), uploaded when missing, and uploaded again with --clobber when
# the draft holds a different or unfinished copy. A published release, or a
# draft for another commit, stops the script. Over SSH, run it under tmux or
# nohup so a dropped connection does not interrupt the upload.
#
DRY_RUN=0
for ARG in "$@"; do
    case "${ARG}" in
        -n|--dry-run) DRY_RUN=1 ;;
        -h|--help) sed -n '2,/^DRY_RUN/p' "$0" | grep '^#'; exit 0 ;;
        *) echo "usage: release.bash [--dry-run]" >&2; exit 2 ;;
    esac
done

# The notes file is removed however the script ends, even if the terminal is lost.
trap 'rm -f release_notes.tmp' EXIT

REPO_ID="${PWD##*/}"
GROUP_ID="$(git config --get remote.origin.url | sed -E 's#.*github\.com[:/]([^/]+)/.*#\1#')"
REPO_URL="https://github.com/${GROUP_ID}/${REPO_ID}"
echo "REPO_URL -> ${REPO_URL}"

#
# Generate a new draft release using jq and gh
#
RELEASE_TAG="v$(jq -r .version codemeta.json)"
if ! printf '%s' "${RELEASE_TAG}" | grep -qE '^v[0-9a-zA-Z._-]+$'; then
    echo "error: version contains unexpected characters: ${RELEASE_TAG}"
    exit 1
fi
echo "tag: ${RELEASE_TAG}, notes:"
jq -r .releaseNotes codemeta.json >release_notes.tmp
cat release_notes.tmp

#
# Generate checksums for distribution zip files
#
CHECKSUM_FILE="${REPO_ID}-${RELEASE_TAG}-checksums.txt"
if command -v sha256sum >/dev/null 2>&1; then
    sha256sum dist/*.zip | sed 's|dist/||' > "dist/${CHECKSUM_FILE}"
else
    shasum -a 256 dist/*.zip | sed 's|dist/||' > "dist/${CHECKSUM_FILE}"
fi
echo "Checksums written to dist/${CHECKSUM_FILE}"

# The release is built from one known commit. Stop unless it is committed and
# pushed.
if ! make -s check-committed; then
    exit 1
fi
RELEASE_COMMIT="$(git rev-parse HEAD)"

# digest prints the SHA-256 of a file.
digest() {
    if command -v sha256sum >/dev/null 2>&1; then
        sha256sum "$1" | cut -d' ' -f1
    else
        shasum -a 256 "$1" | cut -d' ' -f1
    fi
}

# A release that already exists is resumed only if it is a draft for this commit
# (DR-0013); anything else stops the run, so a second release is never created
# and a published one is never touched.
RESUME=0
VIEW=""
if VIEW="$(gh release view "${RELEASE_TAG}" --json isDraft,targetCommitish,assets 2>/dev/null)"; then
    if [ "$(printf '%s' "${VIEW}" | jq -r .isDraft)" != "true" ]; then
        echo "error: ${RELEASE_TAG} is not a draft on GitHub; it is published and is not touched"
        exit 1
    fi
    DRAFT_TARGET="$(printf '%s' "${VIEW}" | jq -r .targetCommitish)"
    if [ "${DRAFT_TARGET}" != "${RELEASE_COMMIT}" ]; then
        echo "error: the draft for ${RELEASE_TAG} targets ${DRAFT_TARGET}, not HEAD ${RELEASE_COMMIT}; delete the draft first"
        exit 1
    fi
    RESUME=1
fi

# plan_for prints what to do with one file: upload, replace or skip.
plan_for() {
    if [ "${RESUME}" != "1" ]; then
        echo upload
        return
    fi
    printf '%s' "${VIEW}" | jq -r --arg n "$(basename "$1")" --arg d "sha256:$(digest "$1")" --argjson s "$(wc -c <"$1" | tr -d ' ')" '
        ([.assets[]? | select(.name == $n)] | .[0]) as $a
        | if $a == null then "upload"
          elif $a.state != "uploaded" then "replace"
          elif ($a.digest // "") != "" then (if $a.digest == $d then "skip" else "replace" end)
          elif $a.size == $s then "skip"
          else "replace" end'
}

# The checksums file first, then the zips.
PLAN=()
TODO_COUNT=0
FILE_COUNT=0
for FILE in "dist/${CHECKSUM_FILE}" dist/*.zip; do
    ACTION="$(plan_for "${FILE}")"
    PLAN+=("${ACTION} ${FILE}")
    FILE_COUNT=$((FILE_COUNT + 1))
    [ "${ACTION}" = "skip" ] || TODO_COUNT=$((TODO_COUNT + 1))
done

if [ "${RESUME}" = "1" ]; then
    echo "Resume: the draft ${RELEASE_TAG} exists and targets ${RELEASE_COMMIT}; ${TODO_COUNT} of ${FILE_COUNT} files to upload"
else
    echo "New release: draft ${RELEASE_TAG} at ${RELEASE_COMMIT}; ${FILE_COUNT} files to upload"
fi
for LINE in "${PLAN[@]}"; do
    echo "  ${LINE%% *}: ${LINE#* }"
done

if [ "${DRY_RUN}" = "1" ]; then
    echo "dry run: nothing changed"
    exit 0
fi
if [ "${RESUME}" = "1" ] && [ "${TODO_COUNT}" = "0" ]; then
    echo "Nothing to upload: the draft holds every file. Finalize it at ${REPO_URL}/releases"
    exit 0
fi
if [ "${RESUME}" = "1" ]; then
    PROMPT="Resume the draft release on GitHub with gh? (y/N) "
else
    PROMPT="Create the draft release on GitHub with gh? (y/N) "
fi
# shellcheck disable=SC2162
read -r -p "${PROMPT}" YES_NO
if [ "${YES_NO}" = "y" ]; then
    if [ "${RESUME}" != "1" ]; then
        echo "Pushing release ${RELEASE_TAG} to GitHub"
        gh release create "${RELEASE_TAG}" \
            --draft \
            --target "${RELEASE_COMMIT}" \
            -F release_notes.tmp \
            --generate-notes || exit 1
    fi
    echo "Uploading distribution files and checksums"
    for LINE in "${PLAN[@]}"; do
        ACTION="${LINE%% *}"
        FILE="${LINE#* }"
        case "${ACTION}" in
            skip)
                echo "  Skipped (already uploaded): ${FILE}"
                continue
                ;;
            replace) CLOBBER="--clobber" ;;
            *) CLOBBER="" ;;
        esac
        echo "  Starting upload: ${FILE}"
        # shellcheck disable=SC2086
        if ! gh release upload "${RELEASE_TAG}" "${FILE}" ${CLOBBER}; then
            echo "error: upload of ${FILE} failed; run release.bash again to resume"
            exit 1
        fi
        echo "  Completed upload: ${FILE}"
    done

    cat <<EOT

Now goto repo release and finalize draft.

	${REPO_URL}/releases

EOT
fi
