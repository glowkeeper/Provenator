import Web3 from 'web3'

class Web3Handler {

  constructor () {
    this.batch = undefined
    this.account = undefined
    this.web3 = undefined
    this._setWeb3()
  }

  async _setWeb3 () {
    // Checking if Web3 has been injected by the new MetaMask web3 instance
    let ethereum = window.ethereum
    this.web3 = window.web3


    if (ethereum) {
      this.web3 = new Web3(ethereum)
      try {
        //console.log('Ethereum')
        await ethereum.enable()
      } catch (error) {
        // User denied account access...
        console.log("Ethereum disabled")
      }
    } else if (this.web3) {
      //console.log('Legacy dapp browsers...')
      this.web3 = new Web3(web3.currentProvider);
    }
    // Non-dapp browsers...
    else {
        //console.log('Non-Ethereum browser detected. You should consider trying MetaMask!')
        let host = '127.0.0.1'
        let port = '8545'
        this.web3 = new Web3(new Web3.providers.HttpProvider('http://' + host + ':' + port))
    }

    setInterval(this._setAccount.bind(this), 3000)
  }

  // metamask sets its account to web3.eth.accounts[0]
  async _setAccount () {
    const accounts = await this.web3.eth.getAccounts()
    if (accounts[0] !== this.account) {
      this.account = accounts[0]
      //console.log("Setting account", this.account)
    }
    //console.log("In set account for ", this.account)
  }

  _callChecker (_func, _cb) {
    // console.log('in func call checker')
    if ((typeof _func === 'function') && (typeof _cb === 'function')) {
      // console.log('Passed call checker!')
      return true
    } else {
      /* console.log('failed call checker!')
      console.log('func ' + _func)
      console.log('cb ' + _cb) */
      return false
    }
  }

  _callParamsChecker (_func, _params, _cb) {
    //console.log('_callParamsChecker', _func, _params, _cb)
    if ((typeof _func === 'function') && (typeof _cb === 'function') && (Array.isArray(_params))) {
      return true
    } else {
      return false
    }
  }

  _call (_caller, _func, _cb, _isCall) {
    if (this._callChecker(_func, _cb)) {
      const transactionObject = {from: this.account}
      if (_isCall) {
        _func().call(transactionObject, function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      } else {
        _func().send(transactionObject, function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      }
    }
  }

  _call1Params (_caller, _func, _params, _cb, _isCall) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      const transactionObject = _params[0]
      if (_isCall) {
        _func().call(transactionObject, function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      } else {
        //console.log('Call1: ', transactionObject)
        _func().send(transactionObject, (err, result) => {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      }
    }
  }

  _call2Params (_caller, _func, _params, _cb, _isCall) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      if (_isCall) {
        _func(_params[0]).call(_params[1], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      } else {
        _func(_params[0]).send(_params[1], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      }
    }
  }

  _call3Params (_caller, _func, _params, _cb, _isCall) {
    //console.log('_call3Params func ' + _func)
    if (this._callParamsChecker(_func, _params, _cb)) {
      //console.log('made it here?')
      if (_isCall) {
        //console.log('am I here?')
        _func(_params[0], _params[1]).call(_params[2], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      } else {
        //console.log('made it here? ', _params)
        //const transactionObject = {from: this.account}
        _func(_params[0], _params[1]).send(_params[2], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      }
    }
  }

  _call4Params (_caller, _func, _params, _cb, _isCall) {
    /* console.log('caller ' + _caller)
    console.log('func ' + _func)
    console.log('params ' + _params)
    console.log('cb ' + _cb)
    console.log('batch ' + _isBatch) */
    if (this._callParamsChecker(_func, _params, _cb)) {
      if (_isCall) {
        _func(_params[0], _params[1], _params[2]).call(_params[3], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      } else {
        _func(_params[0], _params[1], _params[2]).send(_params[3], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      }
    }
  }

  _call5Params (_caller, _func, _params, _cb, _isCall) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      if (_isCall) {
        _func(_params[0], _params[1], _params[2], _params[3]).call(_params[4], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      } else {
        _func(_params[0], _params[1], _params[2], _params[3]).send(_params[4], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      }
    }
  }

  _call6Params (_caller, _func, _params, _cb, _isCall) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      if (_isCall) {
        _func(_params[0], _params[1], _params[2], _params[3], _params[4]).call(_params[5], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      } else {
        _func(_params[0], _params[1], _params[2], _params[3], _params[4]).send(_params[5], function (err, result) {
          if (err) {
            console.log(err)
          } else {
            _cb(_caller, result)
          }
        })
      }
    }
  }

  getAccount () {
    //console.log("Account", this.account)
    return this.account
  }

  getWeb3 () {
    // console.log("Web3", this.web3)
    return this.web3
  }

  callHandler (_caller, _func, _cb, _isCall) {
    // console.log('in grrrr call handler')
    // console.log(_transactionObject)
    // console.log(_func)
    this._call(_caller, _func, _cb, _isCall)
  }

  callParamHandler (_caller, _func, _params, _cb, _isCall) {
    //const transactionObject = {from: this.account, data: _params}
    //console.log("in absolutely new call handler", _func)
    switch (_params.length) {
      case 1:
        this._call1Params(_caller, _func, _params, _cb, _isCall)
        break
      case 2:
        this._call2Params(_caller, _func, _params, _cb, _isCall)
        break
      case 3:
        this._call3Params(_caller, _func, _params, _cb, _isCall)
        break
      case 4:
        this._call4Params(_caller, _func, _params, _cb, _isCall)
        break
      case 5:
        this._call5Params(_caller, _func, _params, _cb, _isCall)
        break
      case 6:
        this._call6Params(_caller, _func, _params, _cb, _isCall)
        break
      default:
        this._call(_caller, _func, _cb, _isCall)
    }
  }
}

export default Web3Handler
