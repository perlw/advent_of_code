@echo off
setlocal ENABLEDELAYEDEXPANSION

for /d %%d in (%cd%) do set base_name=%%~nxd
set bin_name=%base_name%.exe
set pdb_name=%base_name%.pdb

set bold_white=[1;37m
set bold_green=[1;32m
set bold_red=[1;31m
set end=[0m

set compiler_defines=
set common_compiler_flags=-debug -pdb-name:%pdb_name% -show-timings -vet -strict-style -warnings-as-errors -verbose-errors
for %%x in (%*) do (
  if "%%~x" == "-release" (
    set common_compiler_flags=-o:speed -show-timings -vet -strict-style -warnings-as-errors -verbose-errors
  ) else (
    echo %bold_red%Unknown argument "%%~x".%end%
    exit -1
  )
)
set compiler_flags=%compiler_defines% %common_compiler_flags%
echo %bold_white%Using compiler flags:%end% %compiler_flags%

if not exist build mkdir build
pushd build

echo ""
echo %bold_white%Running tests%end%
odin test ../src %compiler_flags%
set result=%errorlevel%
if %result% neq 0 (
  echo %bold_red%FAIL%end%
  exit %result%
)

echo ""
echo %bold_white%Building:%end% %cd%/build/%bin_name%
odin build ../src -out:%bin_name% %compiler_flags%
set result=%errorlevel%
if %result% neq 0 (
  echo %bold_red%FAIL%end%
  exit %result%
)

echo %bold_green%OK%end%

popd
