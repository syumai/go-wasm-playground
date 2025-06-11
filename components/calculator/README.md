# TinyGo Calculator WebAssembly Component

TinyGoで実装した計算機アプリケーションをWasm Component として作成し、jcoでJavaScriptにトランスパイルしてNode.jsから呼び出すサンプルプロジェクトです。

## 概要

このプロジェクトは以下の技術を使用して構築されています：

-   **TinyGo**: Go言語のWebAssembly向けコンパイラ
-   **WebAssembly Component Model**: モジュラーなWebAssemblyコンポーネントの仕様
-   **WIT (WebAssembly Interface Types)**: コンポーネントのインターフェース定義
-   **jco**: WebAssemblyコンポーネントをJavaScriptにトランスパイルするツール
-   **Node.js 22**: TypeScriptのtype stripping機能を使用

## プロジェクト構成

```
.
├── wit/
│   └── world.wit           # WebAssembly Interface Types定義
├── src/
│   └── sample.ts           # TypeScriptサンプルアプリケーション
├── js-output/              # jcoで生成されたJavaScriptファイル
│   ├── calculator.js
│   ├── calculator.d.ts
│   └── interfaces/
├── main.go                 # TinyGo計算機実装
├── calculator.wasm         # ビルドされたWasm Component
├── calculator:calc@0.1.0.wasm # WIT Package
├── sample.js              # JavaScriptサンプル
└── package.json
```

## 必要なツール

-   **TinyGo** v0.34.0+
-   **wasm-tools** CLI
-   **wkg** (WebAssembly Package Manager)
-   **Node.js** v22.6.0+ (type stripping機能のため)
-   **pnpm**

### インストール

```bash
# wasm-tools
brew install wasm-tools

# wkg
cargo binstall wkg

# Node.js 22
nvm install 22
nvm use 22
```

## ビルド手順

### 1. WITファイルからパッケージを生成

```bash
wkg wit build
```

### 2. Goバインディングを生成

```bash
go tool wit-bindgen-go generate --world calculator-world --out internal ./calculator:calc@0.1.0.wasm
```

### 3. TinyGoでWasm Componentをビルド

```bash
tinygo build -target=wasip2 -o calculator.wasm \
  --wit-package ./calculator:calc@0.1.0.wasm \
  --wit-world calculator-world main.go
```

### 4. JavaScriptにトランスパイル

```bash
pnpm exec jco transpile calculator.wasm -o js-output
```

## 実行方法

### JavaScript版

```bash
pnpm run sample
```

### TypeScript版 (Node.js 22 type stripping)

```bash
pnpm run sample:ts
```

## 実装されている機能

計算機コンポーネントは以下の4つの基本的な算術演算をサポートしています：

-   `add(a: f64, b: f64) -> f64` - 加算
-   `subtract(a: f64, b: f64) -> f64` - 減算
-   `multiply(a: f64, b: f64) -> f64` - 乗算
-   `divide(a: f64, b: f64) -> f64` - 除算（ゼロ除算の場合は0を返す）

## 技術的な詳細

### WebAssembly Component Model

このプロジェクトでは、WebAssembly Component Modelを使用してモジュラーなWebAssemblyコンポーネントを作成しています。Component Modelにより、異なる言語で書かれたコンポーネント間での型安全な相互運用が可能になります。

### WIT (WebAssembly Interface Types)

`wit/world.wit`ファイルでコンポーネントのインターフェースを定義しています：

```wit
package calculator:calc@0.1.0;

interface calculator {
    add: func(a: f64, b: f64) -> f64;
    subtract: func(a: f64, b: f64) -> f64;
    multiply: func(a: f64, b: f64) -> f64;
    divide: func(a: f64, b: f64) -> f64;
}

world calculator-world {
    include wasi:cli/imports@0.2.0;
    export calculator;
}
```

### TinyGo実装

TinyGoのエクスポート機能を使用して、WITインターフェースを実装しています：

```go
func init() {
    calculator.Exports.Add = func(a, b float64) float64 {
        return a + b
    }
    // ...
}
```

### jcoトランスパイレーション

jcoにより、WebAssemblyコンポーネントがJavaScriptの非同期モジュールに変換され、Node.jsから直接インポート可能になります。
