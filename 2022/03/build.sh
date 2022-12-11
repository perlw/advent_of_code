#!/bin/sh

bold_white="[1;37m"
bold_green="[1;32m"
bold_red="[1;31m"
end="[0m"

bin_name=$(basename `pwd`)

compiler_defines=""
common_compiler_flags="-debug -vet -strict-style -warnings-as-errors -verbose-errors"
for arg in "$@"; do
  case $arg in
    -release)
      common_compiler_flags="-o:speed -vet -strict-style -warnings-as-errors -verbose-errors"
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

echo ""
echo "${bold_white}Running tests${end}"
odin test ../src $compiler_flags
result=$?
if [ $result -ne 0 ]; then
  echo "${bold_red}FAIL${end}"
  exit $result
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
