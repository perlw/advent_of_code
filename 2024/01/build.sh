#!/bin/sh

bold_white="[1;37m"
bold_green="[1;32m"
bold_red="[1;31m"
end="[0m"

bin_name=$(basename `pwd`)

run_tests=0
compiler_defines=""
common_compiler_flags="-debug -vet -strict-style -warnings-as-errors"
if [[ `uname -s` == 'Darwin' ]]; then
  common_compiler_flags="-minimum-os-version:13.0 -extra-linker-flags:-L/usr/local/opt/openssl@3/lib ${common_compiler_flags}"
fi
for arg in "$@"; do
  case $arg in
    -tests)
      run_tests=1
    ;;
    -release)
      common_compiler_flags="-o:speed -vet -strict-style -warnings-as-errors"
    ;;
    *)
      echo "${bold_red}Unknown parameter \"${arg}\".${end}"
      exit -1
    ;;
  esac
done

compiler_flags="${compiler_defines} ${common_compiler_flags}"
echo "${bold_white}Using compiler flags:${end} ${compiler_flags}"

if [ ! -d build ]; then
  mkdir build
fi
pushd build 1>/dev/null 2>&1

if [ $run_tests -ne 0 ]; then
  echo ""
  echo "${bold_white}Running tests${end}"
  odin test ../src $compiler_flags
  result=$?
  if [ $result -ne 0 ]; then
    echo "${bold_red}FAIL${end}"
    exit $result
  fi
fi

echo ""
echo "${bold_white}Building:${end} $(pwd)/$bin_name"
odin build ../src -out:$bin_name -show-timings $compiler_flags
result=$?
if [ $result -ne 0 ]; then
  echo "${bold_red}FAIL${end}"
  exit $result
fi

echo ""
echo "${bold_green}OK${end}"

popd 1>/dev/null 2>&1
