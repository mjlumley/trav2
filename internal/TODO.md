# To Do Items &amp; Ideas for The Traveller's Tool

> Note: this needs some serious update, to continue the flip over from teh previous version.

This file helps me organise my project. It provides a place to put ideas and todo items.

- When building on windows, refer to <https://medium.com/@yaravind/go-sqlite-on-windows-f91ef2dacfe> for details of building go-sqlite3.

## To Do

- [ ] Design for the Character Generation/Generation State.
- [ ] Complete actual code for star and system generation.
- [ ] Menu options to switch between panes.
- [ ] Status text can display help.
- [ ] Check that exe files are not being uploaded to github.
- [ ] <https://martinheinz.dev/blog/5> - Ultimate setup for your next golang project
- [ ] <https://github.com/thockin/go-build-template> - Go Build Template
- [ ] <https://github.com/spf13/viper> - Go Configuration
- [ ] Chargen: get skills from DB rather than hard-coded.
- [ ] Error codes returned as multiple return parameters - refer to <https://golang.org/doc/effective_go.html#errors>
- [ ] Expand logging, so for instance panics are reported to the log.
- [ ] Log to various levels.
- [ ] Check that defers are used when needed - refer to <https://golang.org/doc/effective_go.html#defer>. Also note that defers can be used for tracing.
- [ ] Use break loop to get out of eternal loops - refer to <https://golang.org/doc/effective_go.html#switch>.
- [ ] Check through all TODO markers.
- [ ] Get some proper organisation of the files in the project.
- [ ] Proper initialisation and creation of objects - refer to <https://golang.org/doc/effective_go.html#composite_literals>. Read particularly the initialisation of arrays, maps and slices.
- [ ] Separate logging out into own package.
- [ ] Implement sort interface (see <https://golang.org/pkg/sort/#Sort>) for custom types. This has already been done for HexLoc.
- [ ] Config can be sourced from database file (sqlite3 possibly)
- [ ] Get testing working.
- [ ] Include a "you rolled a xx for DM" on sLog.
- [ ] All temp screens/pages that work within a single func should have a "thisPage" identifier.
- [ ] Get rid of global variable "smell".
- [ ] New utility to clean traveller.log.
- [ ] Allow loading and re-editing of custom world, systems, sectors.

### To Do (tview) - OLD

- [x] Update to latest version.
- [ ] Check that app.Draw() is only used when necessary. (Note latest update that app.Draw() now have to be app.ForceDraw() - for some reason. This requires investigation.)
- [ ] Fork my own version. Consider contributing.

### General

- Communication with tm.com and ozemail guys about their programs.
- Generate extended star system details for existing canonical world data.
- Handle trading situations.
- Any other table that can be done automatically.
- Dice rolling handled ~~automatically~~ or manually.
- Provide choices for interactive mode so that the user doesn't have to refer to tables in the books (for instance skill selection).
- Allow searching and saving of worlds.

### Git and Golang stuff

- <https://dev.to/loderunner/working-with-forks-in-go-3ab6> - Reasonably recent info on github and go, working with forks.
- <https://kbroman.org/github_tutorial/pages/fork.html> - Contributing to someone's repository.
- <https://blog.sgmansfield.com/2016/06/working-with-forks-in-go/> - Older article though reasonable info.

### Database

- Remove subsectors from DB - but retain A-P link.
- Decent db init script. Revisit the one that is already there.
- Change world table to mainworld?
- Separate table for non-mainworlds (nMW)?
- JSON field for other info that can be stored with MW (or system). Like sophonts or temperature ranges.
- Define a JSON format for this field.
- Get JSON extension working if possible.
  - Recompile sqlite for it.
  - Recompile gosqlite3 for JSON.
  - Get it all working.
  - Blog about it.
- Maps into the database as BLOB.
- Look at alternate DBs:
  - Mariadb for RDB.
  - MongoDB for Document database.
  - What changes would be need for the NoSQL option?
- Grab whole sector from travellermap.com and insert into db (if we don't do that already!).

### World and Star System Generation

- Get Whole sector generation working:
  - [x] New (uncharted) sector.
  - [ ] New Imperial (non-OTU) sector.
- Star System extensions:
  - [ ] Extend basic information to T5SS - generate complete system.
  - [ ] World Builder Handbook extension (things like temperature and orbital eccentricity) -> JSON probably.
  - [ ] Re-generate extended T5SS info after editing a star, system or sector
- [ ] Generate a whole sector .XML file.
- [x] Generate a .tab file for whole sector.
- [ ] All T5SS reference tables into database.
- [ ] All tables into either DB or as functions.
- [ ] Searching for worlds that are already in the database. Database sourced from <https://travellermap.com> or other.
- [ ] Need a way of re-calculating the Extensions from a change in the Trade Classifications/Bases and zone settings.
  We need to be able to do this on a sector/subsector/world basis, but only for home-grown worlds, not the canon.
- [ ] Make sure Sector Generator (and other) code from Challenge #26 is included, incorporated or at least covered.
- [ ] Add options for saving worlds in either the database, ~~or tab file at the end of the generation process~~.
- [ ] Need to be able to save world and star systems that are not part of the OTU (otherwise uncharted)
- [ ] All the following types of Generation systems need to be fully completed, included saving worlds.
- [ ] Compare two worlds by Hex Location.

#### Sector generation

- [x] Have a Sector struct type that contains the collection of worlds when generated.
- [x] Sectors are displayed in a table and when you highlight a row, the underlying object is displayed in the object windows .
- [ ] Possible to edit (either in place or via a dialog) the world object that is displayed. For instance, you may want to change the name. Some fields should be locked and unable to be edited , but may want to add Scout Way stations or change travel zones.
- [ ] Get a list of things that we can edit.
- [ ] Table view of sector has a "Save" button. Perhaps a Flex or Grid for this. Tabbing is possible between the two sections (table and buttons). Ctrl-S to save.
- [x] Saving should save to the sector name, not generic sector.tab.
- [x] ~~Instead of having table containing Sector name, have it at the top (or title). Saves space.~~
  
#### Classic Traveller World Generation

Classic Traveller worlds work for Mongoose Traveller 1e and 2e.

- [x] Classic Traveller - Book 3 - Worlds and Adventures.
- [ ] Classic Traveller - Book 6 - Scouts.

#### MegaTraveller World Generation

- [x] MegaTraveller - Referee's Manual - basic generation.
- [ ] MegaTraveller - Referee's Manual - enhanced generation (of whole star systems).
- [ ] MegaTraveller - World Builder's Handbook.

#### Traveller5 World Generation

- [x] Traveller5 - Core Rules (Second Survey format) Basic mainworld generation.
- [ ] Traveller5 - Whole star system generation.

## Character Generation

The following is the list of sources for character generation, which has been
grabbed from the earlier person's program. Note that most of these are from
MegaTraveller (which may be OK for Classic - TBD):

### Classic Traveller Character Generation

- [ ] Classic Traveller - Basic Imperial (Book 1)
- [ ] Classic Traveller Advanced Generations
  - [ ] Mercenary (Book 4)
  - [ ] High Guard (Book 5)
  - [ ] Scouts (Book 6)
  - [ ] Merchant Prince (Book 7)
  - [ ] Robots (Book 8)
- [ ] CT Supplement Citizens of the Imperium (Book S4)
- [ ] Advanced Scientists: Challenge #29
- [ ] K'kree: Alien Module #2 K'kree by GDW
- [ ] Zhodani: Alien Module #4 Zhohani by GDW
- [ ] Droyne: Alien Module #5 Droyne by GDW
- [ ] Hiver: Alien Module #7 Hivers & Aliens of the Rim both by GDW
- [ ] Darrian: Alien Module #8 Darrians by GDW
- [ ] The Imperial Secret Service - White Dwarf #27
- [ ] Journalist: Travellers Digest #2/Challenge #27
- [ ] Law Enforcers: Challenge #30/Traveller's Digest #4
- [ ] Assassins: Third Imperium Magazine #11
- [ ] Hlanssai: JTAS #22
- [ ] Floriani: Third Imperium #8
- [ ] Newts: JTAS #11
- [ ] Virushi: JTAS #12
- [ ] Sword Worlders: JTAS #18
- [ ] Skyport Authority: JTAS #19 & JTAS #20
- [ ] Girug'kagh: JTAS #21
- [ ] Irklan: JTAS #23
- [ ] Prt': Challenge #26
  
### MegaTraveller Character Generation

- [ ] MegaTraveller Players Manual by GDW.
- [ ] Advanced Flyers: COACC by GDW
- [ ] Advanced IRIS: Challenge #33 &amp; #34
- [ ] NPC Personalities: Challenge #35
- [ ] Vilani & Vargr: MegaTraveller Aliens Vol. 1 by DGP
- [ ] Solomani & Aslan: MegaTraveller Aliens Vol. 2 by DGP
- [ ] Jonkeereen: The MegaTraveller Journal #3 by DGP (and Freelance Traveller #23)
- [ ] Dolphins: Trav. Digest #13, thanks to Fred Schiff
- [ ] Journalist: The Early Adventures by DGP
- [ ] Hhkar: Challenge #52
- [ ] Wet Navy: Challenge #53, #54, #60 <- May not be suitable
- [ ] Answerin: Challenge #54

### Traveller - The New Era Character Generation

- [ ] Traveller The New Era
- [ ] Ithklur: Aliens of the Rim by GDW
- [ ] Quick Start Characters: Challenge #75

### T4 Marc Miller's Traveller Charafcter Generation

- [ ] T4 Marc Miller's Traveller Core Rulebook

### Mongoose Traveller Character Generation

- [ ] Mongoose Traveller Core Rulebook
- [ ] Vegans: The Third Imperium - Solomani Rim
- [ ] Sword Worlders: The Third Imperium - Sword Worlds

### Traveller5 Character Generation

- [ ] Traveller5 Core Rules

### Mongoose Traveller 2nd Edition

- [ ] Mongoose Traveller 2e Core Rulebook
- [ ] Mongoose Traveller 2e Traveller Companion careers (Truther & Believer)

### Can't Find/To be sorted

- [ ] Advanced Belter, Pirate & Spy by Joe Walsh from Goeran's WWW site (defunct)
- [ ] Llellewyloly: TML article by James Kundert
- [ ] Githiaskio: Phil Masters article on WWW
- [ ] Luriani: Andrew Mofatt-Vallance's TML article
- [ ] Cafadi, & Irhadre: James Maliszewski's articles on WWW
- [ ] Girug'kagh: Loren Wiseman's articles on WWW
- [ ] Suerrat: Alien Module #8 & article by Charles Scott Kimball
- [ ] Llamiya: article by Tom H. (?) from TML CD
- [ ] Tirrils: Roger Myhre's "Contact: Tirrils. Updated for TNE" Also Gvurrdon Sector Campaign Book.

## Ideas

### Design Notes for Character Generation

Character generation is relatively easy for a serial asynchronous processes,
you just progress through the gen process, making decisions on the way and
asking for user input when you need it, no matter where in the process you
are. In a event-driven interface lacking true modal dialogs, it becomes very
difficult, if not impossible to break out of the event loop even temporarily
to ask the user to make a decision.

To solve this, we can look (partially) to the GOF State Design Pattern. This
has a virtual state object associated with the context object (as a Friend -
obviously this is not implementable in Go).

This has been (partially) done in the Worlds, where there are a number of different
"World Generation Type"s, and we switch on these. This has been implemented as a
custom type.

#### Questions to ask

- Do we make a generic character interface, and have the various generation or
  character styles separate object types implementing the interface?
- Where are the transitions done? In the State object, or in a function?
- If the State object is an interface, then it can't have actual fields, so how do we save
  state? Do we just do GetState(), SetState() and let the implementers decide state? Similarly
  with generation type/style.
- ~~Are MGT and MGT2 characters identical?~~ The generation process is different.

#### Code

```go
// characterState defines the generation state for a character. It determines
// what type of generation process is in place, and what stage the generation
// is up to.
//
// Transitions between character generation states may require user input,
// such as selection of skill tables or cascading of skills, or input of
// medic skill levels when a character ages. Additionally a transition will
// need to define the next stage of operations in the generation process, and
// this will vary depending on the current state, the type or style of
// generation process, and the stage of the process.
//
// The specialisation of what occurs in character generation for each specific
// generation style is what occurs in the implementing objects
type characterState interface {
    GetState() myState
    SetState(s myState)

    GetGenStyle() myStyle
    SetGenStyle(m myStyle)

    // Move from one State to another. This method decides on the new State.
    Transition()
}

type character struct {
    // ...etc

    // myState contains the current state of the character generation process.
    myState characterState

    // ...etc
}
```
