module.exports = {
    plugins: {
        "autoprefixer": {},
        "postcss-size": {},
        "postcss-media-functions": {
            sizes: {
                sm: '576px',
                md: '768px',
                lg: '992px',
                xl: '1200px',
            },
        },
        "postcss-combine-media-query": {},
        "postcss-csso": {},
        "postcss-write-svg": {}
    }
}
