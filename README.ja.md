# gphotos-auth

[Click here for English version.](https://github.com/Q-Brains/gphotos-auth/blob/master/README.md)

## 概要

[Google Photos API](https://developers.google.com/photos) のOAuth2認証を行う、それだけ。

ちなみに、Google Photos API はOAuth2認証以外の認証方法には対応していない。

## 使い方

1. `.env`(環境変数)を設定する
2. 以下のコマンドを実行する

```command-line
docker run --env-file .env qbrains/gphotos-auth:latest
```

## 環境変数

以下は `.env` の例。

```.env
CLIENT_ID=client_id
CLIENT_SECRET=client_secret
AUTH_SCOPES=read/append sharing
```

### `CLIENT_ID`

Google Cloud Platform で発行される OAuth 2.0 Client ID。  
[こちら](https://console.cloud.google.com/apis/credentials)から確認可能。

### `CLIENT_SECRET`

Google Cloud Platform で発行される OAuth 2.0 Client Secret。  
[こちら](https://console.cloud.google.com/apis/credentials)から確認可能。

### `AUTH_SCOPES`

`OAuth 2.0 scope for the Google Photos Library API` のエイリアス。  
エイリアスについて詳しくは[こちら](#auth-scopes)。

## Auth Scopes

公式リファレンスは[こちら](https://developers.google.com/photos/library/guides/authentication-authorization#OAuth2Authorizing)を参照。

|     Alias     |            OAuth 2.0 scope for the Google Photos Library API            |
| :-----------: | :---------------------------------------------------------------------: |
|    `read`     |        `https://www.googleapis.com/auth/photoslibrary.readonly`         |
|   `append`    |       `https://www.googleapis.com/auth/photoslibrary.appendonly`        |
|   `access`    | `https://www.googleapis.com/auth/photoslibrary.readonly.appcreateddata` |
| `read/append` |             `https://www.googleapis.com/auth/photoslibrary`             |
|   `sharing`   |         `https://www.googleapis.com/auth/photoslibrary.sharing`         |

それぞれのスコープは以上のエイリアスに対応しており、複数宣言する場合は半角スペースで分けて記述する。  
記述例は[環境変数](#環境変数)を参照。
