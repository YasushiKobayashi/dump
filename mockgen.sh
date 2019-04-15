#!/bin/sh
rm -f usecase/mock_usecase/*
for file in `ls -F usecase | grep -v / |grep -v _test.go`; do
  mockgen -source=usecase/${file} -destination usecase/mock_usecase/${file}
done

goimports -w usecase/mock_usecase/*.go
go fmt usecase/mock_usecase/*.go
