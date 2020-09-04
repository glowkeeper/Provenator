import shortid from 'shortid'
import { ethers } from 'ethers'

export const getKey = (): string => shortid.generate()
export const getString = (bytes32Ref: any): string => ethers.utils.parseBytes32String(bytes32Ref)
export const setBytes32 = (key: string): any => ethers.utils.formatBytes32String(key)

export const addressToBytes32 = (address: string): string => {

    if (ethers.utils.isAddress(address)) {
         return address + "000000000000000000000000"
    } else {
        return address
    }
}

export const bytes32ToAddress = (bytesString: any): string => {

  let cutZero = bytesString.length - 1

  // as long as the last character is a 0, remove it
  do {

    if (bytesString[cutZero] === '0') {
      cutZero--
    }

  } while (bytesString[cutZero] === '0')

  return bytesString.substr(0, cutZero + 1)
}

export const md5ToBytes32 = (md5: string): string => "0x" + md5 + "00000000000000000000000000000000"

export const bytes32ToMd5 = (bytesString: any): string => {

  let cutZero = bytesString.length - 1

  // as long as the last character is a 0, remove it
  do {

    if (bytesString[cutZero] === '0') {
      cutZero--
    }

  } while (bytesString[cutZero] === '0')

  return bytesString.substr(2, cutZero + 1)
}
