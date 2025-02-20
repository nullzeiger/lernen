# Copyright 2025 Ivan Guerreschi <ivan.guerreschi.dev@gmail.com>.
# All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all: build

build: main.go
	go build -o build/lernen .

clean: build/lernen
	rm -rf build/
