## hood
>creating template files never be so easy. :kissing:

### Install

- bin

[https://github.com/xuqingfeng/hood/releases](https://github.com/xuqingfeng/hood/releases)

### Usage

`hood -t [go, php, js, ignore...] -n [name]`

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

`demo.php`

```php
<?php
/**
 * Author jsxqf
 * Date 2015/9/22
 */
```

