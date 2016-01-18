#!/usr/bin/env node
'use strict'
let webpack = require('webpack')
let WebpackDevServer = require('webpack-dev-server')
let config = require('./webpack.config')

let devServer = new WebpackDevServer(webpack(config), {
  hot: true,
  publicPath: config.output.publicPath,
  historyApiFallback: true,
  stats: {
    colors: true
  },
  proxy: {
    '*': {
      target: 'http://localhost:3000',
      secure: false,
      bypass: (req, res, proxyOptions) => {
        if (req.headers.accept.indexOf('html') !== -1) {
          return '/index.html';
        }
      }
    }
  },
  headers: {
    'Access-Control-Allow-Origin': '*'
  }
})

devServer.listen(5000, 'localhost', err => {
  if (err) {
    console.error(err)
    return
  }
})
