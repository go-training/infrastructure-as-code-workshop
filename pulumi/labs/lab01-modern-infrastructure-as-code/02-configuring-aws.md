# 設定 AWS 環境

在使用 Pulumi 之前要先把 AWS 環境建立好

## 前置作業

請先將 AWS 環境設定完畢，請用 `aws configure` 完成 profile 設定

```sh
aws configure --profile demo
```

## 步驟一: 設定 AWS Region

可以參考 [AWS 官方的 Available Regions](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions)，並且透過 Pulumi CLI 做調整

```sh
cd demo && pulumi config set aws:region ap-northeast-1
```

切換到 demo 目錄，並執行 `pulumi config`

## 步驟二: 設定 AWS Profile

如果你有很多環境需要設定，請使用 AWS Profile 作切換，不要用 default profile。其中 `demo` 為 profile 名稱

```sh
pulumi config set aws:profile demo
```
