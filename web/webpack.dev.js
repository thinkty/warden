const path = require('path');
const { merge } = require('webpack-merge');
const HtmlWebPackPlugin = require('html-webpack-plugin');
const common = require('./webpack.common');

module.exports = merge(common, {
  mode: 'development',
  plugins: [
    new HtmlWebPackPlugin({
      template: path.resolve(__dirname, 'src/index.html'),
      favicon: path.resolve(__dirname, 'src/favicon.ico'),
      filename: 'index.html',
    }),
  ],
  devtool: 'inline-source-map',
  devServer: {
    static: path.resolve(__dirname, 'src'),
    open: true,
    port: 9000,
    hot: true,
    watchFiles: ['src/**/*'],
  },
});
