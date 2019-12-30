const path = require( 'path' );
const merge = require( 'webpack-merge' );
const webpack = require( 'webpack' );
const commonConfig = require( './webpack.config' );

const SERVER_PORT = process.env.PORT || '2305';
const WEBPACK_HOST = '0.0.0.0';
const WEBPACK_PORT = '3000';

const devConfig = {
    mode: 'development',
    devtool: 'cheap-eval-source-map',

    devServer: {
        port: WEBPACK_PORT,
        host: WEBPACK_HOST,
        stats: "minimal",
        inline: true,
        noInfo: true,
        compress: true,
        contentBase: path.resolve(__dirname, '/dist'),
        historyApiFallback: true,
        proxy: {
            '/v0/api': `http://server:${SERVER_PORT}`
        }
    }
};

module.exports = merge( commonConfig, devConfig );
