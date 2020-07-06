-- Creates the traveller database (rules and rule sets)
-- Version 1.0
-- 
-- This continues on with database tables representing the game rules and tables

-- The ruleset table describes the differing versions of the rules. This may or
-- may not correspond to a "milieu".
CREATE TABLE "ruleset" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"abbreviation"	TEXT NOT NULL,
	"ruleset_name"	TEXT NOT NULL,
	"comments"	TEXT
);

INSERT INTO ruleset ("abbreviation","ruleset_name","comments")
  VALUES
    ("CT","Classic Traveller","Original Traveller ruleset"),
    ("MT","MegaTraveller","2nd ed - collapse of Third Imperium and rebellion"),
    ("TNE","Traveller: The New Era","3rd ed"),
    ("T4","Marc Miller''s Traveller","4th ed - various milieus including 4th Imperium"),
    ("GURPS","GURPS Traveller","5th ed"),
    ("D20","Traveller D20","7th ed"),
    ("IW","GURPS Traveller Interstellar Wars","6th ed - based on the GURPS rules"),
    ("T5","Traveller5","12 ed - The Galaxiad"),
    ("MGT","Mongoose Traveller","Original Mongoose Spinoff set in the CT M1105 milieu"),
    ("MGT2","Mongoose Traveller 2nd Edition","2nd ed of Mongoose Traveller rules set in CT M1105 milieu");

-- The skill table 
CREATE TABLE "skill" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"skill_name"	TEXT NOT NULL,
	"description"	TEXT NOT NULL,
	"parent_skill"	INTEGER,
	"is_virtual"	INTEGER,
	"ruleset"	INTEGER
);

