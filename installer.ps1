#!/usr/bin/env pwsh
# generated with CMTools 0.0.10 b59480b

#
# Set the package name and version to install
#
param(
  [Parameter()]
  [String]$VERSION = "0.0.10"
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
#if ([System.IO.File]::Exists($ZIPFILE)) {
if (!(Test-Path $ZIPFILE)) {
    Write-Output "Failed to download ${ZIPFILE} from ${DOWNLOAD_URL}"
} else {
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
