#!/usr/bin/env bash

# Copyright 2014 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

# Grovels through git to set a set of env variables.
version::get_version_vars() {
  local git=(git --work-tree "${_ROOT}")

  if [[ -n ${_GIT_COMMIT-} ]] || _GIT_COMMIT=$("${git[@]}" rev-parse "HEAD^{commit}" 2>/dev/null); then
    if [[ -z ${_GIT_TREE_STATE-} ]]; then
      # Check if the tree is dirty.  default to dirty
      if git_status=$("${git[@]}" status --porcelain 2>/dev/null) && [[ -z ${git_status} ]]; then
        _GIT_TREE_STATE="clean"
      else
        _GIT_TREE_STATE="dirty"
      fi
    fi

    # Use git describe to find the version based on tags.
    if [[ -n ${_GIT_VERSION-} ]] || _GIT_VERSION=$("${git[@]}" describe --tags --match='v*' --abbrev=14 "${_GIT_COMMIT}^{commit}" 2>/dev/null); then
      # These regexes are painful enough in sed...
      # We don't want to do them in pure shell, so disable SC2001
      # shellcheck disable=SC2001
      DASHES_IN_VERSION=$(echo "${_GIT_VERSION}" | sed "s/[^-]//g")
      if [[ "${DASHES_IN_VERSION}" == "---" ]] ; then
        # shellcheck disable=SC2001
        # We have distance to subversion (v1.1.0-subversion-1-gCommitHash)
        _GIT_VERSION=$(echo "${_GIT_VERSION}" | sed "s/-\([0-9]\{1,\}\)-g\([0-9a-f]\{14\}\)$/.\1\+\2/")
      elif [[ "${DASHES_IN_VERSION}" == "--" ]] ; then
        # shellcheck disable=SC2001
        # We have distance to base tag (v1.1.0-1-gCommitHash)
        _GIT_VERSION=$(echo "${_GIT_VERSION}" | sed "s/-g\([0-9a-f]\{14\}\)$/+\1/")
      fi
      if [[ "${_GIT_TREE_STATE}" == "dirty" ]]; then
        # git describe --dirty only considers changes to existing files, but
        # that is problematic since new untracked .go files affect the build,
        # so use our idea of "dirty" from git status instead.
        _GIT_VERSION+="-dirty"
      fi

      # If _GIT_VERSION is not a valid Semantic Version, then refuse to build.
      if ! [[ "${_GIT_VERSION}" =~ ^v([0-9]+)\.([0-9]+)(\.[0-9]+)?(-[0-9A-Za-z.-]+)?(\+[0-9A-Za-z.-]+)?$ ]]; then
          echo "_GIT_VERSION should be a valid Semantic Version. Current value: ${_GIT_VERSION}"
          echo "Please see more details here: https://semver.org"
          exit 1
      fi
    fi
  fi
}

# Prints the value that needs to be passed to the -ldflags parameter of go build
# in order to set the Kubernetes based on the git tree status.
version::ldflags() {
  version::get_version_vars

  local -a ldflags
  function add_ldflag() {
    local key=${1}
    local val=${2}

    ldflags+=(
      "-X 'pkg/version.${key}=${val}'"
    )
  }

  add_ldflag "buildDate" "$(date ${SOURCE_DATE_EPOCH:+"--date=@${SOURCE_DATE_EPOCH}"} -u +'%Y-%m-%dT%H:%M:%SZ')"
  if [[ -n ${_GIT_COMMIT-} ]]; then
    add_ldflag "gitCommit" "${_GIT_COMMIT}"
    add_ldflag "gitTreeState" "${_GIT_TREE_STATE}"
  fi

  if [[ -n ${_GIT_VERSION-} ]]; then
    add_ldflag "gitVersion" "${_GIT_VERSION}"
  fi

  # The -ldflags parameter takes a single string, so join the output.
  echo "${ldflags[*]-}"
}

echo $(version::ldflags) > "version.txt"

exit 0