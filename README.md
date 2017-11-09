# hood
>creating template files never be so easy :kissing:

### Usage

`hood -t [ansible, c, css, go, html, ignore, java, js, php, python] -n [name]`

### Custom Config

`touch` `.hood.json` in your home directory

`~/.hood.json`

```json
{
  "GIT_IGNORED_LIST": [
    ".DS_Store",
    "idea"
  ],
  "USER": "CUSTOM USER NAME"
}
```

### Demo

`hood -t php -n demo`

```php
// demo.php
<?php
/**
 * Author jsxqf
 * Date 2015/9/22
 */
```

