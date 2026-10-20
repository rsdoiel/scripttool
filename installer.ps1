#!/usr/bin/env pwsh
# generated with CMTools 0.0.12 799ba10

#
# Set the package name and version to install
#
param(
  [Parameter()]
  [String]$VERSION = "0.0.12"
)
[String]$PKG_VERSION = [Environment]::GetEnvironmentVariable("PKG_VERSION")
if ($PKG_VERSION) {
	$VERSION = "${PKG_VERSION}"
	Write-Output "Using '${PKG_VERSION}' for version value '${VERSION}'"
}

$PACKAGE = "scripttool"
$GIT_GROUP = "rsdoiel"
$RELEASE = "https://github.com/${GIT_GROUP}/${PACKAGE}/releases/tag/v${VERSION}"
$SYSTEM_TYPE = Get-ComputerInfo -Property CsSystemType
if ($SYSTEM_TYPE.CsSystemType.Contains("ARM64")) {
    $MACHINE = "arm64"
} else {
    $MACHINE = "x86_64"
}

Write-Output "Using release ${RELEASE}"

# FIGURE OUT Install directory
$BIN_DIR = "${Home}\bin"
Write-Output "${PACKAGE} v${VERSION} will be installed in ${BIN_DIR}"

#
# Figure out what the zip file is named
#
$ZIPFILE = "${PACKAGE}-v${VERSION}-Windows-${MACHINE}.zip"
$CHECKSUM_FILE = "${PACKAGE}-v${VERSION}-checksums.txt"
Write-Output "Fetching Zipfile ${ZIPFILE}"

#
# Check to see if this zip file has been downloaded.
#
$DOWNLOAD_URL = "https://github.com/${GIT_GROUP}/${PACKAGE}/releases/download/v${VERSION}/${ZIPFILE}"
Write-Output "Download URL ${DOWNLOAD_URL}"

if (!(Test-Path $BIN_DIR)) {
  New-Item $BIN_DIR -ItemType Directory | Out-Null
}
curl.exe -Lo "${ZIPFILE}" "${DOWNLOAD_URL}"
if (!(Test-Path $ZIPFILE)) {
    Write-Output "Failed to download ${ZIPFILE} from ${DOWNLOAD_URL}"
} else {
    # Verify checksum
    $CHECKSUM_URL = "https://github.com/${GIT_GROUP}/${PACKAGE}/releases/download/v${VERSION}/${CHECKSUM_FILE}"
    try {
        curl.exe -Lo "${CHECKSUM_FILE}" "${CHECKSUM_URL}"
        if (Test-Path $CHECKSUM_FILE) {
            $expectedLine = Get-Content $CHECKSUM_FILE | Where-Object { $_ -match [regex]::Escape($ZIPFILE) }
            if ($expectedLine) {
                $expectedHash = ($expectedLine -split 's+')[0].ToLower()
                $actualHash = (Get-FileHash -Path $ZIPFILE -Algorithm SHA256).Hash.ToLower()
                if ($expectedHash -eq $actualHash) {
                    Write-Output "Checksum verified: ${ZIPFILE}"
                } else {
                    Write-Warning "Checksum mismatch for ${ZIPFILE}"
                    Write-Warning "  Expected: ${expectedHash}"
                    Write-Warning "  Actual:   ${actualHash}"
                    Write-Warning "  Proceeding anyway — verify the download manually if concerned"
                }
            } else {
                Write-Warning "${ZIPFILE} not found in checksum file, skipping verification"
            }
            Remove-Item $CHECKSUM_FILE -ErrorAction SilentlyContinue
        }
    } catch {
        Write-Warning "Could not download ${CHECKSUM_FILE}, skipping verification"
    }
    # Do we have a zip file or tar.gz file?
    $fileInfo = Get-Item "${ZIPFILE}"

    # Handle zip or tar.gz files
    switch ($fileInfo.Extension) {
        ".zip" {
            Expand-Archive -Force -Path "${ZIPFILE}" "${Home}"
            break
        }
        ".gz" {
            tar.exe xf "${ZIPFILE}" -C "${Home}"
            break
        }
        ".tgz" {
            tar.exe xf "${ZIPFILE}" -C "${Home}"
            break
        }
        default {
            Write-Output "The ${ZIPFILE} from ${DOWNLOAD_URL} is neither a ZIP file nor a gzipped tar file."
            exit 1
        }
    }

    #Remove-Item $ZIPFILE

    $User = [System.EnvironmentVariableTarget]::User
    $Path = [System.Environment]::GetEnvironmentVariable('Path', $User)
    if (!(";${Path};".ToLower() -like "*;${BIN_DIR};*".ToLower())) {
        [System.Environment]::SetEnvironmentVariable('Path', "${Path};${BIN_DIR}", $User)
        $Env:Path += ";${BIN_DIR}"
    }
    Write-Output "${PACKAGE} was installed successfully to ${BIN_DIR}"
	Write-Output "If you get a security warning on Windows or macOS please see INSTALL_NOTES_Windows.md or INSTALL_NOTES_macOS.md"
}
