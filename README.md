# practice-todo-nuxt

GoLang/Revel を使用した学習用デモアプリ API

# 要件

## 機能

- タスクの追加
- タスクの一覧表示
- タスクの完了チェック
- タスクの削除

## 技術スタック

### アプリケーションサーバ

- GoLang/Revel

### データベース

- MySQL

このリポジトリではサーバ API 側のソースコードのみ管理する

## 基本設計

### タスクの作成

ユーザーから送られたタスクをデータベースに保存する。

### タスクの取得

データベースから全てのタスクを取得し、フロントエンドに送信する。

### タスクの更新

タスクの完了状態を更新するために、データベースのタスクを更新する。

### タスクの削除

データベースから指定されたタスクを削除する。

### エラーハンドリング

エラーが発生した場合、適切なエラーメッセージとステータスコードを返す。

# TODO

- [ ] タスク作成の実装
- [ ] タスク詳細取得の実装
- [ ] タスク一覧取得の実装
- [ ] タスク更新の実装
- [ ] タスク削除の実装
- [ ] ユニットテスト導入

## Start the web server:

```
export MYSQL_DATABASE={データベース名}
export MYSQL_USER={データベースユーザ名}
export MYSQL_PASSWORD={データベースパスワード}
revel run
```

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites
