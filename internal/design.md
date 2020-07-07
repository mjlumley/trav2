# Design Decisions

This doc is intended to record any design decisions, diagrams or other design-worthy notes on the Traveller Tool. It is written for my future self. Since this software is not really intended to be used by anyone other than myself, this gives me greater flexibility in making decisions.

## Language &amp; Platform

This project was intended as a learning project in [go](https://golang.org), so the choice of language was pre-ordained. Since I was working in both Linux and Windows, and wished the flexibility of moving between these if needed, I decided early on that everything I wrote had to work on both platforms. The majority, if not all, of my day-to-day coding would be done on my (Windows) Asus ZenBook. If I left the zenbook in the office, I could continue on my Linux platform at home.

## Database

I wanted:

- portable - that is, data easily backed-up or moved.
- lightweight or in-memory - ideally serverless (though not cloud), or runs when needed.
- ease of initial installation - on the very rare chance that someone other than myself was trying to get the software to run.

A further consideration was whether I had experience in it or not. I was willing to learn, but not radically. Speed was not necessarily a factor, unless it was painful to use. These requirements pushed me in the direction of [sqlite](https://www.sqlite.org/) and the [go client](https://github.com/mattn/go-sqlite3), and the fact that you could store JSON in it if needed (how was I going to store whole star systems?). This choice seemed obvious. Other databases considered included:

- [Redis](https://redis.io) - although recommended, and suited to storing JSON data, this seemed a bit of a stretch. And required a server.
- [MongoDB](https://www.mongodb.com/) - I thought that this was a bit heavy, even though I was keen to try a NoSQL approach.
- [H2](https://h2database.com/) - I had used this previously in Struts work. Probably why it was one of the first cut ;-)

The go client was difficult to install on Windows, and so this choice did not fit the last requirement. It required a gcc build environment to be installed, and this was problematic.

## User Interface

I wanted:

- simplicity - I wasn't interested in writing a full-on windows application, and was hoping I could find something that worked on both platforms.
- reasonably kitted out with features.
- stable. Or stable enough, but still under active development.

After a lot of research, I initially settled on [tview](https://github.com/rivo/tview) which is itself based on [tcell](https://github.com/gdamore/tcell). As I got further and further into it, and had committed a LOT of time to it, I realised the limitations of the library were causing more issues than the libraries worth. Tview is a good library, very well-written, if a bit obfuscated in parts, and has very good documentation. My issues with it included the following opions:

- The demo examples were quite complicated for a golang newbie like me.
- While the code was complex, the demo programs presented only incredibly simple examples, they didn't go far enough, and didn't really do anything.
- The library's developer kept a very tight rein on the project, and constraints on his time meant that fixes and updates were very slow in coming along.
- There are warnings about deadlocks, and I was sure I followed all instructions to avoid, but still was plagued by deadlocks. Some updates appeared to fix them, but then there would be a new update and the deadlocks would return.
- A lot of functionality for the widgets is provided by callback functions, and this creates (in my opinion) messy code, and a feeling of never truly knowing where you are in your code.

So in the end, I decided to either use the tcell library directly, or find an alternative. The tcell route seemed a lot of work, so I wanted to avoid that. By chance, while looking through another fellahs project in lua and Love, I discovered the [Dear Imgui](https://github.com/ocornut/imgui) library, which is specifically for C++, however I found [imgui-go](https://github.com/inkyblackness/imgui-go), which provides a wrapper for Imgui. The "immediate mode" paradigm appealed to me, as opposed to the "constrained mode" previously, and re-development began. Imgui-go has its own problems, but once I got OpenGL working and modded the example, I could see that this was going to work. The examples given were comprehensive, and there was a huge range of widgets. The difficulties were:

- imgui-go is still under development, and not all features of the wrapper are implemented yet.
- A *lot* of code had to moved over.

On the other hand, the change 

## Project Layout

Mod files?
cmd and internal?
Testing?
