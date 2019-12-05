module.exports = {
    root: true,
    env: {
        browser: true,
        node: true
    },
    extends: [
        '@nuxtjs/eslint-config-typescript',
        'plugin:prettier/recommended',
        'prettier',
        'prettier/vue'
    ],
    plugins: [
        'vue',
        'nuxt/recommended'
    ],
    rules: {
        'vue/html-closing-bracket-newline': 'off',
        'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        'space-before-function-paren': 0
    }
};