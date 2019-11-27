# Orenotorero

## GitHub運用ルール
- masterへの直接Push＝NG
- ブランチは下記のルールに則って作成する

`feat/[issue番号]/[やること]`

`例）feat/1/createLoginPage`

- 作業を始める時には必ずPullすること！
- プルリクは最低1レビュー以上無いとマージできない

## GitHubの運用の流れ
1. 本スプリントで自分がアサインされているものの中から今から作業するissueを決める
2. そのissue番号を確認する
![](README用画像/issue番号の確認.png)

3. masterブランチでPullしてくる

```
git pull origin master
```

4. 作業ブランチを切り替える

```
git checkout -b feat/14/writeDownREADME
```

5. 1コミットあたりの作業量を意識しながら作業をする

6. 作業が完了したら、変更したファイルが正しいか確認する

```
git status
```

7. 自分の変更した内容が正しいか確認する

```
git diff
```

8. 自分の変更分が正しいことを確認したらリモートへプッシュする

```
git add .

git commit
もしくは、
git commit -m 'hoge'

git push

※初めてプッシュする場合は、訳の分からん英文が出てくるが、落ち着いて、下記のような1行をコピー＆ペーストしよう！

git push --set-upstream origin feat/14/writeDownREADME
```

9. Pull Requestを作成する
![](README用画像/PullRequest1.png)
