#!/bin/bash

set -exuo pipefail

OLD_PACKAGE_NAME="github.com/CodeLieutenant/GoFiber-Boilerplate"
OLD_DOCKERHUB_NAME="codeLieutenant/gofiber-boilerplate"
OLD_BINARY_NAME="GoFiber-Boilerplate"

replace_package_name() {
    local new_package_name="$1"

    find . -type f -exec sed -i "s|$OLD_PACKAGE_NAME|$new_package_name|g" {} +
}

replace_dockerhub_name() {
    local new_dockerhub_name="$1"

    find . -type f -exec sed -i "s|$OLD_DOCKERHUB_NAME|$new_dockerhub_name|g" {} +
}

replace_binary_name() {
    local new_binary_name="$1"

    find . -type f -exec sed -i "s|$OLD_BINARY_NAME|$new_binary_name|g" {} +
}

run_pre_commit_install() {
    if command -v pre-commit &> /dev/null
    then
        pre-commit install
    else
        echo "pre-commit is not installed, skipping installation..."
    fi
}

run_task_build() {
    if command -v task &> /dev/null
    then
        task build
    else
        echo "task is not installed, skipping build..."
    fi
}

run_task_docker_build() {
    if command -v task &> /dev/null
    then
        task docker-build
    else
        echo "task is not installed, skipping docker build..."
    fi
}

read -p "Enter the new package name (e.g., github.com/yourusername/yourproject): " new_package_name
read -p "Enter the new DockerHub package name (e.g., yourusername/yourproject): " new_dockerhub_name
read -p "Enter the new Binary name (e.g., yourproject): " new_binary_name

replace_package_name "$new_package_name"
replace_dockerhub_name "$new_dockerhub_name"
replace_binary_name "$new_binary_name"
run_task_build

rm -rf .git
rm -rf .github/workflows/setup-test.yml
rm -rf .github/FUNDING.yml
rm -rf setup.sh

run_pre_commit_install
