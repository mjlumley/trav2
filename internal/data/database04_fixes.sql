-- Hopefully fixes https://github.com/mjlumley/traveller/issues/12
--
UPDATE subsector SET capital_id = -1;
--
UPDATE subsector SET capital_id = 
  (SELECT world.id FROM world WHERE remarks like('%Cp%') AND 
   subsector.sector_id = world.sector_id AND
   subsector.sector_index = world.subsector_index);

UPDATE subsector SET capital_id=226 WHERE id=8;
UPDATE subsector SET capital_id=251 WHERE id=9;
UPDATE subsector SET capital_id=316 WHERE id=11;
UPDATE subsector SET capital_id=692 WHERE id=26;
UPDATE subsector SET capital_id=853 WHERE id=33;
UPDATE subsector SET capital_id=863 WHERE id=34;
--UPDATE subsector SET capital_id=885 WHERE id=35;
UPDATE subsector SET capital_id=912 WHERE id=37;
UPDATE subsector SET capital_id=999 WHERE id=40;
UPDATE subsector SET capital_id=1039 WHERE id=42;
--UPDATE subsector SET capital_id=1550 WHERE id=75;
UPDATE subsector SET capital_id=1848 WHERE id=87;
UPDATE subsector SET capital_id=2203 WHERE id=97;
UPDATE subsector SET capital_id=2352 WHERE id=102;
UPDATE subsector SET capital_id=2479 WHERE id=106;
--UPDATE subsector SET capital_id=2699 WHERE id=113;
--UPDATE subsector SET capital_id=2704 WHERE id=113;
UPDATE subsector SET capital_id=2781 WHERE id=115;
UPDATE subsector SET capital_id=2850 WHERE id=121;
UPDATE subsector SET capital_id=2948 WHERE id=126;
UPDATE subsector SET capital_id=2962 WHERE id=127;
UPDATE subsector SET capital_id=3168 WHERE id=151;
UPDATE subsector SET capital_id=3780 WHERE id=168;
UPDATE subsector SET capital_id=66225 WHERE id=210;
--UPDATE subsector SET capital_id=66834 WHERE id=235;
UPDATE subsector SET capital_id=68109 WHERE id=267;
UPDATE subsector SET capital_id=69710 WHERE id=325;
UPDATE subsector SET capital_id=70571 WHERE id=358;
UPDATE subsector SET capital_id=70655 WHERE id=361;
--UPDATE subsector SET capital_id=70781 WHERE id=366;
--UPDATE subsector SET capital_id=70818 WHERE id=366;
UPDATE subsector SET capital_id=70825 WHERE id=367;
UPDATE subsector SET capital_id=70867 WHERE id=368;
UPDATE subsector SET capital_id=71249 WHERE id=381;
UPDATE subsector SET capital_id=71754 WHERE id=399;
UPDATE subsector SET capital_id=71919 WHERE id=405;
--UPDATE subsector SET capital_id=73232 WHERE id=435;
UPDATE subsector SET capital_id=73481 WHERE id=450;
UPDATE subsector SET capital_id=73516 WHERE id=452;
UPDATE subsector SET capital_id=73547 WHERE id=453;
UPDATE subsector SET capital_id=73586 WHERE id=454;
UPDATE subsector SET capital_id=73638 WHERE id=456;
--UPDATE subsector SET capital_id=73754 WHERE id=460;
--UPDATE subsector SET capital_id=73770 WHERE id=460;
--UPDATE subsector SET capital_id=73810 WHERE id=462;
--UPDATE subsector SET capital_id=73815 WHERE id=462;
--UPDATE subsector SET capital_id=74257 WHERE id=476;
UPDATE subsector SET capital_id=74448 WHERE id=482;
--UPDATE subsector SET capital_id=75561 WHERE id=514;
--UPDATE subsector SET capital_id=75590 WHERE id=515;
UPDATE subsector SET capital_id=76075 WHERE id=540;
--UPDATE subsector SET capital_id=76820 WHERE id=563;
UPDATE subsector SET capital_id=77475 WHERE id=586;
UPDATE subsector SET capital_id=78666 WHERE id=638;
--UPDATE subsector SET capital_id=79704 WHERE id=671;
UPDATE subsector SET capital_id=80461 WHERE id=708;
UPDATE subsector SET capital_id=80633 WHERE id=714;
UPDATE subsector SET capital_id=80672 WHERE id=715;
--UPDATE subsector SET capital_id=80692 WHERE id=716;
--UPDATE subsector SET capital_id=80706 WHERE id=716;
--UPDATE subsector SET capital_id=80710 WHERE id=716;
--UPDATE subsector SET capital_id=80785 WHERE id=719;
--UPDATE subsector SET capital_id=80790 WHERE id=719;
--UPDATE subsector SET capital_id=80798 WHERE id=719;
--UPDATE subsector SET capital_id=80805 WHERE id=720;
--UPDATE subsector SET capital_id=80811 WHERE id=720;
--UPDATE subsector SET capital_id=80813 WHERE id=720;
--UPDATE subsector SET capital_id=80855 WHERE id=721;
UPDATE subsector SET capital_id=81174 WHERE id=722;
UPDATE subsector SET capital_id=82405 WHERE id=756;
UPDATE subsector SET capital_id=83670 WHERE id=807;
UPDATE subsector SET capital_id=83701 WHERE id=809;
UPDATE subsector SET capital_id=83723 WHERE id=811;
UPDATE subsector SET capital_id=83748 WHERE id=812;
UPDATE subsector SET capital_id=83761 WHERE id=813;
--UPDATE subsector SET capital_id=83812 WHERE id=816;
--UPDATE subsector SET capital_id=83814 WHERE id=816;
--UPDATE subsector SET capital_id=83834 WHERE id=817;
--UPDATE subsector SET capital_id=83851 WHERE id=817;
UPDATE subsector SET capital_id=85356 WHERE id=846;
UPDATE subsector SET capital_id=85401 WHERE id=849;
UPDATE subsector SET capital_id=85457 WHERE id=851;
--UPDATE subsector SET capital_id=85484 WHERE id=853;
--UPDATE subsector SET capital_id=85494 WHERE id=853;
UPDATE subsector SET capital_id=85592 WHERE id=857;
UPDATE subsector SET capital_id=85631 WHERE id=858;
--UPDATE subsector SET capital_id=86138 WHERE id=888;
UPDATE subsector SET capital_id=86506 WHERE id=913;
UPDATE subsector SET capital_id=86964 WHERE id=924;
UPDATE subsector SET capital_id=87012 WHERE id=927;
UPDATE subsector SET capital_id=87031 WHERE id=928;
UPDATE subsector SET capital_id=87053 WHERE id=929;
UPDATE subsector SET capital_id=87101 WHERE id=931;
UPDATE subsector SET capital_id=87136 WHERE id=933;
UPDATE subsector SET capital_id=87255 WHERE id=937;
UPDATE subsector SET capital_id=87263 WHERE id=938;
UPDATE subsector SET capital_id=87299 WHERE id=939;
--UPDATE subsector SET capital_id=87969 WHERE id=965;
UPDATE subsector SET capital_id=88126 WHERE id=970;
UPDATE subsector SET capital_id=90202 WHERE id=1022;
UPDATE subsector SET capital_id=90233 WHERE id=1023;
UPDATE subsector SET capital_id=90248 WHERE id=1024;
UPDATE subsector SET capital_id=90400 WHERE id=1030;
UPDATE subsector SET capital_id=90439 WHERE id=1031;
UPDATE subsector SET capital_id=91765 WHERE id=1093;
UPDATE subsector SET capital_id=93772 WHERE id=1149;
--UPDATE subsector SET capital_id=93855 WHERE id=1151;
--UPDATE subsector SET capital_id=93889 WHERE id=1152;
UPDATE subsector SET capital_id=93932 WHERE id=1153;
UPDATE subsector SET capital_id=94016 WHERE id=1155;
UPDATE subsector SET capital_id=94095 WHERE id=1157;
UPDATE subsector SET capital_id=94267 WHERE id=1161;
--UPDATE subsector SET capital_id=94345 WHERE id=1163;
--UPDATE subsector SET capital_id=96414 WHERE id=1216;
--UPDATE subsector SET capital_id=96420 WHERE id=1216;
--UPDATE subsector SET capital_id=97266 WHERE id=1230;
--UPDATE subsector SET capital_id=97326 WHERE id=1232;
--UPDATE subsector SET capital_id=97386 WHERE id=1234;
UPDATE subsector SET capital_id=99344 WHERE id=1288;
UPDATE subsector SET capital_id=100299 WHERE id=1317;
UPDATE subsector SET capital_id=100967 WHERE id=1345;
UPDATE subsector SET capital_id=101195 WHERE id=1355;
UPDATE subsector SET capital_id=101577 WHERE id=1367;
UPDATE subsector SET capital_id=102184 WHERE id=1376;
UPDATE subsector SET capital_id=104167 WHERE id=1444;
UPDATE subsector SET capital_id=107364 WHERE id=1549;
UPDATE subsector SET capital_id=108038 WHERE id=1571;
UPDATE subsector SET capital_id=109178 WHERE id=1590;
UPDATE subsector SET capital_id=109689 WHERE id=1609;
UPDATE subsector SET capital_id=110119 WHERE id=1624;
UPDATE subsector SET capital_id=110355 WHERE id=1633;
UPDATE subsector SET capital_id=110379 WHERE id=1634;
UPDATE subsector SET capital_id=111080 WHERE id=1659;
UPDATE subsector SET capital_id=111136 WHERE id=1661;
--UPDATE subsector SET capital_id=111171 WHERE id=1662;
--UPDATE subsector SET capital_id=111180 WHERE id=1662;
UPDATE subsector SET capital_id=111205 WHERE id=1663;
--UPDATE subsector SET capital_id=111294 WHERE id=1667;
--UPDATE subsector SET capital_id=111301 WHERE id=1667;
UPDATE subsector SET capital_id=113428 WHERE id=1753;
UPDATE subsector SET capital_id=116250 WHERE id=1851;
--UPDATE subsector SET capital_id=117232 WHERE id=1899;
--UPDATE subsector SET capital_id=118625 WHERE id=1989;
--UPDATE subsector SET capital_id=118640 WHERE id=1989;
UPDATE subsector SET capital_id=118840 WHERE id=1998;
UPDATE subsector SET capital_id=118883 WHERE id=1999;
UPDATE subsector SET capital_id=118961 WHERE id=2003;
UPDATE subsector SET capital_id=118983 WHERE id=2004;
UPDATE subsector SET capital_id=120153 WHERE id=2066;
UPDATE subsector SET capital_id=120205 WHERE id=2068;
--UPDATE subsector SET capital_id=120585 WHERE id=2086;
--UPDATE subsector SET capital_id=120613 WHERE id=2086;
--UPDATE subsector SET capital_id=120661 WHERE id=2088;
--UPDATE subsector SET capital_id=120686 WHERE id=2088;
--UPDATE subsector SET capital_id=120728 WHERE id=2089;
--UPDATE subsector SET capital_id=120793 WHERE id=2091;
--UPDATE subsector SET capital_id=120797 WHERE id=2091;
--UPDATE subsector SET capital_id=120802 WHERE id=2091;
--UPDATE subsector SET capital_id=120805 WHERE id=2091;
UPDATE subsector SET capital_id=120921 WHERE id=2094;
UPDATE subsector SET capital_id=121047 WHERE id=2098;
UPDATE subsector SET capital_id=121141 WHERE id=2100;
UPDATE subsector SET capital_id=121606 WHERE id=2108;
UPDATE subsector SET capital_id=122484 WHERE id=2152;
--UPDATE subsector SET capital_id=123733 WHERE id=2186;
--UPDATE subsector SET capital_id=123768 WHERE id=2188;
--UPDATE subsector SET capital_id=123792 WHERE id=2189;
--UPDATE subsector SET capital_id=123811 WHERE id=2189;
UPDATE subsector SET capital_id=123834 WHERE id=2190;
UPDATE subsector SET capital_id=124578 WHERE id=2206;
--UPDATE subsector SET capital_id=125635 WHERE id=2243;

-- Fix the SECTOR/SUBSECTOR issue
-- Modify the Skill table for the new "rulebook" based way.
BEGIN TRANSACTION;
ALTER TABLE subsector rename to subsector_old;
CREATE TABLE "subsector" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name"	TEXT NOT NULL,
	"lang_id"	INTEGER NOT NULL,
	"sector_id"	INTEGER NOT NULL,
	"subsector_index"	TEXT NOT NULL,
	"capital_id"	INTEGER,
	"remarks"	TEXT
);

INSERT INTO subsector ("id","name","lang_id","sector_id","subsector_index","capital_id","remarks")
SELECT "id","name","lang_id","sector_id","sector_index","capital_id","remarks" FROM subsector_old;

DROP TABLE subsector_old
COMMIT;

-- This code needs to be in there.
INSERT INTO allegiance ("code", "legacy_code", "allegiance_name")
VALUES ("ImXX","Im","Third Imperium")

