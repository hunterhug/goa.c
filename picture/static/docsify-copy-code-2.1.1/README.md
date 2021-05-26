# docsify-copy-code

[![NPM](https://img.shields.io/npm/v/docsify-copy-code.svg?style=flat-square)](https://www.npmjs.com/package/docsify-copy-code)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://github.com/jhildenbiddle/docsify-copy-code/blob/master/LICENSE)

A [docsify](https://docsify.js.org) plugin that adds a button to easily copy code blocks to your clipboard.

## Installation

**Production**

Add following script tag to your `index.html` after docsify. Specifying the `@[version]` in the URL ensures that the release of a major update (v3.x) will not break your production site:

```html
<!-- Latest v2.x.x -->
<script src="https://unpkg.com/docsify-copy-code@2"></script>
```

**Development**

If you prefer to load the latest version of the library, you may do so by omitting the `@[version]` from the above URL.

```html
<!-- Latest (not recommended for production) -->
<script src="https://unpkg.com/docsify-copy-code"></script>
```

## Options

### Button text

Button text can be customized as follows:

```javascript
window.$docsify = {
  // docsify-copy-code (defaults)
  copyCode: {
    buttonText : 'Copy to clipboard',
    errorText  : 'Error',
    successText: 'Copied'
  }
}
```

### Localization (l10n)

Button text can also be customized based on the current URL. Object key/value pairs are processed in the order provided.

```javascript
window.$docsify = {
  copyCode: {
    buttonText: {
      '/zh-cn/': '点击复制',
      '/ru/'   : 'Скопировать в буфер обмена',
      '/de-de/': 'Klicken Sie zum Kopieren',
      '/es/'   : 'Haga clic para copiar',
      '/'      : 'Copy to clipboard'
    },
    errorText: {
      '/zh-cn/': '错误',
      '/ru/'   : 'ошибка',
      '/'      : 'Error'
    },
    successText: {
      '/zh-cn/': '复制',
      '/ru/'   : 'Скопировано',
      '/de-de/': 'Kopiert',
      '/es/'   : 'Copiado',
      '/'      : 'Copied'
    }
  }
}
```

**Note:** Docsify's [alias](https://docsify.js.org/#/configuration?id=alias) option makes it easy to manage local content using separate directories. See the [`/demo/`](https://github.com/jperasmus/docsify-copy-code/tree/master/demo) content in this repo for an example.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jperasmus/docsify-copy-code/blob/master/LICENSE) for details.