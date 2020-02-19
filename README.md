# gphotos-auth

[日本語版はこちら](https://github.com/Q-Brains/gphotos-auth/blob/master/README.ja.md)

Unnatural English writing? Because it is written by Japanese who is not good at English.

## Overview

OAuth2 authentication for [Google Photos API](https://developers.google.com/photos), that's it.

By the way, Google Photos API does not support any method other than OAuth2 authentication.

## Usage

1. Set `.env` (environment variables).
2. Execute the following command.

```command-line
docker run --env-file .env qbrains/gphotos-auth:latest
```

## Environment Variables

The following is an example of `.env`.

```.env
CLIENT_ID=client_id
CLIENT_SECRET=client_secret
AUTH_SCOPES=read/append sharing
```

### `CLIENT_ID`

OAuth 2.0 Client ID issued from Google Cloud Platform.  
You can check from [here](https://console.cloud.google.com/apis/credentials).

### `CLIENT_SECRET`

OAuth 2.0 Client Secret issued from Google Cloud Platform.  
You can check from [here](https://console.cloud.google.com/apis/credentials).

### `AUTH_SCOPES`

The aliases of `OAuth 2.0 scope for the Google Photos Library API`.  
[Learn more about the aliases](#auth-scopes).

## About Scopes

[Click here for the Official Reference.](https://developers.google.com/photos/library/guides/authentication-authorization#OAuth2Authorizing)

|     Alias     |            OAuth 2.0 scope for the Google Photos Library API            |
| :-----------: | :---------------------------------------------------------------------: |
|    `read`     |        `https://www.googleapis.com/auth/photoslibrary.readonly`         |
|   `append`    |       `https://www.googleapis.com/auth/photoslibrary.appendonly`        |
|   `access`    | `https://www.googleapis.com/auth/photoslibrary.readonly.appcreateddata` |
| `read/append` |             `https://www.googleapis.com/auth/photoslibrary`             |
|   `sharing`   |         `https://www.googleapis.com/auth/photoslibrary.sharing`         |

Each scope corresponds to the above character strings. If multiple declarations are made, separate them with spaces.  
See the [Environment Variables](#environment-variables) section for a description example.