-- Classic Traveller skills
-- Virtual skills and basic skills.
INSERT INTO skill ("skill_name","description","parent_skill","is_virtual","ruleset")
  VALUES
    ("Brawling","Brawling is a general skill for hand-to-hand combat. It includes the use of hands, clubs, and bottles as weapons.",0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Blade Combat","Blade combat is a specific skill in the use of blades and polearms.",0,1,(SELECT id from ruleset where abbreviation="CT")),
    ("Gun Combat","Gun combat is a specific skill in the use of firearms.",0,1,(SELECT id from ruleset where abbreviation="CT")),
    ("Vehicle","The individual is skilled in the operation, use, and maintenance of a specific type vehicle commonly available in society",0,1,(SELECT id from ruleset where abbreviation="CT")),
    ("Administration",
     "The individual has had experience with bureaucratic agencies, and understands the requirements of dealing with them and managing them.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Air/Raft",
     "The individual has training and experience in the use and operation of the air/raft, floater, flier, and all types of grav vehicles.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("ATV",
     "The individual is acquainted with modern all terrain vehicles, and has been trained in, or has experience with, their operation. The term ATV (all terrain vehicle) includes AFV (armored fighting vehicle) within its meaning.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Bribery",
     "The individual has experience in bribing petty and not-so-petty officials in order to circumvent regulations or ignore cumbersome laws. Bribery skill does not guarantee success, but does minimize bad effects if the offer is rebuffed.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Computer",
     "The individual is skilled in the programming and operation of electronic and fibre optic computers, both ground and shipboard models.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Electronics",
     "The individual has skill in the use, operation, and repair of electronic devices. The person is considered handy in this field, with the equivalent of a green thumb talent. This skill includes the repair of energy weapons.",
    0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Engineering","The individual is skilled in the operation and maintenance of starship maneuver drives, jump drives, and power plants.",0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Forgery","The individual has a skill at faking documents and papers with a view to deceiving officials, banks, patrons, or other persons.",0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Forward Observer","The individual has been trained (in military service) to call on and adjust artillery (projectile, missile, and laser) fire from distant batteries and from ships in orbit.",0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Gambling",
     "The individual is well informed on games of chance, and wise in their play. He or she has an advantage over non-experts, and is generally capable of winning when engaged in such games. Gambling, however, should not be confused with general risk-taking.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Gunnery",
     "Gunnery is a skill in the use of weapons mounted on board starships (beam and pulse lasers, sandcasters, and missile launchers). This skill entitles the individual to the job title of gunner.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Jack of All Trades",
     "The individual is proven capable of handling a wide variety of situations, and is resourceful in finding solutions and remedies.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Leader",
     "The individual has led troops in battle (or on adventures) and is possessed of a knowledge and self-assurance which will make for a capable emergent or appointed leader",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Mechanical",
     "The individual has skill in the use, operation, and repair of mechanical devices. The person is considered to be handy in this field, with a talent similar to that of a green thumb. This skill specifically excludes the field of engineering; it does include non-energy weapon repair.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Medical","The individual has training and skill in the medical arts and sciences.",0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Navigation","The individual has training and expertise in the art and science of interplanetary and interstellar navigation.",0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Pilot",
     "The individual has training and experience in the operation of starships and large interplanetary ships. This skill encompasses both the interplanetary and the interstellar aspects of large ship operation.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Ship''s Boat",
     "The individual is familiar with the function and operation of small interplanetary craft collectively known as ship''s boats. These craft range in size from five to 100 tons, and include shuttles, life-boats, launches, ship''s boats, and fighters.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Steward","The individual is experienced and capable in the care and feeding of passengers: the duties of the ship''s steward.",0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Streetwise",
     "The individual is acquainted with the ways of local subcultures (which tend to be the same everywhere in human society), and thus is capable of dealing with strangers without alienating them. This skill is not the same as alien contact experience.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Tactics",
     "The individual has training and experience in small unit tactics (up to and including units of 1000 troops or individual spaceships). This skill is not to be confused with strategy, which deals with the reasons for the encounter and the intended results of the encounter; strategy is the realm of the players, rather than the characters.",
     0,0,(SELECT id from ruleset where abbreviation="CT")),
    ("Vacc Suit",
     "The individual has been trained and has experience in the use of the standard vacuum suit (space suit), including armored battle dress and suits for use on various planetary surfaces in the presence of exotic, corrosive, or insidious atmospheres.",
    0,0,(SELECT id from ruleset where abbreviation="CT"));

-- Specific skills for the gun/blade combat and vehicles.
INSERT INTO skill ("skill_name","description","parent_skill","is_virtual","ruleset")
  VALUES
    ("Dagger","A small knife weapon with a flat, two-edged blade approximately 200mm in length.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Blade","A hybrid knife weapon with a heavy, flat two-edged blade nearly 300mm in length, and a semi-basket handguard.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Foil","Also known as the rapier, this weapon is a light, sword-like weapon with a pointed, edged blade 800mm in length, and a basket or cup hilt to protect the hand.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Sword","The standard long-edged weapon, featuring a flat, two-edged blade. It may or may not have a basket hilt or hand protector.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Cutlass","A heavy, flat-bladed, single-edged weapon featuring a full basket hilt to protect the hand. The cutlass is the standard shipboard blade weapon and usually kept in brackets on the bulkhead near important locations; when worn, a belt scabbard is used.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Broadsword","The largest of the sword weapons, also called the two-handed sword because it requires both hands to swing. The blade is extremely heavy, two-edged, and about 1000 to 1200mm in length.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Bayonet","A knife-like weapon similar to a dagger or blade.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Spear","A long (3000mm) polearm with a pointed tip, usually of metal.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Halberd","A quite elaborate polearm featuring a pointed, bladed tip.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Pike","A long (3000 to 4000mm) polearm with some form of flat blade tip.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Cudgel","A basic stick used as a weapon.",(SELECT id from skill where skill_name="Blade Combat"),0,(SELECT id from ruleset where abbreviation="CT")),

    ("Body Pistol","A small, non-metallic semi-automatic pistol designed to evade detection by most weapon detectors.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Auto Pistol","The basic repeating handgun, firing 9mm caliber bullets (each weighing approximately 10 grams) at velocities from 400 to 500 meters per second.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Revolver","An older variety of handgun, the revolver fires 9mm bullets with characteristics similar to those fired by the automatic pistol, but not interchangeable with them.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Carbine","A short type of rifle firing a small caliber round (a 6mm bullet, weighing 5 grams, at a velocity of 900 meters per second).",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Rifle","The standard military arm, firing a 7mm, 10 gram bullet at a velocity of approximately 900 meters per second.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Auto Rifle","A highly refined and tuned version of the rifle, capable of full automatic fire as well as semi-automatic shots.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Shotgun","The basic weapon for maximum shock effect without regard to accuracy.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("SMG","A small automatic weapon designed to fire pistol ammunition.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Laser Carbine","A lightweight version of the laser rifle, firing high energy bolts using current from a backpack battery/power pack.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Laser Rifle","The standard high energy weapon, firing high energy bolts in the same manner as the laser carbine.",(SELECT id from skill where skill_name="Gun Combat"),0,(SELECT id from ruleset where abbreviation="CT")),
    
    ("Ground Car","This class includes all wheeled or tracked vehicles.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Watercraft","Water craft require only one skilled crew member if under 50 tons displacement.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Winged Craft","Winged aircraft generate lift by passing air over wing-surfaces, either fixed or rotating.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Hovercraft","Ground effect vehicles are supported on a cushion of air (at about 1 to 3 meters altitude).",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Grav Belt","Personal anti-gravity transportation using a single null-gravity module and a personal harness.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT"));

-- After consulting Errata for Book 1:
-- This is the complete list of Vehicle skills:Aircraft (select Helicopter, Propeller-driven Fixed Wing, or Jet-driven Fixed Wing), Grav
-- Vehicle, Tracked Vehicle, Wheeled Vehicle, and Watercraft (select Small Watercraft, Large Watercraft, Hovercraft, or Submersible)

DELETE FROM skill where "skill_name" in ("Ground Car","Watercraft","Winged Craft","Grav Belt")

-- The corrections for this are as follows:
INSERT INTO skill ("skill_name","description","parent_skill","is_virtual","ruleset")
  VALUES
    ("Propeller-driven Fixed Wing","This Aircraft class includes all fixed-wing propeller aircraft..",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Jet-driven Fixed Wing","This Aircraft class includes all fixed-wing Jet aircraft.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Grav Vehicle","Includes Grav Belts and other Grav vehicles.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Tracked Vehicle","Included AFVs and tractors.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Wheeled Vehicle","All forms of wheeled ground vehicles.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Small Watercraft","Generally, small pliesure or commercial boats.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Large Watercraft","Large tonnage commercial sea-going vessels.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Submersible","All forms of submersible from one-man submarines to large U-Boats.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT")),
    ("Helicopter","All forms of rotary-winged aircraft.",(SELECT id from skill where skill_name="Vehicle"),0,(SELECT id from ruleset where abbreviation="CT"));

-- The rulebook table describes the differing rule books for the games. Each rulebook belongs
-- to a specific version of the rules (ruleset_id). This can even refer to articles. Tables, 
-- skills, generation sequences, anything that differs with the varying versions of the game
-- should specify a rulebook rather than a ruleset. Some rulesets may only have one rulebook,
-- and not every "book" that was ever published should be a rulebook.
CREATE TABLE "rulebook" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"abbreviation"	TEXT NOT NULL,
	"rulebook_name"	TEXT NOT NULL,
  "ruleset_id" INTEGER NOT NULL,
	"comments"	TEXT
);

INSERT INTO rulebook ("abbreviation","rulebook_name","ruleset_id","comments")
  VALUES
  -- Classic Traveller
    ("CT B01","Book 1",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Book 1 - Characters and Combat"),
    ("CT B02","Book 2",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Book 2 - Starships"),
    ("CT B03","Book 3",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Book 3 - Worlds and Adventures"),
    ("CT B04","Book 4",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Book 4 - Mercenary"),
    ("CT B05","Book 5",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Book 5 - High Guard"),
    ("CT B06","Book 6",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Book 6 - Scouts"),
    ("CT B07","Book 7",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Book 7 - Merchant Prince"),
    ("CT B08","Book 8",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Book 8 - Robots"),
    ("CT AM01","Alien Module 1",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Alien Module 1 - Aslan"),
    ("CT AM02","Alien Module 2",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Alien Module 2 - K''kree"),
    ("CT AM03","Alien Module 3",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Alien Module 3 - Vargr"),
    ("CT AM04","Alien Module 4",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Alien Module 4 - Zhodani"),
    ("CT AM05","Alien Module 5",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Alien Module 5 - Droyne"),
    ("CT AM06","Alien Module 6",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Alien Module 6 - Solomani"),
    ("CT AM07","Alien Module 7",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Alien Module 7 - Hivers"),
    ("CT AM08","Alien Module 8",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Alien Module 8 - Darrians"),
    ("Challenge #28","Challenge Magazine #28",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Challenge Magazine #28"),
    ("WD #27","White Dwarf Magazine #28",(SELECT id FROM ruleset WHERE abbreviation="CT"),"White Dwarf Magazine #28"),
    ("TD #2","Travellers Digest Magazine #2",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Travellers Digest Magazine #2"),
    ("TD #4","Travellers Digest Magazine #4",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Travellers Digest Magazine #4"),
    ("TIM #11","Third Imperium Magazine #11",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Third Imperium Magazine #11"),
    ("CT S04","Supplement 4",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Classic Traveller Supplement 4 - Citizens of the Imperium"),
    ("JTAS #22","Journal of the Travellers Aid Society #22",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Journal of the Travellers Aid Society #22"),
    ("TIM #8","Third Imperium Magazine #8",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Third Imperium Magazine #8")
    ("JTAS #11","Journal of the Travellers Aid Society #11",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Journal of the Travellers Aid Society #11"),
    ("JTAS #12","Journal of the Travellers Aid Society #12",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Journal of the Travellers Aid Society #12"),
    ("JTAS #18","Journal of the Travellers Aid Society #18",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Journal of the Travellers Aid Society #18"),
    ("JTAS #19","Journal of the Travellers Aid Society #19",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Journal of the Travellers Aid Society #19"),
    ("JTAS #21","Journal of the Travellers Aid Society #21",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Journal of the Travellers Aid Society #21"),
    ("JTAS #23","Journal of the Travellers Aid Society #23",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Journal of the Travellers Aid Society #23"),
    ("Challenge #26","Challenge Magazine #26",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Challenge Magazine #26"),
    ("Challenge #27","Challenge Magazine #27",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Challenge Magazine #27"),
    ("Challenge #30","Challenge Magazine #30",(SELECT id FROM ruleset WHERE abbreviation="CT"),"Challenge Magazine #30"),
-- MegaTraveller
    ("MT PM", "MegaTraveller Player''s Manual",(SELECT id FROM ruleset WHERE abbreviation="MT"),"MegaTraveller Player''s Manual"),
    ("MT RM","MegaTraveller Referee''s Manual",(SELECT id FROM ruleset WHERE abbreviation="MT"),"MegaTraveller Referee''s Manual"),
    ("MT COACC", "MegaTraveller COACC",(SELECT id FROM ruleset WHERE abbreviation="MT"),"MegaTraveller Close Orbit and Airspace Control Command"),
    ("MT WBH", "World Builder''s Handbook",(SELECT id FROM ruleset WHERE abbreviation="MT"),"MegaTraveller World Builder''s Handbook - Digest Group Publications"),
    ("MT A01", "MegaTraveller Alien Volume 1",(SELECT id FROM ruleset WHERE abbreviation="MT"),"MegaTraveller Alien Volume 1 - Vilani and Vargr"),
    ("MT A02", "MegaTraveller Alien Volume 2",(SELECT id FROM ruleset WHERE abbreviation="MT"),"MegaTraveller Alien Volume 2 - Solomani and Aslan"),
    ("MT J03", "The MegaTraveller Journal Volume 3",(SELECT id FROM ruleset WHERE abbreviation="MT"),"The MegaTraveller Journal Volume 3"),
    ("Challenge #34","Challenge Magazine #34",(SELECT id FROM ruleset WHERE abbreviation="MT"),"Challenge Magazine #34"),
    ("Challenge #52","Challenge Magazine #52",(SELECT id FROM ruleset WHERE abbreviation="MT"),"Challenge Magazine #52"),
    ("MT TEA", "MegaTraveller The Early Adventures",(SELECT id FROM ruleset WHERE abbreviation="MT"),"MegaTraveller The Early Adventures - Background and Adventures from Travellers'' Digest Issues 1 - 4 - Digest Group Publications"),
    ("TTD #13", "The Travellers'' Digest #13",(SELECT id FROM ruleset WHERE abbreviation="MT"),"The Travellers'' Digest #13"),
    ("Challenge #33","Challenge Magazine #33",(SELECT id FROM ruleset WHERE abbreviation="MT"),"Challenge Magazine #33"),
    ("Challenge #35","Challenge Magazine #35",(SELECT id FROM ruleset WHERE abbreviation="MT"),"Challenge Magazine #35"),
    ("Challenge #53","Challenge Magazine #53",(SELECT id FROM ruleset WHERE abbreviation="MT"),"Challenge Magazine #53"),
    ("Challenge #54","Challenge Magazine #54",(SELECT id FROM ruleset WHERE abbreviation="MT"),"Challenge Magazine #54"),
    ("Challenge #60","Challenge Magazine #60",(SELECT id FROM ruleset WHERE abbreviation="MT"),"Challenge Magazine #60"),    
-- Traveller - The New Era
    ("TNE Core", "Traveller The New Era Core Rulebook",(SELECT id FROM ruleset WHERE abbreviation="TNE"),"Traveller The New Era Core Rulebook"),
    ("TNE AR", "Aliens of the Rim",(SELECT id FROM ruleset WHERE abbreviation="TNE"),"Traveller The New Era - Aliens of the Rim"),
    ("Challenge #75","Challenge Magazine #75",(SELECT id FROM ruleset WHERE abbreviation="TNE"),"Challenge Magazine #75"),
-- T4 Marc Miller's Traveller
    ("T4 Core", "Marc Miller''s Traveller Core Rulebook",(SELECT id FROM ruleset WHERE abbreviation="T4"),"Marc Miller''s Traveller Core Rulebook"),
-- Mongoose Traveller
    ("MGT Core", "Mongoose Traveller Core Rulebook",(SELECT id FROM ruleset WHERE abbreviation="MGT"),"Mongoose Traveller (1st Edition) Core Rulebook"),
-- Traveller5
    ("T5 Core", "Traveller5 Core Rulebook",(SELECT id FROM ruleset WHERE abbreviation="T5"),"Traveller5 Core Rulebook"),
-- Mongoose Traveller 2nd Edition
    ("MGT2 Core", "Mongoose Traveller 2e Core Rulebook",(SELECT id FROM ruleset WHERE abbreviation="MGT2"),"Mongoose Traveller 2nd Edition Core Rulebook")

-- Modify the Skill table for the new "rulebook" based way.
BEGIN TRANSACTION;
ALTER TABLE skill rename to skill_old;
CREATE TABLE "skill" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"skill_name"	TEXT NOT NULL,
	"description"	TEXT NOT NULL,
	"parent_skill"	INTEGER,
	"is_virtual"	INTEGER,
	"rulebook_id"	INTEGER
);

INSERT INTO skill ("id","skill_name","description","parent_skill","is_virtual","rulebook_id")
SELECT "id","skill_name","description","parent_skill","is_virtual","ruleset" FROM skill_old;

DROP TABLE skill_old
COMMIT;

-- Now change the data so it is pointing to Book 1 instead of just to CT
UPDATE skill SET rulebook_id = (SELECT id FROM rulebook WHERE abbreviation='CT B01')

-- I've found this one - an update for the Crenduthaar.
UPDATE RACE 
SET homeworld_id = (SELECT id FROM world WHERE name = 'Ghatsokie')
WHERE race_name = 'Crenduthaar';

-- The Hlanssai
UPDATE RACE 
SET homeworld_id = (SELECT id FROM world WHERE name = 'Vrirhlanz')
WHERE race_name = 'Hlanssai';

-- Let's fix them all while we are here.
UPDATE race
SET homeworld_id = (SELECT id FROM world WHERE name = race.homeworld)
WHERE race.homeworld_id = -1 AND race.homeworld <> 'Unknown';

UPDATE race
SET homeworld_id = -1
WHERE homeworld_id IS NULL;

-- Bad initial data for this one
UPDATE race
SET homeworld = "Ul" WHERE race_name = "Ulane"

UPDATE race
SET homeworld_id = (SELECT id FROM world WHERE name = 'Ul')
WHERE race.homeworld = 'Ul';



