# Traveller Command-line utility

> Note: We are **STILL UNDER CONSTRUCTION** here. You best look away until I'm properly dressed.

A golang command-line utility for Traveller RPG. Also known as **The Traveller's Tool**.

This currently does the following:

- Generates words in the Aslan, Darrian, Droyne, K'kree, Vargr, Vilani and Zhodani languages.
- All the above (partially) implemented using the imgui-go user interface library.
- Tested in Linux shells and Powershell/cmd.

## Usage

    usage: traveller [--help|--version]

    Options:
    With no command-line options, the application runs in UI mode with all choices
    determined by the question/answer session
    
    --help                  displays the help and usage for the application
    --version               displays the application 

## Planned

Upcoming functionality includes the following:

- Retrieve world, subsector and sector data from <https://travellermap.com/>.
- Generate the following types of worlds:
  - Classic Traveller Book 3.
  - MegaTraveller Basic (also suitable for Mongoose Traveller).
  - Traveller5 worlds only (not systems).
- Generate whole sectors of MegaTraveller basic mainworlds.  
- Generate Foreven sector with worlds detailed to T5SS.
- Complete star-system generation for CT Book 6, MegaTraveller and Traveller5.
- Character generation.
- Searching for Official Traveller Universe (OTU) world data stored in the database.
- Saving and searching of custom worlds, star-systems, sectors and characters.

Also, *possibly*

- Robot planning and construction.
- Shipbuilding.
- Trading.
- Date conversion.
- Campaign/scenario assistance for GameMasters/Referees (whatever they call them this year).

## Notes

This is the second major version of this software.

The first version used the [tview](https://github.com/rivo/tview) library for user interface, until I ran into complications with deadlocks and the library itself. I always felt like I was fighting with tview to get it to do what I wanted, and the callbacks were a killer. I was able to implement world and sector generation, and that code will be ported over to the new version.

This second major version uses the [imgui-go](https://github.com/inkyblackness/imgui-go) user interface which is an immediate-mode graphical user interface. It looks nice, and only took me a few hours mucking around to get a basic display up and running. Imgui-go is a golang wrapper around the C++ Dear ImGui library.

## Dependencies

This application uses the following packages:

- <https://github.com/inkyblackness/imgui-go> a Golang wrapper for the (C++) Dear ImGui graphical user interface.
- <https://github.com/dustin/go-humanize> to output numbers etc in a more readable fashion.
- <https://github.com/mattn/go-sqlite3> as the support database is a Sqlite file.
- <https://github.com/atotto/clipboard> to copy screen info to the clipboard. **OLD**
- GFLW  & Open GL3 libraries.
