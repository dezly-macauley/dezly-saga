# React Dev Tutorial
- https://react.dev/learn
_______________________________________________________________________________
## If you have just cloned this repo the run this command

NOTE: This project uses `bun` so you will have to have bun installed.

```sh
bun install
```
_______________________________________________________________________________
## To create this repo from scratch

Make sure that you are in the repository where this command is used and then
run this command:

```sh
bunx create-next-app@latest
```

I then used this command to upgrage to TailwindCSS 4.0
- This command may not be neccessary in the future
```sh
bunx @tailwindcss/upgrade@next
```
_______________________________________________________________________________
## To run your project

Open a new terminal and navigate to your project directory, 
then run this command:

```sh
bun run dev
```

You will get an output that looks like this:
```
$ next dev --turbopack
   ▲ Next.js 15.1.6 (Turbopack)
   - Local:        http://localhost:3000
   - Network:      http://192.168.8.106:3000

 ✓ Starting...
 ✓ Ready in 916ms
```

Leave this terminal running, copy either of the links 
and paste it in your browser.

Use the local one if you are testing your local machine.
E.g. Your laptop

Use the network one if you are testing this one another device 
such as your phone that is connected to the same network

_______________________________________________________________________________
