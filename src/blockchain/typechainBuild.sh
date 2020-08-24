#!/bin/sh

../../node_modules/.bin/typechain --target ethers-v5 --outDir ./typechain "./build/contracts/*json"
