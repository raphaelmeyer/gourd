#!/bin/bash

if [ "${TRAVIS_PULL_REQUEST}" == "false" ]; then
  echo -e "Starting to update gh-pages\n"

  mkdir -p ${HOME}/reports

  #copy data we're interested in to other place
  cp features.html ${HOME}/reports
  cp wip.html ${HOME}/reports

  #go to home and setup git
  cd ${HOME}
  git config --global user.email "travis@travis-ci.org"
  git config --global user.name "Travis"

  #using token clone gh-pages branch
  git clone --quiet --branch=gh-pages https://${GH_TOKEN}@github.com/raphaelmeyer/gourd.git gh-pages > /dev/null

  #copy data we're interested in to that directory
  cp -rf ${HOME}/reports/* ${HOME}/gh-pages

  #go into diractory and add, commit and push files
  cd ${HOME}/gh-pages
  git add -f .
  git commit -m "Travis build ${TRAVIS_BUILD_NUMBER} pushed to gh-pages"
  git push -fq origin gh-pages > /dev/null

  echo -e "Update done.\n"
fi

