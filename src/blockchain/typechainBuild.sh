#!/bin/sh

../../node_modules/.bin/typechain --target ethers-v4 --outDir ./typechain "./build/contracts/*json"
