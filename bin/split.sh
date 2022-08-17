#!/usr/bin/env bash

# Usage:
# ./bin/split.sh [name]
#
# Example:
# ./bin/split.sh support

set -e
set -x

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
BASEPATH=$(cd `dirname $0`; cd ../; pwd)
REPOS=$@

function split()
{
    SHA1=`git subtree split --prefix=$1`
    git push $2 "$SHA1:refs/heads/$CURRENT_BRANCH" -f
}

function remote()
{
    git remote add $1 $2 || true
}

git pull origin $CURRENT_BRANCH

if [[ $# -eq 0 ]]; then
    REPOS=$(ls $BASEPATH)
fi

for REPO in $REPOS ; do
    remote $REPO git@github.com:go-packagist/$REPO.git

    split "$REPO" $REPO
done
