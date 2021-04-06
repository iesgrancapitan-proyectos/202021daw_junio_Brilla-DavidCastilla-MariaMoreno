const { merge } = require('webpack-merge');
const commonConfig = require('./webpack.common');
const { join } = require('path')

const devConfig = {
    optimization: {
        minimize: false,
    },
    devtool: 'source-map',
    devServer: {
        hot: true,
        contentBase: join(__dirname, 'public'),
        publicPath: '/',
        port: 8081,
        historyApiFallback: true
    },
    mode: 'development'
};

module.exports = merge(commonConfig, devConfig);
