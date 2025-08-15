#!/bin/bash

current_branch="$(git branch --show-current)"

git checkout master && git pull origin

last_version="$(
	git branch -a |
		sed -nE 's/.*release\/v([0-9]+\.[0-9]+\.[0-9]+)/\1/p' |
		sort -Vr |
		head -n1
)"

major="$(echo $last_version | awk -F. '{ print $1 }')"
minor="$(echo $last_version | awk -F. '{ print $2 }')"
patch="$(echo $last_version | awk -F. '{ print $3 }')"

echo "current version: $last_version"
echo "which part do you want to increment?"
echo "1 - major. current major version is: $major"
echo "2 - minor. current minor version is: $minor"
echo "3 - patch. current patch version is: $patch"

while read -p "your answer: " input; do
	case "$input" in
	"1")
		((major++))
		break
		;;
	"2")
		((minor++))
		break
		;;
	"3")
		((patch++))
		break
		;;
	*)
		echo "wrong input, try again"
		;;
	esac
done

echo "new version will be $major.$minor.$patch."

while read -p "ok with that? [y/n] " input; do
	case "$input" in
	"y" | "yes")
		branch_name="release/v$major.$minor.$patch"

		{
			git checkout -b "$branch_name" &&
				git push --set-upstream origin "$branch_name"
		} || {
			git checkout "$current_branch"
			git branch -D "$branch_name"
			exit 1
		}

		git checkout "$current_branch"

		exit 0
		;;
	"n" | "no")
		echo "got negative answer, aborting"
		git checkout "$current_branch"

		exit 1
		;;
	*)
		echo "wrong input, try again"
		;;
	esac
done
