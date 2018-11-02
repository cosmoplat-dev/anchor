/*
 * Copyright 2018 COSMOPlat-DEV
 *
 * This program is free software; you can redistribute and/or modify it
 * under the terms of the standard MIT license. See LICENSE for more details
 */

package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/cosmoplat-dev/anchor/internal/app"
)

func main() {
	skel.PluginMain(app.CmdAdd, app.CmdDel, version.All)
}
