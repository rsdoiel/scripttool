#!/bin/bash
# generated with CMTools 0.0.1 7534e43

#
# Release script for scripttool on GitHub using gh cli.
#
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

# Now we're ready to push things.
# shellcheck disable=SC2162
read -r -p "Push release to GitHub with gh? (y/N) " YES_NO
if [ "${YES_NO}" = "y" ]; then
	make save msg="prep for ${RELEASE_TAG}"
	# Now generate a draft release
	echo "Pushing release ${RELEASE_TAG} to GitHub"
	gh release create "${RELEASE_TAG}" \
		--draft \
		-F release_notes.tmp \
		--generate-notes
	echo "Uploading distribution files and checksums"
	echo "Starting upload: dist/${CHECKSUM_FILE}
	"
	gh release upload "${RELEASE_TAG}" "dist/${CHECKSUM_FILE}"
	echo "  Completed upload: dist/${CHECKSUM_FILE}"
	for FILE in dist/*.zip; do
		echo "Starting upload: ${FILE}
	"
		gh release upload "${RELEASE_TAG}" "${FILE}"
		echo "  Completed upload: ${FILE}"
	done

	cat <<EOT

Now goto repo release and finalize draft.

	${REPO_URL}/releases

EOT
	rm release_notes.tmp

fi
