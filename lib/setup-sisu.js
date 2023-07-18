/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

// Node.js core
const fs = require('fs').promises;
const os = require('os');
const path = require('path');

// External
const core = require('@actions/core');
const tc = require('@actions/tool-cache');
const io = require('@actions/io');
const { url } = require('inspector');
// const releases = require('@hashicorp/js-releases');

// arch in [arm, x32, x64...] (https://nodejs.org/api/os.html#os_os_arch)
// return value in [amd64, 386, arm]
function mapArch (arch) {
  const mappings = {
    x32: '386',
    x64: 'amd64'
  };
  return mappings[arch] || arch;
}

// os in [darwin, linux, win32...] (https://nodejs.org/api/os.html#os_os_platform)
// return value in [darwin, linux, windows]
function mapOS (os) {
  const mappings = {
    win32: 'windows'
  };
  return mappings[os] || os;
}

async function downloadCLI (url) {
  core.debug(`Downloading Sisu CLI from ${url}`);
  // const pathToCLIZip = await tc.downloadTool(url);
  
  // let pathToCLI = '';
  const pathToCLI = await tc.downloadTool(url);

  // core.debug('Extracting Sisu CLI zip file');
  // if (os.platform().startsWith('win')) {
  //   core.debug(`Sisu CLI Download Path is ${pathToCLIZip}`);
  //   const fixedPathToCLIZip = `${pathToCLIZip}.zip`;
  //   io.mv(pathToCLIZip, fixedPathToCLIZip);
  //   core.debug(`Moved download to ${fixedPathToCLIZip}`);
  //   pathToCLI = await tc.extractZip(fixedPathToCLIZip);
  // } else {
  //   pathToCLI = await tc.extractZip(pathToCLIZip);
  // }

  // core.debug(`Sisu CLI path is ${pathToCLI}.`);

  // if (!pathToCLIZip || !pathToCLI) {
  //   throw new Error(`Unable to download Sisu from ${url}`);
  // }

  return pathToCLI;
}

async function run () {
  try {
    // Gather GitHub Actions inputs
    const version = core.getInput('sisu_version');

    // Gather OS details
    const osPlatform = os.platform();
    const osArch = os.arch();

    // TODO Download release files
    core.debug(`Finding releases for Sisu version ${version}`);
    // const release = await releases.getRelease('sisu', version, 'GitHub Action: Setup Sisu');

    const platform = mapOS(osPlatform);
    const arch = mapArch(osArch);
    
    // core.debug(`Getting build for Sisu version ${release.version}: ${platform} ${arch}`);
    // const build = release.getBuild(platform, arch);
    // if (!build) {
    //   throw new Error(`Sisu version ${version} not available for ${platform} and ${arch}`);
    // }

    // Download requested version
    // https://github.com/jlrosende/go-action/releases/download/0.0.1/go-action-0.0.1-darwin-amd64

    binUrl = `https://github.com/jlrosende/go-action/releases/download/${version}/go-action-${version}-${platform}-${arch}`
    core.debug(binUrl)
    const pathToCLI = await downloadCLI(binUrl);

    // Add to path
    // core.addPath(pathToCLI);
    core.debug(core.addPath(pathToCLI));

    // return release;
  } catch (error) {
    core.error(error);
    throw error;
  }
}

run()

// module.exports = run;