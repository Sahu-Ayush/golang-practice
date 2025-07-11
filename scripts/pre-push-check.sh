#!/bin/bash

echo "🔍 Checking module name in go.mod..."
go run ./tools/check_module_name.go
STATUS=$?

if [ $STATUS -ne 0 ]; then
  echo "❌ Module name check failed. Push blocked."
  exit 1
fi

echo "✅ All good. Proceeding with push."
exit 0
