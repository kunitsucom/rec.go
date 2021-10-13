#!/usr/bin/env bash
# shellcheck disable=SC2155
set -Eeou pipefail

# const
GOLANGCILINT_VERSION=1.45.0

# MIT License Copyright (c) 2021 newtstat https://github.com/rec-logger/rec.sh
# Common
_recRFC3339() { date "+%Y-%m-%dT%H:%M:%S%z" | sed "s/\(..\)$/:\1/"; }
_recCmd() { for a in "$@"; do if echo "${a:-}" | grep -Eq "[[:blank:]]"; then printf "'%s' " "${a:-}"; else printf "%s " "${a:-}"; fi; done | sed "s/ $//"; }
# Color
RecDefault() { test "${REC_SEVERITY:-0  }" -gt 000 2>/dev/null || echo "$*" | awk "{print   \"$(_recRFC3339) [\\033[0;35m  DEFAULT\\033[0m] \"\$0\"\"}" 1>&2; }
RecDebug() { test "${REC_SEVERITY:-0    }" -gt 100 2>/dev/null || echo "$*" | awk "{print   \"$(_recRFC3339) [\\033[0;34m    DEBUG\\033[0m] \"\$0\"\"}" 1>&2; }
RecInfo() { test "${REC_SEVERITY:-0     }" -gt 200 2>/dev/null || echo "$*" | awk "{print   \"$(_recRFC3339) [\\033[0;32m     INFO\\033[0m] \"\$0\"\"}" 1>&2; }
RecNotice() { test "${REC_SEVERITY:-0   }" -gt 300 2>/dev/null || echo "$*" | awk "{print   \"$(_recRFC3339) [\\033[0;36m   NOTICE\\033[0m] \"\$0\"\"}" 1>&2; }
RecWarning() { test "${REC_SEVERITY:-0  }" -gt 400 2>/dev/null || echo "$*" | awk "{print   \"$(_recRFC3339) [\\033[0;33m  WARNING\\033[0m] \"\$0\"\"}" 1>&2; }
RecError() { test "${REC_SEVERITY:-0    }" -gt 500 2>/dev/null || echo "$*" | awk "{print   \"$(_recRFC3339) [\\033[0;31m    ERROR\\033[0m] \"\$0\"\"}" 1>&2; }
RecCritical() { test "${REC_SEVERITY:-0 }" -gt 600 2>/dev/null || echo "$*" | awk "{print \"$(_recRFC3339) [\\033[0;1;31m CRITICAL\\033[0m] \"\$0\"\"}" 1>&2; }
RecAlert() { test "${REC_SEVERITY:-0    }" -gt 700 2>/dev/null || echo "$*" | awk "{print   \"$(_recRFC3339) [\\033[0;41m    ALERT\\033[0m] \"\$0\"\"}" 1>&2; }
RecEmergency() { test "${REC_SEVERITY:-0}" -gt 800 2>/dev/null || echo "$*" | awk "{print \"$(_recRFC3339) [\\033[0;1;41mEMERGENCY\\033[0m] \"\$0\"\"}" 1>&2; }
RecExec() { RecInfo "$ $(_recCmd "$@")" && "$@"; }
RecRun() { _dlm='####R#E#C#D#E#L#I#M#I#T#E#R####' _all=$({ _out=$("$@") && _rtn=$? || _rtn=$? && printf "\n%s" "${_dlm:?}${_out:-}" && return ${_rtn:-0}; } 2>&1) && _rtn=$? || _rtn=$? && _dlmno=$(echo "${_all:-}" | sed -n "/${_dlm:?}/=") && _cmd=$(_recCmd "$@") && _stdout=$(echo "${_all:-}" | tail -n +"${_dlmno:-1}" | sed "s/^${_dlm:?}//") && _stderr=$(echo "${_all:-}" | head -n "${_dlmno:-1}" | grep -v "^${_dlm:?}") && RecInfo "$ ${_cmd:-}" && RecInfo "${_stdout:-}" && { [ -z "${_stderr:-}" ] || RecWarning "${_stderr:-}"; } && return ${_rtn:-0}; }

# var
repo_root_dir=$(git rev-parse --show-toplevel)
pre_push_dst_path="${repo_root_dir:?}/.git/hooks/pre-push"
local_dir="${repo_root_dir:?}/.local"
local_bin_dir="${local_dir:?}/bin"
local_tmp_dir="${local_dir:?}/tmp"

if [[ ! -d "${local_bin_dir:?}" ]] || [[ ! -d "${local_tmp_dir:?}" ]]; then
  RecExec mkdir -p "${local_bin_dir:?}" "${local_tmp_dir:?}"
fi

if [[ ! -f "${pre_push_dst_path:?}" ]]; then
  RecExec cp -ai "${repo_root_dir:?}/githooks/pre-push" "${pre_push_dst_path:?}"
fi

_arch () {
  local uname_m=$(uname -m)
  if [[ "${uname_m:?}" = x86_64 ]]; then
    echo amd64
  fi
}

_os () {
  local uname_s=$(uname -s)
  if [[ "${uname_s:?}" = Darwin ]]; then
    echo darwin
  fi
}

if [[ ! -x "${local_bin_dir:?}/golangci-lint" ]]; then
  url="https://github.com/golangci/golangci-lint/releases/download/v${GOLANGCILINT_VERSION:?}/golangci-lint-${GOLANGCILINT_VERSION:?}-$(_os)-$(_arch).tar.gz"
  RecExec bash -c "cd ${local_tmp_dir:?} && curl -#fLR ${url:?} -o ./golangci-lint.tar.gz && tar xfz ./golangci-lint.tar.gz && mv -fv golangci-lint-*/golangci-lint ${local_bin_dir:?}"
fi
