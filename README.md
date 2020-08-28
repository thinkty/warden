# Warden
![profile](https://imgur.com/t7AP5n6.png)

Warden is a personal assistant made for my personal use on Discord.
It is made using the [Dialogflow Editor](https://github.com/thinkty/dialogflow-editor) and is based on the [Source Chat](https://github.com/thinkty/source-chat) server with a [Discord](https://discordpy.readthedocs.io/en/latest/discord.html) adapter.
I chose to go with Discord, since Discord provides [Discord.js](https://discord.js.org/#/) which makes interacting with the Discord API a whole lot easier.
The main thing I like about Discord.js is that it uses sockets instead of API endpoints so I can run my flow manager on a left over raspberry pi without having to worry about port forwarding or having to pay for a cloud server.

## System Layout
You can find the explanation of the system layout in the [Source Chat](https://github.com/thinkty/source-chat) repository.

## Installation
I recommend you start with [Source Chat](https://github.com/thinkty/source-chat) instead of this repository as Warden is more personalized (with various action handlers and modifications from the template)

## License
MIT
