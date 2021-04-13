const { merge } = require('webpack-merge');
const commonConfig = require('./webpack.common');
const { join } = require('path')
const webpack = require('webpack')

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
        //host: "0.0.0.0",
        historyApiFallback: true,
        disableHostCheck: true,
    },
    plugins: [
        new webpack.DefinePlugin({
            API_URL: JSON.stringify('http://localhost/api')
        })
    ],
    mode: 'development'
};

module.exports = merge(commonConfig, devConfig);
