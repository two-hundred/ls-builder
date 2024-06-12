#!/bin/bash

function finish {
  echo "staticcheck output:"
  echo ""
  cat staticcheck.out
  echo ""
  echo "govet report output:"
  echo ""
  cat govet-report.out
  echo ""
}

trap finish EXIT

for d in $(go list ./... | grep -v "vendor"); do
    staticcheck $d > staticcheck.out
    exit_code=$?
    if [ $exit_code -ne 0 ]; then
      echo "Exiting for staticcheck with code $exit_code"
      exit $exit_code
    fi

    go vet $d 2> govet-report.out
    exit_code=$?
    if [ $exit_code -ne 0 ]; then
     echo "Exiting for go vet with code $exit_code"
      exit $exit_code
    fi
done
