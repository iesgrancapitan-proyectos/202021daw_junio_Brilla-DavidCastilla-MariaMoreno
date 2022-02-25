const CssMinimizerPlugin = require('css-minimizer-webpack-plugin');
const TerserPlugin = require('terser-webpack-plugin');
const { merge } = require('webpack-merge');
const commonConfig = require('./webpack.common');
const SriPlugin = require('webpack-subresource-integrity');
const webpack = require('webpack');

const prodConf = {
    plugins: [
        new SriPlugin({
            hashFuncNames: ['sha256', 'sha384'],
            enabled: true
        }),
        new webpack.DefinePlugin({
            API_URL: JSON.stringify('http://brilla.iesgrancapitan.org/api')
        })
    ],
    optimization: {
        minimize: true,
        minimizer: [
            new TerserPlugin(),
            new CssMinimizerPlugin(),
        ],
    },
    devtool: false,
    mode: 'production'
};

module.exports = merge(commonConfig, prodConf);
