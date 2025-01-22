Use these command to update Foundry:

rustup update

#### How to install Foundry
cargo install --git https://github.com/foundry-rs/foundry --profile release --locked forge cast chisel anvil

#### How to update Foundry
cargo install --git https://github.com/foundry-rs/foundry --profile release --locked --force forge cast chisel anvil

---

Remember to add this to your path:

Open your .zshrc file and add this line:

```
export PATH="$HOME/.cargo/bin:$PATH"

```

This is needed for Foundry, and other tools installed with Rustup to work
without this you won't be able to use commands like `anvil`

---

#### Install the Solidity LSP:

First install `node`
This is a JavaScript runtime. This LSP was actually made for Hardhat,
which is a JavaScript framework but it will still work in Foundry.

Note: It has to be `node`! The lsp script is designed to look for node.js so
trying to use bun or something else will give you an error.

```
mkdir -p ~/.local/lib/
cd ~/.local/lib/
```

```
npm install @nomicfoundation/solidity-language-server
```

Then open your .zshrc file and add this to your path:

```
export PATH=$HOME/.npm-global/node_modules/.bin:$PATH
```

To check that everything is fine run this:

```
which nomicfoundation-solidity-language-server

```

---
