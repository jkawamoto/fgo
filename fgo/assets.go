/*
 * assets.go
 *
 * Copyright (c) 2016-2019 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package fgo

//go:generate go-assets-builder --output=assets/bindata.go -s="/assets" -p=assets ./assets/Makefile ./assets/formula.rb
